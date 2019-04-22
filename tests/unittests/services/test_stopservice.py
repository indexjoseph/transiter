import datetime
import unittest
import unittest.mock as mock

from transiter import models, exceptions
from transiter.services import stopservice, links
from .. import testutil


class TestDirectionNamesMatcher(unittest.TestCase):
    STOP_PK = 1
    DIRECTION_NAME = 'Direction Name'

    def setUp(self):
        self.stop = models.Stop()
        self.stop.pk = self.STOP_PK

        self.stop_event = models.StopTimeUpdate()
        self.stop_event.track = None
        self.stop_event.stop_id_alias = None
        self.stop_event.trip = models.Trip()
        self.stop_event.trip.direction_id = None
        self.stop_event.stop = self.stop

        self.rule = models.DirectionNameRule()
        self.rule.stop_pk = self.STOP_PK
        self.rule.direction_id = None
        self.rule.track = None
        self.rule.name = self.DIRECTION_NAME

    def test_all_names(self):
        dnm = stopservice._DirectionNameMatcher([self.rule])

        self.assertEqual({self.DIRECTION_NAME}, dnm.all_names())

    def test_no_matching_stop_pk(self):
        self.rule.stop_pk = 2
        dnm = stopservice._DirectionNameMatcher([self.rule])

        direction_name = dnm.match(self.stop_event)

        self.assertEqual(direction_name, None)

    def test_no_matching_direction_id(self):
        self.rule.direction_id = True
        dnm = stopservice._DirectionNameMatcher([self.rule])

        direction_name = dnm.match(self.stop_event)

        self.assertEqual(direction_name, None)

    def test_no_matching_track(self):
        self.rule.track = 'Track'
        dnm = stopservice._DirectionNameMatcher([self.rule])

        direction_name = dnm.match(self.stop_event)

        self.assertEqual(direction_name, None)

    def test_match(self):
        dnm = stopservice._DirectionNameMatcher([self.rule])

        direction_name = dnm.match(self.stop_event)

        self.assertEqual(direction_name, self.DIRECTION_NAME)


class TestTripStopTimeFilter(unittest.TestCase):
    DIRECTION_NAME = '1'
    DATETIME_ONE = datetime.datetime(2000, 1, 1, 1, 0, 0)
    DATETIME_TWO = datetime.datetime(2000, 1, 1, 2, 0, 0)
    ROUTE_ID = '1'

    def setUp(self):
        self.stop_event_filter = stopservice._TripStopTimeFilter()
        self.stop_event = models.StopTimeUpdate()
        self.stop_event.arrival_time = self.DATETIME_ONE
        self.stop_event.trip = models.Trip()
        self.stop_event.trip.route = models.Route()
        self.stop_event.trip.route_id = self.ROUTE_ID

    def test_add_direction_name(self):
        self.stop_event_filter._add_direction_name(self.DIRECTION_NAME)

        self.assertDictEqual(
            self.stop_event_filter._count,
            {self.DIRECTION_NAME: 0}
        )
        self.assertDictEqual(
            self.stop_event_filter._route_ids_so_far,
            {self.DIRECTION_NAME: set()}
        )

    def test_add_direction_name_already_added(self):
        self.stop_event_filter._count[self.DIRECTION_NAME] = 50
        self.stop_event_filter._add_direction_name(self.DIRECTION_NAME)

        self.assertDictEqual(
            self.stop_event_filter._count,
            {self.DIRECTION_NAME: 50}
        )

    @mock.patch.object(stopservice, 'time')
    def test_exclude_route_not_there_yet(self, time):
        self.stop_event_filter._add_direction_name(self.DIRECTION_NAME)
        self.stop_event_filter._count[self.DIRECTION_NAME] = 100
        time.time.return_value = self.DATETIME_ONE.timestamp()
        self.stop_event.departure_time = self.DATETIME_ONE

        exclude = self.stop_event_filter.exclude(
            self.stop_event, self.DIRECTION_NAME)

        self.assertFalse(exclude)


