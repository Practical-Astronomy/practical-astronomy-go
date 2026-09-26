package macros

import (
	"math"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/*
Convert a Civil Time (hours,minutes,seconds) to Decimal Hours.

Original macro name: HMSDH
*/
func HmsDh(hours float64, minutes float64, seconds float64) float64 {
	var fHours float64 = hours
	var fMinutes float64 = minutes
	var fSeconds float64 = seconds

	var a float64 = math.Abs(fSeconds) / 60.0
	var b float64 = (math.Abs(fMinutes) + a) / 60.0
	var c float64 = math.Abs(fHours) + b

	if fHours < 0 || fMinutes < 0 || fSeconds < 0 {
		return -c
	} else {
		return c
	}
}

/* Extract hour part of decimal hours. */
func DecimalHoursHour(decimalHours float64) int {
	var a float64 = math.Abs(decimalHours)
	var b float64 = a * 3600
	var c float64 = pautil.RoundTo(b-60*math.Floor(b/60), 2)

	var e float64
	if c == 60 {
		e = b + 60
	} else {
		e = b
	}

	if decimalHours < 0 {
		return int(-(math.Floor(e / 3600)))
	} else {
		return int(math.Floor(e / 3600))
	}
}

/* Extract minutes part of decimal hours. */
func DecimalHoursMinute(decimalHours float64) int {
	var a float64 = math.Abs(decimalHours)
	var b float64 = a * 3600
	var c float64 = pautil.RoundTo(b-60*math.Floor(b/60), 2)

	var e float64
	if c == 60 {
		e = b + 60
	} else {
		e = b
	}

	return int(math.Floor(e/60)) % 60
}

/* Extract seconds part of decimal hours. */
func DecimalHoursSecond(decimalHours float64) float64 {
	var a float64 = math.Abs(decimalHours)
	var b float64 = a * 3600
	var c float64 = pautil.RoundTo(b-60*math.Floor(b/60), 2)

	var d float64
	if c == 60 {
		d = 0
	} else {
		d = c
	}

	return d
}

/*
Convert a Greenwich Date/Civil Date (day,month,year) to Julian Date

Original macro name: CDJD
*/
func CivilDateToJulianDate(day float64, month float64, year float64) float64 {
	var fDay float64 = day
	var fMonth float64 = month
	var fYear float64 = year

	var y float64
	if fMonth < 3 {
		y = fYear - 1
	} else {
		y = fYear
	}

	var m float64
	if fMonth < 3 {
		m = fMonth + 12
	} else {
		m = fMonth
	}

	var b float64
	if fYear > 1582 {
		var a float64 = math.Floor(y / 100)
		b = 2 - a + math.Floor(a/4)
	} else {
		if fYear == 1582 && fMonth > 10 {
			var a float64 = math.Floor(y / 100)
			b = 2 - a + math.Floor(a/4)
		} else {
			if fYear == 1582 && fMonth == 10 && fDay >= 15 {
				var a float64 = math.Floor(y / 100)
				b = 2 - a + math.Floor(a/4)
			} else {
				b = 0
			}
		}
	}

	var c float64
	if y < 0 {
		c = math.Floor(((365.25 * y) - 0.75))
	} else {
		c = math.Floor(365.25 * y)
	}

	var d float64 = math.Floor(30.6001 * (m + 1.0))

	return b + c + d + fDay + 1720994.5
}

/*
Returns the day part of a Julian Date

Original macro name: JDCDay
*/
func JulianDateDay(julianDate float64) float64 {
	var i float64 = math.Floor(julianDate + 0.5)
	var f float64 = julianDate + 0.5 - i
	var a float64 = math.Floor((i - 1867216.25) / 36524.25)

	var b float64
	if i > 2299160 {
		b = i + 1 + a - math.Floor(a/4)
	} else {
		b = i
	}

	var c float64 = b + 1524
	var d float64 = math.Floor((c - 122.1) / 365.25)
	var e float64 = math.Floor(365.25 * d)
	var g float64 = math.Floor((c - e) / 30.6001)

	return c - e + f - math.Floor(30.6001*g)
}

/*
Returns the month part of a Julian Date

Original macro name: JDCMonth
*/
func JulianDateMonth(julianDate float64) int {
	var i float64 = math.Floor(julianDate + 0.5)
	var a float64 = math.Floor((i - 1867216.25) / 36524.25)

	var b float64
	if i > 2299160 {
		b = i + 1 + a - math.Floor(a/4)
	} else {
		b = i
	}

	var c float64 = b + 1524
	var d float64 = math.Floor((c - 122.1) / 365.25)
	var e float64 = math.Floor(365.25 * d)
	var g float64 = math.Floor((c - e) / 30.6001)

	var returnValue float64
	if g < 13.5 {
		returnValue = g - 1
	} else {
		returnValue = g - 13
	}

	return int(returnValue)
}

/*
Returns the year part of a Julian Date

Original macro name: JDCYear
*/
func JulianDateYear(julianDate float64) int {
	var i float64 = math.Floor(julianDate + 0.5)
	var a float64 = math.Floor((i - 1867216.25) / 36524.25)

	var b float64
	if i > 2299160 {
		b = i + 1.0 + a - math.Floor(a/4.0)
	} else {
		b = i
	}

	var c float64 = b + 1524
	var d float64 = math.Floor((c - 122.1) / 365.25)
	var e float64 = math.Floor(365.25 * d)
	var g float64 = math.Floor((c - e) / 30.6001)

	var h float64
	if g < 13.5 {
		h = g - 1
	} else {
		h = g - 13
	}

	var returnValue float64
	if h > 2.5 {
		returnValue = d - 4716
	} else {
		returnValue = d - 4715
	}

	return int(returnValue)
}

/*
Convert Right Ascension to Hour Angle

Original macro name: RAHA
*/
func RightAscensionToHourAngle(
	raHours float64, raMinutes float64, raSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64,
	daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64,
) float64 {
	var a float64 = LocalCivilTimeToUniversalTime(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var b float64 = LocalCivilTimeGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var c int = int(LocalCivilTimeGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var d int = int(LocalCivilTimeGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var e float64 = UniversalTimeToGreenwichSiderealTime(a, 0, 0, b, c, d)
	var f float64 = GreenwichSiderealTimeToLocalSiderealTime(e, 0, 0, geographicalLongitude)
	var g float64 = HmsDh(raHours, raMinutes, raSeconds)
	var h float64 = f - g

	if h < 0 {
		return 24 + h
	} else {
		return h
	}
}

/*
Convert Hour Angle to Right Ascension

Original macro name: HARA
*/
func HourAngleToRightAscension(
	hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64,
	daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localyear int, geographicalLongitude float64,
) float64 {
	var a float64 = LocalCivilTimeToUniversalTime(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localyear)
	var b float64 = LocalCivilTimeGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localyear)
	var c int = int(LocalCivilTimeGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localyear))
	var d int = int(LocalCivilTimeGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localyear))
	var e float64 = UniversalTimeToGreenwichSiderealTime(a, 0, 0, b, c, d)
	var f float64 = GreenwichSiderealTimeToLocalSiderealTime(e, 0, 00, geographicalLongitude)
	var g float64 = HmsDh(hourAngleHours, hourAngleMinutes, hourAngleSeconds)
	var h float64 = f - g

	if h < 0 {
		return 24 + h
	} else {
		return h
	}
}

/*
Convert Local Civil Time to Universal Time

Original macro name: LctUT
*/
func LocalCivilTimeToUniversalTime(
	lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int,
) float64 {
	var a float64 = HmsDh(lctHours, lctMinutes, lctSeconds)
	var b float64 = a - float64(daylightSaving) - float64(zoneCorrection)
	var c float64 = localDay + (b / 24)
	var d float64 = CivilDateToJulianDate(c, float64(localMonth), float64(localYear))
	var e float64 = JulianDateDay(d)
	var e1 float64 = math.Floor(e)

	return 24 * (e - e1)
}

/*
Convert Universal Time to Local Civil Time

Original macro name: UTLct
*/
func UniversalTimeToLocalCivilTime(uHours float64, uMinutes float64, uSeconds float64, daylightSaving int, zoneCorrection int, greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	var a float64 = HmsDh(uHours, uMinutes, uSeconds)
	var b float64 = a + float64(zoneCorrection)
	var c float64 = b + float64(daylightSaving)
	var d float64 = CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear)) + (c / 24)
	var e float64 = JulianDateDay(d)
	var e1 float64 = math.Floor(e)

	return 24 * (e - e1)
}

/*
Determine Greenwich Day for Local Time

Original macro name: LctGDay
*/
func LocalCivilTimeGreenwichDay(
	lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int,
) float64 {
	var a float64 = HmsDh(lctHours, lctMinutes, lctSeconds)
	var b float64 = a - float64(daylightSaving) - float64(zoneCorrection)
	var c float64 = localDay + (b / 24)
	var d float64 = CivilDateToJulianDate(c, float64(localMonth), float64(localYear))
	var e float64 = JulianDateDay(d)

	return math.Floor(e)
}

