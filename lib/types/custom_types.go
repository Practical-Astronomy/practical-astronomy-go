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

int Hours
int Minutes
float64 Seconds
WarningFlags WarningFlag
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

type PlanetLongLatL4685 struct {
	QA float64
	QB float64
}

type PlanetLongLatL4735 struct {
	QA float64
	QB float64
	QC float64
	QE float64
}

type PlanetLongLatL4810 struct {
	A  float64
	SA float64
	CA float64
	QC float64
	QE float64
	QA float64
	QB float64
}

type PlanetLongLatL4945 struct {
	QA float64
	QB float64
	QC float64
	QD float64
	QE float64
	QF float64
	QG float64
}

type PlanetVisualAspects struct {
	DistanceAu            float64
	AngDiaArcsec          float64
	Phase                 float64
	LightTimeHour         float64
	LightTimeMinutes      float64
	LightTimeSeconds      float64
	PosAngleBrightLimbDeg float64
	ApproximateMagnitude  float64
}

type EllipticalCometPosition struct {
	RaHour    float64
	RaMin     float64
	DecDeg    float64
	DecMin    float64
	DistEarth float64
}

type ParabolicCometPosition struct {
	RaHour    float64
	RaMin     float64
	RaSec     float64
	DecDeg    float64
	DecMin    float64
	DecSec    float64
	DistEarth float64
}

type CometLongLatDist struct {
	LongDeg float64
	LatDeg  float64
	DistAu  float64
}

type BinaryStarOrbitalData struct {
	PositionAngleDeg float64
	SeparationArcsec float64
}

type MoonApproximatePosition struct {
	RaHour float64
	RaMin  float64
	RaSec  float64
	DecDeg float64
	DecMin float64
	DecSec float64
}

type MoonPrecisePosition struct {
	RaHour          float64
	RaMin           float64
	RaSec           float64
	DecDeg          float64
	DecMin          float64
	DecSec          float64
	EarthMoonDistKm float64
	HorParallaxDeg  float64
}

type MoonLongLatHP struct {
	LongDeg float64
	LatDeg  float64
	HorPara float64
}

type MoonPhase struct {
	Phase         float64
	BrightLimbDeg float64
}

type MoonNewFull struct {
	NewLocalTimeHour   float64
	NewLocalTimeMin    float64
	NewLocalDateDay    float64
	NewLocalDateMonth  int
	NewLocalDateYear   int
	FullLocalTimeHour  float64
	FullLocalTimeMin   float64
	FullLocalDateDay   float64
	FullLocalDateMonth int
	FullLocalDateYear  int
}

type NewMoonFullMoonL6855 struct {
	A float64
	B float64
	F float64
}
