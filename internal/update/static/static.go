package static

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jamespfennell/gtfs"
	"github.com/jamespfennell/transiter/internal/apihelpers"
	"github.com/jamespfennell/transiter/internal/gen/db"
	"github.com/jamespfennell/transiter/internal/servicemaps"
)

type UpdateContext struct {
	Querier  db.Querier
	SystemPk int64
	FeedPk   int64
	UpdatePk int64
}

func Update(ctx context.Context, updateCtx UpdateContext, parsedEntities *gtfs.Static) error {
	agencyIdToPk, err := updateAgencies(ctx, updateCtx, parsedEntities.Agencies)
	if err != nil {
		return err
	}
	routeIdToPk, err := updateRoutes(ctx, updateCtx, parsedEntities.Routes, agencyIdToPk)
	if err != nil {
		return err
	}
	stopIdToPk, err := updateStops(ctx, updateCtx, parsedEntities.AllStops())
	if err != nil {
		return err
	}
	if err := updateTransfers(ctx, updateCtx, parsedEntities.Transfers, stopIdToPk); err != nil {
		return err
	}
	if err := servicemaps.UpdateStaticMaps(ctx, updateCtx.Querier, servicemaps.UpdateStaticMapsArgs{
		SystemPk:    updateCtx.SystemPk,
		Trips:       parsedEntities.Trips,
		RouteIdToPk: routeIdToPk,
	}); err != nil {
		return err
	}
	return nil
}

func updateAgencies(ctx context.Context, updateCtx UpdateContext, agencies []gtfs.Agency) (map[string]int64, error) {
	idToPk, err := buildAgencyIdToPkMap(ctx, updateCtx.Querier, updateCtx.SystemPk)
	if err != nil {
		return nil, err
	}
	for _, agency := range agencies {
		var err error
		pk, ok := idToPk[agency.Id]
		if ok {
			err = updateCtx.Querier.UpdateAgency(ctx, db.UpdateAgencyParams{
				Pk:       pk,
				SourcePk: updateCtx.UpdatePk,
				Name:     agency.Name,
				Url:      agency.Url,
				Timezone: agency.Timezone,
				Language: apihelpers.ConvertNullString(agency.Language),
				Phone:    apihelpers.ConvertNullString(agency.Phone),
				FareUrl:  apihelpers.ConvertNullString(agency.FareUrl),
				Email:    apihelpers.ConvertNullString(agency.Email),
			})
		} else {
			pk, err = updateCtx.Querier.InsertAgency(ctx, db.InsertAgencyParams{
				ID:       agency.Id,
				SystemPk: updateCtx.SystemPk,
				SourcePk: updateCtx.UpdatePk,
				Name:     agency.Name,
				Url:      agency.Url,
				Timezone: agency.Timezone,
				Language: apihelpers.ConvertNullString(agency.Language),
				Phone:    apihelpers.ConvertNullString(agency.Phone),
				FareUrl:  apihelpers.ConvertNullString(agency.FareUrl),
				Email:    apihelpers.ConvertNullString(agency.Email),
			})
			idToPk[agency.Id] = pk
		}
		if err != nil {
			return nil, err
		}
	}
	deletedIds, err := updateCtx.Querier.DeleteStaleAgencies(ctx, db.DeleteStaleAgenciesParams{
		FeedPk:   updateCtx.FeedPk,
		UpdatePk: updateCtx.UpdatePk,
	})
	if err != nil {
		return nil, err
	}
	for _, id := range deletedIds {
		delete(idToPk, id)
	}
	return idToPk, nil
}