/*
Determine Greenwich Month for Local Time

Original macro name: LctGMonth
*/
func LocalCivilTimeGreenwichMonth(
	lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int,
) float64 {
	var a float64 = HmsDh(lctHours, lctMinutes, lctSeconds)
	var b float64 = a - float64(daylightSaving) - float64(zoneCorrection)
	var c float64 = localDay + (b / 24)
	var d float64 = CivilDateToJulianDate(c, float64(localMonth), float64(localYear))

	return float64(JulianDateMonth(d))
}

/*
Determine Greenwich Year for Local Time

Original macro name: LctGYear
*/
func LocalCivilTimeGreenwichYear(
	lctHours float64, lctMinutes float64, lctSeconds float64, daylightSaving int, zoneCorrection int, localDay float64, localMonth int, localYear int,
) float64 {
	var a float64 = HmsDh(lctHours, lctMinutes, lctSeconds)
	var b float64 = a - float64(daylightSaving) - float64(zoneCorrection)
	var c float64 = localDay + (b / 24)
	var d float64 = CivilDateToJulianDate(c, float64(localMonth), float64(localYear))

	return float64(JulianDateYear(d))
}

/*
*
Convert Universal Time to Greenwich Sidereal Time

Original macro name: UTGST
*/
func UniversalTimeToGreenwichSiderealTime(
	uHours float64, iMinutes float64, uSeconds float64, greenwichDay float64, greenwichMonth int, greenwichYear int,
) float64 {
	var a float64 = CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear))
	var b float64 = a - 2451545
	var c float64 = b / 36525
	var d float64 = 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c)
	var e float64 = d - (24 * math.Floor(d/24))
	var f float64 = HmsDh(uHours, iMinutes, uSeconds)
	var g float64 = f * 1.002737909
	var h float64 = e + g

	return h - (24 * math.Floor(h/24))
}

/*
Convert Greenwich Sidereal Time to Universal Time

Original macro name: GSTUT
*/
func GreenwichSiderealTimeToUniversalTime(
	greenwichSiderealHours float64, greenwichSiderealMinutes float64, greenwichSiderealSeconds float64,
	greenwichDay float64, greenwichMonth int, greenwichYear int,
) float64 {
	var a float64 = CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear))
	var b float64 = a - 2451545
	var c float64 = b / 36525
	var d float64 = 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c)
	var e float64 = d - (24 * math.Floor(d/24))
	var f float64 = HmsDh(greenwichSiderealHours, greenwichSiderealMinutes, greenwichSiderealSeconds)
	var g float64 = f - e
	var h float64 = g - (24 * math.Floor(g/24))

	return h * 0.9972695663
}

/*
Convert Greenwich Sidereal Time to Local Sidereal Time

Original macro name: GSTLST
*/
func GreenwichSiderealTimeToLocalSiderealTime(
	greenwichHours float64, greenwichMinutes float64, greenwichSeconds float64, geographicalLongitude float64,
) float64 {
	var a float64 = HmsDh(greenwichHours, greenwichMinutes, greenwichSeconds)
	var b float64 = geographicalLongitude / 15
	var c float64 = a + b

	return c - (24 * math.Floor(c/24))
}

/*
Convert Equatorial Coordinates to Azimuth (in decimal degrees)

Original macro name: EQAz
*/
func EquatorialCoordinatesToAzimuth(
	hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64,
	declinationDegrees float64, declinationMinutes float64, declinationSeconds float64,
	geographicalLatitude float64,
) float64 {
	var a float64 = HmsDh(hourAngleHours, hourAngleMinutes, hourAngleSeconds)
	var b float64 = a * 15
	var c float64 = pautil.DegreesToRadians(b)
	var d float64 = DegreesMinutesSecondsToDecimalDegrees(declinationDegrees, declinationMinutes, declinationSeconds)
	var e float64 = pautil.DegreesToRadians(d)
	var f float64 = pautil.DegreesToRadians(geographicalLatitude)
	var g float64 = math.Sin(e)*math.Sin(f) + math.Cos(e)*math.Cos(f)*math.Cos(c)
	var h float64 = -math.Cos(e) * math.Cos(f) * math.Sin(c)
	var i float64 = math.Sin(e) - (math.Sin(f) * g)
	var j float64 = Degrees(math.Atan2(h, i))

	return j - 360.0*math.Floor(j/360)
}

/*
Convert Equatorial Coordinates to Altitude (in decimal degrees)

Original macro name: EQAlt
*/
func EquatorialCoordinatestoAltitude(
	hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64,
	declinationDegrees float64, declinationMinutes float64, declinationSeconds float64,
	geographicalLatitude float64,
) float64 {
	var a float64 = HmsDh(hourAngleHours, hourAngleMinutes, hourAngleSeconds)
	var b float64 = a * 15
	var c float64 = pautil.DegreesToRadians(b)
	var d float64 = DegreesMinutesSecondsToDecimalDegrees(declinationDegrees, declinationMinutes, declinationSeconds)
	var e float64 = pautil.DegreesToRadians(d)
	var f float64 = pautil.DegreesToRadians(geographicalLatitude)
	var g float64 = math.Sin(e)*math.Sin(f) + math.Cos(e)*math.Cos(f)*math.Cos(c)

	return Degrees(math.Asin(g))
}

/*
Convert Degrees Minutes Seconds to Decimal Degrees

Original macro name: DMSDD
*/
func DegreesMinutesSecondsToDecimalDegrees(degrees float64, minutes float64, seconds float64) float64 {
	var a float64 = math.Abs(seconds) / 60
	var b float64 = (math.Abs(minutes) + a) / 60
	var c float64 = math.Abs(degrees) + b

	if degrees < 0 || minutes < 0 || seconds < 0 {
		return -c
	} else {
		return c
	}
}

/*
Convert W to Degrees

Original macro name: Degrees
*/
func Degrees(w float64) float64 {
	return w * 57.29577951
}

/*
Return Degrees part of Decimal Degrees

Original macro name: DDDeg
*/
func DecimalDegreesDegrees(decimalDegrees float64) float64 {
	var a float64 = math.Abs(decimalDegrees)
	var b float64 = a * 3600
	var c float64 = pautil.RoundTo(b-60*math.Floor(b/60), 2)
	var e float64
	if c == 60 {
		e = 60
	} else {
		e = b
	}

	if decimalDegrees < 0 {
		return -(math.Floor(e / 3600))
	} else {
		return math.Floor(e / 3600)
	}
}

/*
Return Minutes part of Decimal Degrees

Original macro name: DDMin
*/
func DecimalDegreesMinutes(decimalDegrees float64) float64 {
	var a float64 = math.Abs(decimalDegrees)
	var b float64 = a * 3600
	var c float64 = pautil.RoundTo(b-60*math.Floor(b/60), 2)
	var e float64
	if c == 60 {
		e = b + 60
	} else {
		e = b
	}

	return float64(int(math.Floor(e/60)) % 60)
}

/*
Return Seconds part of Decimal Degrees

Original macro name: DDSec
*/
func DecimalDegreesSeconds(decimalDegrees float64) float64 {
	var a float64 = math.Abs(decimalDegrees)
	var b float64 = a * 3600
	var c float64 = pautil.RoundTo(b-60*math.Floor(b/60), 2)
	var d float64
	if c == 60 {
		d = 0
	} else {
		d = c
	}

	return d
}

/*
Convert Decimal Degrees to Degree-Hours

Original macro name: DDDH
*/
func DecimalDegreesToDegreeHours(decimal_degrees float64) float64 {
	return decimal_degrees / 15
}

/*
Convert Degree-Hours to Decimal Degrees

Original macro name: DHDD
*/
func DegreeHoursToDecimalDegrees(degree_hours float64) float64 {
	return degree_hours * 15
}

/*
Convert Horizon Coordinates to Declination (in decimal degrees)

Original macro name: HORDec
*/
func HorizonCoordinatesToDeclination(
	azimuthDegrees float64, azimuthMinutes float64, azimuthSeconds float64,
	altitudeDegrees float64, altitudeMinutes float64, altitudeSeconds float64,
	geographicalLatitude float64,
) float64 {
	var a float64 = DegreesMinutesSecondsToDecimalDegrees(azimuthDegrees, azimuthMinutes, azimuthSeconds)
	var b float64 = DegreesMinutesSecondsToDecimalDegrees(altitudeDegrees, altitudeMinutes, altitudeSeconds)
	var c float64 = pautil.DegreesToRadians(a)
	var d float64 = pautil.DegreesToRadians(b)
	var e float64 = pautil.DegreesToRadians(geographicalLatitude)
	var f float64 = math.Sin(d)*math.Sin(e) + math.Cos(d)*math.Cos(e)*math.Cos(c)

	return Degrees(math.Asin(f))
}

