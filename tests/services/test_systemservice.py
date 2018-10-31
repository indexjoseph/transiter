import unittest
import unittest.mock as mock

from transiter.services import systemservice
from transiter.services import exceptions


class TestSystemService(unittest.TestCase):

    SYSTEM_ONE_ID = '1'
    SYSTEM_ONE_HREF = '3'
    SYSTEM_ONE_REPR = {'system_id': SYSTEM_ONE_ID, 'href': SYSTEM_ONE_HREF}
    SYSTEM_TWO_ID = '2'
    SYSTEM_TWO_HREF = '4'
    SYSTEM_TWO_REPR = {'system_id': SYSTEM_TWO_ID, 'href': SYSTEM_TWO_HREF}
    SYSTEM_ONE_NUM_STOPS = 20
    SYSTEM_ONE_NUM_STATIONS = 21
    SYSTEM_ONE_NUM_ROUTES = 22
    SYSTEM_ONE_NUM_FEEDS = 23

    @classmethod
    def setUp(cls):
        cls.system_1 = mock.MagicMock()
        cls.system_1.short_repr.return_value = cls.SYSTEM_ONE_REPR.copy()
        cls.system_2 = mock.MagicMock()
        cls.system_2.short_repr.return_value = cls.SYSTEM_TWO_REPR.copy()

    @mock.patch('transiter.services.systemservice.linksutil')
    @mock.patch('transiter.services.systemservice.system_dao')
    def test_list_all(self, system_dao, linksutil):
        """[System service] List all installed systems"""

        def SystemEntityLink(system):
            if system == self.system_1:
                return self.SYSTEM_ONE_HREF
            if system == self.system_2:
                return self.SYSTEM_TWO_HREF
        linksutil.SystemEntityLink.side_effect = SystemEntityLink

        expected = [self.SYSTEM_ONE_REPR, self.SYSTEM_TWO_REPR]
        system_dao.list_all.return_value = [
            self.system_1,
            self.system_2]

        actual = systemservice.list_all()

        self.assertEqual(actual, expected)
        system_dao.list_all.assert_called_once()
        self.system_1.short_repr.assert_called_once()
        self.system_2.short_repr.assert_called_once()

    @mock.patch('transiter.services.systemservice.linksutil')
    @mock.patch('transiter.services.systemservice.system_dao')
    def test_get_by_id(self, system_dao, linksutil):
        """[System service] Get a specific system"""

        hrefs_dict = {
            'stops': 'href1',
            'stations': 'NI',
            'routes': 'href3',
            'feeds': 'href4'
        }
        child_entities_dict = {
            'stops': self.SYSTEM_ONE_NUM_STOPS,
            'stations': self.SYSTEM_ONE_NUM_STATIONS,
            'routes': self.SYSTEM_ONE_NUM_ROUTES,
            'feeds': self.SYSTEM_ONE_NUM_FEEDS
        }
        expected = {
            name: {
                'count': count,
                'href': hrefs_dict[name]
            } for (name, count) in child_entities_dict.items()
        }
        expected.update(**self.SYSTEM_ONE_REPR)

        system_dao.get_by_id.return_value = self.system_1
        system_dao.count_stops_in_system.return_value = self.SYSTEM_ONE_NUM_STOPS
        system_dao.count_stations_in_system.return_value = self.SYSTEM_ONE_NUM_STATIONS
        system_dao.count_routes_in_system.return_value = self.SYSTEM_ONE_NUM_ROUTES
        system_dao.count_feeds_in_system.return_value = self.SYSTEM_ONE_NUM_FEEDS

        linksutil.StopsInSystemIndexLink.return_value = hrefs_dict['stops']
        linksutil.RoutesInSystemIndexLink.return_value = hrefs_dict['routes']
        linksutil.FeedsInSystemIndexLink.return_value = hrefs_dict['feeds']

        actual = systemservice.get_by_id(self.SYSTEM_ONE_ID)

        self.maxDiff = None
        self.assertDictEqual(actual, expected)
        system_dao.get_by_id.assert_called_once_with(self.SYSTEM_ONE_ID)
        self.system_1.short_repr.assert_called_once()
        system_dao.count_stops_in_system.assert_called_once_with(self.SYSTEM_ONE_ID)
        system_dao.count_stations_in_system.assert_called_once_with(self.SYSTEM_ONE_ID)
        system_dao.count_routes_in_system.assert_called_once_with(self.SYSTEM_ONE_ID)
        system_dao.count_feeds_in_system.assert_called_once_with(self.SYSTEM_ONE_ID)

    @mock.patch('transiter.services.systemservice._import_static_data')
    @mock.patch('transiter.services.systemservice.system_dao')
    def test_install_success(self, system_dao, _import_static_data):
        """[System service] Successfully install a system"""
        new_system = mock.MagicMock()
        system_dao.get_by_id.return_value = None
        system_dao.create.return_value = new_system

        actual = systemservice.install(self.SYSTEM_ONE_ID)

        self.assertEqual(actual, True)
        self.assertEqual(new_system.system_id, self.SYSTEM_ONE_ID)
        system_dao.get_by_id.assert_called_once_with(self.SYSTEM_ONE_ID)
        system_dao.create.assert_called_once_with()
        _import_static_data.assert_called_once_with(new_system)

    @mock.patch('transiter.services.systemservice.system_dao')
    def test_install_already_exists(self, system_dao):
        """[System service] Fail to install because system id already taken"""
        system_dao.get_by_id.return_value = self.system_1

        actual = systemservice.install(self.SYSTEM_ONE_ID)

        self.assertEqual(actual, False)
        system_dao.get_by_id.assert_called_once_with(self.SYSTEM_ONE_ID)

    @mock.patch('transiter.services.systemservice.system_dao')
    def test_delete_success(self, system_dao):
        """[System service] Successfully delete a system"""
        system_dao.delete_by_id.return_value = True

        actual = systemservice.delete_by_id(self.SYSTEM_ONE_ID)

        self.assertEqual(actual, True)
        system_dao.delete_by_id.assert_called_once_with(self.SYSTEM_ONE_ID)

    @mock.patch('transiter.services.systemservice.system_dao')
    def test_delete_failure(self, system_dao):
        """[System service] Fail to delete a nonexistent system"""
        system_dao.delete_by_id.return_value = False

        self.assertRaises(exceptions.IdNotFoundError,
                          systemservice.delete_by_id,
                          self.SYSTEM_ONE_ID)

        system_dao.delete_by_id.assert_called_once_with(self.SYSTEM_ONE_ID)