func updateRoutes(ctx context.Context, updateCtx UpdateContext, routes []gtfs.Route, agencyIdToPk map[string]int64) (map[string]int64, error) {
	idToPk, err := buildRouteIdToPkMap(ctx, updateCtx.Querier, updateCtx.SystemPk)
	if err != nil {
		return nil, err
	}
	for _, route := range routes {
		agencyPk, ok := agencyIdToPk[route.Agency.Id]
		if !ok {
			log.Printf("no agency %q in the database; skipping route %q", route.Agency.Id, route.Id)
			continue
		}
		pk, ok := idToPk[route.Id]
		if ok {
			err = updateCtx.Querier.UpdateRoute(ctx, db.UpdateRouteParams{
				Pk:                pk,
				SourcePk:          updateCtx.UpdatePk,
				Color:             route.Color,
				TextColor:         route.TextColor,
				ShortName:         apihelpers.ConvertNullString(route.ShortName),
				LongName:          apihelpers.ConvertNullString(route.LongName),
				Description:       apihelpers.ConvertNullString(route.Description),
				Url:               apihelpers.ConvertNullString(route.Url),
				SortOrder:         apihelpers.ConvertNullInt32(route.SortOrder),
				Type:              route.Type.String(),
				ContinuousPickup:  route.ContinuousPickup.String(),
				ContinuousDropOff: route.ContinuousDropOff.String(),
				AgencyPk:          agencyPk,
			})
		} else {
			pk, err = updateCtx.Querier.InsertRoute(ctx, db.InsertRouteParams{
				ID:                route.Id,
				SystemPk:          updateCtx.SystemPk,
				SourcePk:          updateCtx.UpdatePk,
				Color:             route.Color,
				TextColor:         route.TextColor,
				ShortName:         apihelpers.ConvertNullString(route.ShortName),
				LongName:          apihelpers.ConvertNullString(route.LongName),
				Description:       apihelpers.ConvertNullString(route.Description),
				Url:               apihelpers.ConvertNullString(route.Url),
				SortOrder:         apihelpers.ConvertNullInt32(route.SortOrder),
				Type:              route.Type.String(),
				ContinuousPickup:  route.ContinuousPickup.String(),
				ContinuousDropOff: route.ContinuousDropOff.String(),
				AgencyPk:          agencyPk,
			})
			idToPk[route.Id] = pk
		}
		if err != nil {
			return nil, err
		}
	}
	deletedIds, err := updateCtx.Querier.DeleteStaleRoutes(ctx, db.DeleteStaleRoutesParams{
		FeedPk:   updateCtx.FeedPk,
		UpdatePk: updateCtx.UpdatePk,
	})
	if err != nil {
		return nil, err
	}
	for _, id := range deletedIds {
		delete(idToPk, id)
	}
	return idToPk, nil
}

func updateStops(ctx context.Context, updateCtx UpdateContext, stops []*gtfs.Stop) (map[string]int64, error) {
	idToPk, err := buildStopIdToPkMap(ctx, updateCtx.Querier, updateCtx.SystemPk)
	if err != nil {
		return nil, err
	}
	for _, stop := range stops {
		pk, ok := idToPk[stop.Id]
		if ok {
			err = updateCtx.Querier.UpdateStop(ctx, db.UpdateStopParams{
				Pk:                 pk,
				SourcePk:           updateCtx.UpdatePk,
				Name:               apihelpers.ConvertNullString(stop.Name),
				Type:               stop.Type.String(),
				Longitude:          convertGpsData(stop.Longitude),
				Latitude:           convertGpsData(stop.Lattitude),
				Url:                apihelpers.ConvertNullString(stop.Url),
				Code:               apihelpers.ConvertNullString(stop.Code),
				Description:        apihelpers.ConvertNullString(stop.Description),
				PlatformCode:       apihelpers.ConvertNullString(stop.PlatformCode),
				Timezone:           apihelpers.ConvertNullString(stop.Timezone),
				WheelchairBoarding: "NOT_SPECIFIED", // TODO, need to make a change to the GTFS package
				ZoneID:             apihelpers.ConvertNullString(stop.ZoneId),
			})
		} else {
			pk, err = updateCtx.Querier.InsertStop(ctx, db.InsertStopParams{
				ID:                 stop.Id,
				SystemPk:           updateCtx.SystemPk,
				SourcePk:           updateCtx.UpdatePk,
				Name:               apihelpers.ConvertNullString(stop.Name),
				Type:               stop.Type.String(),
				Longitude:          convertGpsData(stop.Longitude),
				Latitude:           convertGpsData(stop.Lattitude),
				Url:                apihelpers.ConvertNullString(stop.Url),
				Code:               apihelpers.ConvertNullString(stop.Code),
				Description:        apihelpers.ConvertNullString(stop.Description),
				PlatformCode:       apihelpers.ConvertNullString(stop.PlatformCode),
				Timezone:           apihelpers.ConvertNullString(stop.Timezone),
				WheelchairBoarding: "NOT_SPECIFIED", // TODO, need to make a change to the GTFS package
				ZoneID:             apihelpers.ConvertNullString(stop.ZoneId),
			})
			idToPk[stop.Id] = pk
		}
		if err != nil {
			return nil, err
		}
	}
	deletedIds, err := updateCtx.Querier.DeleteStaleStops(ctx, db.DeleteStaleStopsParams{
		FeedPk:   updateCtx.FeedPk,
		UpdatePk: updateCtx.UpdatePk,
	})
	if err != nil {
		return nil, err
	}
	for _, id := range deletedIds {
		delete(idToPk, id)
	}
	// We now populate the parent stop field
	for _, stop := range stops {
		if stop.Parent == nil {
			continue
		}
		parentStopPk, ok := idToPk[stop.Parent.Id]
		if !ok {
			continue
		}
		if err := updateCtx.Querier.UpdateStopParent(ctx, db.UpdateStopParentParams{
			Pk: idToPk[stop.Id],
			ParentStopPk: sql.NullInt64{
				Int64: parentStopPk,
				Valid: true,
			},
		}); err != nil {
			return nil, err
		}
	}
	return idToPk, nil
}

