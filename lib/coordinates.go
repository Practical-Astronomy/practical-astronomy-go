package lib

import (
	"math"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Convert Angle (degrees,minutes,seconds) to Decimal Degrees */
func AngleToDecimalDegrees(degrees float64, minutes float64, seconds float64) float64 {
	var a float64 = math.Abs(seconds) / 60
	var b float64 = (math.Abs(minutes) + a) / 60
	var c float64 = math.Abs(degrees) + b

	var d float64
	if degrees < 0 || minutes < 0 || seconds < 0 {
		d = -c
	} else {
		d = c
	}

	return d
}

/* Convert Decimal Degrees to an Angle (degrees, minutes, and seconds) */
func DecimalDegreesToAngle(decimalDegrees float64) patype.Angle {
	var unsignedDecimal float64 = math.Abs(decimalDegrees)
	var totalSeconds float64 = unsignedDecimal * 3600
	var seconds2DP float64 = pautil.RoundTo(math.Mod(totalSeconds, 60), 2)

	var correctedSeconds float64
	if seconds2DP == 60 {
		correctedSeconds = 0
	} else {
		correctedSeconds = seconds2DP
	}

	var correctedRemainder float64
	if seconds2DP == 60 {
		correctedRemainder = totalSeconds + 60
	} else {
		correctedRemainder = totalSeconds
	}

	var minutes float64 = math.Mod(math.Floor(correctedRemainder/60), 60)
	var unsignedDegrees float64 = math.Floor(correctedRemainder / 3600)

	var signedDegrees float64
	if decimalDegrees < 0 {
		signedDegrees = -1 * unsignedDegrees
	} else {
		signedDegrees = unsignedDegrees
	}

	return patype.Angle{Degrees: signedDegrees, Minutes: minutes, Seconds: math.Floor(correctedSeconds)}
}

/* Convert Right Ascension to Hour Angle */
func RightAscensionToHourAngle(
	raHours float64, raMinutes float64, raSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64,
	isDaylightSavings bool, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64,
) patype.HourAngle {
	var daylightSaving int
	if isDaylightSavings {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var hourAngle float64 = pamacro.RightAscensionToHourAngle(raHours, raMinutes, raSeconds, lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear, geographicalLongitude)

	var hourAngleHours int = pamacro.DecimalHoursHour(hourAngle)
	var hourAngleMinutes int = pamacro.DecimalHoursMinute(hourAngle)
	var hourAngleSeconds float64 = pamacro.DecimalHoursSecond(hourAngle)

	return patype.HourAngle{Hours: float64(hourAngleHours), Minutes: float64(hourAngleMinutes), Seconds: hourAngleSeconds}
}

/* Convert Hour Angle to Right Ascension */
func HourAngleToRightAscension(
	hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64, lctHours float64, lctMinutes float64, lctSeconds float64,
	isDaylightSavings bool, zoneCorrection int, localDay float64, localMonth int, localYear int, geographicalLongitude float64,
) patype.RightAscension {
	var daylightSaving int
	if isDaylightSavings {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var rightAscension float64 = pamacro.HourAngleToRightAscension(hourAngleHours, hourAngleMinutes, hourAngleSeconds, lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear, geographicalLongitude)

	var rightAscensionHours int = pamacro.DecimalHoursHour(rightAscension)
	var rightAscensionMinutes int = pamacro.DecimalHoursMinute(rightAscension)
	var rightAscensionSeconds int = int(pamacro.DecimalHoursSecond(rightAscension))

	return patype.RightAscension{Hours: float64(rightAscensionHours), Minutes: float64(rightAscensionMinutes), Seconds: float64(rightAscensionSeconds)}
}

/* Convert Equatorial Coordinates to Horizon Coordinates */
func EquatorialCoordinatesToHorizonCoordinates(
	hourAngleHours float64, hourAngleMinutes float64, hourAngleSeconds float64,
	declinationDegrees float64, declinationMinutes float64, declinationSeconds float64,
	geographicalLatitude float64,
) patype.HorizonCoordinates {
	var azimuthInDecimalDegrees float64 = pamacro.EquatorialCoordinatesToAzimuth(hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds, geographicalLatitude)
	var altitudeInDecimalDegrees float64 = pamacro.EquatorialCoordinatestoAltitude(hourAngleHours, hourAngleMinutes, hourAngleSeconds, declinationDegrees, declinationMinutes, declinationSeconds, geographicalLatitude)

	var azimuthDegrees float64 = pamacro.DecimalDegreesDegrees(azimuthInDecimalDegrees)
	var azimuthMinutes float64 = pamacro.DecimalDegreesMinutes(azimuthInDecimalDegrees)
	var azimuthSeconds float64 = pamacro.DecimalDegreesSeconds(azimuthInDecimalDegrees)

	var altitudeDegrees float64 = pamacro.DecimalDegreesDegrees(altitudeInDecimalDegrees)
	var altitudeMinutes float64 = pamacro.DecimalDegreesMinutes(altitudeInDecimalDegrees)
	var altitudeSeconds float64 = pamacro.DecimalDegreesSeconds(altitudeInDecimalDegrees)

	return patype.HorizonCoordinates{
		AzimuthDegrees: azimuthDegrees, AzimuthMinutes: azimuthMinutes, AzimuthSeconds: azimuthSeconds,
		AltitudeDegrees: altitudeDegrees, AltitudeMinutes: altitudeMinutes, AltitudeSeconds: altitudeSeconds,
	}
}

/* Convert Horizon Coordinates to Equatorial Coordinates */
func HorizonCoordinatesToEquatorialCoordinates(
	azimuthDegrees float64, azimuthMinutes float64, azimuthSeconds float64,
	altitudeDegrees float64, altitudeMinutes float64, altitudeSeconds float64,
	geographicalLatitude float64,
) patype.EquatorialCoordinates {
	var hour_angle_in_decimal_degrees float64 = pamacro.HorizonCoordinatesToHourAngle(
		azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds, geographicalLatitude)

	var declination_in_decimal_degrees float64 = pamacro.HorizonCoordinatesToDeclination(
		azimuthDegrees, azimuthMinutes, azimuthSeconds, altitudeDegrees, altitudeMinutes, altitudeSeconds, geographicalLatitude)

	var hour_angle_hours int = pamacro.DecimalHoursHour(hour_angle_in_decimal_degrees)
	var hour_angle_minutes int = pamacro.DecimalHoursMinute(hour_angle_in_decimal_degrees)
	var hour_angle_seconds float64 = pamacro.DecimalDegreesSeconds(hour_angle_in_decimal_degrees)

	var declination_degrees float64 = pamacro.DecimalDegreesDegrees(declination_in_decimal_degrees)
	var declination_minutes float64 = pamacro.DecimalDegreesMinutes(declination_in_decimal_degrees)
	var declination_seconds float64 = pamacro.DecimalDegreesSeconds(declination_in_decimal_degrees)

	return patype.EquatorialCoordinates{
		HourAngleHours: float64(hour_angle_hours), HourAngleMinutes: float64(hour_angle_minutes), HourAngleSeconds: hour_angle_seconds,
		DeclinationDegrees: declination_degrees, DeclinationMinutes: declination_minutes, DeclinationSeconds: declination_seconds,
	}
}

/* Calculate Mean Obliquity of the Ecliptic for a Greenwich Date */
func MeanObliquityOfTheEcliptic(greenwichDay float64, greenwichMonth int, greenwichYear int) float64 {
	var jd float64 = pamacro.CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear))
	var mjd float64 = jd - 2451545
	var t float64 = mjd / 36525
	var de1 float64 = t * (46.815 + t*(0.0006-(t*0.00181)))
	var de2 float64 = de1 / 3600

	return 23.439292 - de2
}

/* Convert Ecliptic Coordinates to Equatorial Coordinates */
func EclipticCoordinatesToEquatorialCoordinates(
	eclipticLongitudeDegrees float64, eclipticLongitudeMinutes float64, eclipticLongitudeSeconds float64,
	eclipticLatitudeDegrees float64, eclipticLatitudeMinutes float64, eclipticLatitudeSeconds float64,
	greenwichDay float64, greenwichMonth int, greenwichYear int,
) patype.EquatorialCoordinates2 {
	var ecLonDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(eclipticLongitudeDegrees, eclipticLongitudeMinutes, eclipticLongitudeSeconds)
	var ecLatDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(eclipticLatitudeDegrees, eclipticLatitudeMinutes, eclipticLatitudeSeconds)
	var ecLonRad float64 = pautil.DegreesToRadians(ecLonDeg)
	var ecLatRad float64 = pautil.DegreesToRadians(ecLatDeg)
	var obliqDeg float64 = pamacro.Obliq(greenwichDay, greenwichMonth, greenwichYear)
	var obliqRad float64 = pautil.DegreesToRadians(obliqDeg)
	var sinDec float64 = math.Sin(ecLatRad)*math.Cos(obliqRad) + math.Cos(ecLatRad)*math.Sin(obliqRad)*math.Sin(ecLonRad)
	var decRad float64 = math.Asin(sinDec)
	var decDeg float64 = pamacro.Degrees(decRad)
	var y float64 = math.Sin(ecLonRad)*math.Cos(obliqRad) - math.Tan(ecLatRad)*math.Sin(obliqRad)
	var x float64 = math.Cos(ecLonRad)
	var raRad float64 = math.Atan2(y, x)
	var raDeg1 float64 = pamacro.Degrees(raRad)
	var raDeg2 float64 = raDeg1 - 360*math.Floor(raDeg1/360)
	var raHours float64 = pamacro.DecimalDegreesToDegreeHours(raDeg2)

	var outRaHours int = pamacro.DecimalHoursHour(raHours)
	var outRaMinutes int = pamacro.DecimalHoursMinute(raHours)
	var outRaSeconds float64 = pamacro.DecimalHoursSecond(raHours)
	var outDecDegrees float64 = pamacro.DecimalDegreesDegrees(decDeg)
	var outDecMinutes float64 = pamacro.DecimalDegreesMinutes(decDeg)
	var outDecSeconds float64 = pamacro.DecimalDegreesSeconds(decDeg)

	return patype.EquatorialCoordinates2{
		RightAscensionHours: float64(outRaHours), RightAscensionMinutes: float64(outRaMinutes), RightAscensionSeconds: outRaSeconds,
		DeclinationDegrees: outDecDegrees, DeclinationMinutes: outDecMinutes, DeclinationSeconds: outDecSeconds,
	}
}

/* Convert Equatorial Coordinates to Ecliptic Coordinates */
func EquatorialCoordinateToEclipticCoordinate(
	raHours float64, raMinutes float64, raSeconds float64,
	decDegrees float64, decMinutes float64, decSeconds float64,
	gwDay float64, gwMonth int, gwYear int,
) patype.EclipticCoordinates {
	var raDeg float64 = pamacro.DegreeHoursToDecimalDegrees(pamacro.HmsDh(raHours, raMinutes, raSeconds))
	var decDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(decDegrees, decMinutes, decSeconds)
	var raRad float64 = pautil.DegreesToRadians(raDeg)
	var decRad float64 = pautil.DegreesToRadians(decDeg)
	var obliqDeg float64 = pamacro.Obliq(gwDay, gwMonth, gwYear)
	var obliqRad float64 = pautil.DegreesToRadians(obliqDeg)
	var sinEclRad float64 = math.Sin(decRad)*math.Cos(obliqRad) - math.Cos(decRad)*math.Sin(obliqRad)*math.Sin(raRad)
	var eclLatRad float64 = math.Asin(sinEclRad)
	var eclLatDeg float64 = pamacro.Degrees(eclLatRad)
	var y float64 = math.Sin(raRad)*math.Cos(obliqRad) + math.Tan(decRad)*math.Sin(obliqRad)
	var x float64 = math.Cos(raRad)
	var eclLongRad float64 = math.Atan2(y, x)
	var eclLongDeg1 float64 = pamacro.Degrees(eclLongRad)
	var eclLongDeg2 float64 = eclLongDeg1 - 360*math.Floor(eclLongDeg1/360)

	var outEclLongDeg float64 = pamacro.DecimalDegreesDegrees(eclLongDeg2)
	var outEclLongMin float64 = pamacro.DecimalDegreesMinutes(eclLongDeg2)
	var outEclLongSec float64 = pamacro.DecimalDegreesSeconds(eclLongDeg2)
	var outEclLatDeg float64 = pamacro.DecimalDegreesDegrees(eclLatDeg)
	var outEclLatMin float64 = pamacro.DecimalDegreesMinutes(eclLatDeg)
	var outEclLatSec float64 = pamacro.DecimalDegreesSeconds(eclLatDeg)

	return patype.EclipticCoordinates{
		LongitudeDegrees: outEclLongDeg, LongitudeMinutes: outEclLongMin, LongitudeSeconds: outEclLongSec,
		LatitudeDegrees: outEclLatDeg, LatitudeMinutes: outEclLatMin, LatitudeSeconds: outEclLatSec,
	}
}

/* Convert Equatorial Coordinates to Galactic Coordinates */
func EquatorialCoordinateToGalacticCoordinate(
	raHours float64, raMinutes float64, raSeconds float64,
	decDegrees float64, decMinutes float64, decSeconds float64,
) patype.GalacticCoordinates {
	var raDeg float64 = pamacro.DegreeHoursToDecimalDegrees(pamacro.HmsDh(raHours, raMinutes, raSeconds))
	var decDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(decDegrees, decMinutes, decSeconds)
	var raRad float64 = pautil.DegreesToRadians(raDeg)
	var decRad float64 = pautil.DegreesToRadians(decDeg)
	var sinB float64 = math.Cos(decRad)*math.Cos(pautil.DegreesToRadians(27.4))*math.Cos(raRad-pautil.DegreesToRadians(192.25)) +
		math.Sin(decRad)*math.Sin(pautil.DegreesToRadians(27.4))
	var bRadians float64 = math.Asin(sinB)
	var bDeg float64 = pamacro.Degrees(bRadians)
	var y float64 = math.Sin(decRad) - sinB*math.Sin(pautil.DegreesToRadians(27.4))
	var x float64 = math.Cos(decRad) * math.Sin(raRad-pautil.DegreesToRadians(192.25)) * math.Cos(pautil.DegreesToRadians(27.4))
	var longDeg1 float64 = pamacro.Degrees(math.Atan2(y, x)) + 33
	var longDeg2 float64 = longDeg1 - 360*math.Floor(longDeg1/360)

	var galLongDeg float64 = pamacro.DecimalDegreesDegrees(longDeg2)
	var galLongMin float64 = pamacro.DecimalDegreesMinutes(longDeg2)
	var galLongSec float64 = pamacro.DecimalDegreesSeconds(longDeg2)
	var galLatDeg float64 = pamacro.DecimalDegreesDegrees(bDeg)
	var galLatMin float64 = pamacro.DecimalDegreesMinutes(bDeg)
	var galLatSec float64 = pamacro.DecimalDegreesSeconds(bDeg)

	return patype.GalacticCoordinates{
		LongitudeDegrees: galLongDeg, LongitudeMinutes: galLongMin, LongitudeSeconds: galLongSec,
		LatitudeDegrees: galLatDeg, LatitudeMinutes: galLatMin, LatitudeSeconds: galLatSec,
	}
}

/* Convert Galactic Coordinates to Equatorial Coordinates */
func GalacticCoordinatesToEquatorialCoordinates(
	galLongDeg float64, galLongMin float64, galLongSec float64,
	galLatDeg float64, galLatMin float64, galLatSec float64,
) patype.EquatorialCoordinates2 {
	var gLongDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(galLongDeg, galLongMin, galLongSec)
	var gLatDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(galLatDeg, galLatMin, galLatSec)
	var gLongRad float64 = pautil.DegreesToRadians(gLongDeg)
	var gLatRad float64 = pautil.DegreesToRadians(gLatDeg)
	var sinDec float64 = math.Cos(gLatRad)*math.Cos(pautil.DegreesToRadians(27.4))*math.Sin(gLongRad-pautil.DegreesToRadians(33.0)) +
		math.Sin(gLatRad)*math.Sin(pautil.DegreesToRadians(27.4))
	var decRadians float64 = math.Asin(sinDec)
	var decDeg float64 = pamacro.Degrees(decRadians)
	var y float64 = math.Cos(gLatRad) * math.Cos(gLongRad-pautil.DegreesToRadians(33.0))
	var x float64 = math.Sin(gLatRad)*math.Cos(pautil.DegreesToRadians(27.4)) -
		math.Cos(gLatRad)*math.Sin(pautil.DegreesToRadians(27.4))*math.Sin(gLongRad-pautil.DegreesToRadians(33.0))

	var raDeg1 float64 = pamacro.Degrees(math.Atan2(y, x)) + 192.25
	var raDeg2 float64 = raDeg1 - 360*math.Floor(raDeg1/360)
	var raHours1 float64 = pamacro.DecimalDegreesToDegreeHours(raDeg2)

	var raHours float64 = float64(pamacro.DecimalHoursHour(raHours1))
	var raMinutes float64 = float64(pamacro.DecimalHoursMinute(raHours1))
	var raSeconds float64 = pamacro.DecimalHoursSecond(raHours1)
	var decDegrees float64 = pamacro.DecimalDegreesDegrees(decDeg)
	var decMinutes float64 = pamacro.DecimalDegreesMinutes(decDeg)
	var decSeconds float64 = pamacro.DecimalDegreesSeconds(decDeg)

	return patype.EquatorialCoordinates2{
		RightAscensionHours: raHours, RightAscensionMinutes: raMinutes, RightAscensionSeconds: raSeconds,
		DeclinationDegrees: decDegrees, DeclinationMinutes: decMinutes, DeclinationSeconds: decSeconds,
	}
}

/* Calculate the angle between two celestial objects */
func AngleBetweenTwoObjects(
	raLong1HourDeg float64, raLong1Min float64, raLong1Sec float64, decLat1Deg float64, decLat1Min float64, decLat1Sec float64,
	raLong2HourDeg float64, raLong2Min float64, raLong2Sec float64, decLat2Deg float64, decLat2Min float64, decLat2Sec float64,
	hourOrDegree patype.AngleMeasurementTypes,
) patype.Angle {
	var raLong1Decimal float64
	if hourOrDegree == patype.AngleMeasurementType_Hours {
		raLong1Decimal = pamacro.HmsDh(raLong1HourDeg, raLong1Min, raLong1Sec)
	} else {
		raLong1Decimal = pamacro.DegreesMinutesSecondsToDecimalDegrees(raLong1HourDeg, raLong1Min, raLong1Sec)
	}
	var raLong1Deg float64
	if hourOrDegree == patype.AngleMeasurementType_Hours {
		raLong1Deg = pamacro.DegreeHoursToDecimalDegrees(raLong1Decimal)
	} else {
		raLong1Deg = raLong1Decimal
	}

	var raLong1Rad float64 = pautil.DegreesToRadians(raLong1Deg)
	var decLat1Deg1 float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(decLat1Deg, decLat1Min, decLat1Sec)
	var decLat1Rad float64 = pautil.DegreesToRadians(decLat1Deg1)

	var raLong2Decimal float64
	if hourOrDegree == patype.AngleMeasurementType_Hours {
		raLong2Decimal = pamacro.HmsDh(raLong2HourDeg, raLong2Min, raLong2Sec)
	} else {
		raLong2Decimal = pamacro.DegreesMinutesSecondsToDecimalDegrees(raLong2HourDeg, raLong2Min, raLong2Sec)
	}
	var raLong2Deg float64
	if hourOrDegree == patype.AngleMeasurementType_Hours {
		raLong2Deg = pamacro.DegreeHoursToDecimalDegrees(raLong2Decimal)
	} else {
		raLong2Deg = raLong2Decimal
	}
	var raLong2Rad float64 = pautil.DegreesToRadians(raLong2Deg)
	var decLat2Deg1 float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(decLat2Deg, decLat2Min, decLat2Sec)
	var decLat2Rad float64 = pautil.DegreesToRadians(decLat2Deg1)

	var cosD float64 = math.Sin(decLat1Rad)*math.Sin(decLat2Rad) + math.Cos(decLat1Rad)*math.Cos(decLat2Rad)*math.Cos(raLong1Rad-raLong2Rad)
	var dRad float64 = math.Acos(cosD)
	var dDeg float64 = pamacro.Degrees(dRad)

	var angleDeg float64 = pamacro.DecimalDegreesDegrees(dDeg)
	var angleMin float64 = pamacro.DecimalDegreesMinutes(dDeg)
	var angleSec float64 = pamacro.DecimalDegreesSeconds(dDeg)

	return patype.Angle{Degrees: angleDeg, Minutes: angleMin, Seconds: angleSec}
}

/* Calculate rising and setting times for an object. */
func RisingAndSetting(
	raHours float64, raMinutes float64, raSeconds float64, decDeg float64, decMin float64, decSec float64,
	gwDateDay float64, gwDateMonth int, gwDateYear int,
	geogLongDeg float64, geogLatDeg float64, vertShiftDeg float64,
) patype.RiseSet {
	var raHours1 float64 = pamacro.HmsDh(raHours, raMinutes, raSeconds)
	var decRad float64 = pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(decDeg, decMin, decSec))
	var verticalDisplRadians float64 = pautil.DegreesToRadians(vertShiftDeg)
	var geoLatRadians float64 = pautil.DegreesToRadians(geogLatDeg)
	var cosH float64 = -(math.Sin(verticalDisplRadians) + math.Sin(geoLatRadians)*math.Sin(decRad)) / (math.Cos(geoLatRadians) * math.Cos(decRad))
	var hHours float64 = pamacro.DecimalDegreesToDegreeHours(pamacro.Degrees(math.Acos(cosH)))
	var lstRiseHours float64 = (raHours1 - hHours) - 24*math.Floor((raHours1-hHours)/24)
	var lstSetHours float64 = (raHours1 + hHours) - 24*math.Floor((raHours1+hHours)/24)
	var aDeg float64 = pamacro.Degrees(math.Acos((math.Sin(decRad) + math.Sin(verticalDisplRadians)*math.Sin(geoLatRadians)) /
		(math.Cos(verticalDisplRadians) * math.Cos(geoLatRadians))))
	var azRiseDeg float64 = aDeg - 360*math.Floor(aDeg/360)
	var azSetDeg float64 = (360 - aDeg) - 360*math.Floor((360-aDeg)/360)
	var utRiseHours1 float64 = pamacro.GreenwichSiderealTimeToUniversalTime(pamacro.LocalSiderealTimeToGreenwichSiderealTime(
		lstRiseHours, 0, 0, geogLongDeg), 0, 0, gwDateDay, gwDateMonth, gwDateYear)
	var utSetHours1 float64 = pamacro.GreenwichSiderealTimeToUniversalTime(pamacro.LocalSiderealTimeToGreenwichSiderealTime(
		lstSetHours, 0, 0, geogLongDeg), 0, 0, gwDateDay, gwDateMonth, gwDateYear)
	var utRiseAdjustedHours float64 = utRiseHours1 + 0.008333
	var utSetAdjustedHours float64 = utSetHours1 + 0.008333

	var rsStatus patype.RiseSetStatus = patype.RiseSetStatus_OK
	if cosH > 1 {
		rsStatus = patype.RiseSetStatus_NeverRises
	}
	if cosH < -1 {
		rsStatus = patype.RiseSetStatus_Circumpolar
	}

	var utRiseHour int
	if rsStatus == patype.RiseSetStatus_OK {
		utRiseHour = pamacro.DecimalHoursHour(utRiseAdjustedHours)
	} else {
		utRiseHour = 0
	}

	var utRiseMin int
	if rsStatus == patype.RiseSetStatus_OK {
		utRiseMin = pamacro.DecimalHoursMinute(utRiseAdjustedHours)
	} else {
		utRiseMin = 0
	}

	var utSetHour int
	if rsStatus == patype.RiseSetStatus_OK {
		utSetHour = pamacro.DecimalHoursHour(utSetAdjustedHours)
	} else {
		utSetHour = 0
	}

	var utSetMin int
	if rsStatus == patype.RiseSetStatus_OK {
		utSetMin = pamacro.DecimalHoursMinute(utSetAdjustedHours)
	} else {
		utSetMin = 0
	}

	var azRise float64
	if rsStatus == patype.RiseSetStatus_OK {
		azRise = pautil.RoundTo(azRiseDeg, 2)
	} else {
		azRise = 0
	}

	var azSet float64
	if rsStatus == patype.RiseSetStatus_OK {
		azSet = pautil.RoundTo(azSetDeg, 2)
	} else {
		azSet = 0
	}

	return patype.RiseSet{
		RiseSetStatusCurrent: rsStatus,
		UtRiseHour:           float64(utRiseHour), UtRiseMinute: float64(utRiseMin),
		UtSetHour: float64(utSetHour), UtSetMinute: float64(utSetMin),
		AzRise: azRise, AzSet: azSet,
	}
}