class TestStopService(testutil.TestCase(stopservice), unittest.TestCase):
    SYSTEM_ID = '1'
    STOP_ONE_ID = '2'
    STOP_ONE_PK = 3
    STOP_ONE_REPR = {'stop_id': STOP_ONE_ID}
    ALL_DIRECTION_NAMES = ['A', 'B']
    STOP_ONE_HREF = '10'
    TRIP_HREF = '11'
    ROUTE_HREF = '12'
    TRIP_REPR = {'21': '22'}
    ROUTE_REPR = {'23': '24'}
    STOP_EVENT_REPR = {'25': '26'}
    DEFAULT_TRIPS = ['31', '32', '33']

    def setUp(self):
        self.stop_dao = self.mockImportedModule(stopservice.stopdam)
        self.systemdam = self.mockImportedModule(stopservice.systemdam)
        self.tripdam = self.mockImportedModule(stopservice.tripdam)

        self.stop_one = models.Stop()
        self.stop_one.pk = self.STOP_ONE_PK
        self.stop_one.id = self.STOP_ONE_ID
        self.stop_dao.list_all_in_system.return_value = [
            self.stop_one]

        self.usual_routes = []
        for route_id in self.DEFAULT_TRIPS:
            route = models.Route()
            route.id = route_id
            self.usual_routes.append(route)

    def test_list_all_in_system(self):
        """[Stop service] List all in system"""
        expected = [
            {
                'href': links.StopEntityLink(self.stop_one),
                **self.stop_one.short_repr(),
            }
        ]

        actual = stopservice.list_all_in_system(self.SYSTEM_ID, True)

        self.assertListEqual(actual, expected)
        self.stop_dao.list_all_in_system.assert_called_once_with(self.SYSTEM_ID)

    def test_list_all_in_system__system_not_found(self):
        """[Stop service] List all in system - system not found"""
        self.systemdam.get_by_id.return_value = None

        self.assertRaises(
            exceptions.IdNotFoundError,
            lambda: stopservice.list_all_in_system(self.SYSTEM_ID)
        )

    def test_get_in_system_by_id__stop_not_found(self):
        """[Stop service] Get stop - stop not found"""
        self.stop_dao.get_in_system_by_id.return_value = None

        self.assertRaises(
            exceptions.IdNotFoundError,
            lambda: stopservice.get_in_system_by_id(self.SYSTEM_ID, self.STOP_ONE_ID)
        )

    def test_build_stop_tree_response(self):
        """[Stop service] Build stop tree response"""
        # Stops tree in this test:
        #     0
        #    / \
        #   1   2
        #  /
        # 3*
        # * not a station
        stops = [models.Stop()] * 4
        for i in range(4):
            stops[i] = models.Stop()
            stops[i].pk = i
            stops[i].id = str(i)
            stops[i].is_station = True
            stops[i].system_id = 'system'
        stops[0].child_stops = [stops[1], stops[2]]
        stops[1].child_stops = [stops[3]]
        stops[3].is_station = False

        stop_pk_to_service_maps_response = {pk: pk for pk in range(4)}

        expected = {
            **stops[1].short_repr(),
            'service_maps': 1,
            'href': links.StopEntityLink(stops[1]),
            'parent_stop': {
                **stops[0].short_repr(),
                'service_maps': 0,
                'href': links.StopEntityLink(stops[0]),
                'parent_stop': None,
                'child_stops': [
                    {
                        **stops[2].short_repr(),
                        'service_maps': 2,
                        'href': links.StopEntityLink(stops[2]),
                        'child_stops': []
                    }
                ]
            },
            'child_stops': []
        }

        actual = stopservice._build_stop_tree_response(
            stops[1],
            stop_pk_to_service_maps_response,
            True,
            True
        )

        self.maxDiff = None
        self.assertDictEqual(expected, actual)

    def test_get_stop_descendants(self):
        """[Stop service] Get all stop descendants"""
        # Stops tree in this test:
        #     0
        #    / \
        #   1   2
        #  /
        # 3*
        # * not a station
        stops = [models.Stop()] * 4
        for i in range(4):
            stops[i] = models.Stop()
            stops[i].pk = i
            stops[i].id = str(i)
            stops[i].is_station = True
        stops[0].child_stops = [stops[1], stops[2]]
        stops[1].child_stops = [stops[3]]
        stops[3].is_station = False

        expected_pks = {0, 1, 2, 3}

        actual_pks = {stop.pk for stop in stopservice._get_stop_descendants(stops[0])}

        self.assertEqual(expected_pks, actual_pks)

    def test_get_stop_ancestors(self):
        """[Stop service] Get all stop ancestors and ancestor siblings that are stations"""
        # Stops tree in this test:
        #      2
        #    / | \
        #   1  3  4
        #  /   |
        # 0    5*
        # * not a station
        stops = [None] * 6
        for i in range(6):
            stops[i] = models.Stop()
            stops[i].pk = i
            stops[i].id = str(i)
            stops[i].is_station = True
        stops[0].parent_stop = stops[1]
        stops[1].parent_stop = stops[2]
        stops[2].child_stops = [stops[1], stops[3], stops[4]]
        stops[3].child_stops = [stops[5]]
        stops[4].is_station = False

        expected_pks = {0, 1, 2, 3, 5}

        for stop in stops:
            if not stop.is_station:
                continue
            actual_pks = {stop.pk for stop in stopservice._get_all_stations(stop)}

            self.assertEqual(expected_pks, actual_pks)

    @mock.patch('transiter.services.stopservice.servicepatterndam')
    @mock.patch('transiter.services.stopservice._TripStopTimeFilter')
    @mock.patch('transiter.services.stopservice._DirectionNameMatcher')
    def _test_get_in_system_by_id(self, _DirectionNameMatcher, _StopEventFilter,
                                  service_pattern_dao):
        self.stop_one.parent_stop = None
        self.stop_one.child_stops = []
        self.stop_dao.get_in_system_by_id.return_value = self.stop_one
        service_pattern_dao.get_default_routes_at_stops_map.return_value = {
            self.STOP_ONE_PK: self.usual_routes
        }

        stop_event_one = mock.MagicMock()
        stop_event_two = mock.MagicMock()
        stop_event_two.stop = self.stop_one
        stop_event_two.short_repr.return_value = self.STOP_EVENT_REPR
        stop_event_two.trip.long_repr.return_value = self.TRIP_REPR
        stop_event_two.trip.route.short_repr.return_value = self.ROUTE_REPR
        self.stop_dao.list_stop_time_updates_at_stops.return_value = [stop_event_one, stop_event_two]

        direction_name_matcher = mock.MagicMock()
        _DirectionNameMatcher.return_value = direction_name_matcher
        direction_name_matcher.all_names.return_value = self.ALL_DIRECTION_NAMES
        direction_name_matcher.match.side_effect = self.ALL_DIRECTION_NAMES

        stop_event_filter = mock.MagicMock()
        _StopEventFilter.return_value = stop_event_filter
        stop_event_filter.exclude.side_effect = [True, False]

        expected_response = {
            **self.STOP_ONE_REPR,
            'usual_routes': self.DEFAULT_TRIPS,
            'direction_names': self.ALL_DIRECTION_NAMES,
            'stop_time_updates': [
                {
                    'direction_name': self.ALL_DIRECTION_NAMES[1],
                    'stop_id': self.STOP_ONE_ID,
                    **self.STOP_EVENT_REPR,
                    'trip': {
                        **self.TRIP_REPR,
                        'route': {
                            **self.ROUTE_REPR,
                            'href': self.ROUTE_HREF
                        },
                        'last_stop': {'href': '10'},
                        'href': self.TRIP_HREF
                    }
                }
            ],
            'child_stops': [],
            'parent_stop': None
        }

        actual_response = stopservice.get_in_system_by_id(self.SYSTEM_ID, self.STOP_ONE_ID)

        print(expected_response)
        print(actual_response)
        self.assertDictEqual(expected_response, actual_response)
