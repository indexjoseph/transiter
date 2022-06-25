// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CalculatePeriodicityForRoute(ctx context.Context, arg CalculatePeriodicityForRouteParams) (int32, error)
	CountAgenciesInSystem(ctx context.Context, systemPk int64) (int64, error)
	CountFeedsInSystem(ctx context.Context, systemPk int64) (int64, error)
	CountRoutesInSystem(ctx context.Context, systemPk int64) (int64, error)
	CountStopsInSystem(ctx context.Context, systemPk int64) (int64, error)
	CountSystems(ctx context.Context) (int64, error)
	CountTransfersInSystem(ctx context.Context, systemPk sql.NullInt64) (int64, error)
	DeleteAlerts(ctx context.Context, alertPks []int64) error
	DeleteFeed(ctx context.Context, pk int64) error
	DeleteServiceMap(ctx context.Context, arg DeleteServiceMapParams) error
	DeleteServiceMapConfig(ctx context.Context, pk int64) error
	DeleteStaleAgencies(ctx context.Context, arg DeleteStaleAgenciesParams) ([]string, error)
	// TODO: These DeleteStaleT queries can be simpler and just take the update_pk
	DeleteStaleAlerts(ctx context.Context, arg DeleteStaleAlertsParams) error
	DeleteStaleRoutes(ctx context.Context, arg DeleteStaleRoutesParams) ([]string, error)
	DeleteStaleStops(ctx context.Context, arg DeleteStaleStopsParams) ([]string, error)
	DeleteStaleTransfers(ctx context.Context, arg DeleteStaleTransfersParams) error
	// TODO: These DeleteStaleT queries can be simpler and just take the update_pk
	DeleteStaleTrips(ctx context.Context, arg DeleteStaleTripsParams) ([]int64, error)
	DeleteStopHeadsignRules(ctx context.Context, sourcePk int64) error
	DeleteSystem(ctx context.Context, pk int64) error
	DeleteTripStopTimes(ctx context.Context, pks []int64) error
	FinishFeedUpdate(ctx context.Context, arg FinishFeedUpdateParams) error
	GetAgencyInSystem(ctx context.Context, arg GetAgencyInSystemParams) (Agency, error)
	GetFeedForUpdate(ctx context.Context, updatePk int64) (Feed, error)
	GetFeedInSystem(ctx context.Context, arg GetFeedInSystemParams) (Feed, error)
	GetLastFeedUpdateContentHash(ctx context.Context, feedPk int64) (sql.NullString, error)
	GetLastStopsForTrips(ctx context.Context, tripPks []int64) ([]GetLastStopsForTripsRow, error)
	GetRouteInSystem(ctx context.Context, arg GetRouteInSystemParams) (GetRouteInSystemRow, error)
	GetStopInSystem(ctx context.Context, arg GetStopInSystemParams) (GetStopInSystemRow, error)
	GetSystem(ctx context.Context, id string) (System, error)
	GetTrip(ctx context.Context, arg GetTripParams) (GetTripRow, error)
	InsertAgency(ctx context.Context, arg InsertAgencyParams) (int64, error)
	InsertAlert(ctx context.Context, arg InsertAlertParams) (int64, error)
	InsertAlertActivePeriod(ctx context.Context, arg InsertAlertActivePeriodParams) error
	InsertAlertAgency(ctx context.Context, arg InsertAlertAgencyParams) error
	InsertAlertRoute(ctx context.Context, arg InsertAlertRouteParams) error
	InsertAlertStop(ctx context.Context, arg InsertAlertStopParams) error
	InsertFeed(ctx context.Context, arg InsertFeedParams) error
	InsertFeedUpdate(ctx context.Context, arg InsertFeedUpdateParams) (int64, error)
	InsertRoute(ctx context.Context, arg InsertRouteParams) (int64, error)
	InsertServiceMap(ctx context.Context, arg InsertServiceMapParams) (int64, error)
	InsertServiceMapConfig(ctx context.Context, arg InsertServiceMapConfigParams) error
	InsertServiceMapStop(ctx context.Context, arg InsertServiceMapStopParams) error
	InsertStop(ctx context.Context, arg InsertStopParams) (int64, error)
	InsertStopHeadSignRule(ctx context.Context, arg InsertStopHeadSignRuleParams) error
	InsertSystem(ctx context.Context, arg InsertSystemParams) (int64, error)
	InsertTransfer(ctx context.Context, arg InsertTransferParams) error
	InsertTrip(ctx context.Context, arg InsertTripParams) (int64, error)
	InsertTripStopTime(ctx context.Context, arg InsertTripStopTimeParams) error
	ListActiveAlertsForAgencies(ctx context.Context, arg ListActiveAlertsForAgenciesParams) ([]ListActiveAlertsForAgenciesRow, error)
	ListActiveAlertsForRoutes(ctx context.Context, arg ListActiveAlertsForRoutesParams) ([]ListActiveAlertsForRoutesRow, error)
	ListActiveAlertsForStops(ctx context.Context, arg ListActiveAlertsForStopsParams) ([]ListActiveAlertsForStopsRow, error)
	ListAgenciesInSystem(ctx context.Context, systemPk int64) ([]Agency, error)
	ListAlertPksAndHashes(ctx context.Context, arg ListAlertPksAndHashesParams) ([]ListAlertPksAndHashesRow, error)
	ListAutoUpdateFeedsForSystem(ctx context.Context, systemID string) ([]ListAutoUpdateFeedsForSystemRow, error)
	ListFeedsInSystem(ctx context.Context, systemPk int64) ([]Feed, error)
	ListRoutesByPk(ctx context.Context, routePks []int64) ([]Route, error)
	ListRoutesInAgency(ctx context.Context, agencyPk int64) ([]ListRoutesInAgencyRow, error)
	ListRoutesInSystem(ctx context.Context, systemPk int64) ([]Route, error)
	ListServiceMapConfigsInSystem(ctx context.Context, systemPk int64) ([]ServiceMapConfig, error)
	ListServiceMapsConfigIDsForStops(ctx context.Context, stopPks []int64) ([]ListServiceMapsConfigIDsForStopsRow, error)
	ListServiceMapsForRoute(ctx context.Context, routePk int64) ([]ListServiceMapsForRouteRow, error)
	ListServiceMapsForStops(ctx context.Context, stopPks []int64) ([]ListServiceMapsForStopsRow, error)
	ListStopHeadsignRulesForStops(ctx context.Context, stopPks []int64) ([]StopHeadsignRule, error)
	ListStopTimesAtStops(ctx context.Context, stopPks []int64) ([]ListStopTimesAtStopsRow, error)
	ListStopsInStopTree(ctx context.Context, pk int64) ([]Stop, error)
	ListStopsInSystem(ctx context.Context, systemPk int64) ([]Stop, error)
	ListStopsTimesForTrip(ctx context.Context, tripPk int64) ([]ListStopsTimesForTripRow, error)
	ListSystems(ctx context.Context) ([]System, error)
	ListTransfersFromStops(ctx context.Context, fromStopPks []int64) ([]ListTransfersFromStopsRow, error)
	ListTransfersInSystem(ctx context.Context, systemPk sql.NullInt64) ([]ListTransfersInSystemRow, error)
	ListTripStopTimesForUpdate(ctx context.Context, tripPks []int64) ([]ListTripStopTimesForUpdateRow, error)
	ListTripsForUpdate(ctx context.Context, routePks []int64) ([]ListTripsForUpdateRow, error)
	ListTripsInRoute(ctx context.Context, routePk int64) ([]ListTripsInRouteRow, error)
	ListUpdatesInFeed(ctx context.Context, feedPk int64) ([]FeedUpdate, error)
	MapAgencyPkToIdInSystem(ctx context.Context, systemPk int64) ([]MapAgencyPkToIdInSystemRow, error)
	MapRoutePkToIdInSystem(ctx context.Context, systemPk int64) ([]MapRoutePkToIdInSystemRow, error)
	MapRoutesInSystem(ctx context.Context, arg MapRoutesInSystemParams) ([]MapRoutesInSystemRow, error)
	MapStopIDToStationPk(ctx context.Context, systemPk int64) ([]MapStopIDToStationPkRow, error)
	MapStopPkToIdInSystem(ctx context.Context, systemPk int64) ([]MapStopPkToIdInSystemRow, error)
	MapStopPkToStationPk(ctx context.Context, stopPks []int64) ([]MapStopPkToStationPkRow, error)
	MapStopsInSystem(ctx context.Context, arg MapStopsInSystemParams) ([]MapStopsInSystemRow, error)
	MarkAlertsFresh(ctx context.Context, arg MarkAlertsFreshParams) error
	MarkTripStopTimesPast(ctx context.Context, arg MarkTripStopTimesPastParams) error
	UpdateAgency(ctx context.Context, arg UpdateAgencyParams) error
	UpdateFeed(ctx context.Context, arg UpdateFeedParams) error
	UpdateRoute(ctx context.Context, arg UpdateRouteParams) error
	UpdateServiceMapConfig(ctx context.Context, arg UpdateServiceMapConfigParams) error
	UpdateStop(ctx context.Context, arg UpdateStopParams) error
	UpdateStopParent(ctx context.Context, arg UpdateStopParentParams) error
	UpdateSystem(ctx context.Context, arg UpdateSystemParams) error
	UpdateSystemStatus(ctx context.Context, arg UpdateSystemStatusParams) error
	UpdateTrip(ctx context.Context, arg UpdateTripParams) error
	UpdateTripStopTime(ctx context.Context, arg UpdateTripStopTimeParams) error
}

var _ Querier = (*Queries)(nil)