/* Calculate precession (corrected coordinates between two epochs) */
func CorrectForPrecession(
	raHour float64, raMinutes float64, raSeconds float64, decDeg float64, decMinutes float64, decSeconds float64,
	epoch1Day float64, epoch1Month int, epoch1Year int, epoch2Day float64, epoch2Month int, epoch2Year int,
) patype.CorrectedPrecession {
	var ra1Rad float64 = pautil.DegreesToRadians(pamacro.DegreeHoursToDecimalDegrees(pamacro.HmsDh(raHour, raMinutes, raSeconds)))
	var dec1Rad float64 = pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(decDeg, decMinutes, decSeconds))
	var tCenturies float64 = (pamacro.CivilDateToJulianDate(epoch1Day, float64(epoch1Month), float64(epoch1Year)) - 2415020) / 36525
	var mSec float64 = 3.07234 + (0.00186 * tCenturies)
	var nArcsec float64 = 20.0468 - (0.0085 * tCenturies)
	var nYears float64 = (pamacro.CivilDateToJulianDate(epoch2Day, float64(epoch2Month), float64(epoch2Year)) -
		pamacro.CivilDateToJulianDate(epoch1Day, float64(epoch1Month), float64(epoch1Year))) / 365.25
	var s1Hours float64 = ((mSec + (nArcsec * math.Sin(ra1Rad) * math.Tan(dec1Rad) / 15)) * nYears) / 3600
	var ra2Hours float64 = pamacro.HmsDh(raHour, raMinutes, raSeconds) + s1Hours
	var s2Deg float64 = (nArcsec * math.Cos(ra1Rad) * nYears) / 3600
	var dec2Deg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(decDeg, decMinutes, decSeconds) + s2Deg

	var correctedRaHour int = pamacro.DecimalHoursHour(ra2Hours)
	var correctedRaMinutes int = pamacro.DecimalHoursMinute(ra2Hours)
	var correctedRaSeconds float64 = pamacro.DecimalHoursSecond(ra2Hours)
	var correctedDecDeg float64 = pamacro.DecimalDegreesDegrees(dec2Deg)
	var correctedDecMinutes float64 = pamacro.DecimalDegreesMinutes(dec2Deg)
	var correctedDecSeconds float64 = pamacro.DecimalDegreesSeconds(dec2Deg)

	return patype.CorrectedPrecession{
		RightAscensionHours: float64(correctedRaHour), RightAscensionMinutes: float64(correctedRaMinutes), RightAscensionSeconds: correctedRaSeconds,
		DeclinationDegrees: correctedDecDeg, DeclinationMinutes: correctedDecMinutes, DeclinationSeconds: correctedDecSeconds,
	}
}

