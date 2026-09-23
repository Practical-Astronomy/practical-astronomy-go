package types

/* Date information with individual month, day, and year properties. */
type FullDate struct {
	Month int
	Day   float64
	Year  int
}

/* Time information with individual hours, minutes, and seconds properties */
type FullTime struct {
	Hours   int
	Minutes int
	Seconds float64
}

/*
Structure to hold a Time value, along with a calculation warning:

int hours
int minutes
double seconds
WarningFlags warning_flag
*/
type FullTimeWithWarning struct {
	Hours       int
	Minutes     int
	Seconds     float64
	WarningFlag WarningFlags
}

/* Date and Time information with individual month, day, year, hours, minutes, and seconds properties */
type FullDateTime struct {
	Month   int
	Day     int
	Year    int
	Hours   int
	Minutes int
	Seconds float64
}

/* Angle value */
type Angle struct {
	Degrees float64
	Minutes float64
	Seconds float64
}

/* Hour Angle value */
type HourAngle struct {
	Hours   float64
	Minutes float64
	Seconds float64
}

/* Right Ascension value */
type RightAscension struct {
	Hours   float64
	Minutes float64
	Seconds float64
}

/* Horizon Coordinate value */
type HorizonCoordinates struct {
	AzimuthDegrees  float64
	AzimuthMinutes  float64
	AzimuthSeconds  float64
	AltitudeDegrees float64
	AltitudeMinutes float64
	AltitudeSeconds float64
}

/*
Equatorial Coordinate value

Hour Angle: hours, minutes, and seconds
Declination: degrees, minutes, and seconds
*/
type EquatorialCoordinates struct {
	HourAngleHours     float64
	HourAngleMinutes   float64
	HourAngleSeconds   float64
	DeclinationDegrees float64
	DeclinationMinutes float64
	DeclinationSeconds float64
}

/*
Equatorial Coordinate value

Right Ascension: hours, minutes, and seconds
Declination: degrees, minutes, and seconds
*/
type EquatorialCoordinates2 struct {
	RightAscensionHours   float64
	RightAscensionMinutes float64
	RightAscensionSeconds float64
	DeclinationDegrees    float64
	DeclinationMinutes    float64
	DeclinationSeconds    float64
}

/* These types have the same structure as EquatorialCoordinates2 */
type (
	CorrectedPrecession = EquatorialCoordinates2
	CorrectedRefraction = EquatorialCoordinates2
	CorrectedParallax   = EquatorialCoordinates2
)

type EclipticGalacticCoordinates struct {
	LongitudeDegrees float64
	LongitudeMinutes float64
	LongitudeSeconds float64
	LatitudeDegrees  float64
	LatitudeMinutes  float64
	LatitudeSeconds  float64
}

type (
	EclipticCoordinates          = EclipticGalacticCoordinates
	CorrectedEclipticCoordinates = EclipticGalacticCoordinates
	GalacticCoordinates          = EclipticGalacticCoordinates
)

type RiseSet struct {
	RiseSetStatusCurrent RiseSetStatus
	UtRiseHour           float64
	UtRiseMinute         float64
	UtSetHour            float64
	UtSetMinute          float64
	AzRise               float64
	AzSet                float64
}

type Nutation struct {
	NutationInEcliptionLongitude float64
	NutationInObliquity          float64
}

type ParallaxHelper struct {
	P float64
	Q float64
}

type HeliographicCoordinates struct {
	LongitudeDegrees float64
	LatitudeDegrees  float64
}

type SelenographicSubEarthCoordinates struct {
	Longitude           float64
	Latitude            float64
	PositionAngleOfPole float64
}

type SelenographicSubSolarCoordinates struct {
	Longitude   float64
	CoLongitude float64
	Latitude    float64
}

type SunPosition struct {
	RightAscensionHour    float64
	RightAscensionMinutes float64
	RightAscensionSeconds float64
	DeclinationDegrees    float64
	DeclinationMinutes    float64
	DeclinationSeconds    float64
}

type SunDistanceSize struct {
	DistanceInKilometers float64
	AngularSize_Degrees  float64
	AngularSize_Minutes  float64
	AngularSize_Seconds  float64
}

type SunriseSunsetInfo struct {
	LocalSunriseHour    float64
	LocalSunriseMinute  float64
	LocalSunsetHour     float64
	LocalSunsetMinute   float64
	AzimuthOfSunriseDeg float64
	AzimuthOfSunsetDeg  float64
	Status              RiseSetStatus
}

type SunriseLctHelper struct {
	A  float64
	X  float64
	Y  float64
	La float64
	S  RiseSetStatus
}

type (
	SunsetLctHelper = SunriseLctHelper
)

type TwilightInfo struct {
	AmTwilightBeginsHour float64
	AmTwilightBeginsMin  float64
	PmTwilightEndsHour   float64
	PmTwilightEndsMin    float64
	Status               TwilightStatus
}

type TwilightLctHelper struct {
	A  float64
	X  float64
	Y  float64
	La float64
	S  RiseSetStatus
}

type TwilightLctHelper2 struct {
	A  float64
	X  float64
	Y  float64
	La float64
	S  TwilightStatus
}

type EquationOfTime struct {
	Minutes float64
	Seconds float64
}

type PlanetPosition struct {
	RightAscensionHour    float64
	RightAscensionMinutes float64
	RightAscensionSeconds float64
	DeclinationDegrees    float64
	DeclinationMinutes    float64
	DeclinationSeconds    float64
}

type PlanetCoordinates struct {
	Longitude  float64
	Latitude   float64
	DistanceAu float64
	HLong1     float64
	HLong2     float64
	HLat       float64
	RVect      float64
}
