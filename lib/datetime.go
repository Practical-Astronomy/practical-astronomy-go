package lib

import (
	"math"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Determine the date of Easter for a given year. */
func GetDateOfEaster(inputYear int) patype.FullDate {
	var year float64 = float64(inputYear)

	var a float64 = float64(int(year) % 19)
	var b float64 = math.Floor(year / 100)
	var c float64 = float64(int(year) % 100)
	var d float64 = math.Floor(b / 4)
	var e float64 = float64(int(b) % 4)
	var f float64 = float64(math.Floor((b + 8) / 25))
	var g float64 = float64(math.Floor((b - f + 1) / 3))
	var h float64 = float64(int(((19.0 * a) + b - d - g + 15.0)) % 30)
	var i float64 = math.Floor(c / 4)
	var k float64 = float64(int(c) % 4)
	var l float64 = float64(int(32.0+2.0*(e+i)-h-k) % 7)
	var m float64 = math.Floor((a + (11.0 * h) + (22.0 * l)) / 451.0)
	var n float64 = math.Floor((h + l - (7.0 * m) + 114.0) / 31.0)
	var p float64 = float64(int(h+l-(7.0*m)+114.0) % 31)

	var day float64 = p + 1
	var month int = int(n)

	returnValue := patype.FullDate{Month: month, Day: day, Year: int(year)}

	return returnValue
}

/* Convert a civil date (month, day, and year) to a day number (the number of days from the beginning of the year upon which the civil date falls) */
func CivilDateToDayNumber(month int, day int, year int) int {
	if month <= 2 {
		month = month - 1
		if pautil.IsLeapYear(year) {
			month = month * 62
		} else {
			month = month * 63
		}
		month = int(math.Floor(float64(month) / 2))
	} else {
		month = int(math.Floor((float64(month) + 1) * 30.6))
		if pautil.IsLeapYear(year) {
			month = month - 62
		} else {
			month = month - 63
		}
	}

	return month + day
}

/*  Convert civil time (HH:MM:SS) to decimal hours (HH.########) */
func CivilTimeToDecimalHours(hours float64, minutes float64, seconds float64) float64 {
	return pamacro.HmsDh(hours, minutes, seconds)
}

/* Convert decimal hours (HH.########) to civil time (HH:MM:SS) */
func DecimalHoursToCivilTime(decimalHours float64) patype.FullTime {
	var hours int = pamacro.DecimalHoursHour(decimalHours)
	var minutes int = pamacro.DecimalHoursMinute(decimalHours)
	var seconds float64 = pamacro.DecimalHoursSecond(decimalHours)

	var returnValue patype.FullTime = patype.FullTime{Hours: hours, Minutes: minutes, Seconds: seconds}

	return returnValue
}

/**
 * Convert local civil time to universal time.
 */
func LocalCivilTimeToUniversalTime(
	lctHours float64, lctMinutes float64, lctSeconds float64,
	isDaylightSavings bool, zoneCorrection int, localDay float64, localMonth int, localYear int,
) patype.FullDateTime {
	var lct float64 = CivilTimeToDecimalHours(lctHours, lctMinutes, lctSeconds)

	var daylight_savings_offset int
	if isDaylightSavings {
		daylight_savings_offset = 1
	} else {
		daylight_savings_offset = 0
	}

	var ut_interim float64 = lct - float64(daylight_savings_offset) - float64(zoneCorrection)
	var g_day_interim float64 = localDay + (ut_interim / 24)

	var jd float64 = pamacro.CivilDateToJulianDate(g_day_interim, float64(localMonth), float64(localYear))

	var g_day float64 = pamacro.JulianDateDay(jd)
	var g_month int = pamacro.JulianDateMonth(jd)
	var g_year int = pamacro.JulianDateYear(jd)

	var ut float64 = 24 * (g_day - math.Floor(g_day))

	return patype.FullDateTime{
		Month:   g_month,
		Day:     int(math.Floor(g_day)),
		Year:    g_year,
		Hours:   pamacro.DecimalHoursHour(ut),
		Minutes: pamacro.DecimalHoursMinute(ut),
		Seconds: pamacro.DecimalHoursSecond(ut),
	}
}

/* Convert universal time to local civil time. */
func UniversalTimeToLocalCivilTime(utHours float64, utMinutes float64, utSeconds float64, isDaylightSavings bool, zoneCorrection int,
	gwDay int, gwMonth int, gwYear int) patype.FullDateTime {
	var dstValue int
	if isDaylightSavings {
		dstValue = 1
	} else {
		dstValue = 0
	}

	var ut float64 = pamacro.HmsDh(utHours, utMinutes, utSeconds)
	var zoneTime float64 = ut + float64(zoneCorrection)
	var localTime float64 = zoneTime + float64(dstValue)
	var localJdPlusLocalTime float64 = pamacro.CivilDateToJulianDate(float64(gwDay), float64(gwMonth), float64(gwYear)) + (localTime / 24)
	var localDay float64 = pamacro.JulianDateDay(localJdPlusLocalTime)
	var integerDay float64 = math.Floor(localDay)
	var localMonth int = pamacro.JulianDateMonth(localJdPlusLocalTime)
	var localYear int = pamacro.JulianDateYear(localJdPlusLocalTime)

	var lct float64 = 24 * (localDay - integerDay)

	return patype.FullDateTime{
		Month:   localMonth,
		Day:     int(integerDay),
		Year:    localYear,
		Hours:   pamacro.DecimalHoursHour(lct),
		Minutes: pamacro.DecimalHoursMinute(lct),
		Seconds: pamacro.DecimalHoursSecond(lct),
	}
}

/* Convert Universal Time to Greenwich Sidereal Time */
func UniversalTimeToGreenwichSiderealTime(utHours float64, utMinutes float64, utSeconds float64, gwDay float64, gwMonth int, gwYear int) patype.FullTime {
	var jd float64 = pamacro.CivilDateToJulianDate(gwDay, float64(gwMonth), float64(gwYear))
	var s float64 = jd - 2451545.0
	var t float64 = s / 36525.0
	var t01 float64 = 6.697374558 + (2400.051336 * t) + (0.000025862 * t * t)
	var t02 float64 = t01 - (24.0 * math.Floor(t01/24.0))
	var ut float64 = pamacro.HmsDh(utHours, utMinutes, utSeconds)
	var a float64 = ut * 1.002737909
	var gst1 float64 = t02 + a
	var gst2 float64 = gst1 - (24.0 * math.Floor(gst1/24.0))

	var gstHours int = pamacro.DecimalHoursHour(gst2)
	var gstMinutes int = pamacro.DecimalHoursMinute(gst2)
	var gstSeconds float64 = pamacro.DecimalHoursSecond(gst2)

	return patype.FullTime{Hours: gstHours, Minutes: gstMinutes, Seconds: gstSeconds}
}

/* Convert Greenwich Sidereal Time to Universal Time */
func GreenwichSiderealTimeToUniversalTime(
	gstHours float64, gstMinutes float64, gstSeconds float64, gwDay float64, gwMonth int, gwYear int) patype.FullTimeWithWarning {
	var jd float64 = pamacro.CivilDateToJulianDate(gwDay, float64(gwMonth), float64(gwYear))
	var s float64 = jd - 2451545
	var t float64 = s / 36525
	var t01 float64 = 6.697374558 + (2400.051336 * t) + (0.000025862 * t * t)
	var t02 float64 = t01 - (24 * math.Floor(t01/24))
	var gstHours1 float64 = pamacro.HmsDh(gstHours, gstMinutes, gstSeconds)

	var a float64 = gstHours1 - t02
	var b float64 = a - (24 * math.Floor(a/24))
	var ut float64 = b * 0.9972695663
	var utHours int = pamacro.DecimalHoursHour(ut)
	var utMinutes int = pamacro.DecimalHoursMinute(ut)
	var utSeconds float64 = pamacro.DecimalHoursSecond(ut)

	var warningFlag patype.WarningFlags
	if ut < 0.065574 {
		warningFlag = patype.WarningFlag_Warning
	} else {
		warningFlag = patype.WarningFlag_OK
	}

	return patype.FullTimeWithWarning{Hours: utHours, Minutes: utMinutes, Seconds: utSeconds, WarningFlag: warningFlag}
}

/* Convert Greenwich Sidereal Time to Local Sidereal Time */
func GreenwichSiderealTimeToLocalSiderealTime(gstHours float64, gstMinutes float64, gstSeconds float64, geographicalLongitude float64) patype.FullTime {
	var gst float64 = pamacro.HmsDh(gstHours, gstMinutes, gstSeconds)
	var offset float64 = geographicalLongitude / 15
	var lstHours1 float64 = gst + offset
	var lstHours2 float64 = lstHours1 - (24 * math.Floor(lstHours1/24))

	var lstHours int = pamacro.DecimalHoursHour(lstHours2)
	var lstMinutes int = pamacro.DecimalHoursMinute(lstHours2)
	var lstSeconds float64 = pamacro.DecimalHoursSecond(lstHours2)

	return patype.FullTime{Hours: lstHours, Minutes: lstMinutes, Seconds: lstSeconds}
}

/* Convert Local Sidereal Time to Greenwich Sidereal Time */
func LocalSiderealTimeToGreenwichSiderealTime(lstHours float64, lstMinutes float64, lstSeconds float64, geographicalLongitude float64) patype.FullTime {
	var gst float64 = pamacro.HmsDh(lstHours, lstMinutes, lstSeconds)
	var longHours float64 = geographicalLongitude / 15
	var gst1 float64 = gst - longHours
	var gst2 float64 = gst1 - (24 * math.Floor(gst1/24))

	var gstHours int = pamacro.DecimalHoursHour(gst2)
	var gstMinutes int = pamacro.DecimalHoursMinute(gst2)
	var gstSeconds float64 = pamacro.DecimalHoursSecond(gst2)

	return patype.FullTime{Hours: gstHours, Minutes: gstMinutes, Seconds: gstSeconds}
}