/* Calculate nutation for two values: ecliptic longitude and obliquity, for a Greenwich date. */
func NutationInEclipticLongitudeAndObliquity(greenwichDay float64, greenwichMonth int, greenwichYear int) patype.Nutation {
	var jdDays float64 = pamacro.CivilDateToJulianDate(greenwichDay, float64(greenwichMonth), float64(greenwichYear))
	var tCenturies float64 = (jdDays - 2415020) / 36525
	var aDeg float64 = 100.0021358 * tCenturies
	var l1Deg float64 = 279.6967 + (0.000303 * tCenturies * tCenturies)
	var lDeg1 float64 = l1Deg + 360*(aDeg-math.Floor(aDeg))
	var lDeg2 float64 = lDeg1 - 360*math.Floor(lDeg1/360)
	var lRad float64 = pautil.DegreesToRadians(lDeg2)
	var bDeg float64 = 5.372617 * tCenturies
	var nDeg1 float64 = 259.1833 - 360*(bDeg-math.Floor(bDeg))
	var nDeg2 float64 = nDeg1 - 360*(math.Floor(nDeg1/360))
	var nRad float64 = pautil.DegreesToRadians(nDeg2)
	var nutInLongArcsec float64 = -17.2*math.Sin(nRad) - 1.3*math.Sin(2*lRad)
	var nutInOblArcsec float64 = 9.2*math.Cos(nRad) + 0.5*math.Cos(2*lRad)

	var nutInLongDeg float64 = nutInLongArcsec / 3600
	var nutInOblDeg float64 = nutInOblArcsec / 3600

	return patype.Nutation{NutationInEcliptionLongitude: nutInLongDeg, NutationInObliquity: nutInOblDeg}
}