/*
Convert Horizon Coordinates to Hour Angle (in decimal degrees)

Original macro name: HORHa
*/
func HorizonCoordinatesToHourAngle(
	azimuthDegrees float64, azimuthMinutes float64, azimuthSeconds float64,
	altitudeDegrees float64, altitudeMinutes float64, altitudeSeconds float64,
	geographicalLatitude float64,
) float64 {
	var a float64 = DegreesMinutesSecondsToDecimalDegrees(azimuthDegrees, azimuthMinutes, azimuthSeconds)
	var b float64 = DegreesMinutesSecondsToDecimalDegrees(altitudeDegrees, altitudeMinutes, altitudeSeconds)
	var c float64 = pautil.DegreesToRadians(a)
	var d float64 = pautil.DegreesToRadians(b)
	var e float64 = pautil.DegreesToRadians(geographicalLatitude)
	var f float64 = math.Sin(d)*math.Sin(e) + math.Cos(d)*math.Cos(e)*math.Cos(c)
	var g float64 = -math.Cos(d) * math.Cos(e) * math.Sin(c)
	var h float64 = math.Sin(d) - math.Sin(e)*f
	var i float64 = DecimalDegreesToDegreeHours(Degrees(math.Atan2(g, h)))

	return i - 24*math.Floor(i/24)
}

/*
Obliquity of the Ecliptic for a Greenwich Date

Original macro name: Obliq
*/
func Obliq(greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	var a float64 = CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear))
	var b float64 = a - 2415020
	var c float64 = (b / 36525) - 1
	var d float64 = c * (46.815 + c*(0.0006-(c*0.00181)))
	var e float64 = d / 3600

	return 23.43929167 - e + NutatObl(greenwichDay, greenwichMonth, greenwichYear)
}

/*
Nutation amount to be added in ecliptic longitude, in degrees.

Original macro name: NutatLong
*/
func NutatLong(gd float64, gm int, gy int) float64 {
	var dj float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020
	var t float64 = dj / 36525
	var t2 float64 = t * t

	var a float64 = 100.0021358 * t
	var b float64 = 360 * (a - math.Floor(a))

	var l1 float64 = 279.6967 + 0.000303*t2 + b
	var l2 float64 = 2 * pautil.DegreesToRadians(l1)

	a = 1336.855231 * t
	b = 360 * (a - math.Floor(a))

	var d1 float64 = 270.4342 - 0.001133*t2 + b
	var d2 float64 = 2 * pautil.DegreesToRadians(d1)

	a = 99.99736056 * t
	b = 360 * (a - math.Floor(a))

	var m1 float64 = 358.4758 - 0.00015*t2 + b
	m1 = pautil.DegreesToRadians(m1)

	a = 1325.552359 * t
	b = 360 * (a - math.Floor(a))

	var m2 float64 = 296.1046 + 0.009192*t2 + b
	m2 = pautil.DegreesToRadians(m2)

	a = 5.372616667 * t
	b = 360 * (a - math.Floor(a))

	var n1 float64 = 259.1833 + 0.002078*t2 - b
	n1 = pautil.DegreesToRadians(n1)

	var n2 float64 = 2.0 * n1

	var dp float64 = (-17.2327 - 0.01737*t) * math.Sin(n1)
	dp = dp + (-1.2729-0.00013*t)*math.Sin(l2) + 0.2088*math.Sin(n2)
	dp = dp - 0.2037*math.Sin(d2) + (0.1261-0.00031*t)*math.Sin(m1)
	dp = dp + 0.0675*math.Sin(m2) - (0.0497-0.00012*t)*math.Sin(l2+m1)
	dp = dp - 0.0342*math.Sin(d2-n1) - 0.0261*math.Sin(d2+m2)
	dp = dp + 0.0214*math.Sin(l2-m1) - 0.0149*math.Sin(l2-d2+m2)
	dp = dp + 0.0124*math.Sin(l2-n1) + 0.0114*math.Sin(d2-m2)

	return dp / 3600
}

/*
Nutation of Obliquity

Original macro name: NutatObl
*/
func NutatObl(greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	var dj float64 = CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear)) - 2415020
	var t float64 = dj / 36525
	var t2 float64 = t * t

	var a float64 = 100.0021358 * t
	var b float64 = 360 * (a - math.Floor(a))

	var l1 float64 = 279.6967 + 0.000303*t2 + b
	var l2 float64 = 2 * pautil.DegreesToRadians(l1)

	a = 1336.855231 * t
	b = 360 * (a - math.Floor(a))

	var d1 float64 = 270.4342 - 0.001133*t2 + b
	var d2 float64 = 2 * pautil.DegreesToRadians(d1)

	a = 99.99736056 * t
	b = 360 * (a - math.Floor(a))

	var m1 float64 = pautil.DegreesToRadians(358.4758 - 0.00015*t2 + b)

	a = 1325.552359 * t
	b = 360 * (a - math.Floor(a))

	var m2 float64 = pautil.DegreesToRadians(296.1046 + 0.009192*t2 + b)

	a = 5.372616667 * t
	b = 360 * (a - math.Floor(a))

	var n1 float64 = pautil.DegreesToRadians(259.1833 + 0.002078*t2 - b)

	var n2 float64 = 2 * n1

	var ddo float64 = (9.21 + 0.00091*t) * math.Cos(n1)
	ddo = ddo + (0.5522-0.00029*t)*math.Cos(l2) - 0.0904*math.Cos(n2)
	ddo = ddo + 0.0884*math.Cos(d2) + 0.0216*math.Cos(l2+m1)
	ddo = ddo + 0.0183*math.Cos(d2-n1) + 0.0113*math.Cos(d2+m2)
	ddo = ddo - 0.0093*math.Cos(l2-m1) - 0.0066*math.Cos(l2-n1)

	return ddo / 3600
}

/* Convert Local Sidereal Time to Greenwich Sidereal Time */
func LocalSiderealTimeToGreenwichSiderealTime(localHours float64, localMinutes float64, localSeconds float64, longitude float64) float64 {
	var a float64 = HmsDh(localHours, localMinutes, localSeconds)
	var b float64 = longitude / 15
	var c float64 = a - b

	return c - (24 * math.Floor(c/24))
}

