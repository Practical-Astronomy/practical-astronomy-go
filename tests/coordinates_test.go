package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAngleToDecimalDegrees(t *testing.T) {
	var degrees float64 = 182
	var minutes float64 = 31
	var seconds float64 = 27

	var expectedDecimalDegrees float64 = 182.524167
	var actualDecimalDegrees float64 = pautil.RoundTo(palib.AngleToDecimalDegrees(degrees, minutes, seconds), 6)

	require.Equal(t, expectedDecimalDegrees, actualDecimalDegrees, "Mismatch in Decimal Degrees")
}

func TestDecimalDegreesToAngle(t *testing.T) {
	var decimalDegrees float64 = 182.524167

	var expectedAngle patype.Angle = patype.Angle{Degrees: 182, Minutes: 31, Seconds: 27}
	var actualAngle patype.Angle = palib.DecimalDegreesToAngle(decimalDegrees)

	require.Equal(t, expectedAngle, actualAngle, "Mismatch in Angle")
}
func TestRightAscensionToHourAngle(t *testing.T) {
	var raHours float64 = 18
	var raMinutes float64 = 32
	var raSeconds float64 = 21
	var lctHours float64 = 14
	var lctMinutes float64 = 36
	var lctSeconds float64 = 51.67
	var isDaylightSavings bool = false
	var zoneCorrection int = -4
	var localDay float64 = 22
	var localMonth int = 4
	var localYear int = 1980
	var geographicalLongitude float64 = -64

	var expectedHourAngle patype.HourAngle = patype.HourAngle{Hours: 9, Minutes: 52, Seconds: 23.66}
	var actualHourAngle patype.HourAngle = palib.RightAscensionToHourAngle(
		raHours, raMinutes, raSeconds, lctHours, lctMinutes, lctSeconds, isDaylightSavings, zoneCorrection,
		localDay, localMonth, localYear, geographicalLongitude,
	)

	require.Equal(t, expectedHourAngle, actualHourAngle, "Mismatch in Hour Angle")
}

func TestHourAngleToRightAscension(t *testing.T) {
	var hourAngleHours float64 = 9
	var hourAngleMinutes float64 = 52
	var hourAngleSeconds float64 = 23.66
	var lctHours float64 = 14
	var lctMinutes float64 = 36
	var lctSeconds float64 = 51.67
	var isDaylightSavings bool = false
	var zoneCorrection int = -4
	var localDay float64 = 22
	var localMonth int = 4
	var localYear int = 1980
	var geographicalLongitude float64 = -64

	var expectedRightAscension patype.RightAscension = patype.RightAscension{Hours: 18, Minutes: 32, Seconds: 21}
	var actualRightAscension patype.RightAscension = palib.HourAngleToRightAscension(
		hourAngleHours, hourAngleMinutes, hourAngleSeconds, lctHours, lctMinutes, lctSeconds, isDaylightSavings, zoneCorrection,
		localDay, localMonth, localYear, geographicalLongitude,
	)

	require.Equal(t, expectedRightAscension, actualRightAscension, "Mismatch in Right Ascension")
}
func TestEquatorialCoordinatesToHorizonCoordinates(t *testing.T) {
	var hourAngleHours float64 = 5
	var hourAngleMinutes float64 = 51
	var hourAngleSeconds float64 = 44
	var declinationDegrees float64 = 23
	var declinationMinutes float64 = 13
	var declinationSeconds float64 = 10
	var geographicalLatitude float64 = 52

	var expectedHorizonCoordinates patype.HorizonCoordinates = patype.HorizonCoordinates{
		AzimuthDegrees: 283, AzimuthMinutes: 16, AzimuthSeconds: 15.7, AltitudeDegrees: 19, AltitudeMinutes: 20, AltitudeSeconds: 3.64,
	}
	var actualHorizonCoordinates patype.HorizonCoordinates = palib.EquatorialCoordinatesToHorizonCoordinates(
		hourAngleHours, hourAngleMinutes, hourAngleSeconds,
		declinationDegrees, declinationMinutes, declinationSeconds,
		geographicalLatitude,
	)

	require.Equal(t, expectedHorizonCoordinates, actualHorizonCoordinates, "Mismatch in Horizon Coordinates")
}
func TestHorizonCoordinatesToEquatorialCoordinates(t *testing.T) {
	var azimuthDegrees float64 = 283
	var azimuthMinutes float64 = 16
	var azimuthSeconds float64 = 15.7
	var altitudeDegrees float64 = 19
	var altitudeMinutes float64 = 20
	var altitudeSeconds float64 = 3.64
	var geographicalLatitude float64 = 52

	var expectedEquatorialCoordinates patype.EquatorialCoordinates = patype.EquatorialCoordinates{
		HourAngleHours: 5, HourAngleMinutes: 51, HourAngleSeconds: 44, DeclinationDegrees: 23, DeclinationMinutes: 13, DeclinationSeconds: 10,
	}
	var actualEquatorialCoordinates patype.EquatorialCoordinates = palib.HorizonCoordinatesToEquatorialCoordinates(
		azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds, geographicalLatitude)

	require.Equal(t, expectedEquatorialCoordinates, actualEquatorialCoordinates, "Mismatch in Equatorial Coordinates")
}
func TestMeanObliquityOfTheEcliptic(t *testing.T) {
	var greenwichDay float64 = 6
	var greenwichMonth int = 7
	var greenwichYear = 2009

	var expectedMeanObliquity float64 = 23.43805531
	var actualMeanObliquity = palib.MeanObliquityOfTheEcliptic(greenwichDay, greenwichMonth, greenwichYear)

	require.Equal(t, expectedMeanObliquity, pautil.RoundTo(actualMeanObliquity, 8), "Mismatch in Mean Obliquity")
}

