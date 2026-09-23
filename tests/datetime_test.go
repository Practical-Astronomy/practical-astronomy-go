package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDateOfEaster(t *testing.T) {
	var year int = 2026

	var expectedDateOfEaster patype.FullDate = patype.FullDate{Month: 4, Day: 5, Year: year}
	var actualDateOfEaster patype.FullDate = palib.GetDateOfEaster(year)

	require.Equal(t, expectedDateOfEaster, actualDateOfEaster, "Mismatch in Date of Easter")
}

func TestCivilDateToDayNumber(t *testing.T) {
	var dayNumber int

	dayNumber = palib.CivilDateToDayNumber(1, 1, 2000)
	require.Equal(t, 1, dayNumber, "Day Number for 1/1/2000")

	dayNumber = palib.CivilDateToDayNumber(3, 1, 2000)
	require.Equal(t, 61, dayNumber, "Day Number for 3/1/2000")

	dayNumber = palib.CivilDateToDayNumber(6, 1, 2003)
	require.Equal(t, 152, dayNumber, "Day Number for 6/1/2003")

	dayNumber = palib.CivilDateToDayNumber(11, 27, 2009)
	require.Equal(t, 331, dayNumber, "Day Number for 11/27/2009")
}

func TestCivilTimeToDecimalHours(t *testing.T) {
	var inputHours float64 = 18
	var inputMinutes float64 = 31
	var inputSeconds float64 = 27

	var expectedDecimalHours float64 = 18.52416667
	var actualDecimalHours float64 = palib.CivilTimeToDecimalHours(inputHours, inputMinutes, inputSeconds)

	require.Equal(t, expectedDecimalHours, pautil.RoundTo(actualDecimalHours, 8), "Mismatch in Decimal Hours")
}

func TestDecimalHoursToCivilTime(t *testing.T) {
	var inputDecimalHours float64 = 18.52416667

	var expectedCivilTime patype.FullTime = patype.FullTime{Hours: 18, Minutes: 31, Seconds: 27}
	var actualCivilTime patype.FullTime = palib.DecimalHoursToCivilTime(inputDecimalHours)

	require.Equal(t, expectedCivilTime, actualCivilTime, "Mismatch in Civil Time")
}

func TestLocalCivilTimeToUniversalTime(t *testing.T) {
	var lctHours float64 = 3
	var lctMinutes float64 = 37
	var lctSeconds float64 = 0
	var isDaylightSavings bool = true
	var zoneCorrection int = 4
	var localDay float64 = 1
	var localMonth int = 7
	var localYear int = 2013

	var expectedUniversalTime patype.FullDateTime = patype.FullDateTime{Month: 6, Day: 30, Year: 2013, Hours: 22, Minutes: 37, Seconds: 0}
	var actualUniversalTime patype.FullDateTime = palib.LocalCivilTimeToUniversalTime(
		lctHours, lctMinutes, lctSeconds, isDaylightSavings, zoneCorrection, localDay, localMonth, localYear,
	)

	require.Equal(t, expectedUniversalTime, actualUniversalTime, "Mismatch in Universal Time")
}

func TestUniversalTimeToLocalCivilTime(t *testing.T) {
	var utHours float64 = 22
	var utMinutes float64 = 37
	var utSeconds float64 = 0
	var isDaylightSavings bool = true
	var zoneCorrection int = 4
	var gwDay int = 30
	var gwMonth int = 6
	var gwYear int = 2013

	var expectedLocalCivilTime patype.FullDateTime = patype.FullDateTime{Month: 7, Day: 1, Year: 2013, Hours: 3, Minutes: 37, Seconds: 0}
	var actualLocalCivilTime patype.FullDateTime = palib.UniversalTimeToLocalCivilTime(
		utHours, utMinutes, utSeconds, isDaylightSavings, zoneCorrection, gwDay, gwMonth, gwYear,
	)

	require.Equal(t, expectedLocalCivilTime, actualLocalCivilTime, "Mismatch in Local Civil Time")
}

func TestUniversalTimeToGreenwichSiderealTime(t *testing.T) {
	var utHours float64 = 14
	var utMinutes float64 = 36
	var utSeconds float64 = 51.67
	var gwDay float64 = 22
	var gwMonth int = 4
	var gwYear int = 1980

	var expectedGreenwichSiderealTime patype.FullTime = patype.FullTime{Hours: 4, Minutes: 40, Seconds: 5.23}
	var actualGreenwichSiderealTime patype.FullTime = palib.UniversalTimeToGreenwichSiderealTime(utHours, utMinutes, utSeconds, gwDay, gwMonth, gwYear)

	require.Equal(t, expectedGreenwichSiderealTime, actualGreenwichSiderealTime, "Mismatch in Greenwich Sidereal Time")
}
func TestGreenwichSiderealTimeToUniversalTime(t *testing.T) {
	var gstHours float64 = 4
	var gstMinutes float64 = 40
	var gstSeconds float64 = 5.23
	var gwDay float64 = 22
	var gwMonth int = 4
	var gwYear int = 1980

	var expectedUniversalTime patype.FullTimeWithWarning = patype.FullTimeWithWarning{Hours: 14, Minutes: 36, Seconds: 51.67, WarningFlag: patype.WarningFlag_OK}
	var actualUniversalTime patype.FullTimeWithWarning = palib.GreenwichSiderealTimeToUniversalTime(gstHours, gstMinutes, gstSeconds, gwDay, gwMonth, gwYear)

	require.Equal(t, expectedUniversalTime, actualUniversalTime, "Mismatch in Universal Time")
}
func TestGreenwichSiderealTimeToLocalSiderealTime(t *testing.T) {
	var gstHours float64 = 4
	var gstMinutes float64 = 40
	var gstSeconds float64 = 5.23
	var geographicalLongitude float64 = -64

	var expectedLocalSiderealTime patype.FullTime = patype.FullTime{Hours: 0, Minutes: 24, Seconds: 5.23}
	var actualLocalSiderealTime patype.FullTime = palib.GreenwichSiderealTimeToLocalSiderealTime(gstHours, gstMinutes, gstSeconds, geographicalLongitude)

	require.Equal(t, expectedLocalSiderealTime, actualLocalSiderealTime, "Mismatch in Local Sidereal Time")
}
func TestLocalSiderealTimeToGreenwichSiderealTime(t *testing.T) {
	var lstHours float64 = 0
	var lstMinutes float64 = 24
	var lstSeconds float64 = 5.23
	var geographicalLongitude float64 = -64

	var expectedGreenwichSiderealTime patype.FullTime = patype.FullTime{Hours: 4, Minutes: 40, Seconds: 5.23}
	var actualGreenwichSiderealTime patype.FullTime = palib.LocalSiderealTimeToGreenwichSiderealTime(lstHours, lstMinutes, lstSeconds, geographicalLongitude)

	require.Equal(t, expectedGreenwichSiderealTime, actualGreenwichSiderealTime, "Mismatch in Greenwich Sidereal Time")
}
