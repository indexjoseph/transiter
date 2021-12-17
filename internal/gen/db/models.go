// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Agency struct {
	Pk       int32
	ID       string
	SystemPk int32
	SourcePk int32
	Name     string
	Url      sql.NullString
	Timezone string
	Language sql.NullString
	Phone    sql.NullString
	FareUrl  sql.NullString
	Email    sql.NullString
}

type Alert struct {
	Pk        int32
	ID        string
	SourcePk  int32
	SystemPk  int32
	Cause     string
	Effect    string
	CreatedAt sql.NullTime
	SortOrder sql.NullInt32
	UpdatedAt sql.NullTime
}

type AlertActivePeriod struct {
	Pk       int32
	AlertPk  int32
	StartsAt sql.NullTime
	EndsAt   sql.NullTime
}

type AlertAgency struct {
	AlertPk  int32
	AgencyPk int32
}

type AlertMessage struct {
	Pk          int32
	AlertPk     int32
	Header      string
	Description string
	Url         sql.NullString
	Language    sql.NullString
}

type AlertRoute struct {
	AlertPk int32
	RoutePk int32
}

type AlertStop struct {
	AlertPk int32
	StopPk  int32
}

type AlertTrip struct {
	AlertPk int32
	TripPk  int32
}

type DirectionNameRule struct {
	Pk          int32
	ID          sql.NullString
	StopPk      int32
	SourcePk    int32
	Priority    int32
	DirectionID sql.NullBool
	Track       sql.NullString
	Name        string
}

type Feed struct {
	Pk                int32
	ID                string
	SystemPk          int32
	AutoUpdateEnabled bool
	AutoUpdatePeriod  sql.NullInt32
	Config            string
}

type FeedUpdate struct {
	Pk                 int32
	FeedPk             int32
	ContentLength      sql.NullInt32
	CompletedAt        sql.NullTime
	ContentCreatedAt   sql.NullTime
	ContentHash        sql.NullString
	DownloadDuration   sql.NullFloat64
	Result             sql.NullString
	NumParsedEntities  sql.NullInt32
	NumAddedEntities   sql.NullInt32
	NumUpdatedEntities sql.NullInt32
	NumDeletedEntities sql.NullInt32
	ResultMessage      sql.NullString
	ScheduledAt        sql.NullTime
	TotalDuration      sql.NullFloat64
	UpdateType         string
	Status             sql.NullString
}

type Route struct {
	Pk                int32
	ID                string
	SystemPk          int32
	SourcePk          int32
	Color             sql.NullString
	TextColor         sql.NullString
	ShortName         sql.NullString
	LongName          sql.NullString
	Description       sql.NullString
	Url               sql.NullString
	SortOrder         sql.NullInt32
	Type              sql.NullString
	AgencyPk          int32
	ContinuousDropOff string
	ContinuousPickup  string
}

type ScheduledService struct {
	Pk        int32
	ID        string
	SystemPk  int32
	SourcePk  int32
	Monday    sql.NullBool
	Tuesday   sql.NullBool
	Wednesday sql.NullBool
	Thursday  sql.NullBool
	Friday    sql.NullBool
	Saturday  sql.NullBool
	Sunday    sql.NullBool
	EndDate   sql.NullTime
	StartDate sql.NullTime
}

type ScheduledServiceAddition struct {
	Pk        int32
	ServicePk int32
	Date      time.Time
}

type ScheduledServiceRemoval struct {
	Pk        int32
	ServicePk int32
	Date      time.Time
}

type ScheduledTrip struct {
	Pk                   int32
	ID                   string
	RoutePk              int32
	ServicePk            int32
	DirectionID          sql.NullBool
	BikesAllowed         string
	BlockID              sql.NullString
	Headsign             sql.NullString
	ShortName            sql.NullString
	WheelchairAccessible string
}

type ScheduledTripFrequency struct {
	Pk             int32
	TripPk         int32
	StartTime      time.Time
	EndTime        time.Time
	Headway        int32
	FrequencyBased bool
}