/*
Calculate Sun's ecliptic longitude

Original macro name: SunLong
*/
func SunLong(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	var aa float64 = LocalCivilTimeGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var bb int = int(LocalCivilTimeGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var cc int = int(LocalCivilTimeGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var ut float64 = LocalCivilTimeToUniversalTime(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var dj float64 = CivilDateToJulianDate(aa, float64(bb), float64(cc)) - 2415020
	var t float64 = (dj / 36525) + (ut / 876600)
	var t2 float64 = t * t
	var a float64 = 100.0021359 * t
	var b float64 = 360.0 * (a - math.Floor(a))

	var l float64 = 279.69668 + 0.0003025*t2 + b
	a = 99.99736042 * t
	b = 360 * (a - math.Floor(a))

	var m1 float64 = 358.47583 - (0.00015+0.0000033*t)*t2 + b
	var ec float64 = 0.01675104 - 0.0000418*t - 0.000000126*t2

	var am float64 = pautil.DegreesToRadians(m1)
	var at float64 = TrueAnomaly(am, ec)

	a = 62.55209472 * t
	b = 360 * (a - math.Floor(a))

	var a1 float64 = pautil.DegreesToRadians(153.23 + b)
	a = 125.1041894 * t
	b = 360 * (a - math.Floor(a))

	var b1 float64 = pautil.DegreesToRadians(216.57 + b)
	a = 91.56766028 * t
	b = 360 * (a - math.Floor(a))

	var c1 float64 = pautil.DegreesToRadians(312.69 + b)
	a = 1236.853095 * t
	b = 360 * (a - math.Floor(a))

	var d1 float64 = pautil.DegreesToRadians(350.74 - 0.00144*t2 + b)
	var e1 float64 = pautil.DegreesToRadians(231.19 + 20.2*t)
	a = 183.1353208 * t
	b = 360 * (a - math.Floor(a))
	// var h1 float64 = pautil.DegreesToRadians(353.4 + b)

	var d2 float64 = 0.00134*math.Cos(a1) + 0.00154*math.Cos(b1) + 0.002*math.Cos(c1)
	d2 = d2 + 0.00179*math.Sin(d1) + 0.00178*math.Sin(e1)
	var d3 float64 = 0.00000543*math.Sin(a1) + 0.00001575*math.Sin(b1)
	d3 = d3 + 0.00001627*math.Sin(c1) + 0.00003076*math.Cos(d1)

	var sr float64 = at + pautil.DegreesToRadians(l-m1+d2)
	var tp float64 = 6.283185308

	sr = sr - tp*math.Floor(sr/tp)

	return Degrees(sr)
}

/*
Calculate Sun's angular diameter in decimal degrees

Original macro name: SunDia
*/
func SunDia(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	var a float64 = SunDist(lch, lcm, lcs, ds, zc, ld, lm, ly)

	return 0.533128 / a
}

/*
Calculate Sun's distance from the Earth in astronomical units

Original macro name: SunDist
*/
func SunDist(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	var aa float64 = LocalCivilTimeGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var bb int = int(LocalCivilTimeGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var cc int = int(LocalCivilTimeGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var ut float64 = LocalCivilTimeToUniversalTime(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var dj float64 = CivilDateToJulianDate(aa, float64(bb), float64(cc)) - 2415020

	var t float64 = (dj / 36525) + (ut / 876600)
	var t2 float64 = t * t

	var a float64 = 100.0021359 * t
	var b float64 = 360 * (a - math.Floor(a))
	a = 99.99736042 * t
	b = 360 * (a - math.Floor(a))
	var m1 float64 = 358.47583 - (0.00015+0.0000033*t)*t2 + b
	var ec float64 = 0.01675104 - 0.0000418*t - 0.000000126*t2

	var am float64 = pautil.DegreesToRadians(m1)
	var ae float64 = EccentricAnomaly(am, ec)

	a = 62.55209472 * t
	b = 360 * (a - math.Floor(a))
	var a1 float64 = pautil.DegreesToRadians(153.23 + b)
	a = 125.1041894 * t
	b = 360 * (a - math.Floor(a))
	var b1 float64 = pautil.DegreesToRadians(216.57 + b)
	a = 91.56766028 * t
	b = 360 * (a - math.Floor(a))
	var c1 float64 = pautil.DegreesToRadians(312.69 + b)
	a = 1236.853095 * t
	b = 360 * (a - math.Floor(a))
	var d1 float64 = pautil.DegreesToRadians(350.74 - 0.00144*t2 + b)
	a = 183.1353208 * t
	b = 360 * (a - math.Floor(a))
	var h1 float64 = pautil.DegreesToRadians(353.4 + b)

	var d3 float64 = (0.00000543*math.Sin(a1) + 0.00001575*math.Sin(b1)) + (0.00001627*math.Sin(c1) + 0.00003076*math.Cos(d1)) + (0.00000927 * math.Sin(h1))

	return 1.0000002*(1-ec*math.Cos(ae)) + d3
}

/*
Solve Kepler's equation, and return value of the true anomaly in radians

Original macro name: TrueAnomaly
*/
func TrueAnomaly(am float64, ec float64) float64 {
	var tp float64 = 6.283185308
	var m float64 = am - tp*math.Floor(am/tp)
	var ae float64 = m

	for true {
		var d float64 = ae - (ec * math.Sin(ae)) - m
		if math.Abs(d) < 0.000001 {
			break
		}
		d = d / (1.0 - (ec * math.Cos(ae)))
		ae = ae - d
	}
	var a float64 = math.Sqrt((1+ec)/(1-ec)) * math.Tan(ae/2)
	var at float64 = 2.0 * math.Atan(a)

	return at
}

/*
Solve Kepler's equation, and return value of the eccentric anomaly in radians

Original macro name: EccentricAnomaly
*/
func EccentricAnomaly(am float64, ec float64) float64 {
	var tp float64 = 6.283185308
	var m float64 = am - tp*math.Floor(am/tp)
	var ae float64 = m

	for true {
		var d float64 = ae - (ec * math.Sin(ae)) - m

		if math.Abs(d) < 0.000001 {
			break
		}

		d = d / (1 - (ec * math.Cos(ae)))
		ae = ae - d
	}

	return ae
}

/*
Calculate effects of refraction

Original macro name: Refract
*/
func Refract(y2 float64, sw patype.CoordinateType, pr float64, tr float64) float64 {
	var y float64 = pautil.DegreesToRadians(y2)

	var d float64
	if sw == patype.CoordinateType_Actual {
		d = -1.0
	} else {
		d = 1.0
	}

	if d == -1 {
		var y3 float64 = y
		var y1 float64 = y
		var r1 float64 = 0.0

		for true {
			var y_new float64 = y1 + r1
			var rf_new float64 = Refract_L3035(pr, tr, y_new, d)

			if y < -0.087 {
				return 0
			}

			var r2 float64 = rf_new

			if (r2 == 0) || (math.Abs(r2-r1) < 0.000001) {
				var q_new float64 = y3

				return Degrees(q_new + rf_new)
			}

			r1 = r2
		}
	}

	var rf float64 = Refract_L3035(pr, tr, y, d)

	if y < -0.087 {
		return 0
	}

	var q float64 = y

	return Degrees(q + rf)
}

/* Helper function for Refract */
func Refract_L3035(pr float64, tr float64, y float64, d float64) float64 {
	if y < 0.2617994 {
		if y < -0.087 {
			return 0
		}

		var yd float64 = Degrees(y)
		var a float64 = ((0.00002*yd+0.0196)*yd + 0.1594) * pr
		var b float64 = (273.0 + tr) * ((0.0845*yd+0.505)*yd + 1)

		return pautil.DegreesToRadians(-(a / b) * d)
	}

	return -d * 0.00007888888 * pr / ((273.0 + tr) * math.Tan(y))
}

/*
Calculate corrected hour angle in decimal hours

Original macro name: ParallaxHA
*/
func ParallaxHa(hh float64, hm float64, hs float64, dd float64, dm float64, ds float64, sw patype.CoordinateType,
	gp float64, ht float64, hp float64,
) float64 {
	var a float64 = pautil.DegreesToRadians(gp)
	var c1 float64 = math.Cos(a)
	var s1 float64 = math.Sin(a)

	var u float64 = math.Atan(0.996647 * s1 / c1)
	var c2 float64 = math.Cos(u)
	var s2 float64 = math.Sin(u)
	var b float64 = ht / 6378160

	var rs float64 = (0.996647 * s2) + (b * s1)

	var rc float64 = c2 + (b * c1)
	var tp float64 = 6.283185308

	var rp float64 = 1.0 / math.Sin(pautil.DegreesToRadians(hp))

	var x float64 = pautil.DegreesToRadians(DegreeHoursToDecimalDegrees(HmsDh(hh, hm, hs)))
	var x1 float64 = x
	var y float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var y1 float64 = y

	var d float64
	if sw == patype.CoordinateType_Actual {
		d = 1.0
	} else {
		d = -1.0
	}

	if d == 1 {
		var result patype.ParallaxHelper = ParallaxHa_L2870(x, y, rc, rp, rs, tp)

		return DecimalDegreesToDegreeHours(Degrees(result.P))
	}

	var p1 float64 = 0.0
	var q1 float64 = 0.0
	var xLoop float64 = x
	var yLoop float64 = y

	for true {
		var result patype.ParallaxHelper = ParallaxHa_L2870(xLoop, yLoop, rc, rp, rs, tp)

		var p2 float64 = result.P - xLoop
		var q2 float64 = result.Q - yLoop

		var aa float64 = math.Abs(p2 - p1)
		var bb float64 = math.Abs(q2 - q1)

		if (aa < 0.000001) && (bb < 0.000001) {
			var p float64 = x1 - p2

			return DecimalDegreesToDegreeHours(Degrees(p))
		}

		xLoop = x1 - p2
		yLoop = y1 - q2
		p1 = p2
		q1 = q2
	}

	return 0 // silence the compiler error about a return
}

/* Helper function for ParallaxHa */
func ParallaxHa_L2870(x float64, y float64, rc float64, rp float64, rs float64, tp float64) patype.ParallaxHelper {
	var cx float64 = math.Cos(x)
	var sy float64 = math.Sin(y)
	var cy float64 = math.Cos(y)

	var aa float64 = (rc * math.Sin(x)) / ((rp * cy) - (rc * cx))

	var dx float64 = math.Atan(aa)
	var p float64 = x + dx
	var cp float64 = math.Cos(p)

	p = p - tp*math.Floor(p/tp)
	var q float64 = math.Atan(cp * (rp*sy - rs) / (rp*cy*cx - rc))

	return patype.ParallaxHelper{P: p, Q: q}
}

/*
Calculate corrected declination in decimal degrees

Original macro name: ParallaxDec
*/
func ParallaxDec(hh float64, hm float64, hs float64, dd float64, dm float64, ds float64, sw patype.CoordinateType, gp float64, ht float64, hp float64,
) float64 {
	var a float64 = pautil.DegreesToRadians(gp)
	var c1 float64 = math.Cos(a)
	var s1 float64 = math.Sin(a)

	var u float64 = math.Atan(0.996647 * s1 / c1)

	var c2 float64 = math.Cos(u)
	var s2 float64 = math.Sin(u)
	var b float64 = ht / 6378160
	var rs float64 = (0.996647 * s2) + (b * s1)

	var rc float64 = c2 + (b * c1)
	var tp float64 = 6.283185308

	var rp float64 = 1.0 / math.Sin(pautil.DegreesToRadians(hp))

	var x float64 = pautil.DegreesToRadians(DegreeHoursToDecimalDegrees(HmsDh(hh, hm, hs)))
	var x1 float64 = x

	var y float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var y1 float64 = y

	var d float64
	if sw == patype.CoordinateType_Actual {
		d = 1.0
	} else {
		d = -1.0
	}

	if d == 1 {
		var result patype.ParallaxHelper = ParallaxDec_L2870(x, y, rc, rp, rs, tp)

		return Degrees(result.Q)
	}

	var p1 float64 = 0.0

	var xLoop float64 = x
	var yLoop float64 = y

	for true {
		var result patype.ParallaxHelper = ParallaxDec_L2870(xLoop, yLoop, rc, rp, rs, tp)
		var p2 float64 = result.P - xLoop
		var q2 float64 = result.Q - yLoop
		var aa float64 = math.Abs(p2 - p1)

		if (aa < 0.000001) && (b < 0.000001) {
			var q float64 = y1 - q2

			return Degrees(q)
		}
		xLoop = x1 - p2
		yLoop = y1 - q2
		p1 = p2
	}

	return 0 // silence the compiler error about a return
}

/* Helper function for parallax_dec */
func ParallaxDec_L2870(x float64, y float64, rc float64, rp float64, rs float64, tp float64) patype.ParallaxHelper {
	var cx float64 = math.Cos(x)
	var sy float64 = math.Sin(y)
	var cy float64 = math.Cos(y)

	var aa float64 = (rc * math.Sin(x)) / ((rp * cy) - (rc * cx))
	var dx float64 = math.Atan(aa)
	var p float64 = x + dx
	var cp float64 = math.Cos(p)

	p = p - tp*math.Floor(p/tp)
	var q float64 = math.Atan(cp * (rp*sy - rs) / (rp*cy*cx - rc))

	return patype.ParallaxHelper{P: p, Q: q}
}

/*
Calculate geocentric ecliptic longitude for the Moon

Original macro name: MoonLong
*/
func MoonLongitude(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var ut float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var t float64 = ((CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020) / 36525) + (ut / 876600)
	var t2 float64 = t * t

	var m1 float64 = 27.32158213
	var m2 float64 = 365.2596407
	var m3 float64 = 27.55455094
	var m4 float64 = 29.53058868
	var m5 float64 = 27.21222039
	var m6 float64 = 6798.363307
	var q float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020 + (ut / 24)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360 * (m1 - math.Floor(m1))
	m2 = 360 * (m2 - math.Floor(m2))
	m3 = 360 * (m3 - math.Floor(m3))
	m4 = 360 * (m4 - math.Floor(m4))
	m5 = 360 * (m5 - math.Floor(m5))
	m6 = 360 * (m6 - math.Floor(m6))

	var ml float64 = 270.434164 + m1 - (0.001133-0.0000019*t)*t2
	var ms float64 = 358.475833 + m2 - (0.00015+0.0000033*t)*t2
	var md float64 = 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	var me1 float64 = 350.737486 + m4 - (0.001436-0.0000019*t)*t2
	var mf float64 = 11.250889 + m5 - (0.003211+0.0000003*t)*t2
	var na float64 = 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	var a float64 = pautil.DegreesToRadians(51.2 + 20.2*t)
	var s1 float64 = math.Sin(a)
	var s2 float64 = math.Sin(pautil.DegreesToRadians(na))
	var b float64 = 346.56 + (132.87-0.0091731*t)*t
	var s3 float64 = 0.003964 * math.Sin(pautil.DegreesToRadians(b))
	var c float64 = pautil.DegreesToRadians(na + 275.05 - 2.3*t)
	var s4 float64 = math.Sin(c)
	ml = ml + 0.000233*s1 + s3 + 0.001964*s2
	ms = ms - 0.001778*s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	var e float64 = 1.0 - (0.002495+0.00000752*t)*t
	var e2 float64 = e * e
	ml = pautil.DegreesToRadians(ml)
	ms = pautil.DegreesToRadians(ms)
	me1 = pautil.DegreesToRadians(me1)
	mf = pautil.DegreesToRadians(mf)
	md = pautil.DegreesToRadians(md)

	var l float64 = 6.28875*math.Sin(md) + 1.274018*math.Sin(2*me1-md)
	l = l + 0.658309*math.Sin(2*me1) + 0.213616*math.Sin(2*md)
	l = l - e*0.185596*math.Sin(ms) - 0.114336*math.Sin(2*mf)
	l = l + 0.058793*math.Sin(2*(me1-md))
	l = l + 0.057212*e*math.Sin(2*me1-ms-md) + 0.05332*math.Sin(2*me1+md)
	l = l + 0.045874*e*math.Sin(2*me1-ms) + 0.041024*e*math.Sin(md-ms)
	l = l - 0.034718*math.Sin(me1) - e*0.030465*math.Sin(ms+md)
	l = l + 0.015326*math.Sin(2*(me1-mf)) - 0.012528*math.Sin(2*mf+md)
	l = l - 0.01098*math.Sin(2*mf-md) + 0.010674*math.Sin(4*me1-md)
	l = l + 0.010034*math.Sin(3*md) + 0.008548*math.Sin(4*me1-2*md)
	l = l - e*0.00791*math.Sin(ms-md+2*me1) - e*0.006783*math.Sin(2*me1+ms)
	l = l + 0.005162*math.Sin(md-me1) + e*0.005*math.Sin(ms+me1)
	l = l + 0.003862*math.Sin(4*me1) + e*0.004049*math.Sin(md-ms+2*me1)
	l = l + 0.003996*math.Sin(2*(md+me1)) + 0.003665*math.Sin(2*me1-3*md)
	l = l + e*0.002695*math.Sin(2*md-ms) + 0.002602*math.Sin(md-2*(mf+me1))
	l = l + e*0.002396*math.Sin(2*(me1-md)-ms) - 0.002349*math.Sin(md+me1)
	l = l + e2*0.002249*math.Sin(2*(me1-ms)) - e*0.002125*math.Sin(2*md+ms)
	l = l - e2*0.002079*math.Sin(2*ms) + e2*0.002059*math.Sin(2*(me1-ms)-md)
	l = l - 0.001773*math.Sin(md+2*(me1-mf)) - 0.001595*math.Sin(2*(mf+me1))
	l = l + e*0.00122*math.Sin(4*me1-ms-md) - 0.00111*math.Sin(2*(md+mf))
	l = l + 0.000892*math.Sin(md-3*me1) - e*0.000811*math.Sin(ms+md+2*me1)
	l = l + e*0.000761*math.Sin(4*me1-ms-2*md)
	l = l + e2*0.000704*math.Sin(md-2*(ms+me1))
	l = l + e*0.000693*math.Sin(ms-2*(md-me1))
	l = l + e*0.000598*math.Sin(2*(me1-mf)-ms)
	l = l + 0.00055*math.Sin(md+4*me1) + 0.000538*math.Sin(4*md)
	l = l + e*0.000521*math.Sin(4*me1-ms) + 0.000486*math.Sin(2*md-me1)
	l = l + e2*0.000717*math.Sin(md-2*ms)

	var mm float64 = Unwind(ml + pautil.DegreesToRadians(l))

	return Degrees(mm)
}

/*
Calculate geocentric ecliptic latitude for the Moon

Original macro name: MoonLat
*/
func MoonLatitude(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var ut float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var t float64 = ((CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020) / 36525) + (ut / 876600)
	var t2 float64 = t * t

	var m1 float64 = 27.32158213
	var m2 float64 = 365.2596407
	var m3 float64 = 27.55455094
	var m4 float64 = 29.53058868
	var m5 float64 = 27.21222039
	var m6 float64 = 6798.363307
	var q float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020 + (ut / 24)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360 * (m1 - math.Floor(m1))
	m2 = 360 * (m2 - math.Floor(m2))
	m3 = 360 * (m3 - math.Floor(m3))
	m4 = 360 * (m4 - math.Floor(m4))
	m5 = 360 * (m5 - math.Floor(m5))
	m6 = 360 * (m6 - math.Floor(m6))

	var ml float64 = 270.434164 + m1 - (0.001133-0.0000019*t)*t2
	var ms float64 = 358.475833 + m2 - (0.00015+0.0000033*t)*t2
	var md float64 = 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	var me1 float64 = 350.737486 + m4 - (0.001436-0.0000019*t)*t2
	var mf float64 = 11.250889 + m5 - (0.003211+0.0000003*t)*t2
	var na float64 = 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	var a float64 = pautil.DegreesToRadians(51.2 + 20.2*t)
	var s1 float64 = math.Sin(a)
	var s2 float64 = math.Sin(pautil.DegreesToRadians(na))
	var b float64 = 346.56 + (132.87-0.0091731*t)*t
	var s3 float64 = 0.003964 * math.Sin(pautil.DegreesToRadians(b))
	var c float64 = pautil.DegreesToRadians(na + 275.05 - 2.3*t)
	var s4 float64 = math.Sin(c)
	ml = ml + 0.000233*s1 + s3 + 0.001964*s2
	ms = ms - 0.001778*s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	var e float64 = 1.0 - (0.002495+0.00000752*t)*t
	var e2 float64 = e * e
	ms = pautil.DegreesToRadians(ms)
	na = pautil.DegreesToRadians(na)
	me1 = pautil.DegreesToRadians(me1)
	mf = pautil.DegreesToRadians(mf)
	md = pautil.DegreesToRadians(md)

	var g float64 = 5.128189*math.Sin(mf) + 0.280606*math.Sin(md+mf)
	g = g + 0.277693*math.Sin(md-mf) + 0.173238*math.Sin(2*me1-mf)
	g = g + 0.055413*math.Sin(2*me1+mf-md) + 0.046272*math.Sin(2*me1-mf-md)
	g = g + 0.032573*math.Sin(2*me1+mf) + 0.017198*math.Sin(2*md+mf)
	g = g + 0.009267*math.Sin(2*me1+md-mf) + 0.008823*math.Sin(2*md-mf)
	g = g + e*0.008247*math.Sin(2*me1-ms-mf) + 0.004323*math.Sin(2*(me1-md)-mf)
	g = g + 0.0042*math.Sin(2*me1+mf+md) + e*0.003372*math.Sin(mf-ms-2*me1)
	g = g + e*0.002472*math.Sin(2*me1+mf-ms-md)
	g = g + e*0.002222*math.Sin(2*me1+mf-ms)
	g = g + e*0.002072*math.Sin(2*me1-mf-ms-md)
	g = g + e*0.001877*math.Sin(mf-ms+md) + 0.001828*math.Sin(4*me1-mf-md)
	g = g - e*0.001803*math.Sin(mf+ms) - 0.00175*math.Sin(3*mf)
	g = g + e*0.00157*math.Sin(md-ms-mf) - 0.001487*math.Sin(mf+me1)
	g = g - e*0.001481*math.Sin(mf+ms+md) + e*0.001417*math.Sin(mf-ms-md)
	g = g + e*0.00135*math.Sin(mf-ms) + 0.00133*math.Sin(mf-me1)
	g = g + 0.001106*math.Sin(mf+3*md) + 0.00102*math.Sin(4*me1-mf)
	g = g + 0.000833*math.Sin(mf+4*me1-md) + 0.000781*math.Sin(md-3*mf)
	g = g + 0.00067*math.Sin(mf+4*me1-2*md) + 0.000606*math.Sin(2*me1-3*mf)
	g = g + 0.000597*math.Sin(2*(me1+md)-mf)
	g = g + e*0.000492*math.Sin(2*me1+md-ms-mf) + 0.00045*math.Sin(2*(md-me1)-mf)
	g = g + 0.000439*math.Sin(3*md-mf) + 0.000423*math.Sin(mf+2*(me1+md))
	g = g + 0.000422*math.Sin(2*me1-mf-3*md) - e*0.000367*math.Sin(ms+mf+2*me1-md)
	g = g - e*0.000353*math.Sin(ms+mf+2*me1) + 0.000331*math.Sin(mf+4*me1)
	g = g + e*0.000317*math.Sin(2*me1+mf-ms+md)
	g = g + e2*0.000306*math.Sin(2*(me1-ms)-mf) - 0.000283*math.Sin(md+3*mf)

	var w1 float64 = 0.0004664 * math.Cos(na)
	var w2 float64 = 0.0000754 * math.Cos(c)
	var bm float64 = pautil.DegreesToRadians(g) * (1.0 - w1 - w2)

	return Degrees(bm)
}

/*
Calculate horizontal parallax for the Moon

Original macro name: MoonHP
*/
func MoonHorizontalParallax(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var ut float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var t float64 = ((CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020) / 36525) + (ut / 876600)
	var t2 float64 = t * t

	var m1 float64 = 27.32158213
	var m2 float64 = 365.2596407
	var m3 float64 = 27.55455094
	var m4 float64 = 29.53058868
	var m5 float64 = 27.21222039
	var m6 float64 = 6798.363307
	var q float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020 + (ut / 24)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360 * (m1 - math.Floor(m1))
	m2 = 360 * (m2 - math.Floor(m2))
	m3 = 360 * (m3 - math.Floor(m3))
	m4 = 360 * (m4 - math.Floor(m4))
	m5 = 360 * (m5 - math.Floor(m5))
	m6 = 360 * (m6 - math.Floor(m6))

	var ml float64 = 270.434164 + m1 - (0.001133-0.0000019*t)*t2
	var ms float64 = 358.475833 + m2 - (0.00015+0.0000033*t)*t2
	var md float64 = 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	var me1 float64 = 350.737486 + m4 - (0.001436-0.0000019*t)*t2
	var mf float64 = 11.250889 + m5 - (0.003211+0.0000003*t)*t2
	var na float64 = 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	var a float64 = pautil.DegreesToRadians(51.2 + 20.2*t)
	var s1 float64 = math.Sin(a)
	var s2 float64 = math.Sin(pautil.DegreesToRadians(na))
	var b float64 = 346.56 + (132.87-0.0091731*t)*t
	var s3 float64 = 0.003964 * math.Sin(pautil.DegreesToRadians(b))
	var c float64 = pautil.DegreesToRadians(na + 275.05 - 2.3*t)
	var s4 float64 = math.Sin(c)
	ml = ml + 0.000233*s1 + s3 + 0.001964*s2
	ms = ms - 0.001778*s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	var e float64 = 1.0 - (0.002495+0.00000752*t)*t
	var e2 float64 = e * e
	ms = pautil.DegreesToRadians(ms)
	me1 = pautil.DegreesToRadians(me1)
	mf = pautil.DegreesToRadians(mf)
	md = pautil.DegreesToRadians(md)

	var pm float64 = 0.950724 + 0.051818*math.Cos(md) + 0.009531*math.Cos(2*me1-md)
	pm = pm + 0.007843*math.Cos(2*me1) + 0.002824*math.Cos(2*md)
	pm = pm + 0.000857*math.Cos(2*me1+md) + e*0.000533*math.Cos(2*me1-ms)
	pm = pm + e*0.000401*math.Cos(2*me1-md-ms)
	pm = pm + e*0.00032*math.Cos(md-ms) - 0.000271*math.Cos(me1)
	pm = pm - e*0.000264*math.Cos(ms+md) - 0.000198*math.Cos(2*mf-md)
	pm = pm + 0.000173*math.Cos(3*md) + 0.000167*math.Cos(4*me1-md)
	pm = pm - e*0.000111*math.Cos(ms) + 0.000103*math.Cos(4*me1-2*md)
	pm = pm - 0.000084*math.Cos(2*md-2*me1) - e*0.000083*math.Cos(2*me1+ms)
	pm = pm + 0.000079*math.Cos(2*me1+2*md) + 0.000072*math.Cos(4*me1)
	pm = pm + e*0.000064*math.Cos(2*me1-ms+md) - e*0.000063*math.Cos(2*me1+ms-md)
	pm = pm + e*0.000041*math.Cos(ms+me1) + e*0.000035*math.Cos(2*md-ms)
	pm = pm - 0.000033*math.Cos(3*md-2*me1) - 0.00003*math.Cos(md+me1)
	pm = pm - 0.000029*math.Cos(2*(mf-me1)) - e*0.000029*math.Cos(2*md+ms)
	pm = pm + e2*0.000026*math.Cos(2*(me1-ms)) - 0.000023*math.Cos(2*(mf-me1)+md)
	pm = pm + e*0.000019*math.Cos(4*me1-ms-md)

	return pm
}

/*
Convert angle in radians to equivalent angle in degrees.

Original macro name: Unwind
*/
func Unwind(w float64) float64 {
	return w - 6.283185308*math.Floor(w/6.283185308)
}

/*
Convert angle in degrees to equivalent angle in the range 0 to 360 degrees.

Original macro name: UnwindDeg
*/
func ma_unwind_deg(w float64) float64 {
	return w - 360*math.Floor(w/360)
}

/*
Mean ecliptic longitude of the Sun at the epoch

Original macro name: SunElong
*/
func SunEclipticLongitude(gd float64, gm int, gy int) float64 {
	var t float64 = (CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020) / 36525
	var t2 float64 = t * t
	var x float64 = 279.6966778 + 36000.76892*t + 0.0003025*t2

	return x - 360*math.Floor(x/360)
}

/*
Longitude of the Sun at perigee

Original macro name: SunPeri
*/
func SunPerigee(gd float64, gm int, gy int) float64 {
	var t float64 = (CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020) / 36525
	var t2 float64 = t * t
	var x float64 = 281.2208444 + 1.719175*t + 0.000452778*t2

	return x - 360*math.Floor(x/360)
}

/*
Eccentricity of the Sun-Earth orbit

Original macro name: SunEcc
*/
func SunEccentricity(gd float64, gm int, gy int) float64 {
	var t float64 = (CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020) / 36525
	var t2 float64 = t * t

	return 0.01675104 - 0.0000418*t - 0.000000126*t2
}

/*
Ecliptic - Declination (degrees)

Original macro name: ECDec
*/
func EclipticDeclination(eld float64, elm float64, els float64, bd float64, bm float64, bs float64, gd float64, gm int, gy int) float64 {
	var a float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(eld, elm, els))
	var b float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(bd, bm, bs))
	var c float64 = pautil.DegreesToRadians(Obliq(gd, gm, gy))
	var d float64 = math.Sin(b)*math.Cos(c) + math.Cos(b)*math.Sin(c)*math.Sin(a)

	return Degrees(math.Asin(d))
}

/*
Ecliptic - Right Ascension (degrees)

Original macro name: ECRA
*/
func EclipticRightAscension(eld float64, elm float64, els float64, bd float64, bm float64, bs float64, gd float64, gm int, gy int) float64 {
	var a float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(eld, elm, els))
	var b float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(bd, bm, bs))
	var c float64 = pautil.DegreesToRadians(Obliq(gd, gm, gy))
	var d float64 = math.Sin(a)*math.Cos(c) - math.Tan(b)*math.Sin(c)
	var e float64 = math.Cos(a)
	var f float64 = Degrees(math.Atan2(d, e))

	return f - 360*math.Floor(f/360)
}

/*
Calculate Sun's true anomaly, i.e., how much its orbit deviates from a true circle to an ellipse.

Original macro name: SunTrueAnomaly
*/
func SunTrueAnomaly(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	var aa float64 = LocalCivilTimeGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var bb int = int(LocalCivilTimeGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var cc int = int(LocalCivilTimeGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var ut float64 = LocalCivilTimeToUniversalTime(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var dj float64 = CivilDateToJulianDate(aa, float64(bb), float64(cc)) - 2415020

	var t float64 = (dj / 36525) + (ut / 876600)
	var t2 float64 = t * t

	var a float64 = 99.99736042 * t
	var b float64 = 360 * (a - math.Floor(a))

	var m1 float64 = 358.47583 - (0.00015+0.0000033*t)*t2 + b
	var ec float64 = 0.01675104 - 0.0000418*t - 0.000000126*t2

	var am float64 = pautil.DegreesToRadians(m1)

	return Degrees(TrueAnomaly(am, ec))
}

/*
Calculate local civil time of sunrise.

Original macro name: SunriseLCT
*/
func SunriseLct(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64) float64 {
	var di float64 = 0.8333333
	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.SunriseLctHelper = SunriseLct_L3710(gd, gm, gy, sr, di, gp)

	var xx float64
	if result1.S != patype.RiseSetStatus_OK {
		xx = -99.0
	} else {
		var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
		var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

		if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
			xx = -99.0
		} else {
			sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)
			var result2 patype.SunriseLctHelper = SunriseLct_L3710(gd, gm, gy, sr, di, gp)

			if result2.S != patype.RiseSetStatus_OK {
				xx = -99.0
			} else {
				x = LocalSiderealTimeToGreenwichSiderealTime(result2.La, 0, 0, gl)
				ut = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)
				xx = UniversalTimeToLocalCivilTime(ut, 0, 0, ds, zc, gd, gm, gy)
			}
		}
	}

	return xx
}

/* Helper function for sunrise_lct() */
func SunriseLct_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.SunriseLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeRise(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0.0, 0.0, y, 0.0, 0.0, di, gp)

	return patype.SunriseLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/* Calculate local civil time of sunset. */
func SunsetLct(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64) float64 {
	var di float64 = 0.8333333
	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.SunsetLctHelper = SunsetLct_L3710(gd, gm, gy, sr, di, gp)

	var xx float64
	if result1.S != patype.RiseSetStatus_OK {
		xx = -99.0
	} else {
		var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
		var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

		if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
			xx = -99.0
		} else {
			sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)
			var result2 patype.SunsetLctHelper = SunsetLct_L3710(gd, gm, gy, sr, di, gp)

			if result2.S != patype.RiseSetStatus_OK {
				xx = -99
			} else {
				x = LocalSiderealTimeToGreenwichSiderealTime(result2.La, 0, 0, gl)
				ut = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)
				xx = UniversalTimeToLocalCivilTime(ut, 0, 0, ds, zc, gd, gm, gy)
			}
		}
	}

	return xx
}