/* Correct ecliptic coordinates for the effects of aberration. */
func CorrectForAberration(
	utHour float64, utMinutes float64, utSeconds float64, gwDay float64, gwMonth int, gwYear int,
	trueEclLongDeg float64, trueEclLongMin float64, trueEclLongSec float64,
	trueEclLatDeg float64, trueEclLatMin float64, trueEclLatSec float64,
) patype.CorrectedEclipticCoordinates {
	var trueLongDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(trueEclLongDeg, trueEclLongMin, trueEclLongSec)
	var trueLatDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(trueEclLatDeg, trueEclLatMin, trueEclLatSec)
	var sunTrueLongDeg float64 = pamacro.SunLong(utHour, utMinutes, utSeconds, 0, 0, gwDay, gwMonth, gwYear)
	var dLongArcsec float64 = -20.5 * math.Cos(pautil.DegreesToRadians(sunTrueLongDeg-trueLongDeg)) / math.Cos(pautil.DegreesToRadians(trueLatDeg))
	var dLatArcsec float64 = -20.5 * math.Sin(pautil.DegreesToRadians(sunTrueLongDeg-trueLongDeg)) * math.Sin(pautil.DegreesToRadians(trueLatDeg))
	var apparentLongDeg float64 = trueLongDeg + (dLongArcsec / 3600)
	var apparentLatDeg float64 = trueLatDeg + (dLatArcsec / 3600)

	var apparentEclLongDeg float64 = pamacro.DecimalDegreesDegrees(apparentLongDeg)
	var apparentEclLongMin float64 = pamacro.DecimalDegreesMinutes(apparentLongDeg)
	var apparentEclLongSec float64 = pamacro.DecimalDegreesSeconds(apparentLongDeg)
	var apparentEclLatDeg float64 = pamacro.DecimalDegreesDegrees(apparentLatDeg)
	var apparentEclLatMin float64 = pamacro.DecimalDegreesMinutes(apparentLatDeg)
	var apparentEclLatSec float64 = pamacro.DecimalDegreesSeconds(apparentLatDeg)

	return patype.CorrectedEclipticCoordinates{
		LongitudeDegrees: apparentEclLongDeg, LongitudeMinutes: apparentEclLongMin, LongitudeSeconds: apparentEclLongSec,
		LatitudeDegrees: apparentEclLatDeg, LatitudeMinutes: apparentEclLatMin, LatitudeSeconds: apparentEclLatSec,
	}
}