type ScheduledTripStopTime struct {
	Pk                    int32
	TripPk                int32
	StopPk                int32
	ArrivalTime           sql.NullTime
	DepartureTime         sql.NullTime
	StopSequence          int32
	ContinuousDropOff     string
	ContinuousPickup      string
	DropOffType           string
	ExactTimes            bool
	Headsign              sql.NullString
	PickupType            string
	ShapeDistanceTraveled sql.NullFloat64
}

type ServiceMap struct {
	Pk      int32
	RoutePk int32
	GroupPk int32
}

type ServiceMapGroup struct {
	Pk                 int32
	ID                 string
	SystemPk           int32
	Conditions         sql.NullString
	Threshold          float64
	UseForRoutesAtStop bool
	UseForStopsInRoute bool
	Source             string
}

type ServiceMapVertex struct {
	Pk       int32
	StopPk   int32
	MapPk    int32
	Position int32
}

type Stop struct {
	Pk                 int32
	ID                 string
	SystemPk           int32
	SourcePk           int32
	ParentStopPk       sql.NullInt32
	Name               string
	Longitude          sql.NullString
	Latitude           sql.NullString
	Url                sql.NullString
	Code               sql.NullString
	Description        sql.NullString
	PlatformCode       sql.NullString
	Timezone           sql.NullString
	Type               string
	WheelchairBoarding sql.NullString
	ZoneID             sql.NullString
}

type System struct {
	Pk       int32
	ID       string
	Name     string
	Timezone sql.NullString
	Status   string
}

type SystemUpdate struct {
	Pk               int32
	SystemPk         int32
	Status           string
	StatusMessage    sql.NullString
	TotalDuration    sql.NullFloat64
	ScheduledAt      sql.NullTime
	CompletedAt      sql.NullTime
	Config           sql.NullString
	ConfigTemplate   sql.NullString
	ConfigParameters sql.NullString
	ConfigSourceUrl  sql.NullString
	TransiterVersion sql.NullString
}

type Transfer struct {
	Pk              int32
	SourcePk        sql.NullInt32
	ConfigSourcePk  sql.NullInt32
	SystemPk        sql.NullInt32
	FromStopPk      int32
	ToStopPk        int32
	Type            string
	MinTransferTime sql.NullInt32
	Distance        sql.NullInt32
}

type TransfersConfig struct {
	Pk       int32
	Distance string
}

type TransfersConfigSystem struct {
	TransfersConfigPk sql.NullInt32
	SystemPk          sql.NullInt32
}

type Trip struct {
	Pk                  int32
	ID                  string
	RoutePk             int32
	SourcePk            int32
	DirectionID         sql.NullBool
	Delay               sql.NullInt32
	StartedAt           sql.NullTime
	UpdatedAt           sql.NullTime
	CurrentStopSequence sql.NullInt32
}

type TripStopTime struct {
	Pk                   int32
	StopPk               int32
	TripPk               int32
	ArrivalTime          sql.NullTime
	ArrivalDelay         sql.NullInt32
	ArrivalUncertainty   sql.NullInt32
	DepartureTime        sql.NullTime
	DepartureDelay       sql.NullInt32
	DepartureUncertainty sql.NullInt32
	StopSequence         int32
	Track                sql.NullString
}

type Vehicle struct {
	Pk                  int32
	ID                  sql.NullString
	SourcePk            int32
	SystemPk            int32
	TripPk              sql.NullInt32
	Label               sql.NullString
	LicensePlate        sql.NullString
	CurrentStatus       string
	Latitude            sql.NullFloat64
	Longitude           sql.NullFloat64
	Bearing             sql.NullFloat64
	Odometer            sql.NullFloat64
	Speed               sql.NullFloat64
	CongestionLevel     string
	UpdatedAt           sql.NullTime
	CurrentStopPk       sql.NullInt32
	CurrentStopSequence sql.NullInt32
	OccupancyStatus     string
}