/* Helper function for sunset_lct() */
func SunsetLct_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.SunsetLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0.0, 0.0, 0.0, 0.0, 0.0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0.0, 0.0, 0.0, 0.0, 0.0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	return patype.SunsetLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/*
Status of conversion of Greenwich Sidereal Time to Universal Time.

Original macro name: eGSTUT
*/
func EGstUt(gsh float64, gsm float64, gss float64, gd float64, gm int, gy int) patype.WarningFlags {
	var a float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy))
	var b float64 = a - 2451545
	var c float64 = b / 36525
	var d float64 = 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c)
	var e float64 = d - (24 * math.Floor(d/24))
	var f float64 = HmsDh(gsh, gsm, gss)
	var g float64 = f - e
	var h float64 = g - (24 * math.Floor(g/24))

	if (h * 0.9972695663) < (4.0 / 60.0) {
		return patype.WarningFlag_Warning
	} else {
		return patype.WarningFlag_OK
	}
}

/*
Local sidereal time of rise, in hours.

Original macro name: RSLSTR
*/
func RiseSetLocalSiderealTimeRise(rah float64, ram float64, ras float64, dd float64, dm float64, ds float64, vd float64, g float64) float64 {
	var a float64 = HmsDh(rah, ram, ras)
	var b float64 = pautil.DegreesToRadians(DegreeHoursToDecimalDegrees(a))
	var c float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var d float64 = pautil.DegreesToRadians(vd)
	var e float64 = pautil.DegreesToRadians(g)
	var f float64 = -(math.Sin(d) + math.Sin(e)*math.Sin(c)) / (math.Cos(e) * math.Cos(c))

	var h float64
	if math.Abs(f) < 1 {
		h = math.Acos(f)
	} else {
		h = 0
	}

	var i float64 = DecimalDegreesToDegreeHours(Degrees(b - h))

	return i - 24*math.Floor(i/24)
}

