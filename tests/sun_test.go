package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApproximatePositionOfSun(t *testing.T) {
	var lctHours float64 = 0
	var lctMinutes float64 = 0
	var lctSeconds float64 = 0
	var localDay float64 = 27
	var localMonth int = 7
	var localYear int = 2003
	var isDaylightSaving bool = false
	var zoneCorrection int = 0

	var expectedApproximatePositionofSun patype.SunPosition = patype.SunPosition{
		RightAscensionHour: 8, RightAscensionMinutes: 23, RightAscensionSeconds: 33.73,
		DeclinationDegrees: 19, DeclinationMinutes: 21, DeclinationSeconds: 14.33,
	}
	var actualApproximatePositionofSun patype.SunPosition = palib.ApproximatePositionOfSun(
		lctHours, lctMinutes, lctSeconds, localDay, localMonth, localYear, isDaylightSaving, zoneCorrection,
	)

	require.Equal(t, expectedApproximatePositionofSun, actualApproximatePositionofSun, "Mismatch in Approximate Position of Sun")
}

func TestPrecisePositionOfSun(t *testing.T) {
	var lctHours float64 = 0
	var lctMinutes float64 = 0
	var lctSeconds float64 = 0
	var localDay float64 = 27
	var localMonth int = 7
	var localYear int = 1988
	var isDaylightSaving bool = false
	var zoneCorrection int = 0

	var expectedPrecisePositionofSun patype.SunPosition = patype.SunPosition{
		RightAscensionHour: 8, RightAscensionMinutes: 26, RightAscensionSeconds: 3.83,
		DeclinationDegrees: 19, DeclinationMinutes: 12, DeclinationSeconds: 49.72,
	}
	var actualPrecisePositionofSun patype.SunPosition = palib.PrecisePositionOfSun(
		lctHours, lctMinutes, lctSeconds, localDay, localMonth, localYear, isDaylightSaving, zoneCorrection,
	)

	require.Equal(t, expectedPrecisePositionofSun, actualPrecisePositionofSun, "Mismatch in Precise Position of Sun")
}

func TestSunDistanceAndAngularSize(t *testing.T) {
	var lctHours float64 = 0
	var lctMinutes float64 = 0
	var lctSeconds float64 = 0
	var localDay float64 = 27
	var localMonth int = 7
	var localYear int = 1988
	var isDaylightSaving bool = false
	var zoneCorrection int = 0

	var expectedSunDistanceAndAngularSize patype.SunDistanceSize = patype.SunDistanceSize{
		DistanceInKilometers: 151920130, AngularSize_Degrees: 0, AngularSize_Minutes: 31, AngularSize_Seconds: 29.93}
	var actualSunDistanceAndAngularSize patype.SunDistanceSize = palib.SunDistanceAndAngularSize(
		lctHours, lctMinutes, lctSeconds, localDay, localMonth, localYear, isDaylightSaving, zoneCorrection)

	require.Equal(t, expectedSunDistanceAndAngularSize, actualSunDistanceAndAngularSize, "Mismatch in Sun Distance and Angular Size")
}

func TestSunriseAndSunset(t *testing.T) {
	var localDay float64 = 10
	var localMonth int = 3
	var localYear int = 1986
	var isDaylightSaving bool = false
	var zoneCorrection int = -5
	var geographicalLongDeg float64 = -71.05
	var geographicalLatDeg float64 = 42.37

	var expectedSunriseAndSunset patype.SunriseSunsetInfo = patype.SunriseSunsetInfo{
		LocalSunriseHour: 6, LocalSunriseMinute: 5, LocalSunsetHour: 17, LocalSunsetMinute: 45,
		AzimuthOfSunriseDeg: 94.83, AzimuthOfSunsetDeg: 265.43, Status: patype.RiseSetStatus_OK,
	}
	var actualSunriseAndSunset patype.SunriseSunsetInfo = palib.SunriseAndSunset(
		localDay, localMonth, localYear, isDaylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg)

	require.Equal(t, expectedSunriseAndSunset, actualSunriseAndSunset, "Mismatch in Sunrise/Sunset")
}

func TestMorningAndEveningTwilight(t *testing.T) {
	var localDay float64 = 7
	var localMonth int = 9
	var localYear int = 1979
	var isDaylightSaving bool = false
	var zoneCorrection int = 0
	var geographicalLongDeg float64 = 0
	var geographicalLatDeg float64 = 52
	var twilightType patype.TwilightType = patype.TwilightType_ASTRONOMICAL

	var expectedTwilightInfo patype.TwilightInfo = patype.TwilightInfo{
		AmTwilightBeginsHour: 3, AmTwilightBeginsMin: 17, PmTwilightEndsHour: 20, PmTwilightEndsMin: 37, Status: patype.TwilightStatus_OK}
	var actualTwilightInfo patype.TwilightInfo = palib.MorningAndEveningTwilight(
		localDay, localMonth, localYear, isDaylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg, twilightType)

	require.Equal(t, expectedTwilightInfo, actualTwilightInfo, "Mismatch in Twilight Info")
}

func TestEquationOfTime(t *testing.T) {
	var gwDateDay float64 = 27
	var gwDateMonth int = 7
	var gwDateYear int = 2010

	var expectedEquationOfTime patype.EquationOfTime = patype.EquationOfTime{Minutes: 6, Seconds: 31.52}
	var actualEquationOfTime patype.EquationOfTime = palib.EquationOfTime(gwDateDay, gwDateMonth, gwDateYear)

	require.Equal(t, expectedEquationOfTime, actualEquationOfTime, "Mismatch in Equation of Time")
}

func TestSolarElongation(t *testing.T) {
	var raHour float64 = 10
	var raMin float64 = 6
	var raSec float64 = 45
	var decDeg float64 = 11
	var decMin float64 = 57
	var decSec float64 = 27
	var gwDateDay float64 = 27.8333333
	var gwDateMonth int = 7
	var gwDateYear int = 2010

	var expectedSolarElongation float64 = 24.78
	var actualSolarElongation float64 = palib.SolarElongation(raHour, raMin, raSec, decDeg, decMin, decSec, gwDateDay, gwDateMonth, gwDateYear)

	require.Equal(t, expectedSolarElongation, actualSolarElongation, "Mismatch in Solar Elongation")
}
