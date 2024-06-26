package endpoints

import (
	"context"
	"math"

	"github.com/jamespfennell/transiter/internal/convert"
	"github.com/jamespfennell/transiter/internal/gen/api"
	"github.com/jamespfennell/transiter/internal/gen/db"
	"github.com/jamespfennell/transiter/internal/public/errors"
)

func ListShapes(ctx context.Context, r *Context, req *api.ListShapesRequest) (*api.ListShapesReply, error) {
	system, err := getSystem(ctx, r.Querier, req.SystemId)
	if err != nil {
		return nil, err
	}
	filterByID := req.GetFilterById()
	//lint:ignore SA1019 we still consult the deprecated field as it's not been removed yet
	if req.GetOnlyReturnSpecifiedIds() {
		r.Monitoring.RecordDeprecatedFeatureUse("ListShapesRequest.only_return_specified_ids")
		filterByID = true
	}
	if !filterByID && len(req.Id) > 0 {
		return nil, errors.NewInvalidArgumentError("filter_by_id is false but IDs were provided")
	}

	numShapes := r.EndpointOptions.MaxEntitiesPerRequest
	if numShapes <= 0 {
		// Avoid overflow since pagination over-fetches by one
		numShapes = math.MaxInt32 - 1
	}
	if req.Limit != nil && *req.Limit < numShapes {
		numShapes = *req.Limit
	}
	var firstID string
	if req.FirstId != nil {
		firstID = *req.FirstId
	}

	dbShapes, err := r.Querier.ListShapes(ctx, db.ListShapesParams{
		SystemPk:     system.Pk,
		FirstShapeID: firstID,
		NumShapes:    numShapes + 1,
		FilterByID:   filterByID,
		ShapeIds:     req.Id,
	})
	if err != nil {
		return nil, err
	}

	apiShapes, err := buildApiShapes(dbShapes)
	if err != nil {
		return nil, err
	}

	var nextID *string
	if len(apiShapes) == int(numShapes+1) {
		nextID = &apiShapes[len(apiShapes)-1].Id
		apiShapes = apiShapes[:len(apiShapes)-1]
	}

	return &api.ListShapesReply{
		Shapes: apiShapes,
		NextId: nextID,
	}, nil
}

func GetShape(ctx context.Context, r *Context, req *api.GetShapeRequest) (*api.Shape, error) {
	system, err := getSystem(ctx, r.Querier, req.SystemId)
	if err != nil {
		return nil, err
	}

	dbShape, err := r.Querier.GetShape(ctx, db.GetShapeParams{
		SystemPk: system.Pk,
		ShapeID:  req.ShapeId,
	})
	if err != nil {
		return nil, err
	}

	apiShapes, err := buildApiShapes([]db.Shape{dbShape})
	if err != nil {
		return nil, err
	}

	return apiShapes[0], nil
}

func buildApiShapes(dbShapes []db.Shape) ([]*api.Shape, error) {
	apiShapes := make([]*api.Shape, len(dbShapes))
	for i, shapeRow := range dbShapes {
		apiShape, err := convert.JSONShapeToApiShape(shapeRow.Shape)
		if err != nil {
			return nil, err
		}

		apiShapes[i] = apiShape
	}
	return apiShapes, nil
}
