// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PublicClient is the client API for Public service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicClient interface {
	// API entrypoint
	//
	// `GET /`
	//
	// Provides basic information about this Transiter instance and the transit systems it contains.
	Entrypoint(ctx context.Context, in *EntrypointRequest, opts ...grpc.CallOption) (*EntrypointReply, error)
	// List systems
	//
	// `GET /systems`
	//
	// List all transit systems that are installed in this Transiter instance.
	ListSystems(ctx context.Context, in *ListSystemsRequest, opts ...grpc.CallOption) (*ListSystemsReply, error)
	// Get system
	//
	// `GET /systems/<system_id>`
	//
	// Get a system by its ID.
	GetSystem(ctx context.Context, in *GetSystemRequest, opts ...grpc.CallOption) (*System, error)
	// List agencies
	//
	// `GET /systems/<system_id>/agencies`
	//
	// List all agencies in a system.
	ListAgencies(ctx context.Context, in *ListAgenciesRequest, opts ...grpc.CallOption) (*ListAgenciesReply, error)
	// Get agency
	//
	// `GET /systems/<system_id>/agencies/<agency_id>`
	//
	// Get an agency in a system by its ID.
	GetAgency(ctx context.Context, in *GetAgencyRequest, opts ...grpc.CallOption) (*Agency, error)
	// List stops
	//
	// `GET /systems/<system_id>/stops`
	//
	// List all stops in a system.
	//
	// This endpoint is paginated.
	// If there are more results, the `next_id` field of the response will be populated.
	// To get more results, make the same request with the `first_id` field set to the value of `next_id` in the response.
	ListStops(ctx context.Context, in *ListStopsRequest, opts ...grpc.CallOption) (*ListStopsReply, error)
	// Get stop
	//
	// `GET /systems/<system_id>/stops/<stop_id>`
	//
	// Get a stop in a system by its ID.
	GetStop(ctx context.Context, in *GetStopRequest, opts ...grpc.CallOption) (*Stop, error)
	// List routes
	//
	// `GET /systems/<system_id>/routes`
	//
	// List all routes in a system.
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesReply, error)
	// Get route
	//
	// `GET /systems/<system_id>/routes/<route_id>`
	//
	// Get a route in a system by its ID.
	GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*Route, error)
	// List trips
	//
	// `GET /systems/<system_id>/routes/<route_id>/trips`
	//
	// List all trips in a route.
	ListTrips(ctx context.Context, in *ListTripsRequest, opts ...grpc.CallOption) (*ListTripsReply, error)
	// Get trip
	//
	// `GET /systems/<system_id>/routes/<route_id>/trips/<trip_id>`
	//
	// Get a trip by its ID.
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*Trip, error)
	// List alerts
	//
	// `GET /systems/<system_id>/alerts`
	//
	// List all alerts in a system.
	// By default this endpoint returns both active alerts
	//
	//	(alerts which have an active period containing the current time) and non-active alerts.
	ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsReply, error)
	// Get alert
	//
	// `GET /systems/<system_id>/alerts/<alert_id>`
	//
	// Get an alert by its ID.
	GetAlert(ctx context.Context, in *GetAlertRequest, opts ...grpc.CallOption) (*Alert, error)
	// List transfers
	//
	// `GET /systems/<system_id>/transfers`
	//
	// List all transfers in a system.
	ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersReply, error)
	// Get transfer
	//
	// `GET /systems/<system_id>/transfers/<transfer_id>`
	//
	// Get a transfer by its ID.
	GetTransfer(ctx context.Context, in *GetTransferRequest, opts ...grpc.CallOption) (*Transfer, error)
	// List feeds
	//
	// `GET /systems/<system_id>/feeds`
	//
	// List all feeds for a system.
	ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsReply, error)
	// Get feed
	//
	// `GET /systems/<system_id>/feeds/<feed_id>`
	//
	// Get a feed in a system by its ID.
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*Feed, error)
	// List vehicles
	//
	// `GET /systems/<system_id>/vehicles`
	//
	// List all feeds for a system.
	ListVehicles(ctx context.Context, in *ListVehiclesRequest, opts ...grpc.CallOption) (*ListVehiclesReply, error)
	// Get vehicle
	//
	// `GET /systems/<system_id>/vehicles/<vehicle_id>`
	//
	// Get a vehicle in a system by its ID.
	GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*Vehicle, error)
	// List shapes
	//
	// `GET /systems/<system_id>/shapes`
	//
	// List all shapes in a system.
	ListShapes(ctx context.Context, in *ListShapesRequest, opts ...grpc.CallOption) (*ListShapesReply, error)
	// Get shape
	//
	// `GET /systems/<system_id>/shapes/<shape_id>`
	//
	// Get a shape in a system by its ID.
	GetShape(ctx context.Context, in *GetShapeRequest, opts ...grpc.CallOption) (*Shape, error)
}

type publicClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicClient(cc grpc.ClientConnInterface) PublicClient {
	return &publicClient{cc}
}

func (c *publicClient) Entrypoint(ctx context.Context, in *EntrypointRequest, opts ...grpc.CallOption) (*EntrypointReply, error) {
	out := new(EntrypointReply)
	err := c.cc.Invoke(ctx, "/Public/Entrypoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListSystems(ctx context.Context, in *ListSystemsRequest, opts ...grpc.CallOption) (*ListSystemsReply, error) {
	out := new(ListSystemsReply)
	err := c.cc.Invoke(ctx, "/Public/ListSystems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetSystem(ctx context.Context, in *GetSystemRequest, opts ...grpc.CallOption) (*System, error) {
	out := new(System)
	err := c.cc.Invoke(ctx, "/Public/GetSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListAgencies(ctx context.Context, in *ListAgenciesRequest, opts ...grpc.CallOption) (*ListAgenciesReply, error) {
	out := new(ListAgenciesReply)
	err := c.cc.Invoke(ctx, "/Public/ListAgencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetAgency(ctx context.Context, in *GetAgencyRequest, opts ...grpc.CallOption) (*Agency, error) {
	out := new(Agency)
	err := c.cc.Invoke(ctx, "/Public/GetAgency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListStops(ctx context.Context, in *ListStopsRequest, opts ...grpc.CallOption) (*ListStopsReply, error) {
	out := new(ListStopsReply)
	err := c.cc.Invoke(ctx, "/Public/ListStops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetStop(ctx context.Context, in *GetStopRequest, opts ...grpc.CallOption) (*Stop, error) {
	out := new(Stop)
	err := c.cc.Invoke(ctx, "/Public/GetStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesReply, error) {
	out := new(ListRoutesReply)
	err := c.cc.Invoke(ctx, "/Public/ListRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := c.cc.Invoke(ctx, "/Public/GetRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListTrips(ctx context.Context, in *ListTripsRequest, opts ...grpc.CallOption) (*ListTripsReply, error) {
	out := new(ListTripsReply)
	err := c.cc.Invoke(ctx, "/Public/ListTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/Public/GetTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsReply, error) {
	out := new(ListAlertsReply)
	err := c.cc.Invoke(ctx, "/Public/ListAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetAlert(ctx context.Context, in *GetAlertRequest, opts ...grpc.CallOption) (*Alert, error) {
	out := new(Alert)
	err := c.cc.Invoke(ctx, "/Public/GetAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListTransfers(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersReply, error) {
	out := new(ListTransfersReply)
	err := c.cc.Invoke(ctx, "/Public/ListTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetTransfer(ctx context.Context, in *GetTransferRequest, opts ...grpc.CallOption) (*Transfer, error) {
	out := new(Transfer)
	err := c.cc.Invoke(ctx, "/Public/GetTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsReply, error) {
	out := new(ListFeedsReply)
	err := c.cc.Invoke(ctx, "/Public/ListFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*Feed, error) {
	out := new(Feed)
	err := c.cc.Invoke(ctx, "/Public/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListVehicles(ctx context.Context, in *ListVehiclesRequest, opts ...grpc.CallOption) (*ListVehiclesReply, error) {
	out := new(ListVehiclesReply)
	err := c.cc.Invoke(ctx, "/Public/ListVehicles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*Vehicle, error) {
	out := new(Vehicle)
	err := c.cc.Invoke(ctx, "/Public/GetVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) ListShapes(ctx context.Context, in *ListShapesRequest, opts ...grpc.CallOption) (*ListShapesReply, error) {
	out := new(ListShapesReply)
	err := c.cc.Invoke(ctx, "/Public/ListShapes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetShape(ctx context.Context, in *GetShapeRequest, opts ...grpc.CallOption) (*Shape, error) {
	out := new(Shape)
	err := c.cc.Invoke(ctx, "/Public/GetShape", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServer is the server API for Public service.
// All implementations should embed UnimplementedPublicServer
// for forward compatibility
type PublicServer interface {
	// API entrypoint
	//
	// `GET /`
	//
	// Provides basic information about this Transiter instance and the transit systems it contains.
	Entrypoint(context.Context, *EntrypointRequest) (*EntrypointReply, error)
	// List systems
	//
	// `GET /systems`
	//
	// List all transit systems that are installed in this Transiter instance.
	ListSystems(context.Context, *ListSystemsRequest) (*ListSystemsReply, error)
	// Get system
	//
	// `GET /systems/<system_id>`
	//
	// Get a system by its ID.
	GetSystem(context.Context, *GetSystemRequest) (*System, error)
	// List agencies
	//
	// `GET /systems/<system_id>/agencies`
	//
	// List all agencies in a system.
	ListAgencies(context.Context, *ListAgenciesRequest) (*ListAgenciesReply, error)
	// Get agency
	//
	// `GET /systems/<system_id>/agencies/<agency_id>`
	//
	// Get an agency in a system by its ID.
	GetAgency(context.Context, *GetAgencyRequest) (*Agency, error)
	// List stops
	//
	// `GET /systems/<system_id>/stops`
	//
	// List all stops in a system.
	//
	// This endpoint is paginated.
	// If there are more results, the `next_id` field of the response will be populated.
	// To get more results, make the same request with the `first_id` field set to the value of `next_id` in the response.
	ListStops(context.Context, *ListStopsRequest) (*ListStopsReply, error)
	// Get stop
	//
	// `GET /systems/<system_id>/stops/<stop_id>`
	//
	// Get a stop in a system by its ID.
	GetStop(context.Context, *GetStopRequest) (*Stop, error)
	// List routes
	//
	// `GET /systems/<system_id>/routes`
	//
	// List all routes in a system.
	ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesReply, error)
	// Get route
	//
	// `GET /systems/<system_id>/routes/<route_id>`
	//
	// Get a route in a system by its ID.
	GetRoute(context.Context, *GetRouteRequest) (*Route, error)
	// List trips
	//
	// `GET /systems/<system_id>/routes/<route_id>/trips`
	//
	// List all trips in a route.
	ListTrips(context.Context, *ListTripsRequest) (*ListTripsReply, error)
	// Get trip
	//
	// `GET /systems/<system_id>/routes/<route_id>/trips/<trip_id>`
	//
	// Get a trip by its ID.
	GetTrip(context.Context, *GetTripRequest) (*Trip, error)
	// List alerts
	//
	// `GET /systems/<system_id>/alerts`
	//
	// List all alerts in a system.
	// By default this endpoint returns both active alerts
	//
	//	(alerts which have an active period containing the current time) and non-active alerts.
	ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsReply, error)
	// Get alert
	//
	// `GET /systems/<system_id>/alerts/<alert_id>`
	//
	// Get an alert by its ID.
	GetAlert(context.Context, *GetAlertRequest) (*Alert, error)
	// List transfers
	//
	// `GET /systems/<system_id>/transfers`
	//
	// List all transfers in a system.
	ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersReply, error)
	// Get transfer
	//
	// `GET /systems/<system_id>/transfers/<transfer_id>`
	//
	// Get a transfer by its ID.
	GetTransfer(context.Context, *GetTransferRequest) (*Transfer, error)
	// List feeds
	//
	// `GET /systems/<system_id>/feeds`
	//
	// List all feeds for a system.
	ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsReply, error)
	// Get feed
	//
	// `GET /systems/<system_id>/feeds/<feed_id>`
	//
	// Get a feed in a system by its ID.
	GetFeed(context.Context, *GetFeedRequest) (*Feed, error)
	// List vehicles
	//
	// `GET /systems/<system_id>/vehicles`
	//
	// List all feeds for a system.
	ListVehicles(context.Context, *ListVehiclesRequest) (*ListVehiclesReply, error)
	// Get vehicle
	//
	// `GET /systems/<system_id>/vehicles/<vehicle_id>`
	//
	// Get a vehicle in a system by its ID.
	GetVehicle(context.Context, *GetVehicleRequest) (*Vehicle, error)
	// List shapes
	//
	// `GET /systems/<system_id>/shapes`
	//
	// List all shapes in a system.
	ListShapes(context.Context, *ListShapesRequest) (*ListShapesReply, error)
	// Get shape
	//
	// `GET /systems/<system_id>/shapes/<shape_id>`
	//
	// Get a shape in a system by its ID.
	GetShape(context.Context, *GetShapeRequest) (*Shape, error)
}

// UnimplementedPublicServer should be embedded to have forward compatible implementations.
type UnimplementedPublicServer struct {
}

func (UnimplementedPublicServer) Entrypoint(context.Context, *EntrypointRequest) (*EntrypointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entrypoint not implemented")
}
func (UnimplementedPublicServer) ListSystems(context.Context, *ListSystemsRequest) (*ListSystemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystems not implemented")
}
func (UnimplementedPublicServer) GetSystem(context.Context, *GetSystemRequest) (*System, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystem not implemented")
}
func (UnimplementedPublicServer) ListAgencies(context.Context, *ListAgenciesRequest) (*ListAgenciesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgencies not implemented")
}
func (UnimplementedPublicServer) GetAgency(context.Context, *GetAgencyRequest) (*Agency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgency not implemented")
}
func (UnimplementedPublicServer) ListStops(context.Context, *ListStopsRequest) (*ListStopsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStops not implemented")
}
func (UnimplementedPublicServer) GetStop(context.Context, *GetStopRequest) (*Stop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStop not implemented")
}
func (UnimplementedPublicServer) ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}
func (UnimplementedPublicServer) GetRoute(context.Context, *GetRouteRequest) (*Route, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoute not implemented")
}
func (UnimplementedPublicServer) ListTrips(context.Context, *ListTripsRequest) (*ListTripsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrips not implemented")
}
func (UnimplementedPublicServer) GetTrip(context.Context, *GetTripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (UnimplementedPublicServer) ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlerts not implemented")
}
func (UnimplementedPublicServer) GetAlert(context.Context, *GetAlertRequest) (*Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlert not implemented")
}
func (UnimplementedPublicServer) ListTransfers(context.Context, *ListTransfersRequest) (*ListTransfersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransfers not implemented")
}
func (UnimplementedPublicServer) GetTransfer(context.Context, *GetTransferRequest) (*Transfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransfer not implemented")
}
func (UnimplementedPublicServer) ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeeds not implemented")
}
func (UnimplementedPublicServer) GetFeed(context.Context, *GetFeedRequest) (*Feed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedPublicServer) ListVehicles(context.Context, *ListVehiclesRequest) (*ListVehiclesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicles not implemented")
}
func (UnimplementedPublicServer) GetVehicle(context.Context, *GetVehicleRequest) (*Vehicle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicle not implemented")
}
func (UnimplementedPublicServer) ListShapes(context.Context, *ListShapesRequest) (*ListShapesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShapes not implemented")
}
func (UnimplementedPublicServer) GetShape(context.Context, *GetShapeRequest) (*Shape, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShape not implemented")
}

// UnsafePublicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicServer will
// result in compilation errors.
type UnsafePublicServer interface {
	mustEmbedUnimplementedPublicServer()
}

func RegisterPublicServer(s grpc.ServiceRegistrar, srv PublicServer) {
	s.RegisterService(&Public_ServiceDesc, srv)
}

func _Public_Entrypoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntrypointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).Entrypoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/Entrypoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).Entrypoint(ctx, req.(*EntrypointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListSystems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListSystems(ctx, req.(*ListSystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetSystem(ctx, req.(*GetSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListAgencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListAgencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListAgencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListAgencies(ctx, req.(*ListAgenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetAgency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetAgency(ctx, req.(*GetAgencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListStops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStopsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListStops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListStops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListStops(ctx, req.(*ListStopsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetStop(ctx, req.(*GetStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListRoutes(ctx, req.(*ListRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetRoute(ctx, req.(*GetRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListTrips(ctx, req.(*ListTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListAlerts(ctx, req.(*ListAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetAlert(ctx, req.(*GetAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListTransfers(ctx, req.(*ListTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetTransfer(ctx, req.(*GetTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListFeeds(ctx, req.(*ListFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVehiclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListVehicles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListVehicles(ctx, req.(*ListVehiclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetVehicle(ctx, req.(*GetVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_ListShapes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShapesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).ListShapes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/ListShapes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).ListShapes(ctx, req.(*ListShapesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetShape_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShapeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetShape(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Public/GetShape",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetShape(ctx, req.(*GetShapeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Public_ServiceDesc is the grpc.ServiceDesc for Public service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Public_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Public",
	HandlerType: (*PublicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Entrypoint",
			Handler:    _Public_Entrypoint_Handler,
		},
		{
			MethodName: "ListSystems",
			Handler:    _Public_ListSystems_Handler,
		},
		{
			MethodName: "GetSystem",
			Handler:    _Public_GetSystem_Handler,
		},
		{
			MethodName: "ListAgencies",
			Handler:    _Public_ListAgencies_Handler,
		},
		{
			MethodName: "GetAgency",
			Handler:    _Public_GetAgency_Handler,
		},
		{
			MethodName: "ListStops",
			Handler:    _Public_ListStops_Handler,
		},
		{
			MethodName: "GetStop",
			Handler:    _Public_GetStop_Handler,
		},
		{
			MethodName: "ListRoutes",
			Handler:    _Public_ListRoutes_Handler,
		},
		{
			MethodName: "GetRoute",
			Handler:    _Public_GetRoute_Handler,
		},
		{
			MethodName: "ListTrips",
			Handler:    _Public_ListTrips_Handler,
		},
		{
			MethodName: "GetTrip",
			Handler:    _Public_GetTrip_Handler,
		},
		{
			MethodName: "ListAlerts",
			Handler:    _Public_ListAlerts_Handler,
		},
		{
			MethodName: "GetAlert",
			Handler:    _Public_GetAlert_Handler,
		},
		{
			MethodName: "ListTransfers",
			Handler:    _Public_ListTransfers_Handler,
		},
		{
			MethodName: "GetTransfer",
			Handler:    _Public_GetTransfer_Handler,
		},
		{
			MethodName: "ListFeeds",
			Handler:    _Public_ListFeeds_Handler,
		},
		{
			MethodName: "GetFeed",
			Handler:    _Public_GetFeed_Handler,
		},
		{
			MethodName: "ListVehicles",
			Handler:    _Public_ListVehicles_Handler,
		},
		{
			MethodName: "GetVehicle",
			Handler:    _Public_GetVehicle_Handler,
		},
		{
			MethodName: "ListShapes",
			Handler:    _Public_ListShapes_Handler,
		},
		{
			MethodName: "GetShape",
			Handler:    _Public_GetShape_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/public.proto",
}