func TestEclipticCoordinatesToEquatorialCoordinates(t *testing.T) {
	var eclipticLongitudeDegrees float64 = 139
	var eclipticLongitudeMinutes float64 = 41
	var eclipticLongitudeSeconds float64 = 10
	var eclipticLatitudeDegrees float64 = 4
	var eclipticLatitudeMinutes float64 = 52
	var eclipticLatitudeSeconds float64 = 31
	var greenwichDay float64 = 6
	var greenwichMonth int = 7
	var greenwichYear int = 2009

	var expectedEquatorialCoordinates patype.EquatorialCoordinates2 = patype.EquatorialCoordinates2{
		RightAscensionHours: 9, RightAscensionMinutes: 34, RightAscensionSeconds: 53.4,
		DeclinationDegrees: 19, DeclinationMinutes: 32, DeclinationSeconds: 8.52,
	}
	var actualEquatorialCoordinates patype.EquatorialCoordinates2 = palib.EclipticCoordinatesToEquatorialCoordinates(
		eclipticLongitudeDegrees, eclipticLongitudeMinutes, eclipticLongitudeSeconds,
		eclipticLatitudeDegrees, eclipticLatitudeMinutes, eclipticLatitudeSeconds,
		greenwichDay, greenwichMonth, greenwichYear,
	)

	require.Equal(t, expectedEquatorialCoordinates, actualEquatorialCoordinates, "Mismatch in Equatorial Coordinates")
}
func TestEquatorialCoordinateToEclipticCoordinate(t *testing.T) {
	var raHours float64 = 9
	var raMinutes float64 = 34
	var raSeconds float64 = 53.4
	var decDegrees float64 = 19
	var decMinutes float64 = 32
	var decSeconds float64 = 8.52
	var gwDay float64 = 6
	var gwMonth int = 7
	var gwYear int = 2009

	var expectedEclipticCoordinates patype.EclipticCoordinates = patype.EclipticCoordinates{
		LongitudeDegrees: 139, LongitudeMinutes: 41, LongitudeSeconds: 9.97,
		LatitudeDegrees: 4, LatitudeMinutes: 52, LatitudeSeconds: 30.99,
	}
	var actualEclipticCoordinates patype.EclipticCoordinates = palib.EquatorialCoordinateToEclipticCoordinate(
		raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds, gwDay, gwMonth, gwYear)

	require.Equal(t, expectedEclipticCoordinates, actualEclipticCoordinates, "Mismatch in Ecliptic Coordinates")
}
func TestEquatorialCoordinateToGalacticCoordinate(t *testing.T) {
	var raHours float64 = 10
	var raMinutes float64 = 21
	var raSeconds float64 = 0
	var decDegrees float64 = 10
	var decMinutes float64 = 3
	var decSeconds float64 = 11

	var expectedGalacticCoordinates patype.GalacticCoordinates = patype.GalacticCoordinates{
		LongitudeDegrees: 232, LongitudeMinutes: 14, LongitudeSeconds: 52.38,
		LatitudeDegrees: 51, LatitudeMinutes: 7, LatitudeSeconds: 20.16,
	}
	var actualGalacticCoordinates patype.GalacticCoordinates = palib.EquatorialCoordinateToGalacticCoordinate(
		raHours, raMinutes, raSeconds, decDegrees, decMinutes, decSeconds)

	require.Equal(t, expectedGalacticCoordinates, actualGalacticCoordinates, "Mismatch in Galactic Coordinates")

}

