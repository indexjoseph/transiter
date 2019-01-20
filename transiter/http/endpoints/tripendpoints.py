from flask import Blueprint

from transiter.services import tripservice
from transiter.http.responsemanager import http_get_response

trip_endpoints = Blueprint('trip_endpoints', __name__)


@trip_endpoints.route('')
@http_get_response
def list_all_in_route(system_id, route_id):
    """List all trips for a specific system

    .. :quickref: Trip; List all trips for a specific system

    :param system_id: The system's ID
    :param route_id: The route's ID
    :status 200: the system was found
    :status 404: a system with that ID does not exist
    :return: If successful, a JSON response like the following:

    .. code-block:: json

        [
            {
                "id" : "LN1537314760",
                "last_stop": {
                    "id": "L20",
                    "system_id": "nycsubway",
                    "name": "Canarsie",
                    "href": "..."
                }
            }
        ]

    """
    return tripservice.list_all_in_route(system_id, route_id)


@trip_endpoints.route('/<trip_id>')
@http_get_response
def get_in_route_by_id(system_id, route_id, trip_id):
    """Retrieve a specific trip in a specific system.

    .. :quickref: Trip; Retrieve a specific trip

    In version 0.2 this will accept a bunch of GET parameters for
    customizing the precise trip events to return (direction,
    status, etc..)

    :param system_id:  The system's ID
    :param route_id: The route's ID
    :param trip_id: The trip's ID
    :status 200: the stop was found
    :status 404: a stop with that ID does not exist within
        a system with that ID
    :return: If successful, a JSON response like the following:

    .. code-block:: json

        {
            "trip_id" : "LN1537314960",
            "direction": "NORTH",
            "start_time": 100000000,
            "last_update_time": 100000000,
            "feed_update_time": 100000100,
            "status": "RUNNING",
            "train_id" : "0L 1956+RPY/8AV",
            "stop_events" : [
                {
                    "stop": {
                        "stop_id": "F01",
                        "name": "14th st",
                        "location": "NI",
                        "usual_service": ["4", "5", "6"],
                        "href": "https://transiter.io/systems/nycsubway/stops/F01"
                    },
                    "arrival_time" : 1537316801,
                    "departure_time" : 1537316816,
                    "scheduled_track" : "2",
                    "track" : "2",
                    "status": "FUTURE"
                },
            ],
        }

    Note that the stop event item here is the same as the stop event
    in a stop response, except stop database is returned instead of
    trip database.
    """
    return tripservice.get_in_route_by_id(system_id, route_id, trip_id)