/*
Calculate corrected RA/Dec, accounting for atmospheric refraction.

NOTE: Valid values for coordinate_type are "TRUE" and "APPARENT".
*/
func AtmosphericRefraction(
	trueRaHour float64, trueRaMin float64, trueRaSec float64,
	trueDecDeg float64, trueDecMin float64, trueDecSec float64,
	coordinateType1 patype.CoordinateType, geogLongDeg float64, geogLatDeg float64,
	daylightSavingHours int, timezoneHours int,
	lcdDay float64, lcdMonth int, lcdYear int,
	lctHour float64, lctMin float64, lctSec float64,
	atmosphericPressureMbar float64, atmosphericTemperatureCelsius float64,
) patype.CorrectedRefraction {
	var haHour float64 = pamacro.RightAscensionToHourAngle(
		trueRaHour, trueRaMin, trueRaSec, lctHour, lctMin, lctSec, daylightSavingHours, timezoneHours, lcdDay, lcdMonth, lcdYear, geogLongDeg)
	var azimuthDeg float64 = pamacro.EquatorialCoordinatesToAzimuth(haHour, 0, 0, trueDecDeg, trueDecMin, trueDecSec, geogLatDeg)
	var altitudeDeg float64 = pamacro.EquatorialCoordinatestoAltitude(haHour, 0, 0, trueDecDeg, trueDecMin, trueDecSec, geogLatDeg)
	var correctedAltitudeDeg float64 = pamacro.Refract(altitudeDeg, coordinateType1, atmosphericPressureMbar, atmosphericTemperatureCelsius)

	var correctedHaHour float64 = pamacro.HorizonCoordinatesToHourAngle(azimuthDeg, 0, 0, correctedAltitudeDeg, 0, 0, geogLatDeg)
	var correctedRaHour1 float64 = pamacro.HourAngleToRightAscension(
		correctedHaHour, 0, 0, lctHour, lctMin, lctSec, daylightSavingHours, timezoneHours, lcdDay, lcdMonth, lcdYear, geogLongDeg)
	var correctedDecDeg1 float64 = pamacro.HorizonCoordinatesToDeclination(azimuthDeg, 0, 0, correctedAltitudeDeg, 0, 0, geogLatDeg)

	var correctedRaHour int = pamacro.DecimalHoursHour(correctedRaHour1)
	var correctedRaMin int = pamacro.DecimalHoursMinute(correctedRaHour1)
	var correctedRaSec float64 = pamacro.DecimalHoursSecond(correctedRaHour1)
	var correctedDecDeg float64 = pamacro.DecimalDegreesDegrees(correctedDecDeg1)
	var correctedDecMin float64 = pamacro.DecimalDegreesMinutes(correctedDecDeg1)
	var correctedDecSec float64 = pamacro.DecimalDegreesSeconds(correctedDecDeg1)

	return patype.CorrectedRefraction{
		RightAscensionHours: float64(correctedRaHour), RightAscensionMinutes: float64(correctedRaMin), RightAscensionSeconds: correctedRaSec,
		DeclinationDegrees: correctedDecDeg, DeclinationMinutes: correctedDecMin, DeclinationSeconds: correctedDecSec,
	}
}