func TestGalacticCoordinatesToEquatorialCoordinates(t *testing.T) {
	var galLongDeg float64 = 232
	var galLongMin float64 = 14
	var galLongSec float64 = 52.38
	var galLatDeg float64 = 51
	var galLatMin float64 = 7
	var galLatSec float64 = 20.16

	var expectedEquatorialCoordinates patype.EquatorialCoordinates2 = patype.EquatorialCoordinates2{
		RightAscensionHours: 10, RightAscensionMinutes: 21, RightAscensionSeconds: 0,
		DeclinationDegrees: 10, DeclinationMinutes: 3, DeclinationSeconds: 11,
	}
	var actualEquatorialCoordinates patype.EquatorialCoordinates2 = palib.GalacticCoordinatesToEquatorialCoordinates(
		galLongDeg, galLongMin, galLongSec, galLatDeg, galLatMin, galLatSec)

	require.Equal(t, expectedEquatorialCoordinates, actualEquatorialCoordinates, "Mismatch in Equatorial Coordinates")
}

func TestAngleBetweenTwoObjects(t *testing.T) {
	var raLong1HourDeg float64 = 5
	var raLong1Min float64 = 13
	var raLong1Sec float64 = 31.7
	var decLat1Deg float64 = -8
	var decLat1Min float64 = 13
	var decLat1Sec float64 = 30
	var raLong2HourDeg float64 = 6
	var raLong2Min float64 = 44
	var raLong2Sec float64 = 13.4
	var decLat2Deg float64 = -16
	var decLat2Min float64 = 41
	var decLat2Sec float64 = 11
	var hourOrDegree patype.AngleMeasurementTypes = patype.AngleMeasurementType_Hours

	var expectedAngleBetweenTwoObjects patype.Angle = patype.Angle{Degrees: 23, Minutes: 40, Seconds: 25.86}
	var actualAngleBetweenTwoObjects patype.Angle = palib.AngleBetweenTwoObjects(
		raLong1HourDeg, raLong1Min, raLong1Sec, decLat1Deg, decLat1Min, decLat1Sec,
		raLong2HourDeg, raLong2Min, raLong2Sec, decLat2Deg, decLat2Min, decLat2Sec, hourOrDegree,
	)

	require.Equal(t, expectedAngleBetweenTwoObjects, actualAngleBetweenTwoObjects, "Mismatch in Angle Between Two Objects")
}

func TestRisingAndSetting(t *testing.T) {
	var raHours float64 = 23
	var raMinutes float64 = 39
	var raSeconds float64 = 20
	var decDeg float64 = 21
	var decMin float64 = 42
	var decSec float64 = 0
	var gwDateDay float64 = 24
	var gwDateMonth int = 8
	var gwDateYear int = 2010
	var geogLongDeg float64 = 64
	var geogLatDeg float64 = 30
	var vertShiftDeg float64 = 0.5667

	var expectedRiseSet patype.RiseSet = patype.RiseSet{
		RiseSetStatusCurrent: patype.RiseSetStatus_OK, UtRiseHour: 14, UtRiseMinute: 16, UtSetHour: 4, UtSetMinute: 10, AzRise: 64.36, AzSet: 295.64}
	var actualRiseSet patype.RiseSet = palib.RisingAndSetting(
		raHours, raMinutes, raSeconds, decDeg, decMin, decSec, gwDateDay, gwDateMonth, gwDateYear, geogLongDeg, geogLatDeg, vertShiftDeg)

	require.Equal(t, expectedRiseSet, actualRiseSet, "Mismatch in Rise/Set")
}