/*
Local sidereal time of setting, in hours.

Original macro name: RSLSTS
*/
func RiseSetLocalSiderealTimeSet(rah float64, ram float64, ras float64, dd float64, dm float64, ds float64, vd float64, g float64) float64 {
	var a float64 = HmsDh(rah, ram, ras)
	var b float64 = pautil.DegreesToRadians(DegreeHoursToDecimalDegrees(a))
	var c float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var d float64 = pautil.DegreesToRadians(vd)
	var e float64 = pautil.DegreesToRadians(g)
	var f float64 = -(math.Sin(d) + math.Sin(e)*math.Sin(c)) / (math.Cos(e) * math.Cos(c))
	var h float64
	if math.Abs(f) < 1 {
		h = math.Acos(f)
	} else {
		h = 0
	}
	var i float64 = DecimalDegreesToDegreeHours(Degrees(b + h))

	return i - 24*math.Floor(i/24)
}

/*
Azimuth of rising, in degrees.

Original macro name: RSAZR
*/
func RiseSetAzimuthRise(rah float64, ram float64, ras float64, dd float64, dm float64, ds float64, vd float64, g float64) float64 {
	var c float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var d float64 = pautil.DegreesToRadians(vd)
	var e float64 = pautil.DegreesToRadians(g)
	var f float64 = (math.Sin(c) + math.Sin(d)*math.Sin(e)) / (math.Cos(d) * math.Cos(e))

	var h float64
	if ERiseSet(rah, ram, ras, dd, dm, ds, vd, g) == patype.RiseSetStatus_OK {
		h = math.Acos(f)
	} else {
		h = 0
	}

	var i float64 = Degrees(h)

	return i - 360*math.Floor(i/360)
}

