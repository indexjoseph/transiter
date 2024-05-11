import pytest
from . import gtfs_realtime_pb2 as gtfs
from . import shared
from . import client

STOP_1 = "stop-1"
STOP_2 = "stop-2"
STOP_3 = "stop-3"
STOP_4 = "stop-4"
STOP_5 = "stop-5"
STOP_6 = "stop-6"
STOP_7 = "stop-7"
ROUTE_ID = "A"
TRIP_ID = "trip_id"

GTFS_STATIC_TXTAR = f"""
-- stops.txt --
stop_id
{STOP_1}
{STOP_2}
{STOP_3}
{STOP_4}
{STOP_5}
{STOP_6}
{STOP_7}
-- routes.txt --
route_id,route_type
{ROUTE_ID},2
"""


TRIP_INITIAL_TIMETABLE = {
    STOP_1: 300,
    STOP_2: 600,
    STOP_3: 800,
    STOP_4: 900,
    STOP_5: 1800,
    STOP_6: 2500,
}


@pytest.mark.parametrize("use_stop_sequences", [True, False])
@pytest.mark.parametrize("current_time", [0, 10, 700, 4000])
@pytest.mark.parametrize(
    "stop_id_to_time_2",
    [
        # Basic case where the second update does nothing.
        TRIP_INITIAL_TIMETABLE,
        # Change the stop times - change time at STOP_3 to before the update at t=700
        {
            STOP_1: 300,
            STOP_2: 800,
            STOP_3: 850,
            STOP_4: 900,
            STOP_5: 1800,
            STOP_6: 2500,
        },
        # Change the stop times - change time at STOP_3 to after the update at t=700
        {
            STOP_1: 300,
            STOP_2: 600,
            STOP_3: 650,
            STOP_4: 900,
            STOP_5: 1800,
            STOP_6: 2500,
        },
        # Add a new stop at the end
        {
            STOP_1: 200,
            STOP_2: 600,
            STOP_3: 800,
            STOP_4: 900,
            STOP_5: 1800,
            STOP_6: 2500,
            STOP_7: 2600,
        },
        # Delete the last stop
        {STOP_1: 200, STOP_2: 600, STOP_3: 800, STOP_4: 900, STOP_5: 1800},
        # Swap the ordering of the last two stops
        {
            STOP_1: 300,
            STOP_2: 600,
            STOP_3: 800,
            STOP_4: 900,
            STOP_5: 1800,
            STOP_7: 2500,
            STOP_6: 3000,
        },
    ],
)
class TestTrip:
    def test_stop_view(
        self,
        install_system,
        system_id,
        transiter_client: client.TransiterClient,
        source_server: shared.SourceServerClient,
        stop_id_to_time_2,
        current_time,
        use_stop_sequences,
    ):
        __, realtime_feed_url = install_system(system_id, GTFS_STATIC_TXTAR)

        for stop_id_to_time, time_at_update in [
            (TRIP_INITIAL_TIMETABLE, 0),
            (stop_id_to_time_2, current_time),
        ]:
            source_server.put(
                realtime_feed_url,
                build_gtfs_rt_message(
                    time_at_update,
                    stop_id_to_time,
                    use_stop_sequences,
                ).SerializeToString(),
            )
            transiter_client.perform_feed_update(
                system_id, shared.GTFS_REALTIME_FEED_ID
            )

        stop_id_to_stop_sequence = {
            stop_id: stop_sequence + 25
            for stop_sequence, stop_id in enumerate(stop_id_to_time_2.keys())
        }
        all_stop_ids = set(TRIP_INITIAL_TIMETABLE.keys()).union(
            stop_id_to_time_2.keys()
        )
        for stop_id in all_stop_ids:
            got_stop = transiter_client.get_stop(system_id, stop_id)

            time = stop_id_to_time_2.get(stop_id)
            if time is None or time < current_time:
                want_stop_times = []
            else:
                want_stop_times = [
                    client.StopTime(
                        trip=client.TripReference(
                            id=TRIP_ID,
                            directionId=True,
                            route=client.RouteReference(id=ROUTE_ID),
                        ),
                        # destination=
                        # vehicle=
                        arrival=client.EstimatedTime(
                            time=time,
                        ),
                        departure=client.EstimatedTime(
                            time=time + 15,
                        ),
                        stopSequence=(
                            stop_id_to_stop_sequence[stop_id]
                            if use_stop_sequences
                            else None
                        ),
                        future=True,
                    )
                ]
            # We don't know the exact values of the Transiter generated stop sequences
            if not use_stop_sequences:
                for stop_time in got_stop.stopTimes:
                    stop_time.stopSequence = None
            assert got_stop.stopTimes == want_stop_times

    def test_trip_view(
        self,
        install_system,
        system_id,
        transiter_client: client.TransiterClient,
        source_server: shared.SourceServerClient,
        stop_id_to_time_2,
        current_time,
        use_stop_sequences,
    ):
        __, realtime_feed_url = install_system(system_id, GTFS_STATIC_TXTAR)

        for stop_id_to_time, time_at_update in [
            (TRIP_INITIAL_TIMETABLE, 0),
            (stop_id_to_time_2, current_time),
        ]:
            source_server.put(
                realtime_feed_url,
                build_gtfs_rt_message(
                    time_at_update,
                    stop_id_to_time,
                    use_stop_sequences,
                ).SerializeToString(),
            )
            transiter_client.perform_feed_update(
                system_id, shared.GTFS_REALTIME_FEED_ID
            )

        want_stop_times = []
        stop_ids_in_second_update = {
            stop_id
            for stop_id, time in stop_id_to_time_2.items()
            if time >= current_time
        }
        for stop_id, time in TRIP_INITIAL_TIMETABLE.items():
            if stop_id in stop_ids_in_second_update:
                break
            want_stop_times.append(
                client.StopTime(
                    trip=None,
                    arrival=client.EstimatedTime(
                        time=time,
                    ),
                    departure=client.EstimatedTime(
                        time=time + 15,
                    ),
                    stopSequence=None,
                    future=False,
                )
            )

        for stop_id, time in stop_id_to_time_2.items():
            if time < current_time:
                continue
            want_stop_times.append(
                client.StopTime(
                    trip=None,
                    arrival=client.EstimatedTime(
                        time=time,
                    ),
                    departure=client.EstimatedTime(
                        time=time + 15,
                    ),
                    stopSequence=None,
                    future=True,
                )
            )
        want_trip = client.Trip(
            id=TRIP_ID,
            shape=None,
            vehicle=None,
            stopTimes=want_stop_times,
        )

        got_trip = transiter_client.get_trip(system_id, ROUTE_ID, TRIP_ID)
        # We don't know the exact values of the Transiter generated stop sequences
        for stop_time in got_trip.stopTimes:
            stop_time.stopSequence = None

        assert got_trip == want_trip


def build_gtfs_rt_message(
    current_time,
    stop_id_to_time,
    use_stop_sequences,
):
    return gtfs.FeedMessage(
        header=gtfs.FeedHeader(gtfs_realtime_version="2.0", timestamp=current_time),
        entity=[
            gtfs.FeedEntity(
                id="1",
                trip_update=gtfs.TripUpdate(
                    trip=gtfs.TripDescriptor(
                        trip_id=TRIP_ID, route_id=ROUTE_ID, direction_id=True
                    ),
                    stop_time_update=[
                        gtfs.TripUpdate.StopTimeUpdate(
                            arrival=gtfs.TripUpdate.StopTimeEvent(time=time),
                            departure=gtfs.TripUpdate.StopTimeEvent(time=time + 15),
                            stop_id=stop_id,
                            stop_sequence=(
                                stop_sequence + 25 if use_stop_sequences else None
                            ),
                        )
                        for stop_sequence, (stop_id, time) in enumerate(
                            stop_id_to_time.items()
                        )
                        if time >= current_time
                    ],
                ),
            )
        ],
    )