func TestCorrectForPrecession(t *testing.T) {
	var raHour float64 = 9
	var raMinutes float64 = 10
	var raSeconds float64 = 43
	var decDeg float64 = 14
	var decMinutes float64 = 23
	var decSeconds float64 = 25
	var epoch1Day float64 = 0.923
	var epoch1Month int = 1
	var epoch1Year int = 1950
	var epoch2Day float64 = 1
	var epoch2Month int = 6
	var epoch2Year int = 1979

	var expectedCorrection patype.CorrectedPrecession = patype.CorrectedPrecession{RightAscensionHours: 9, RightAscensionMinutes: 12, RightAscensionSeconds: 20.18, DeclinationDegrees: 14, DeclinationMinutes: 16, DeclinationSeconds: 9.12}
	var actualCorrection patype.CorrectedPrecession = palib.CorrectForPrecession(raHour, raMinutes, raSeconds, decDeg, decMinutes, decSeconds, epoch1Day, epoch1Month, epoch1Year, epoch2Day, epoch2Month, epoch2Year)

	require.Equal(t, expectedCorrection, actualCorrection, "Mismatch in Correction for Precession")
}

func TestNutationInEclipticLongitudeAndObliquity(t *testing.T) {
	var greenwichDay float64 = 1
	var greenwichMonth int = 9
	var greenwichYear int = 1988

	var expectedNutation patype.Nutation = patype.Nutation{NutationInEcliptionLongitude: .001525808, NutationInObliquity: .0025671}
	var actualNutation patype.Nutation = palib.NutationInEclipticLongitudeAndObliquity(greenwichDay, greenwichMonth, greenwichYear)

	require.Equal(t, expectedNutation.NutationInEcliptionLongitude, pautil.RoundTo(actualNutation.NutationInEcliptionLongitude, 9), "Mismatch in Nutation")
	require.Equal(t, expectedNutation.NutationInObliquity, pautil.RoundTo(actualNutation.NutationInObliquity, 7), "Mismatch in Nutation")
}

func TestCorrectForAberration(t *testing.T) {
	var utHour float64 = 0
	var utMinutes float64 = 0
	var utSeconds float64 = 0
	var gwDay float64 = 8
	var gwMonth int = 9
	var gwYear int = 1988
	var trueEclLongDeg float64 = 352
	var trueEclLongMin float64 = 37
	var trueEclLongSec float64 = 10.1
	var trueEclLatDeg float64 = -1
	var trueEclLatMin float64 = 32
	var trueEclLatSec float64 = 56.4

	var expectedAberration patype.CorrectedEclipticCoordinates = patype.CorrectedEclipticCoordinates{
		LongitudeDegrees: 352, LongitudeMinutes: 37, LongitudeSeconds: 30.45,
		LatitudeDegrees: -1, LatitudeMinutes: 32, LatitudeSeconds: 56.33,
	}
	var actualAberration patype.CorrectedEclipticCoordinates = palib.CorrectForAberration(
		utHour, utMinutes, utSeconds, gwDay, gwMonth, gwYear,
		trueEclLongDeg, trueEclLongMin, trueEclLongSec,
		trueEclLatDeg, trueEclLatMin, trueEclLatSec,
	)

	require.Equal(t, expectedAberration, actualAberration, "Mismatched in Corrected Aberration")
}

func TestAtmosphericRefraction(t *testing.T) {
	var trueRaHour float64 = 23
	var trueRaMin float64 = 14
	var trueRaSec float64 = 0
	var trueDecDeg float64 = 40
	var trueDecMin float64 = 10
	var trueDecSec float64 = 0
	var coordinateType1 patype.CoordinateType = patype.CoordinateType_Actual
	var geogLongDeg float64 = 0.17
	var geogLatDeg float64 = 51.2036110
	var daylightSavingHours int = 0
	var timezoneHours int = 0
	var lcdDay float64 = 23
	var lcdMonth int = 3
	var lcdYear int = 1987
	var lctHour float64 = 1
	var lctMin float64 = 1
	var lctSec float64 = 24
	var atmosphericPressureMbar float64 = 1012
	var atmosphericTemperatureCelsius float64 = 21.7

	var expectedRefraction patype.CorrectedRefraction = patype.CorrectedRefraction{
		RightAscensionHours: 23, RightAscensionMinutes: 13, RightAscensionSeconds: 44.74,
		DeclinationDegrees: 40, DeclinationMinutes: 19, DeclinationSeconds: 45.76,
	}
	var actualRefraction patype.CorrectedRefraction = palib.AtmosphericRefraction(
		trueRaHour, trueRaMin, trueRaSec, trueDecDeg, trueDecMin, trueDecSec,
		coordinateType1, geogLongDeg, geogLatDeg, daylightSavingHours, timezoneHours,
		lcdDay, lcdMonth, lcdYear, lctHour, lctMin, lctSec,
		atmosphericPressureMbar, atmosphericTemperatureCelsius,
	)

	require.Equal(t, expectedRefraction, actualRefraction, "Mismatch in Corrected Refraction")
}

