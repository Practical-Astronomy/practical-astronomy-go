package macros

import (
	"math"
	padata "practicalastro/lib/data"
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
func DecimalDegreesToDegreeHours(decimalDegrees float64) float64 {
	return decimalDegrees / 15
}

/*
Convert Degree-Hours to Decimal Degrees

Original macro name: DHDD
*/
func DegreeHoursToDecimalDegrees(degreeHours float64) float64 {
	return degreeHours * 15
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
			var yNew float64 = y1 + r1
			var rfNew float64 = Refract_L3035(pr, tr, yNew, d)

			if y < -0.087 {
				return 0
			}

			var r2 float64 = rfNew

			if (r2 == 0) || (math.Abs(r2-r1) < 0.000001) {
				var qNew float64 = y3

				return Degrees(qNew + rfNew)
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

/* Helper function for ParallaxDec */
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
func UnwindDeg(w float64) float64 {
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
Calculate the Sun's mean anomaly.

Original macro name: SunMeanAnomaly
*/
func SunMeanAnomaly(lch float64, lcm float64, lcs float64, ds int, zc int, ld float64, lm int, ly int) float64 {
	var aa float64 = LocalCivilTimeGreenwichDay(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var bb int = int(LocalCivilTimeGreenwichMonth(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var cc int = int(LocalCivilTimeGreenwichYear(lch, lcm, lcs, ds, zc, ld, lm, ly))
	var ut float64 = LocalCivilTimeToUniversalTime(lch, lcm, lcs, ds, zc, ld, lm, ly)
	var dj float64 = CivilDateToJulianDate(aa, float64(bb), float64(cc)) - 2415020
	var t float64 = (dj / 36525) + (ut / 876600)
	var t2 float64 = t * t
	var a float64 = 100.0021359 * t
	var b float64 = 360 * (a - math.Floor(a))
	var m1 float64 = 358.47583 - (0.00015+0.0000033*t)*t2 + b
	var am float64 = Unwind(pautil.DegreesToRadians(m1))

	return am
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

/* Helper function for SunriseLct() */
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

/* Helper function for SunsetLct() */
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

	var returnValue patype.RiseSetStatus = patype.RiseSetStatus_OK
	if f >= 1 {
		returnValue = patype.RiseSetStatus_NeverRises
	}
	if f <= -1 {
		returnValue = patype.RiseSetStatus_Circumpolar
	}

	return returnValue
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

/* Helper function for ESunRiseSet() */
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

/* Helper function for SunriseAz() */
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

/* Helper function for SunsetAz() */
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

/* Helper function for TwilightAmLct() */
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

/* Helper function for TwilightPmLct() */
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

/* Helper function for ETwilight() */
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

/* Calculate several planetary properties. */
func PlanetCoordinates(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int, s string) patype.PlanetCoordinates {
	var a11 float64 = 178.179078
	var a12 float64 = 415.2057519
	var a13 float64 = 0.0003011
	var a14 float64 = 0.0
	var a21 float64 = 75.899697
	var a22 float64 = 1.5554889
	var a23 float64 = 0.0002947
	var a24 float64 = 0.0
	var a31 float64 = 0.20561421
	var a32 float64 = 0.00002046
	var a33 float64 = -0.00000003
	var a34 float64 = 0.0
	var a41 float64 = 7.002881
	var a42 float64 = 0.0018608
	var a43 float64 = -0.0000183
	var a44 float64 = 0.0
	var a51 float64 = 47.145944
	var a52 float64 = 1.1852083
	var a53 float64 = 0.0001739
	var a54 float64 = 0.0
	var a61 float64 = 0.3870986
	var a62 float64 = 6.74
	var a63 float64 = -0.42

	var b11 float64 = 342.767053
	var b12 float64 = 162.5533664
	var b13 float64 = 0.0003097
	var b14 float64 = 0.0
	var b21 float64 = 130.163833
	var b22 float64 = 1.4080361
	var b23 float64 = -0.0009764
	var b24 float64 = 0.0
	var b31 float64 = 0.00682069
	var b32 float64 = -0.00004774
	var b33 float64 = 0.000000091
	var b34 float64 = 0.0
	var b41 float64 = 3.393631
	var b42 float64 = 0.0010058
	var b43 float64 = -0.000001
	var b44 float64 = 0.0
	var b51 float64 = 75.779647
	var b52 float64 = 0.89985
	var b53 float64 = 0.00041
	var b54 float64 = 0.0
	var b61 float64 = 0.7233316
	var b62 float64 = 16.92
	var b63 float64 = -4.4

	var c11 float64 = 293.737334
	var c12 float64 = 53.17137642
	var c13 float64 = 0.0003107
	var c14 float64 = 0.0
	var c21 float64 = 334.218203
	var c22 float64 = 1.8407584
	var c23 float64 = 0.0001299
	var c24 float64 = -0.00000119
	var c31 float64 = 0.0933129
	var c32 float64 = 0.000092064
	var c33 float64 = -0.000000077
	var c34 float64 = 0.0
	var c41 float64 = 1.850333
	var c42 float64 = -0.000675
	var c43 float64 = 0.0000126
	var c44 float64 = 0.0
	var c51 float64 = 48.786442
	var c52 float64 = 0.7709917
	var c53 float64 = -0.0000014
	var c54 float64 = -0.00000533
	var c61 float64 = 1.5236883
	var c62 float64 = 9.36
	var c63 float64 = -1.52

	var d11 float64 = 238.049257
	var d12 float64 = 8.434172183
	var d13 float64 = 0.0003347
	var d14 float64 = -0.00000165
	var d21 float64 = 12.720972
	var d22 float64 = 1.6099617
	var d23 float64 = 0.00105627
	var d24 float64 = -0.00000343
	var d31 float64 = 0.04833475
	var d32 float64 = 0.00016418
	var d33 float64 = -0.0000004676
	var d34 float64 = -0.0000000017
	var d41 float64 = 1.308736
	var d42 float64 = -0.0056961
	var d43 float64 = 0.0000039
	var d44 float64 = 0.0
	var d51 float64 = 99.443414
	var d52 float64 = 1.01053
	var d53 float64 = 0.00035222
	var d54 float64 = -0.00000851
	var d61 float64 = 5.202561
	var d62 float64 = 196.74
	var d63 float64 = -9.4

	var e11 float64 = 266.564377
	var e12 float64 = 3.398638567
	var e13 float64 = 0.0003245
	var e14 float64 = -0.0000058
	var e21 float64 = 91.098214
	var e22 float64 = 1.9584158
	var e23 float64 = 0.00082636
	var e24 float64 = 0.00000461
	var e31 float64 = 0.05589232
	var e32 float64 = -0.0003455
	var e33 float64 = -0.000000728
	var e34 float64 = 0.00000000074
	var e41 float64 = 2.492519
	var e42 float64 = -0.0039189
	var e43 float64 = -0.00001549
	var e44 float64 = 0.00000004
	var e51 float64 = 112.790414
	var e52 float64 = 0.8731951
	var e53 float64 = -0.00015218
	var e54 float64 = -0.00000531
	var e61 float64 = 9.554747
	var e62 float64 = 165.6
	var e63 float64 = -8.88

	var f11 float64 = 244.19747
	var f12 float64 = 1.194065406
	var f13 float64 = 0.000316
	var f14 float64 = -0.0000006
	var f21 float64 = 171.548692
	var f22 float64 = 1.4844328
	var f23 float64 = 0.0002372
	var f24 float64 = -0.00000061
	var f31 float64 = 0.0463444
	var f32a float64 = -0.00002658
	var f33 float64 = 0.000000077
	var f34 float64 = 0.0
	var f41 float64 = 0.772464
	var f42 float64 = 0.0006253
	var f43 float64 = 0.0000395
	var f44 float64 = 0.0
	var f51 float64 = 73.477111
	var f52 float64 = 0.4986678
	var f53 float64 = 0.0013117
	var f54 float64 = 0.0
	var f61 float64 = 19.21814
	var f62 float64 = 65.8
	var f63 float64 = -7.19

	var g11 float64 = 84.457994
	var g12 float64 = 0.6107942056
	var g13 float64 = 0.0003205
	var g14 float64 = -0.0000006
	var g21 float64 = 46.727364
	var g22 float64 = 1.4245744
	var g23 float64 = 0.00039082
	var g24 float64 = -0.000000605
	var g31 float64 = 0.00899704
	var g32 float64 = 0.00000633
	var g33 float64 = -0.000000002
	var g34 float64 = 0.0
	var g41 float64 = 1.779242
	var g42 float64 = -0.0095436
	var g43 float64 = -0.0000091
	var g44 float64 = 0.0
	var g51 float64 = 130.681389
	var g52 float64 = 1.098935
	var g53 float64 = 0.00024987
	var g54 float64 = -0.000004718
	var g61 float64 = 30.10957
	var g62 float64 = 62.2
	var g63 float64 = -6.87

	var pl [9]padata.PlanetDataPrecise

	pl[0] = padata.PopulatePrecisePlanetData("", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	var ip int = 0
	var b float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var a float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy))
	var t float64 = ((a - 2415020.0) / 36525.0) + (b / 876600.0)

	var a0 float64 = a11
	var a1 float64 = a12
	var a2 float64 = a13
	var a3 float64 = a14
	var b0 float64 = a21
	var b1 float64 = a22
	var b2 float64 = a23
	var b3 float64 = a24
	var c0 float64 = a31
	var c1 float64 = a32
	var c2 float64 = a33
	var c3 float64 = a34
	var d0 float64 = a41
	var d1 float64 = a42
	var d2 float64 = a43
	var d3 float64 = a44
	var e0 float64 = a51
	var e1 float64 = a52
	var e2 float64 = a53
	var e3 float64 = a54
	var f float64 = a61
	var g float64 = a62
	var h float64 = a63
	var aa float64 = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	var c float64 = a0 + b + (a3*t+a2)*t*t

	pl[1] = padata.PopulatePrecisePlanetData(
		"Mercury", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	a0 = b11
	a1 = b12
	a2 = b13
	a3 = b14
	b0 = b21
	b1 = b22
	b2 = b23
	b3 = b24
	c0 = b31
	c1 = b32
	c2 = b33
	c3 = b34
	d0 = b41
	d1 = b42
	d2 = b43
	d3 = b44
	e0 = b51
	e1 = b52
	e2 = b53
	e3 = b54
	f = b61
	g = b62
	h = b63
	aa = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	c = a0 + b + (a3*t+a2)*t*t

	pl[2] = padata.PopulatePrecisePlanetData(
		"Venus", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	a0 = c11
	a1 = c12
	a2 = c13
	a3 = c14
	b0 = c21
	b1 = c22
	b2 = c23
	b3 = c24
	c0 = c31
	c1 = c32
	c2 = c33
	c3 = c34
	d0 = c41
	d1 = c42
	d2 = c43
	d3 = c44
	e0 = c51
	e1 = c52
	e2 = c53
	e3 = c54
	f = c61
	g = c62
	h = c63

	aa = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	c = a0 + b + (a3*t+a2)*t*t

	pl[3] = padata.PopulatePrecisePlanetData(
		"Mars", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	a0 = d11
	a1 = d12
	a2 = d13
	a3 = d14
	b0 = d21
	b1 = d22
	b2 = d23
	b3 = d24
	c0 = d31
	c1 = d32
	c2 = d33
	c3 = d34
	d0 = d41
	d1 = d42
	d2 = d43
	d3 = d44
	e0 = d51
	e1 = d52
	e2 = d53
	e3 = d54
	f = d61
	g = d62
	h = d63

	aa = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	c = a0 + b + (a3*t+a2)*t*t

	pl[4] = padata.PopulatePrecisePlanetData(
		"Jupiter", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	a0 = e11
	a1 = e12
	a2 = e13
	a3 = e14
	b0 = e21
	b1 = e22
	b2 = e23
	b3 = e24
	c0 = e31
	c1 = e32
	c2 = e33
	c3 = e34
	d0 = e41
	d1 = e42
	d2 = e43
	d3 = e44
	e0 = e51
	e1 = e52
	e2 = e53
	e3 = e54
	f = e61
	g = e62
	h = e63

	aa = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	c = a0 + b + (a3*t+a2)*t*t

	pl[5] = padata.PopulatePrecisePlanetData(
		"Saturn", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	a0 = f11
	a1 = f12
	a2 = f13
	a3 = f14
	b0 = f21
	b1 = f22
	b2 = f23
	b3 = f24
	c0 = f31
	c1 = f32a
	c2 = f33
	c3 = f34
	d0 = f41
	d1 = f42
	d2 = f43
	d3 = f44
	e0 = f51
	e1 = f52
	e2 = f53
	e3 = f54
	f = f61
	g = f62
	h = f63

	aa = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	c = a0 + b + (a3*t+a2)*t*t

	pl[6] = padata.PopulatePrecisePlanetData(
		"Uranus", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	a0 = g11
	a1 = g12
	a2 = g13
	a3 = g14
	b0 = g21
	b1 = g22
	b2 = g23
	b3 = g24
	c0 = g31
	c1 = g32
	c2 = g33
	c3 = g34
	d0 = g41
	d1 = g42
	d2 = g43
	d3 = g44
	e0 = g51
	e1 = g52
	e2 = g53
	e3 = g54
	f = g61
	g = g62
	h = g63

	aa = a1 * t
	b = 360.0 * (aa - math.Floor(aa))
	c = a0 + b + (a3*t+a2)*t*t

	pl[7] = padata.PopulatePrecisePlanetData(
		"Neptune", c-360.0*math.Floor(c/360.0), (a1*0.009856263)+(a2+a3)/36525.0, ((b3*t+b2)*t+b1)*t+b0, ((c3*t+c2)*t+c1)*t+c0,
		((d3*t+d2)*t+d1)*t+d0, ((e3*t+e2)*t+e1)*t+e0, f, g, h, 0,
	)

	var checkPlanet padata.PlanetDataPrecise = padata.GetPrecisePlanetData(s, pl[:])

	if checkPlanet.Name == "NOTFOUND" {
		return patype.PlanetCoordinates{
			Longitude: Degrees(Unwind(0)), Latitude: Degrees(Unwind(0)), DistanceAu: Degrees(Unwind(0)),
			HLong1: Degrees(Unwind(0)), HLong2: Degrees(Unwind(0)), HLat: Degrees(Unwind(0)), RVect: Degrees(Unwind(0)),
		}
	}

	var li float64 = 0.0
	var ms float64 = SunMeanAnomaly(lh, lm, ls, ds, zc, dy, mn, yr)
	var sr float64 = pautil.DegreesToRadians(SunLong(lh, lm, ls, ds, zc, dy, mn, yr))
	var re float64 = SunDist(lh, lm, ls, ds, zc, dy, mn, yr)
	var lg float64 = sr + math.Pi

	var l0 float64 = 0.0
	var s0 float64 = 0.0
	var p0 float64 = 0.0
	var vo float64 = 0.0
	var lp1 float64 = 0.0
	var ll float64 = 0.0
	var rd float64 = 0.0
	var pd float64 = 0.0
	var sp float64 = 0.0
	var ci float64 = 0.0

	for k := 1; k <= 3; k++ {
		for i := range len(pl) {
			pl[i].ApValue = pautil.DegreesToRadians(pl[i].Value1 - pl[i].Value3 - li*pl[i].Value2)
		}

		var qa float64 = 0.0
		var qb float64 = 0.0
		var qc float64 = 0.0
		var qd float64 = 0.0
		var qe float64 = 0.0
		var qf float64 = 0.0
		var qg float64 = 0.0

		if s == "Mercury" {
			var tempResult patype.PlanetLongLatL4685 = PlanetLongL4685(pl[:])

			qa = tempResult.QA
			qb = tempResult.QB
		}

		if s == "Venus" {
			var tempResult patype.PlanetLongLatL4735 = PlanetLongL4735(pl[:], ms, t)

			qa = tempResult.QA
			qb = tempResult.QB
			qc = tempResult.QC
			qe = tempResult.QE
		}

		if s == "Mars" {
			var tempResult patype.PlanetLongLatL4810 = PlanetLongL4810(pl[:], ms)

			qc = tempResult.QC
			qe = tempResult.QE
			qa = tempResult.QA
			qb = tempResult.QB
		}

		var matchPlanet padata.PlanetDataPrecise = padata.GetPrecisePlanetData(s, pl[:])

		if s == "Jupiter" || s == "Saturn" || s == "Uranus" || s == "Neptune" {
			var tempResult patype.PlanetLongLatL4945 = PlanetLongL4945(t, matchPlanet)

			qa = tempResult.QA
			qb = tempResult.QB
			qc = tempResult.QC
			qd = tempResult.QD
			qe = tempResult.QE
			qf = tempResult.QF
			qg = tempResult.QG
		}

		var ec float64 = matchPlanet.Value4 + qd
		var am float64 = matchPlanet.ApValue + qe
		var at float64 = TrueAnomaly(am, ec)
		var pvv float64 = (matchPlanet.Value7 + qf) * (1.0 - ec*ec) / (1.0 + ec*math.Cos(at))
		var lp float64 = Degrees(at) + matchPlanet.Value3 + Degrees(qc-qe)
		lp = pautil.DegreesToRadians(lp)
		var om float64 = pautil.DegreesToRadians(matchPlanet.Value6)
		var lo float64 = lp - om
		var so float64 = math.Sin(lo)
		var co float64 = math.Cos(lo)
		var inn float64 = pautil.DegreesToRadians(matchPlanet.Value5)
		pvv += qb
		sp = so * math.Sin(inn)
		var y float64 = so * math.Cos(inn)
		var ps float64 = math.Asin(sp) + qg
		sp = math.Sin(ps)
		pd = math.Atan2(y, co) + om + pautil.DegreesToRadians(qa)
		pd = Unwind(pd)
		ci = math.Cos(ps)
		rd = pvv * ci
		ll = pd - lg
		var rh float64 = re*re + pvv*pvv - 2.0*re*pvv*ci*math.Cos(ll)
		rh = math.Sqrt(rh)
		li = rh * 0.005775518

		if k == 1 {
			l0 = pd
			s0 = ps
			p0 = pvv
			vo = rh
			lp1 = lp
		}
	}

	var l1 float64 = math.Sin(ll)
	var l2 float64 = math.Cos(ll)

	var ep float64 = 0
	if ip < 3 {
		ep = math.Atan(-1.0*rd*l1/(re-rd*l2)) + lg + math.Pi
	} else {
		ep = math.Atan(re*l1/(rd-re*l2)) + pd
	}

	ep = Unwind(ep)

	var bp float64 = math.Atan(rd * sp * math.Sin(ep-pd) / (ci * re * l1))

	var planetLongitude float64 = Degrees(Unwind(ep))
	var planetLatitude float64 = Degrees(Unwind(bp))
	var planetDistanceAU float64 = vo
	var planetHLong1 float64 = Degrees(lp1)
	var planetHLong2 float64 = Degrees(l0)
	var planetHLat float64 = Degrees(s0)
	var planetRVect float64 = p0

	return patype.PlanetCoordinates{
		Longitude: planetLongitude, Latitude: planetLatitude, DistanceAu: planetDistanceAU,
		HLong1: planetHLong1, HLong2: planetHLong2, HLat: planetHLat, RVect: planetRVect,
	}
}

/* Helper function for PlanetCoordinates() */
func PlanetLongL4685(pl []padata.PlanetDataPrecise) patype.PlanetLongLatL4685 {
	var qa float64 = 0.00204 * math.Cos(5.0*pl[2].ApValue-2.0*pl[1].ApValue+0.21328)
	qa += 0.00103 * math.Cos(2.0*pl[2].ApValue-pl[1].ApValue-2.8046)
	qa += 0.00091 * math.Cos(2.0*pl[4].ApValue-pl[1].ApValue-0.64582)
	qa += 0.00078 * math.Cos(5.0*pl[2].ApValue-3.0*pl[1].ApValue+0.17692)

	var qb float64 = 0.000007525 * math.Cos(2.0*pl[4].ApValue-pl[1].ApValue+0.925251)
	qb += 0.000006802 * math.Cos(5.0*pl[2].ApValue-3.0*pl[1].ApValue-4.53642)
	qb += 0.000005457 * math.Cos(2.0*pl[2].ApValue-2.0*pl[1].ApValue-1.24246)
	qb += 0.000003569 * math.Cos(5.0*pl[2].ApValue-pl[1].ApValue-1.35699)

	return patype.PlanetLongLatL4685{QA: qa, QB: qb}
}

/* Helper function for PlanetCoordinates() */
func PlanetLongL4735(pl []padata.PlanetDataPrecise, ms float64, t float64) patype.PlanetLongLatL4735 {
	var qc float64 = 0.00077 * math.Sin(4.1406+t*2.6227)
	qc = pautil.DegreesToRadians(qc)
	var qe float64 = qc

	var qa float64 = 0.00313 * math.Cos(2.0*ms-2.0*pl[2].ApValue-2.587)
	qa += 0.00198 * math.Cos(3.0*ms-3.0*pl[2].ApValue+0.044768)
	qa += 0.00136 * math.Cos(ms-pl[2].ApValue-2.0788)
	qa += 0.00096 * math.Cos(3.0*ms-2.0*pl[2].ApValue-2.3721)
	qa += 0.00082 * math.Cos(pl[4].ApValue-pl[2].ApValue-3.6318)

	var qb float64 = 0.000022501 * math.Cos(2.0*ms-2.0*pl[2].ApValue-1.01592)
	qb += 0.000019045 * math.Cos(3.0*ms-3.0*pl[2].ApValue+1.61577)
	qb += 0.000006887 * math.Cos(pl[4].ApValue-pl[2].ApValue-2.06106)
	qb += 0.000005172 * math.Cos(ms-pl[2].ApValue-0.508065)
	qb += 0.00000362 * math.Cos(5.0*ms-4.0*pl[2].ApValue-1.81877)
	qb += 0.000003283 * math.Cos(4.0*ms-4.0*pl[2].ApValue+1.10851)
	qb += 0.000003074 * math.Cos(2.0*pl[4].ApValue-2.0*pl[2].ApValue-0.962846)

	return patype.PlanetLongLatL4735{QA: qa, QB: qb, QC: qc, QE: qe}
}

/* Helper function for PlanetCoordinates() */
func PlanetLongL4810(pl []padata.PlanetDataPrecise, ms float64) patype.PlanetLongLatL4810 {
	var a float64 = 3.0*pl[4].ApValue - 8.0*pl[3].ApValue + 4.0*ms
	var sa float64 = math.Sin(a)
	var ca float64 = math.Cos(a)
	var qc float64 = -(0.01133*sa + 0.00933*ca)
	qc = pautil.DegreesToRadians(qc)
	var qe float64 = qc

	var qa float64 = 0.00705 * math.Cos(pl[4].ApValue-pl[3].ApValue-0.85448)
	qa += 0.00607 * math.Cos(2.0*pl[4].ApValue-pl[3].ApValue-3.2873)
	qa += 0.00445 * math.Cos(2.0*pl[4].ApValue-2.0*pl[3].ApValue-3.3492)
	qa += 0.00388 * math.Cos(ms-2.0*pl[3].ApValue+0.35771)
	qa += 0.00238 * math.Cos(ms-pl[3].ApValue+0.61256)
	qa += 0.00204 * math.Cos(2.0*ms-3.0*pl[3].ApValue+2.7688)
	qa += 0.00177 * math.Cos(3.0*pl[3].ApValue-pl[2].ApValue-1.0053)
	qa += 0.00136 * math.Cos(2.0*ms-4.0*pl[3].ApValue+2.6894)
	qa += 0.00104 * math.Cos(pl[4].ApValue+0.30749)

	var qb float64 = 0.000053227 * math.Cos(pl[4].ApValue-pl[3].ApValue+0.717864)
	qb += 0.000050989 * math.Cos(2.0*pl[4].ApValue-2.0*pl[3].ApValue-1.77997)
	qb += 0.000038278 * math.Cos(2.0*pl[4].ApValue-pl[3].ApValue-1.71617)
	qb += 0.000015996 * math.Cos(ms-pl[3].ApValue-0.969618)
	qb += 0.000014764 * math.Cos(2.0*ms-3.0*pl[3].ApValue+1.19768)
	qb += 0.000008966 * math.Cos(pl[4].ApValue-2.0*pl[3].ApValue+0.761225)
	qb += 0.000007914 * math.Cos(3.0*pl[4].ApValue-2.0*pl[3].ApValue-2.43887)
	qb += 0.000007004 * math.Cos(2.0*pl[4].ApValue-3.0*pl[3].ApValue-1.79573)
	qb += 0.00000662 * math.Cos(ms-2.0*pl[3].ApValue+1.97575)
	qb += 0.00000493 * math.Cos(3.0*pl[4].ApValue-3.0*pl[3].ApValue-1.33069)
	qb += 0.000004693 * math.Cos(3.0*ms-5.0*pl[3].ApValue+3.32665)
	qb += 0.000004571 * math.Cos(2.0*ms-4.0*pl[3].ApValue+4.27086)
	qb += 0.000004409 * math.Cos(3.0*pl[4].ApValue-pl[3].ApValue-2.02158)

	return patype.PlanetLongLatL4810{A: a, SA: sa, CA: ca, QC: qc, QE: qe, QA: qa, QB: qb}
}

/* Helper function for PlanetCoordinates() */
func PlanetLongL4945(t float64, planet padata.PlanetDataPrecise) patype.PlanetLongLatL4945 {
	var qa float64 = 0.0
	var qb float64 = 0.0
	var qc float64 = 0.0
	var qd float64 = 0.0
	var qe float64 = 0.0
	var qf float64 = 0.0
	var qg float64 = 0.0
	var vk float64 = 0.0
	var ja float64 = 0.0
	var jb float64 = 0.0
	var jc float64 = 0.0

	var j1 float64 = t/5.0 + 0.1
	var j2 float64 = Unwind(4.14473 + 52.9691*t)
	var j3 float64 = Unwind(4.641118 + 21.32991*t)
	var j4 float64 = Unwind(4.250177 + 7.478172*t)
	var j5 float64 = 5.0*j3 - 2.0*j2
	var j6 float64 = 2.0*j2 - 6.0*j3 + 3.0*j4

	if planet.Name == "Mercury" || planet.Name == "Venus" || planet.Name == "Mars" {
		return patype.PlanetLongLatL4945{QA: qa, QB: qb, QC: qc, QD: qd, QE: qe, QF: qf, QG: qg}
	}

	if planet.Name == "Jupiter" || planet.Name == "Saturn" {
		var j7 float64 = j3 - j2
		var u1 float64 = math.Sin(j3)
		var u2 float64 = math.Cos(j3)
		var u3 float64 = math.Sin(2.0 * j3)
		var u4 float64 = math.Cos(2.0 * j3)
		var u5 float64 = math.Sin(j5)
		var u6 float64 = math.Cos(j5)
		var u7 float64 = math.Sin(2.0 * j5)
		var u8a float64 = math.Sin(j6)
		var u9 float64 = math.Sin(j7)
		var ua float64 = math.Cos(j7)
		var ub float64 = math.Sin(2.0 * j7)
		var uc float64 = math.Cos(2.0 * j7)
		var ud float64 = math.Sin(3.0 * j7)
		var ue float64 = math.Cos(3.0 * j7)
		var uf float64 = math.Sin(4.0 * j7)
		var ug float64 = math.Cos(4.0 * j7)
		var vh float64 = math.Cos(5.0 * j7)

		if planet.Name == "Saturn" {
			var ui float64 = math.Sin(3.0 * j3)
			var uj float64 = math.Cos(3.0 * j3)
			var uk float64 = math.Sin(4.0 * j3)
			var ul float64 = math.Cos(4.0 * j3)
			var vi float64 = math.Cos(2.0 * j5)
			var un float64 = math.Sin(5.0 * j7)
			var j8 float64 = j4 - j3
			var uo float64 = math.Sin(2.0 * j8)
			var up float64 = math.Cos(2.0 * j8)
			var uq float64 = math.Sin(3.0 * j8)
			var ur float64 = math.Cos(3.0 * j8)

			qc = 0.007581*u7 - 0.007986*u8a - 0.148811*u9
			qc -= (0.814181 - (0.01815-0.016714*j1)*j1) * u5
			qc -= (0.010497 - (0.160906-0.0041*j1)*j1) * u6
			qc = qc - 0.015208*ud - 0.006339*uf - 0.006244*u1
			qc = qc - 0.0165*ub*u1 - 0.040786*ub
			qc = qc + (0.008931+0.002728*j1)*u9*u1 - 0.005775*ud*u1
			qc = qc + (0.081344+0.003206*j1)*ua*u1 + 0.015019*uc*u1
			qc = qc + (0.085581+0.002494*j1)*u9*u2 + 0.014394*uc*u2
			qc = qc + (0.025328-0.003117*j1)*ua*u2 + 0.006319*ue*u2
			qc = qc + 0.006369*u9*u3 + 0.009156*ub*u3 + 0.007525*uq*u3
			qc = qc - 0.005236*ua*u4 - 0.007736*uc*u4 - 0.007528*ur*u4
			qc = pautil.DegreesToRadians(qc)

			qd = (-7927.0 + (2548.0+91.0*j1)*j1) * u5
			qd = qd + (13381.0+(1226.0-253.0*j1)*j1)*u6 + (248.0-121.0*j1)*u7
			qd = qd - (305.0+91.0*j1)*vi + 412.0*ub + 12415.0*u1
			qd = qd + (390.0-617.0*j1)*u9*u1 + (165.0-204.0*j1)*ub*u1
			qd = qd + 26599.0*ua*u1 - 4687.0*uc*u1 - 1870.0*ue*u1 - 821.0*ug*u1
			qd = qd - 377.0*vh*u1 + 497.0*up*u1 + (163.0-611.0*j1)*u2
			qd = qd - 12696.0*u9*u2 - 4200.0*ub*u2 - 1503.0*ud*u2 - 619.0*uf*u2
			qd = qd - 268.0*un*u2 - (282.0+1306.0*j1)*ua*u2
			qd = qd + (-86.0+230.0*j1)*uc*u2 + 461.0*uo*u2 - 350.0*u3
			qd = qd + (2211.0-286.0*j1)*u9*u3 - 2208.0*ub*u3 - 568.0*ud*u3
			qd = qd - 346.0*uf*u3 - (2780.0+222.0*j1)*ua*u3
			qd = qd + (2022.0+263.0*j1)*uc*u3 + 248.0*ue*u3 + 242.0*uq*u3
			qd = qd + 467.0*ur*u3 - 490.0*u4 - (2842.0+279.0*j1)*u9*u4
			qd = qd + (128.0+226.0*j1)*ub*u4 + 224.0*ud*u4
			qd = qd + (-1594.0+282.0*j1)*ua*u4 + (2162.0-207.0*j1)*uc*u4
			qd = qd + 561.0*ue*u4 + 343.0*ug*u4 + 469.0*uq*u4 - 242.0*ur*u4
			qd = qd - 205.0*u9*ui + 262.0*ud*ui + 208.0*ua*uj - 271.0*ue*uj
			qd = qd - 382.0*ue*uk - 376.0*ud*ul
			qd *= 0.0000001

			vk = (0.077108 + (0.007186-0.001533*j1)*j1) * u5
			vk -= 0.007075 * u9
			vk += (0.045803 - (0.014766+0.000536*j1)*j1) * u6
			vk = vk - 0.072586*u2 - 0.075825*u9*u1 - 0.024839*ub*u1
			vk = vk - 0.008631*ud*u1 - 0.150383*ua*u2
			vk = vk + 0.026897*uc*u2 + 0.010053*ue*u2
			vk = vk - (0.013597+0.001719*j1)*u9*u3 + 0.011981*ub*u4
			vk -= (0.007742 - 0.001517*j1) * ua * u3
			vk += (0.013586 - 0.001375*j1) * uc * u3
			vk -= (0.013667 - 0.001239*j1) * u9 * u4
			vk += (0.014861 + 0.001136*j1) * ua * u4
			vk -= (0.013064 + 0.001628*j1) * uc * u4
			qe = qc - (pautil.DegreesToRadians(vk) / planet.Value4)

			qf = 572.0*u5 - 1590.0*ub*u2 + 2933.0*u6 - 647.0*ud*u2
			qf = qf + 33629.0*ua - 344.0*uf*u2 - 3081.0*uc + 2885.0*ua*u2
			qf = qf - 1423.0*ue + (2172.0+102.0*j1)*uc*u2 - 671.0*ug
			qf = qf + 296.0*ue*u2 - 320.0*vh - 267.0*ub*u3 + 1098.0*u1
			qf = qf - 778.0*ua*u3 - 2812.0*u9*u1 + 495.0*uc*u3 + 688.0*ub*u1
			qf = qf + 250.0*ue*u3 - 393.0*ud*u1 - 856.0*u9*u4 - 228.0*uf*u1
			qf = qf + 441.0*ub*u4 + 2138.0*ua*u1 + 296.0*uc*u4 - 999.0*uc*u1
			qf = qf + 211.0*ue*u4 - 642.0*ue*u1 - 427.0*u9*ui - 325.0*ug*u1
			qf = qf + 398.0*ud*ui - 890.0*u2 + 344.0*ua*uj + 2206.0*u9*u2
			qf -= 427.0 * ue * uj
			qf *= 0.000001

			qg = 0.000747*ua*u1 + 0.001069*ua*u2 + 0.002108*ub*u3
			qg = qg + 0.001261*uc*u3 + 0.001236*ub*u4 - 0.002075*uc*u4
			qg = pautil.DegreesToRadians(qg)

			return patype.PlanetLongLatL4945{QA: qa, QB: qb, QC: qc, QD: qd, QE: qe, QF: qf, QG: qg}
		}

		qc = (0.331364 - (0.010281+0.004692*j1)*j1) * u5
		qc += (0.003228 - (0.064436-0.002075*j1)*j1) * u6
		qc -= (0.003083 + (0.000275-0.000489*j1)*j1) * u7
		qc = qc + 0.002472*u8a + 0.013619*u9 + 0.018472*ub
		qc = qc + 0.006717*ud + 0.002775*uf + 0.006417*ub*u1
		qc = qc + (0.007275-0.001253*j1)*u9*u1 + 0.002439*ud*u1
		qc = qc - (0.035681+0.001208*j1)*u9*u2 - 0.003767*uc*u1
		qc = qc - (0.033839+0.001125*j1)*ua*u1 - 0.004261*ub*u2
		qc = qc + (0.001161*j1-0.006333)*ua*u2 + 0.002178*u2
		qc = qc - 0.006675*uc*u2 - 0.002664*ue*u2 - 0.002572*u9*u3
		qc = qc - 0.003567*ub*u3 + 0.002094*ua*u4 + 0.003342*uc*u4
		qc = pautil.DegreesToRadians(qc)

		qd = (3606.0+(130.0-43.0*j1)*j1)*u5 + (1289.0-580.0*j1)*u6
		qd = qd - 6764.0*u9*u1 - 1110.0*ub*u1 - 224.0*ud*u1 - 204.0*u1
		qd = qd + (1284.0+116.0*j1)*ua*u1 + 188.0*uc*u1
		qd = qd + (1460.0+130.0*j1)*u9*u2 + 224.0*ub*u2 - 817.0*u2
		qd = qd + 6074.0*u2*ua + 992.0*uc*u2 + 508.0*ue*u2 + 230.0*ug*u2
		qd = qd + 108.0*vh*u2 - (956.0+73.0*j1)*u9*u3 + 448.0*ub*u3
		qd = qd + 137.0*ud*u3 + (108.0*j1-997.0)*ua*u3 + 480.0*uc*u3
		qd = qd + 148.0*ue*u3 + (99.0*j1-956.0)*u9*u4 + 490.0*ub*u4
		qd = qd + 158.0*ud*u4 + 179.0*u4 + (1024.0+75.0*j1)*ua*u4
		qd = qd - 437.0*uc*u4 - 132.0*ue*u4
		qd *= 0.0000001

		vk = (0.007192-0.003147*j1)*u5 - 0.004344*u1
		vk += (j1*(0.000197*j1-0.000675) - 0.020428) * u6
		vk = vk + 0.034036*ua*u1 + (0.007269+0.000672*j1)*u9*u1
		vk = vk + 0.005614*uc*u1 + 0.002964*ue*u1 + 0.037761*u9*u2
		vk = vk + 0.006158*ub*u2 - 0.006603*ua*u2 - 0.005356*u9*u3
		vk = vk + 0.002722*ub*u3 + 0.004483*ua*u3
		vk = vk - 0.002642*uc*u3 + 0.004403*u9*u4
		vk = vk - 0.002536*ub*u4 + 0.005547*ua*u4 - 0.002689*uc*u4
		qe = qc - (pautil.DegreesToRadians(vk) / planet.Value4)

		qf = 205.0*ua - 263.0*u6 + 693.0*uc + 312.0*ue + 147.0*ug + 299.0*u9*u1
		qf = qf + 181.0*uc*u1 + 204.0*ub*u2 + 111.0*ud*u2 - 337.0*ua*u2
		qf -= 111.0 * uc * u2
		qf *= 0.000001

		return patype.PlanetLongLatL4945{QA: qa, QB: qb, QC: qc, QD: qd, QE: qe, QF: qf, QG: qg}
	}

	if planet.Name == "Uranus" || planet.Name == "Neptune" {
		var j8 float64 = Unwind(1.46205 + 3.81337*t)
		var j9 float64 = 2.0*j8 - j4
		var vj float64 = math.Sin(j9)
		var uu float64 = math.Cos(j9)
		var uv float64 = math.Sin(2.0 * j9)
		var uw float64 = math.Cos(2.0 * j9)

		if planet.Name == "Neptune" {
			ja = j8 - j2
			jb = j8 - j3
			jc = j8 - j4
			qc = (0.001089*j1 - 0.589833) * vj
			qc = qc + (0.004658*j1-0.056094)*uu - 0.024286*uv
			qc = pautil.DegreesToRadians(qc)

			vk = 0.024039*vj - 0.025303*uu + 0.006206*uv
			vk -= 0.005992 * uw
			qe = qc - (pautil.DegreesToRadians(vk) / planet.Value4)

			qd = 4389.0*vj + 1129.0*uv + 4262.0*uu + 1089.0*uw
			qd *= 0.0000001

			qf = 8189.0*uu - 817.0*vj + 781.0*uw
			qf *= 0.000001

			var vd float64 = math.Sin(2.0 * jc)
			var ve float64 = math.Cos(2.0 * jc)
			var vf float64 = math.Sin(j8)
			var vg float64 = math.Cos(j8)
			qa = -0.009556*math.Sin(ja) - 0.005178*math.Sin(jb)
			qa = qa + 0.002572*vd - 0.002972*ve*vf - 0.002833*vd*vg

			qg = 0.000336*ve*vf + 0.000364*vd*vg
			qg = pautil.DegreesToRadians(qg)

			qb = -40596.0 + 4992.0*math.Cos(ja) + 2744.0*math.Cos(jb)
			qb = qb + 2044.0*math.Cos(jc) + 1051.0*ve
			qb *= 0.000001

			return patype.PlanetLongLatL4945{QA: qa, QB: qb, QC: qc, QD: qd, QE: qe, QF: qf, QG: qg}
		}

		ja = j4 - j2
		jb = j4 - j3
		jc = j8 - j4
		qc = (0.864319 - 0.001583*j1) * vj
		qc = qc + (0.082222-0.006833*j1)*uu + 0.036017*uv
		qc = qc - 0.003019*uw + 0.008122*math.Sin(j6)
		qc = pautil.DegreesToRadians(qc)

		vk = 0.120303*vj + 0.006197*uv
		vk += (0.019472 - 0.000947*j1) * uu
		qe = qc - (pautil.DegreesToRadians(vk) / planet.Value4)

		qd = (163.0*j1-3349.0)*vj + 20981.0*uu + 1311.0*uw
		qd *= 0.0000001

		qf = -0.003825 * uu

		qa = (-0.038581 + (0.002031-0.00191*j1)*j1) * math.Cos(j4+jb)
		qa += (0.010122 - 0.000988*j1) * math.Sin(j4+jb)
		var a float64 = (0.034964 - (0.001038-0.000868*j1)*j1) * math.Cos(2.0*j4+jb)
		qa = a + qa + 0.005594*math.Sin(j4+3.0*jc) - 0.014808*math.Sin(ja)
		qa = qa - 0.005794*math.Sin(jb) + 0.002347*math.Cos(jb)
		qa = qa + 0.009872*math.Sin(jc) + 0.008803*math.Sin(2.0*jc)
		qa -= 0.004308 * math.Sin(3.0*jc)

		var ux float64 = math.Sin(jb)
		var uy float64 = math.Cos(jb)
		var uz float64 = math.Sin(j4)
		var va float64 = math.Cos(j4)
		var vb float64 = math.Sin(2.0 * j4)
		var vc float64 = math.Cos(2.0 * j4)
		qg = (0.000458*ux - 0.000642*uy - 0.000517*math.Cos(4.0*jc)) * uz
		qg -= (0.000347*ux + 0.000853*uy + 0.000517*math.Sin(4.0*jb)) * va
		qg += 0.000403 * (math.Cos(2.0*jc)*vb + math.Sin(2.0*jc)*vc)
		qg = pautil.DegreesToRadians(qg)

		qb = -25948.0 + 4985.0*math.Cos(ja) - 1230.0*va + 3354.0*uy
		qb = qb + 904.0*math.Cos(2.0*jc) + 894.0*(math.Cos(jc)-math.Cos(3.0*jc))
		qb += (5795.0*va - 1165.0*uz + 1388.0*vc) * ux
		qb += (1351.0*va + 5702.0*uz + 1388.0*vb) * uy
		qb *= 0.000001

		return patype.PlanetLongLatL4945{QA: qa, QB: qb, QC: qc, QD: qd, QE: qe, QF: qf, QG: qg}
	}

	return patype.PlanetLongLatL4945{QA: qa, QB: qb, QC: qc, QD: qd, QE: qe, QF: qf, QG: qg}
}

/*
Calculate longitude, latitude, and distance of parabolic-orbit comet.

	Original macro names: PcometLong, PcometLat, PcometDist
*/
func PCometLongLatDist(lh float64, /* Local civil time, hour part. */
	lm float64, /* Local civil time, minutes part. */
	ls float64, /* Local civil time, seconds part. */
	ds int, /* Daylight Savings offset. */
	zc int, /* Time zone correction, in hours. */
	dy float64, /* Local date, day part. */
	mn int, /* Local date, month part. */
	yr int, /* Local date, year part. */
	td float64, /* Perihelion epoch (day) */
	tm int, /* Perihelion epoch (month) */
	ty int, /* Perihelion epoch (year) */
	q float64, /* a (AU) */
	i float64, /* Inclination (degrees) */
	p float64, /* Perihelion (degrees) */
	n float64, /* Node (degrees) */
) patype.CometLongLatDist {
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var ut float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var tpe float64 = (ut / 365.242191) + CivilDateToJulianDate(gd, float64(gm), float64(gy)) - CivilDateToJulianDate(td, float64(tm), float64(ty))
	var lg float64 = pautil.DegreesToRadians(SunLong(lh, lm, ls, ds, zc, dy, mn, yr) + 180.0)
	var re float64 = SunDist(lh, lm, ls, ds, zc, dy, mn, yr)

	var rh2 float64 = 0.0
	var rd float64 = 0.0
	var s3 float64 = 0.0
	var c3 float64 = 0.0
	var lc float64 = 0.0
	var s2 float64 = 0.0
	var c2 float64 = 0.0

	for k := 1; k < 3; k++ {
		var s float64 = SolveCubic(0.0364911624 * tpe / (q * math.Sqrt(q)))
		var nu float64 = 2.0 * math.Atan(s)
		var r float64 = q * (1.0 + s*s)
		var l float64 = nu + pautil.DegreesToRadians(p)
		var s1 float64 = math.Sin(l)
		var c1 float64 = math.Cos(l)
		var i1 float64 = pautil.DegreesToRadians(i)
		s2 = s1 * math.Sin(i1)
		var ps float64 = math.Asin(s2)
		var y float64 = s1 * math.Cos(i1)
		lc = math.Atan2(y, c1) + pautil.DegreesToRadians(n)
		c2 = math.Cos(ps)
		rd = r * c2
		var ll float64 = lc - lg
		c3 = math.Cos(ll)
		s3 = math.Sin(ll)
		// var rh float64 = math.Sqrt((re * re) + (r * r) - (2.0 * re * rd * c3 * math.Cos(ps)))  // not used?
		if k == 1 {
			rh2 = math.Sqrt((re * re) + (r * r) - (2.0 * re * r * math.Cos(ps) * math.Cos(l+pautil.DegreesToRadians(n)-lg)))
		}
	}

	var ep float64

	if rd < re {
		ep = math.Atan(-rd*s3/(re-(rd*c3))) + lg + 3.141592654
	} else {
		ep = math.Atan(re*s3/(rd-(re*c3))) + lc
	}

	ep = Unwind(ep)

	var tb float64 = rd * s2 * math.Sin(ep-lc) / (c2 * re * s3)
	var bp float64 = math.Atan(tb)

	var comet_long_deg float64 = Degrees(ep)
	var comet_lat_deg float64 = Degrees(bp)
	var comet_dist_au float64 = rh2

	return patype.CometLongLatDist{LongDeg: comet_long_deg, LatDeg: comet_lat_deg, DistAu: comet_dist_au}
}

/*
For W, in radians, return S, also in radians.

Original macro name: SolveCubic
*/
func SolveCubic(w float64) float64 {
	var s float64 = w / 3.0

	for 1 == 1 {
		var s2 float64 = s * s
		var d float64 = (s2+3.0)*s - w

		if math.Abs(d) < 0.000001 {
			return s
		}

		s = ((2.0 * s * s2) + w) / (3.0 * (s2 + 1.0))
	}

	return 0
}
