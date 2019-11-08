from transiter import models
from transiter.data import dbconnection
from transiter.data.dams import scheduledam, genericqueries


def sync_trips(feed_update, services):
    """

    :param feed_update:
    :param services:
    :return: True if the schedule was updated
    """

    num_entities_deleted = delete_trips_associated_to_feed(feed_update.feed.pk)
    if len(services) == 0:
        return num_entities_deleted > 0

    session = dbconnection.get_session()

    service_id_to_pk = genericqueries.get_id_to_pk_map_by_feed_pk(
        models.ScheduledService, feed_update.feed.pk
    )
    route_id_to_pk = genericqueries.get_id_to_pk_map(
        models.Route, feed_update.feed.system.id
    )
    trips = []
    for service in services:
        for trip in service.trips:
            trip.route_pk = route_id_to_pk[trip.route_id]
            trip.service_pk = service_id_to_pk[service.id]
        trips.extend(service.trips)
    session.bulk_save_objects(trips)

    trip_id_to_pk = scheduledam.get_trip_id_to_pk_map_by_feed_pk(feed_update.feed.pk)
    stop_id_to_pk = genericqueries.get_id_to_pk_map(
        models.Stop, feed_update.feed.system.id
    )
    stop_time_mappings = []
    for trip in trips:
        for stop_time in trip.stop_times_light:
            if stop_time.stop_id not in stop_id_to_pk:
                continue
            stop_time_mappings.append(
                {
                    "trip_pk": trip_id_to_pk[trip.id],
                    "stop_pk": stop_id_to_pk[stop_time.stop_id],
                    "arrival_time": stop_time.arrival_time,
                    "departure_time": stop_time.departure_time,
                    "stop_sequence": stop_time.stop_sequence,
                }
            )
    session.bulk_insert_mappings(models.ScheduledTripStopTime, stop_time_mappings)

    return num_entities_deleted > 0 or len(trips) > 0 or len(stop_time_mappings) > 0


def delete_trips_associated_to_feed(feed_pk):

    session = dbconnection.get_session()
    num_stop_times_deleted = (
        session.query(models.ScheduledTripStopTime)
        .filter(
            models.ScheduledTrip.pk == models.ScheduledTripStopTime.trip_pk,
            models.ScheduledService.pk == models.ScheduledTrip.service_pk,
            models.FeedUpdate.pk == models.ScheduledService.source_pk,
            models.FeedUpdate.feed_pk == feed_pk,
        )
        .delete(synchronize_session=False)
    )
    num_trips_deleted = (
        session.query(models.ScheduledTrip)
        .filter(
            models.ScheduledService.pk == models.ScheduledTrip.service_pk,
            models.FeedUpdate.pk == models.ScheduledService.source_pk,
            models.FeedUpdate.feed_pk == feed_pk,
        )
        .delete(synchronize_session=False)
    )
    return num_stop_times_deleted + num_trips_deleted