func TestCorrectionsForGeocentricParallax(t *testing.T) {
	var raHour float64 = 22
	var raMin float64 = 35
	var raSec float64 = 19
	var decDeg float64 = -7
	var decMin float64 = 41
	var decSec float64 = 13
	var coordinateType patype.CoordinateType = patype.CoordinateType_Actual
	var equatorialHorParallaxDeg float64 = 1.019167
	var geogLongDeg float64 = -100
	var geogLatDeg float64 = 50
	var heightM float64 = 60
	var daylightSaving int = 0
	var timezoneHours int = -6
	var lcdDay float64 = 26
	var lcdMonth int = 2
	var lcdYear int = 1979
	var lctHour float64 = 10
	var lctMin float64 = 45
	var lctSec float64 = 0

	var expectedParallax patype.CorrectedParallax = patype.CorrectedParallax{
		RightAscensionHours: 22, RightAscensionMinutes: 36, RightAscensionSeconds: 43.22,
		DeclinationDegrees: -8, DeclinationMinutes: 32, DeclinationSeconds: 17.4,
	}
	var actualParallax patype.CorrectedParallax = palib.CorrectionsForGeocentricParallax(
		raHour, raMin, raSec, decDeg, decMin, decSec,
		coordinateType, equatorialHorParallaxDeg, geogLongDeg, geogLatDeg, heightM, daylightSaving, timezoneHours,
		lcdDay, lcdMonth, lcdYear, lctHour, lctMin, lctSec,
	)

	require.Equal(t, expectedParallax, actualParallax, "Mismatch in Corrected Parallax")
}

func TestHeliographicCoordinates(t *testing.T) {
	var helioPositionAngleDeg float64 = 220
	var helioDisplacementArcmin float64 = 10.5
	var gwdateDay float64 = 1
	var gwDateMonth int = 5
	var gwdateYear int = 1988

	var expectedHeliographicCoordinates patype.HeliographicCoordinates = patype.HeliographicCoordinates{LongitudeDegrees: 142.59, LatitudeDegrees: -19.94}
	var actualHeliographicCoordinates patype.HeliographicCoordinates = palib.HeliographicCoordinates(
		helioPositionAngleDeg, helioDisplacementArcmin, gwdateDay, gwDateMonth, gwdateYear)

	require.Equal(t, expectedHeliographicCoordinates, actualHeliographicCoordinates, "Mismatch in Heliographic Coordinates")
}

func TestCarringtonRotationNumber(t *testing.T) {
	var gwdateDay float64 = 27
	var gwdateMonth int = 1
	var gwdateYear int = 1975

	var expectedCarringtonRotationNumber int = 1624
	var actualCarringtonRotationNumber int = palib.CarringtonRotationNumber(gwdateDay, gwdateMonth, gwdateYear)

	require.Equal(t, expectedCarringtonRotationNumber, actualCarringtonRotationNumber, "Mismatch in Carrington Rotation Number")
}

func TestSelenographicCoordinates1(t *testing.T) {
	var gwdateDay float64 = 1
	var gwdateMonth int = 5
	var gwdateYear int = 1988

	var expectedSelenegraphicCoordinates patype.SelenographicSubEarthCoordinates = patype.SelenographicSubEarthCoordinates{
		Longitude: -4.88, Latitude: 4.04, PositionAngleOfPole: 19.78}
	var actualSelenegraphicCoordinates patype.SelenographicSubEarthCoordinates = palib.SelenographicCoordinates1(gwdateDay, gwdateMonth, gwdateYear)

	require.Equal(t, expectedSelenegraphicCoordinates, actualSelenegraphicCoordinates, "Mismatch in Selenographic Subearth Coordinates")
}

func TestSelenographicCoordinates2(t *testing.T) {
	var gwdateDay float64 = 1
	var gwdateMonth int = 5
	var gwdateYear int = 1988

	var expectedSelenegraphicCoordinates patype.SelenographicSubSolarCoordinates = patype.SelenographicSubSolarCoordinates{
		Longitude: 6.81, CoLongitude: 83.19, Latitude: 1.19}
	var actualSelenegraphicCoordinates patype.SelenographicSubSolarCoordinates = palib.SelenographicCoordinates2(gwdateDay, gwdateMonth, gwdateYear)

	require.Equal(t, expectedSelenegraphicCoordinates, actualSelenegraphicCoordinates, "Mismatch in Selenographic Subsolar Coordinates")
}
