package types

type WarningFlags int

const (
	WarningFlag_OK WarningFlags = iota
	WarningFlag_Warning
)

type AngleMeasurementTypes int

const (
	AngleMeasurementType_Degrees AngleMeasurementTypes = iota
	AngleMeasurementType_Hours
)

type RiseSetStatus int

const (
	RiseSetStatus_OK RiseSetStatus = iota
	RiseSetStatus_NeverRises
	RiseSetStatus_Circumpolar
	RiseSetStatus_GstToUtConversionWarning
)

type TwilightStatus int

const (
	TwilightStatus_OK TwilightStatus = iota
	TwilightStatus_LastsAllNight
	TwilightStatus_SunTooFarBelowHorizon
	TwilightStatus_GstToUtConversionWarning
)

type TwilightType int

const (
	_                         TwilightStatus = iota     // skip 0
	TwilightType_Civil                       = iota * 6 // 1*6 = 6
	TwilightType_NAUTICAL                    = iota * 6 // 2*6 = 12
	TwilightType_ASTRONOMICAL                = iota * 6 // 3*6 = 18
)

type CoordinateType int

const (
	CoordinateType_Actual CoordinateType = iota
	CoordinateType_Apparent
)
