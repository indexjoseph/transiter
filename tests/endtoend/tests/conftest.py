import io
import os
import zipfile

import pytest
import requests


class SourceServerClient:
    def __init__(self, base_url):
        self._created_urls = []
        self._base_url = base_url

    def put(self, url, content):
        requests.put(self._base_url + "/" + url, data=content).raise_for_status()

    def create(self, prefix="", suffix=""):
        response = requests.post(
            self._base_url, params={"prefix": prefix, "suffix": suffix}
        )
        response.raise_for_status()
        created_url = response.text
        self._created_urls.append(created_url)
        return created_url

    def __del__(self):
        for url in self._created_urls:
            requests.delete(self._base_url + "/" + url)


@pytest.fixture
def source_server():
    return SourceServerClient(
        os.environ.get("SOURCE_SERVER_HOST", "http://localhost:5001")
    )


@pytest.fixture
def transiter_host():
    return os.environ.get("TRANSITER_HOST", "http://localhost:8000")


@pytest.fixture
def source_server_host_within_transiter():
    return os.environ.get(
        "SOURCE_SERVER_HOST_WITHIN_TRANSITER",
        os.environ.get("SOURCE_SERVER_HOST", "http://localhost:5001"),
    )


def get_zip():
    output_bytes = io.BytesIO()
    # writing files to a zipfile
    data_dir = os.path.join(os.path.dirname(__file__), "data", "gtfsstatic")
    with zipfile.ZipFile(output_bytes, "w") as zip_file:
        for file_name in os.listdir(data_dir):
            full_path = os.path.join(data_dir, file_name)
            zip_file.write(full_path, arcname=file_name)
    return output_bytes.getvalue()


def get_config():
    file_path = os.path.join(os.path.dirname(__file__), "data", "system-config.yaml")
    with open(file_path, "r") as f:
        config = f.read()
    return config


@pytest.fixture
def install_system_1(
    request,
    source_server: SourceServerClient,
    transiter_host,
    source_server_host_within_transiter,
):
    def install(system_id, realtime_auto_update_period="1 day"):
        def delete():
            requests.delete(transiter_host + "/systems/" + system_id)

        delete()

        system_config_url = source_server.create(
            "", "/" + system_id + "/system-config.yaml.jinja"
        )
        static_feed_url = source_server.create("", "/" + system_id + "/gtfs-static.zip")
        realtime_feed_url = source_server.create(
            "", "/" + system_id + "/gtfs-realtime.gtfs"
        )

        source_server.put(static_feed_url, get_zip())
        source_server.put(
            system_config_url,
            get_config().format(
                static_feed_url=source_server_host_within_transiter
                + "/"
                + static_feed_url,
                realtime_feed_url=source_server_host_within_transiter
                + "/"
                + realtime_feed_url,
                realtime_auto_update_period=realtime_auto_update_period,
            ),
        )

        response = requests.put(
            transiter_host + "/systems/" + system_id,
            data={
                "config_file": source_server_host_within_transiter
                + "/"
                + system_config_url
            },
        )
        response.raise_for_status()

        request.addfinalizer(delete)

        return static_feed_url, realtime_feed_url

    return install