/*
Azimuth of setting, in degrees.

Original macro name: RSAZS
*/
func RiseSetAzimuthSet(rah float64, ram float64, ras float64, dd float64, dm float64, ds float64, vd float64, g float64) float64 {
	var c float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var d float64 = pautil.DegreesToRadians(vd)
	var e float64 = pautil.DegreesToRadians(g)
	var f float64 = (math.Sin(c) + math.Sin(d)*math.Sin(e)) / (math.Cos(d) * math.Cos(e))

	var h float64
	if ERiseSet(rah, ram, ras, dd, dm, ds, vd, g) == patype.RiseSetStatus_OK {
		h = math.Acos(f)
	} else {
		h = 0
	}

	var i float64 = 360 - Degrees(h)

	return i - 360*math.Floor(i/360)
}

/* Rise/Set status */
func ERiseSet(rah float64, ram float64, ras float64, dd float64, dm float64, ds float64, vd float64, g float64) patype.RiseSetStatus {
	var c float64 = pautil.DegreesToRadians(DegreesMinutesSecondsToDecimalDegrees(dd, dm, ds))
	var d float64 = pautil.DegreesToRadians(vd)
	var e float64 = pautil.DegreesToRadians(g)
	var f float64 = -(math.Sin(d) + math.Sin(e)*math.Sin(c)) / (math.Cos(e) * math.Cos(c))

	var return_value patype.RiseSetStatus = patype.RiseSetStatus_OK
	if f >= 1 {
		return_value = patype.RiseSetStatus_NeverRises
	}
	if f <= -1 {
		return_value = patype.RiseSetStatus_Circumpolar
	}

	return return_value
}