/* Calculate corrected RA/Dec, accounting for geocentric parallax. */
func CorrectionsForGeocentricParallax(
	raHour float64, raMin float64, raSec float64, decDeg float64, decMin float64, decSec float64,
	coordinateType patype.CoordinateType, equatorialHorParallaxDeg float64,
	geogLongDeg float64, geogLatDeg float64, heightM float64, daylightSaving int, timezoneHours int,
	lcdDay float64, lcdMonth int, lcdYear int, lctHour float64, lctMin float64, lctSec float64,
) patype.CorrectedParallax {
	var haHours float64 = pamacro.RightAscensionToHourAngle(
		raHour, raMin, raSec, lctHour, lctMin, lctSec, daylightSaving, timezoneHours, lcdDay, lcdMonth, lcdYear, geogLongDeg)

	var correctedHaHours float64 = pamacro.ParallaxHa(
		haHours, 0, 0, decDeg, decMin, decSec, coordinateType, geogLatDeg, heightM, equatorialHorParallaxDeg)
	var correctedRaHours float64 = pamacro.HourAngleToRightAscension(
		correctedHaHours, 0, 0, lctHour, lctMin, lctSec, daylightSaving, timezoneHours, lcdDay, lcdMonth, lcdYear, geogLongDeg)
	var correctedDecDeg1 float64 = pamacro.ParallaxDec(
		haHours, 0, 0, decDeg, decMin, decSec, coordinateType, geogLatDeg, heightM, equatorialHorParallaxDeg)

	var cRaHour int = pamacro.DecimalHoursHour(correctedRaHours)
	var cRaMin int = pamacro.DecimalHoursMinute(correctedRaHours)
	var cRaSec float64 = pamacro.DecimalHoursSecond(correctedRaHours)
	var cDecDeg float64 = pamacro.DecimalDegreesDegrees(correctedDecDeg1)
	var cDecMin float64 = pamacro.DecimalDegreesMinutes(correctedDecDeg1)
	var cDecSec float64 = pamacro.DecimalDegreesSeconds(correctedDecDeg1)

	return patype.CorrectedParallax{
		RightAscensionHours: float64(cRaHour), RightAscensionMinutes: float64(cRaMin), RightAscensionSeconds: cRaSec,
		DeclinationDegrees: cDecDeg, DeclinationMinutes: cDecMin, DeclinationSeconds: cDecSec,
	}
}