func updateTransfers(ctx context.Context, updateCtx UpdateContext, transfers []gtfs.Transfer, stopIdToPk map[string]int64) error {
	if err := updateCtx.Querier.DeleteStaleTransfers(ctx, db.DeleteStaleTransfersParams{
		FeedPk:   updateCtx.FeedPk,
		UpdatePk: updateCtx.UpdatePk,
	}); err != nil {
		return err
	}
	for _, transfer := range transfers {
		fromPk, ok := stopIdToPk[transfer.From.Id]
		if !ok {
			continue
		}
		toPk, ok := stopIdToPk[transfer.To.Id]
		if !ok {
			continue
		}
		if err := updateCtx.Querier.InsertTransfer(ctx, db.InsertTransferParams{
			SystemPk:        apihelpers.ConvertNullInt64(&updateCtx.SystemPk),
			SourcePk:        apihelpers.ConvertNullInt64(&updateCtx.UpdatePk),
			FromStopPk:      fromPk,
			ToStopPk:        toPk,
			Type:            transfer.Type.String(),
			MinTransferTime: apihelpers.ConvertNullInt32(transfer.MinTransferTime),
		}); err != nil {
			return err
		}
	}
	return nil
}

func buildAgencyIdToPkMap(ctx context.Context, querier db.Querier, systemPk int64) (map[string]int64, error) {
	idToPk := map[string]int64{}
	rows, err := querier.MapAgencyPkToIdInSystem(ctx, systemPk)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		idToPk[row.ID] = row.Pk
	}
	return idToPk, nil
}

func buildRouteIdToPkMap(ctx context.Context, querier db.Querier, systemPk int64) (map[string]int64, error) {
	idToPk := map[string]int64{}
	rows, err := querier.MapRoutePkToIdInSystem(ctx, systemPk)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		idToPk[row.ID] = row.Pk
	}
	return idToPk, nil
}

func buildStopIdToPkMap(ctx context.Context, querier db.Querier, systemPk int64) (map[string]int64, error) {
	idToPk := map[string]int64{}
	rows, err := querier.MapStopPkToIdInSystem(ctx, systemPk)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		idToPk[row.ID] = row.Pk
	}
	return idToPk, nil
}

func convertGpsData(f *float64) sql.NullString {
	if f == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		Valid:  true,
		String: fmt.Sprintf("%f", *f),
	}
}