/*
Sunrise/Sunset calculation status.

Original macro name: eSunRS
*/
func ESunRiseSet(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64) patype.RiseSetStatus {
	var di float64 = 0.8333333
	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.SunriseLctHelper = ESunRiseSet_L3710(gd, gm, gy, sr, di, gp)

	if result1.S != patype.RiseSetStatus_OK {
		return result1.S
	} else {
		var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
		var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)
		sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)
		var result2 patype.SunriseLctHelper = ESunRiseSet_L3710(gd, gm, gy, sr, di, gp)
		if result2.S != patype.RiseSetStatus_OK {
			return result2.S
		} else {
			x = LocalSiderealTimeToGreenwichSiderealTime(result2.La, 0, 0, gl)

			if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
				var s patype.RiseSetStatus = patype.RiseSetStatus_GstToUtConversionWarning

				return s
			}

			return result2.S
		}
	}
}

/* Helper function for e_sun_rs() */
func ESunRiseSet_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.SunriseLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeRise(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	return patype.SunriseLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/*
Calculate azimuth of sunrise.

Original macro name: SunriseAz
*/
func SunriseAz(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64) float64 {
	var di float64 = 0.8333333
	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.SunriseLctHelper = SunriseAz_L3710(gd, gm, gy, sr, di, gp)

	if result1.S != patype.RiseSetStatus_OK {
		return -99.0
	}

	var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
	var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

	if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
		return -99.0
	}

	sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)
	var result2 patype.SunriseLctHelper = SunriseAz_L3710(gd, gm, gy, sr, di, gp)

	if result2.S != patype.RiseSetStatus_OK {
		return -99.0
	}

	return RiseSetAzimuthRise(DecimalDegreesToDegreeHours(x), 0, 0, result2.Y, 0.0, 0.0, di, gp)
}

/* Helper function for sunrise_az() */
func SunriseAz_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.SunriseLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeRise(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	return patype.SunriseLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/*
Calculate azimuth of sunset.

Original macro name: SunsetAz
*/
func SunsetAz(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64) float64 {
	var di float64 = 0.8333333
	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.SunsetLctHelper = SunsetAz_L3710(gd, gm, gy, sr, di, gp)

	if result1.S != patype.RiseSetStatus_OK {
		return -99.0
	}

	var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
	var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

	if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
		return -99.0
	}

	sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)

	var result2 patype.SunsetLctHelper = SunsetAz_L3710(gd, gm, gy, sr, di, gp)

	if result2.S != patype.RiseSetStatus_OK {
		return -99.0
	}

	return RiseSetAzimuthSet(DecimalDegreesToDegreeHours(x), 0, 0, result2.Y, 0, 0, di, gp)
}

/* Helper function for sunset_az() */
func SunsetAz_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.SunsetLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	return patype.SunsetLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/*
Calculate morning twilight start, in local time.

Original macro name: TwilightAMLCT
*/
func TwilightAmLct(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64, tt patype.TwilightType) float64 {
	var di float64 = float64(tt)

	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.TwilightLctHelper = TwilightAmLct_L3710(gd, gm, gy, sr, di, gp)

	if result1.S != patype.RiseSetStatus_OK {
		return -99.0
	}

	var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
	var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

	if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
		return -99.0
	}

	sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)

	var result2 patype.TwilightLctHelper = TwilightAmLct_L3710(gd, gm, gy, sr, di, gp)

	if result2.S != patype.RiseSetStatus_OK {
		return -99.0
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(result2.La, 0, 0, gl)
	ut = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

	var xx float64 = UniversalTimeToLocalCivilTime(ut, 0, 0, ds, zc, gd, gm, gy)

	return xx
}

/* Helper function for twilight_am_lct() */
func TwilightAmLct_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.TwilightLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeRise(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	return patype.TwilightLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/* Calculate evening twilight end, in local time. */
func TwilightPmLct(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64, tt patype.TwilightType) float64 {
	var di float64 = float64(tt)

	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.TwilightLctHelper = TwilightPmLct_L3710(gd, gm, gy, sr, di, gp)

	if result1.S != patype.RiseSetStatus_OK {
		return 0.0
	}

	var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
	var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

	if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
		return 0.0
	}

	sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)

	var result2 patype.TwilightLctHelper = TwilightPmLct_L3710(gd, gm, gy, sr, di, gp)

	if result2.S != patype.RiseSetStatus_OK {
		return 0.0
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(result2.La, 0, 0, gl)
	ut = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)

	return UniversalTimeToLocalCivilTime(ut, 0, 0, ds, zc, gd, gm, gy)
}

/* Helper function for twilight_pm_lct() */
func TwilightPmLct_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.TwilightLctHelper {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	return patype.TwilightLctHelper{A: a, X: x, Y: y, La: la, S: s}
}

/*
Twilight calculation status.

Original macro name: eTwilight
*/
func ETwilight(ld float64, lm int, ly int, ds int, zc int, gl float64, gp float64, tt patype.TwilightType) patype.TwilightStatus {
	var di float64 = float64(tt)

	var gd float64 = LocalCivilTimeGreenwichDay(12, 0, 0, ds, zc, ld, lm, ly)
	var gm int = int(LocalCivilTimeGreenwichMonth(12, 0, 0, ds, zc, ld, lm, ly))
	var gy int = int(LocalCivilTimeGreenwichYear(12, 0, 0, ds, zc, ld, lm, ly))
	var sr float64 = SunLong(12, 0, 0, ds, zc, ld, lm, ly)

	var result1 patype.TwilightLctHelper2 = ETwilight_L3710(gd, gm, gy, sr, di, gp)

	if result1.S != patype.TwilightStatus_OK {
		return result1.S
	}

	var x float64 = LocalSiderealTimeToGreenwichSiderealTime(result1.La, 0, 0, gl)
	var ut float64 = GreenwichSiderealTimeToUniversalTime(x, 0, 0, gd, gm, gy)
	sr = SunLong(ut, 0, 0, 0, 0, gd, gm, gy)

	var result2 patype.TwilightLctHelper2 = ETwilight_L3710(gd, gm, gy, sr, di, gp)

	if result2.S != patype.TwilightStatus_OK {
		return result1.S
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(result2.La, 0, 0, gl)

	if EGstUt(x, 0, 0, gd, gm, gy) != patype.WarningFlag_OK {
		return patype.TwilightStatus_GstToUtConversionWarning
	}

	return result2.S
}

/* Helper function for e_twilight() */
func ETwilight_L3710(gd float64, gm int, gy int, sr float64, di float64, gp float64) patype.TwilightLctHelper2 {
	var a float64 = sr + NutatLong(gd, gm, gy) - 0.005694
	var x float64 = EclipticRightAscension(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var y float64 = EclipticDeclination(a, 0, 0, 0, 0, 0, gd, gm, gy)
	var la float64 = RiseSetLocalSiderealTimeRise(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)
	var s patype.RiseSetStatus = ERiseSet(DecimalDegreesToDegreeHours(x), 0, 0, y, 0, 0, di, gp)

	var ts patype.TwilightStatus = patype.TwilightStatus_OK
	if s == patype.RiseSetStatus_Circumpolar {
		ts = patype.TwilightStatus_LastsAllNight
	}
	if s == patype.RiseSetStatus_NeverRises {
		ts = patype.TwilightStatus_SunTooFarBelowHorizon
	}

	return patype.TwilightLctHelper2{A: a, X: x, Y: y, La: la, S: ts}
}

/*
Calculate the angle between two celestial objects

Original macro name: Angle
*/
func Angle(
	xx1 float64, xm1 float64, xs1 float64, dd1 float64, dm1 float64, ds1 float64,
	xx2 float64, xm2 float64, xs2 float64, dd2 float64, dm2 float64, ds2 float64,
	s patype.AngleMeasurementTypes,
) float64 {
	var a float64
	if s == patype.AngleMeasurementType_Hours {
		a = DegreeHoursToDecimalDegrees(HmsDh(xx1, xm1, xs1))
	} else {
		a = DegreesMinutesSecondsToDecimalDegrees(xx1, xm1, xs1)
	}

	var b float64 = pautil.DegreesToRadians(a)
	var c float64 = DegreesMinutesSecondsToDecimalDegrees(dd1, dm1, ds1)
	var d float64 = pautil.DegreesToRadians(c)

	var e float64
	if s == patype.AngleMeasurementType_Hours {
		e = DegreeHoursToDecimalDegrees(HmsDh(xx2, xm2, xs2))
	} else {
		e = DegreesMinutesSecondsToDecimalDegrees(xx2, xm2, xs2)
	}

	var f float64 = pautil.DegreesToRadians(e)
	var g float64 = DegreesMinutesSecondsToDecimalDegrees(dd2, dm2, ds2)
	var h float64 = pautil.DegreesToRadians(g)
	var i float64 = math.Acos(math.Sin(d)*math.Sin(h) + math.Cos(d)*math.Cos(h)*math.Cos(b-f))

	return Degrees(i)
}