/* Calculate heliographic coordinates for a given Greenwich date, with a given heliographic position angle and heliographic displacement in arc minutes. */
func HeliographicCoordinates(
	helioPositionAngleDeg float64, helioDisplacementArcmin float64,
	gwdateDay float64, gwDateMonth int, gwdateYear int,
) patype.HeliographicCoordinates {
	var julianDateDays float64 = pamacro.CivilDateToJulianDate(gwdateDay, float64(gwDateMonth), float64(gwdateYear))
	var tCenturies float64 = (julianDateDays - 2415020) / 36525
	var longAscNodeDeg float64 = pamacro.DegreesMinutesSecondsToDecimalDegrees(74, 22, 0) + (84 * tCenturies / 60)
	var sunLongDeg float64 = pamacro.SunLong(0, 0, 0, 0, 0, gwdateDay, gwDateMonth, gwdateYear)
	var y float64 = math.Sin(pautil.DegreesToRadians(longAscNodeDeg-sunLongDeg)) *
		math.Cos(pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(7, 15, 0)))
	var x float64 = -math.Cos(pautil.DegreesToRadians(longAscNodeDeg - sunLongDeg))
	var aDeg float64 = pamacro.Degrees(math.Atan2(y, x))
	var mDeg1 float64 = 360 - (360 * (julianDateDays - 2398220) / 25.38)
	var mDeg2 float64 = mDeg1 - 360*math.Floor(mDeg1/360)
	var l0Deg1 float64 = mDeg2 + aDeg
	var b0Rad float64 = math.Asin(math.Sin(pautil.DegreesToRadians(sunLongDeg-longAscNodeDeg)) *
		math.Sin(pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(7, 15, 0))))
	var theta1Rad float64 = math.Atan(-math.Cos(pautil.DegreesToRadians(sunLongDeg)) *
		math.Tan(pautil.DegreesToRadians(pamacro.Obliq(gwdateDay, gwDateMonth, gwdateYear))))
	var theta2Rad float64 = math.Atan(-math.Cos(pautil.DegreesToRadians(longAscNodeDeg-sunLongDeg)) *
		math.Tan(pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(7, 15, 0))))
	var pDeg float64 = pamacro.Degrees(theta1Rad + theta2Rad)
	var rho1Deg float64 = helioDisplacementArcmin / 60
	var rhoRad float64 = math.Asin(2*rho1Deg/pamacro.SunDia(0, 0, 0, 0, 0, gwdateDay, gwDateMonth, gwdateYear)) - pautil.DegreesToRadians(rho1Deg)
	var bRad float64 = math.Asin(math.Sin(b0Rad)*math.Cos(rhoRad) +
		math.Cos(b0Rad)*math.Sin(rhoRad)*math.Cos(pautil.DegreesToRadians(pDeg-helioPositionAngleDeg)))
	var bDeg float64 = pamacro.Degrees(bRad)
	var lDeg1 float64 = pamacro.Degrees(math.Asin(math.Sin(rhoRad)*math.Sin(pautil.DegreesToRadians(pDeg-helioPositionAngleDeg))/math.Cos(bRad))) + l0Deg1
	var lDeg2 float64 = lDeg1 - 360*math.Floor(lDeg1/360)

	var helioLongDeg float64 = pautil.RoundTo(lDeg2, 2)
	var helioLatDeg float64 = pautil.RoundTo(bDeg, 2)

	return patype.HeliographicCoordinates{LongitudeDegrees: helioLongDeg, LatitudeDegrees: helioLatDeg}
}

/* Calculate carrington rotation number for a Greenwich date */
func CarringtonRotationNumber(gwdateDay float64, gwdateMonth int, gwdateYear int) int {
	var julianDateDays float64 = pamacro.CivilDateToJulianDate(gwdateDay, float64(gwdateMonth), float64(gwdateYear))

	var crn int = 1690 + int(pautil.RoundTo((julianDateDays-2444235.34)/27.2753, 0))

	return crn
}

