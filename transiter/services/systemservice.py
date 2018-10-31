from ..database import connection
from ..database import models
from transiter.database.daos import direction_name_dao, feed_dao, route_dao, stop_dao, station_dao, system_dao
from . import exceptions
from transiter.utils import gtfsstaticutil
import csv
import os
from transiter.utils import linksutil
from transiter.utils import servicepatternmanager
import yaml

@connection.unit_of_work
def list_all():
    """
    List all installed systems.
    :return: A list of short representation of systems
    """
    response = []
    for system in system_dao.list_all():
        system_response = system.short_repr()
        system_response.update({
            'href': linksutil.SystemEntityLink(system)
        })
        response.append(system_response)
    return response


@connection.unit_of_work
def get_by_id(system_id):
    system = system_dao.get_by_id(system_id)
    if system is None:
        raise exceptions.IdNotFoundError
    response = system.short_repr()
    response.update({
        "stops": {
            "count": system_dao.count_stops_in_system(system_id),
            "href": linksutil.StopsInSystemIndexLink(system)
        },
        "stations": {
            "count": system_dao.count_stations_in_system(system_id),
            "href": "NI"
        },
        "routes": {
            "count": system_dao.count_routes_in_system(system_id),
            "href": linksutil.RoutesInSystemIndexLink(system)
        },
        "feeds": {
            "count": system_dao.count_feeds_in_system(system_id),
            "href": linksutil.FeedsInSystemIndexLink(system)
        }
    })
    return response


@connection.unit_of_work
def install(system_id):
    if system_dao.get_by_id(system_id) is not None:
        return False

    system = system_dao.create()
    system.system_id = system_id

    _import_static_data(system)
    return True


@connection.unit_of_work
def delete_by_id(system_id):
    deleted = system_dao.delete_by_id(system_id)
    if not deleted:
        raise exceptions.IdNotFoundError
    return True



def _read_csv_file(file_path):
    with open(file_path, mode='r') as csv_file:
        csv_reader = csv.DictReader(csv_file)
        yield from csv_reader


def _import_static_data(system):




    system_base_dir = os.path.join(
        os.path.dirname(__file__),
        '../systems',
        system.system_id
        )
    agency_data_dir = os.path.join(system_base_dir, 'agencydata')
    custom_data_dir = os.path.join(system_base_dir, 'customdata')
    print(agency_data_dir)

    config_file_path = os.path.join(system_base_dir, 'config.yaml')
    system_config = SystemConfig(config_file_path)

    gtfs_static_parser = gtfsstaticutil.GtfsStaticParser()
    gtfs_static_parser.parse_from_directory(agency_data_dir)

    for route in gtfs_static_parser.route_id_to_route.values():
        route.system = system

    station_sets_by_stop_id = {}
    for stop in gtfs_static_parser.stop_id_to_stop.values():
        stop.system = system
        station_sets_by_stop_id[stop.stop_id] = {stop.stop_id}

    for (stop_id_1, stop_id_2) in gtfs_static_parser.transfer_tuples:
        updated_station_set = station_sets_by_stop_id[stop_id_1].union(
            station_sets_by_stop_id[stop_id_2])
        for stop_id in updated_station_set:
            station_sets_by_stop_id[stop_id] = updated_station_set

    for station_set in station_sets_by_stop_id.values():
        # TODO: option to make this 1 also so stations only multistop
        if len(station_set) == 0:
            continue
        station = models.Station()
        for stop_id in station_set:
            gtfs_static_parser.stop_id_to_stop[stop_id].station = station
        station.system = system
        station_set.clear()

    for stop_alias in gtfs_static_parser.stop_id_alias_to_stop_alias.values():
        stop_id = stop_alias.stop_id
        stop = gtfs_static_parser.stop_id_to_stop[stop_id]
        stop_alias.stop = stop

    servicepatternmanager.construct_sps_from_gtfs_static_data(
        gtfs_static_parser,
        system_config.static_route_sps,
        system_config.static_other_sps,
    )

    direction_name_rules_files = system_config.direction_name_rules_files
    priority = 0
    for direction_name_rules_file_path in direction_name_rules_files:
        full_path = os.path.join(custom_data_dir, direction_name_rules_file_path)
        with open(full_path) as csv_file:
            csv_reader = csv.DictReader(csv_file)
            for row in csv_reader:
                # TODO: allow either stop_id or stop_id alias
                stop_id = row['stop_id']
                stop = gtfs_static_parser.stop_id_to_stop.get(stop_id, None)
                if stop is None:
                    continue
                direction_id = row.get('direction_id', None)
                if direction_id is not None:
                    direction_id = (direction_id == '0')
                direction_name_rule = models.DirectionNameRule()
                direction_name_rule.stop = stop
                direction_name_rule.priority = priority
                direction_name_rule.direction_id = direction_id
                direction_name_rule.track = row.get('track', None)
                direction_name_rule.stop_id_alias = row.get('stop_id_alias', None)
                direction_name_rule.name = row['direction_name']
                priority += 1

    for feed_config in system_config.feeds:
        feed = feed_dao.create()
        feed.system = system
        feed.feed_id = feed_config['name']
        feed.url = feed_config['url']
        feed.parser_module = feed_config['parser_module']
        feed.parser_function = feed_config['parser_function']


class SystemConfig:

    def __init__(self, config_file_path):
        with open(config_file_path, 'r') as f:
            self.config = yaml.load(f)
        self.feeds = self.config.get('feeds', None)
        self.static_route_sps = self.config.get(
            'static_route_service_patterns', [])
        self.static_other_sps = self.config.get(
            'static_other_service_patterns', [])
        self.realtime_route_sps = self.config.get(
            'realtime_route_service_patterns',
            {'enabled': False})
        self.direction_name_rules_files = self.config.get(
            'direction_name_rules_files', [])