/* Calculate selenographic (lunar) coordinates (sub-Earth) */
func SelenographicCoordinates1(gwdateDay float64, gwdateMonth int, gwdateYear int) patype.SelenographicSubEarthCoordinates {
	var julianDateDays float64 = pamacro.CivilDateToJulianDate(gwdateDay, float64(gwdateMonth), float64(gwdateYear))
	var tCenturies float64 = (julianDateDays - 2451545) / 36525
	var longAscNodeDeg float64 = 125.044522 - 1934.136261*tCenturies
	var f1 float64 = 93.27191 + 483202.0175*tCenturies
	var f2 float64 = f1 - 360*math.Floor(f1/360)
	var geocentricMoonLongDeg float64 = pamacro.MoonLongitude(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear)
	var geocentricMoonLatRad float64 = pautil.DegreesToRadians(pamacro.MoonLatitude(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear))
	var inclinationRad float64 = pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(1, 32, 32.7))
	var nodeLongRad float64 = pautil.DegreesToRadians(longAscNodeDeg - geocentricMoonLongDeg)
	var sinBe float64 = -math.Cos(inclinationRad)*math.Sin(geocentricMoonLatRad) +
		math.Sin(inclinationRad)*math.Cos(geocentricMoonLatRad)*math.Sin(nodeLongRad)
	var subEarthLatDeg float64 = pamacro.Degrees(math.Asin(sinBe))
	var aRad float64 = math.Atan2(-math.Sin(geocentricMoonLatRad)*math.Sin(inclinationRad)-math.Cos(geocentricMoonLatRad)*
		math.Cos(inclinationRad)*math.Sin(nodeLongRad), math.Cos(geocentricMoonLatRad)*math.Cos(nodeLongRad))
	var aDeg float64 = pamacro.Degrees(aRad)
	var subEarthLongDeg1 float64 = aDeg - f2
	var subEarthLongDeg2 float64 = subEarthLongDeg1 - 360*math.Floor(subEarthLongDeg1/360)
	var subEarthLongDeg3 float64
	if subEarthLongDeg2 > 180 {
		subEarthLongDeg3 = (subEarthLongDeg2 - 360)
	} else {
		subEarthLongDeg3 = subEarthLongDeg2
	}
	var c1Rad float64 = math.Atan(math.Cos(nodeLongRad) * math.Sin(inclinationRad) / (math.Cos(geocentricMoonLatRad)*math.Cos(inclinationRad) +
		math.Sin(geocentricMoonLatRad)*math.Sin(inclinationRad)*math.Sin(nodeLongRad)))
	var obliquityRad float64 = pautil.DegreesToRadians(pamacro.Obliq(gwdateDay, gwdateMonth, gwdateYear))
	var c2Rad float64 = math.Atan(math.Sin(obliquityRad) * math.Cos(pautil.DegreesToRadians(geocentricMoonLongDeg)) /
		(math.Sin(obliquityRad)*math.Sin(geocentricMoonLatRad)*math.Sin(pautil.DegreesToRadians(geocentricMoonLongDeg)) -
			math.Cos(obliquityRad)*math.Cos(geocentricMoonLatRad)))
	var cDeg float64 = pamacro.Degrees(c1Rad + c2Rad)

	var subEarthLongitude float64 = pautil.RoundTo(subEarthLongDeg3, 2)
	var subEarthLatitude float64 = pautil.RoundTo(subEarthLatDeg, 2)
	var positionAngleOfPole float64 = pautil.RoundTo(cDeg, 2)

	return patype.SelenographicSubEarthCoordinates{Longitude: subEarthLongitude, Latitude: subEarthLatitude, PositionAngleOfPole: positionAngleOfPole}
}

/* Calculate selenographic (lunar) coordinates (sub-Solar) */
func SelenographicCoordinates2(gwdateDay float64, gwdateMonth int, gwdateYear int) patype.SelenographicSubSolarCoordinates {
	var julianDateDays float64 = pamacro.CivilDateToJulianDate(gwdateDay, float64(gwdateMonth), float64(gwdateYear))
	var tCenturies float64 = (julianDateDays - 2451545) / 36525
	var longAscNodeDeg float64 = 125.044522 - 1934.136261*tCenturies
	var f1 float64 = 93.27191 + 483202.0175*tCenturies
	var f2 float64 = f1 - 360*math.Floor(f1/360)
	var sunGeocentricLongDeg float64 = pamacro.SunLong(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear)
	var moonEquHorParallaxArcMin float64 = pamacro.MoonHorizontalParallax(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear) * 60
	var sunEarthDistAU float64 = pamacro.SunDist(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear)
	var geocentricMoonLatRad float64 = pautil.DegreesToRadians(pamacro.MoonLatitude(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear))
	var geocentricMoonLongDeg float64 = pamacro.MoonLongitude(0, 0, 0, 0, 0, gwdateDay, gwdateMonth, gwdateYear)
	var adjustedMoonLongDeg float64 = sunGeocentricLongDeg + 180 + (26.4 * math.Cos(geocentricMoonLatRad) *
		math.Sin(pautil.DegreesToRadians(sunGeocentricLongDeg-geocentricMoonLongDeg)) / (moonEquHorParallaxArcMin * sunEarthDistAU))
	var adjustedMoonLatRad float64 = 0.14666 * geocentricMoonLatRad / (moonEquHorParallaxArcMin * sunEarthDistAU)
	var inclinationRad float64 = pautil.DegreesToRadians(pamacro.DegreesMinutesSecondsToDecimalDegrees(1, 32, 32.7))
	var nodeLongRad float64 = pautil.DegreesToRadians(longAscNodeDeg - adjustedMoonLongDeg)
	var sinBs float64 = -math.Cos(inclinationRad)*math.Sin(adjustedMoonLatRad) + math.Sin(inclinationRad)*math.Cos(adjustedMoonLatRad)*math.Sin(nodeLongRad)
	var subSolarLatDeg float64 = pamacro.Degrees(math.Asin(sinBs))
	var aRad float64 = math.Atan2(-math.Sin(adjustedMoonLatRad)*math.Sin(inclinationRad)-math.Cos(adjustedMoonLatRad)*
		math.Cos(inclinationRad)*math.Sin(nodeLongRad), math.Cos(adjustedMoonLatRad)*math.Cos(nodeLongRad))
	var aDeg float64 = pamacro.Degrees(aRad)
	var subSolarLongDeg1 float64 = aDeg - f2
	var subSolarLongDeg2 float64 = subSolarLongDeg1 - 360*math.Floor(subSolarLongDeg1/360)
	var subSolarLongDeg3 float64
	if subSolarLongDeg2 > 180 {
		subSolarLongDeg3 = subSolarLongDeg2 - 360
	} else {
		subSolarLongDeg3 = subSolarLongDeg2
	}
	var subSolarColongDeg float64 = 90 - subSolarLongDeg3

	var subSolarLongitude float64 = pautil.RoundTo(subSolarLongDeg3, 2)
	var subSolarColongitude float64 = pautil.RoundTo(subSolarColongDeg, 2)
	var subSolarLatitude float64 = pautil.RoundTo(subSolarLatDeg, 2)

	return patype.SelenographicSubSolarCoordinates{Longitude: subSolarLongitude, CoLongitude: subSolarColongitude, Latitude: subSolarLatitude}
}
