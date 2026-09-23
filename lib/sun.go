package lib

import (
	"math"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Calculate approximate position of the sun for a local date and time. */
func ApproximatePositionOfSun(
	lctHours float64, lctMinutes float64, lctSeconds float64, localDay float64, localMonth int, localYear int,
	isDaylightSaving bool, zoneCorrection int,
) patype.SunPosition {
	var daylightSaving int
	if isDaylightSaving == true {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var greenwichDateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var greenwichDateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var greenwichDateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var utHours float64 = pamacro.LocalCivilTimeToUniversalTime(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var utDays float64 = utHours / 24
	var jdDays float64 = pamacro.CivilDateToJulianDate(greenwichDateDay, float64(greenwichDateMonth), float64(greenwichDateYear)) + utDays
	var dDays float64 = jdDays - pamacro.CivilDateToJulianDate(0, 1, 2010)
	var nDeg float64 = 360 * dDays / 365.242191
	var mDeg1 float64 = nDeg + pamacro.SunEclipticLongitude(0, 1, 2010) - pamacro.SunPerigee(0, 1, 2010)
	var mDeg2 float64 = mDeg1 - 360*math.Floor(mDeg1/360)
	var ecDeg float64 = 360 * pamacro.SunEccentricity(0, 1, 2010) * math.Sin(pautil.DegreesToRadians(mDeg2)) / math.Pi
	var lsDeg1 float64 = nDeg + ecDeg + pamacro.SunEclipticLongitude(0, 1, 2010)
	var lsDeg2 float64 = lsDeg1 - 360*math.Floor(lsDeg1/360)
	var raDeg float64 = pamacro.EclipticRightAscension(lsDeg2, 0, 0, 0, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear)
	var raHours float64 = pamacro.DecimalDegreesToDegreeHours(raDeg)
	var decDeg float64 = pamacro.EclipticDeclination(lsDeg2, 0, 0, 0, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear)

	var sunRaHour int = pamacro.DecimalHoursHour(raHours)
	var sunRaMin int = pamacro.DecimalHoursMinute(raHours)
	var sunRaSec float64 = pamacro.DecimalHoursSecond(raHours)
	var sunDecDeg float64 = pamacro.DecimalDegreesDegrees(decDeg)
	var sunDecMin float64 = pamacro.DecimalDegreesMinutes(decDeg)
	var sunDecSec float64 = pamacro.DecimalDegreesSeconds(decDeg)

	return patype.SunPosition{
		RightAscensionHour: float64(sunRaHour), RightAscensionMinutes: float64(sunRaMin), RightAscensionSeconds: sunRaSec,
		DeclinationDegrees: sunDecDeg, DeclinationMinutes: sunDecMin, DeclinationSeconds: sunDecSec,
	}
}

/* Calculate precise position of the sun for a local date and time. */
func PrecisePositionOfSun(
	lctHours float64, lctMinutes float64, lctSeconds float64, localDay float64, localMonth int, localYear int,
	isDaylightSaving bool, zoneCorrection int,
) patype.SunPosition {
	var daylightSaving int
	if isDaylightSaving == true {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var gDay float64 = pamacro.LocalCivilTimeGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var gMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var gYear int = int(pamacro.LocalCivilTimeGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var sunEclipticLongitudeDeg float64 = pamacro.SunLong(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var raDeg float64 = pamacro.EclipticRightAscension(sunEclipticLongitudeDeg, 0, 0, 0, 0, 0, gDay, gMonth, gYear)
	var raHours float64 = pamacro.DecimalDegreesToDegreeHours(raDeg)
	var decDeg float64 = pamacro.EclipticDeclination(sunEclipticLongitudeDeg, 0, 0, 0, 0, 0, gDay, gMonth, gYear)

	var sunRaHour int = pamacro.DecimalHoursHour(raHours)
	var sunRaMinutes int = pamacro.DecimalHoursMinute(raHours)
	var sunRaSeconds float64 = pamacro.DecimalHoursSecond(raHours)
	var sunDecDeg float64 = pamacro.DecimalDegreesDegrees(decDeg)
	var sunDecMinutes float64 = pamacro.DecimalDegreesMinutes(decDeg)
	var sunDecSeconds float64 = pamacro.DecimalDegreesSeconds(decDeg)

	return patype.SunPosition{
		RightAscensionHour: float64(sunRaHour), RightAscensionMinutes: float64(sunRaMinutes), RightAscensionSeconds: sunRaSeconds,
		DeclinationDegrees: sunDecDeg, DeclinationMinutes: sunDecMinutes, DeclinationSeconds: sunDecSeconds,
	}
}

/* Calculate distance to the Sun (in km), and angular size. */
func SunDistanceAndAngularSize(
	lctHours float64, lctMinutes float64, lctSeconds float64, localDay float64, localMonth int, localYear int,
	isDaylightSaving bool, zoneCorrection int,
) patype.SunDistanceSize {
	var daylightSaving int
	if isDaylightSaving {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var gDay float64 = pamacro.LocalCivilTimeGreenwichDay(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var gMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var gYear int = int(pamacro.LocalCivilTimeGreenwichYear(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear))
	var trueAnomalyDeg float64 = pamacro.SunTrueAnomaly(lctHours, lctMinutes, lctSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	var trueAnomalyRad float64 = pautil.DegreesToRadians(trueAnomalyDeg)
	var eccentricity float64 = pamacro.SunEccentricity(gDay, gMonth, gYear)
	var f float64 = (1 + eccentricity*math.Cos(trueAnomalyRad)) / (1 - eccentricity*eccentricity)
	var rKm float64 = 149598500 / f
	var thetaDeg float64 = f * 0.533128

	var sunDistKm float64 = pautil.RoundTo(rKm, 0)
	var sunAngSizeDeg float64 = pamacro.DecimalDegreesDegrees(thetaDeg)
	var sunAngSizeMin float64 = pamacro.DecimalDegreesMinutes(thetaDeg)
	var sunAngSizeSec float64 = pamacro.DecimalDegreesSeconds(thetaDeg)

	return patype.SunDistanceSize{
		DistanceInKilometers: sunDistKm, AngularSize_Degrees: sunAngSizeDeg, AngularSize_Minutes: sunAngSizeMin, AngularSize_Seconds: sunAngSizeSec}
}

/* Calculate local sunrise and sunset. */
func SunriseAndSunset(
	localDay float64, localMonth int, localYear int, isDaylightSaving bool, zoneCorrection int,
	geographicalLongDeg float64, geographicalLatDeg float64,
) patype.SunriseSunsetInfo {
	var daylightSaving int
	if isDaylightSaving {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var localSunriseHours float64 = pamacro.SunriseLct(
		localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg)
	var localSunsetHours float64 = pamacro.SunsetLct(
		localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg)

	var sunRiseSetStatus patype.RiseSetStatus = pamacro.ESunRiseSet(
		localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg)

	var adjustedSunriseHours float64 = localSunriseHours + 0.008333
	var adjustedSunsetHours float64 = localSunsetHours + 0.008333

	var azimuthOfSunriseDeg1 float64 = pamacro.SunriseAz(localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg)
	var azimuthOfSunsetDeg1 float64 = pamacro.SunsetAz(localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg)

	var localSunriseHour int
	if sunRiseSetStatus == patype.RiseSetStatus_OK {
		localSunriseHour = pamacro.DecimalHoursHour(adjustedSunriseHours)
	} else {
		localSunriseHour = 0
	}
	var localSunriseMinute int
	if sunRiseSetStatus == patype.RiseSetStatus_OK {
		localSunriseMinute = pamacro.DecimalHoursMinute(adjustedSunriseHours)
	} else {
		localSunriseMinute = 0
	}

	var localSunsetHour int
	if sunRiseSetStatus == patype.RiseSetStatus_OK {
		localSunsetHour = pamacro.DecimalHoursHour(adjustedSunsetHours)
	} else {
		localSunsetHour = 0
	}

	var localSunsetMinute int
	if sunRiseSetStatus == patype.RiseSetStatus_OK {
		localSunsetMinute = pamacro.DecimalHoursMinute(adjustedSunsetHours)
	} else {
		localSunsetMinute = 0
	}

	var azimuthOfSunriseDeg float64
	if sunRiseSetStatus == patype.RiseSetStatus_OK {
		azimuthOfSunriseDeg = pautil.RoundTo(azimuthOfSunriseDeg1, 2)
	} else {
		azimuthOfSunriseDeg = 0
	}

	var azimuthOfSunsetDeg float64
	if sunRiseSetStatus == patype.RiseSetStatus_OK {
		azimuthOfSunsetDeg = pautil.RoundTo(azimuthOfSunsetDeg1, 2)
	} else {
		azimuthOfSunsetDeg = 0
	}

	var status patype.RiseSetStatus = sunRiseSetStatus

	return patype.SunriseSunsetInfo{
		LocalSunriseHour: float64(localSunriseHour), LocalSunriseMinute: float64(localSunriseMinute),
		LocalSunsetHour: float64(localSunsetHour), LocalSunsetMinute: float64(localSunsetMinute),
		AzimuthOfSunriseDeg: azimuthOfSunriseDeg, AzimuthOfSunsetDeg: azimuthOfSunsetDeg, Status: status,
	}
}

/* Calculate times of morning and evening twilight. */
func MorningAndEveningTwilight(
	localDay float64, localMonth int, localYear int, isDaylightSaving bool, zoneCorrection int,
	geographicalLongDeg float64, geographicalLatDeg float64, twilightType patype.TwilightType,
) patype.TwilightInfo {
	var daylightSaving int
	if isDaylightSaving {
		daylightSaving = 1
	} else {
		daylightSaving = 0
	}

	var startOfAmTwilightHours float64 = pamacro.TwilightAmLct(
		localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg, twilightType)

	var endOfPmTwilightHours float64 = pamacro.TwilightPmLct(
		localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg, twilightType)

	var twilightStatus patype.TwilightStatus = pamacro.ETwilight(
		localDay, localMonth, localYear, daylightSaving, zoneCorrection, geographicalLongDeg, geographicalLatDeg, twilightType)

	var adjustedAmStartTime float64 = startOfAmTwilightHours + 0.008333
	var adjustedPmStartTime float64 = endOfPmTwilightHours + 0.008333

	var amTwilightBeginsHour float64
	if twilightStatus == patype.TwilightStatus_OK {
		amTwilightBeginsHour = float64(pamacro.DecimalHoursHour(adjustedAmStartTime))
	} else {
		amTwilightBeginsHour = -99
	}

	var amTwilightBeginsMin float64
	if twilightStatus == patype.TwilightStatus_OK {
		amTwilightBeginsMin = float64(pamacro.DecimalHoursMinute(adjustedAmStartTime))
	} else {
		amTwilightBeginsMin = -99
	}

	var pmTwilightEndsHour float64
	if twilightStatus == patype.TwilightStatus_OK {
		pmTwilightEndsHour = float64(pamacro.DecimalHoursHour(adjustedPmStartTime))
	} else {
		pmTwilightEndsHour = -99
	}

	var pmTwilightEndsMin float64
	if twilightStatus == patype.TwilightStatus_OK {
		pmTwilightEndsMin = float64(pamacro.DecimalHoursMinute(adjustedPmStartTime))
	} else {
		pmTwilightEndsMin = -99
	}

	var status patype.TwilightStatus = twilightStatus

	return patype.TwilightInfo{
		AmTwilightBeginsHour: amTwilightBeginsHour, AmTwilightBeginsMin: amTwilightBeginsMin,
		PmTwilightEndsHour: pmTwilightEndsHour, PmTwilightEndsMin: pmTwilightEndsMin,
		Status: status,
	}
}

/* Calculate the equation of time. (The difference between the real Sun time and the mean Sun time.) */
func EquationOfTime(gwDateDay float64, gwDateMonth int, gwDateYear int) patype.EquationOfTime {
	var sunLongitudeDeg float64 = pamacro.SunLong(12, 0, 0, 0, 0, gwDateDay, gwDateMonth, gwDateYear)
	var sunRaHours float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(sunLongitudeDeg, 0, 0, 0, 0, 0, gwDateDay, gwDateMonth, gwDateYear))
	var equivalentUtHours float64 = pamacro.GreenwichSiderealTimeToUniversalTime(sunRaHours, 0, 0, gwDateDay, gwDateMonth, gwDateYear)
	var equationOfTimeHours float64 = equivalentUtHours - 12

	var equationOfTimeMin int = pamacro.DecimalHoursMinute(equationOfTimeHours)
	var equationOfTimeSec float64 = pamacro.DecimalHoursSecond(equationOfTimeHours)

	return patype.EquationOfTime{Minutes: float64(equationOfTimeMin), Seconds: equationOfTimeSec}
}

/*
Calculate solar elongation for a celestial body.

Solar elongation is the angle between the lines of sight from the Earth to the Sun and from the Earth to the celestial body.
*/
func SolarElongation(
	raHour float64, raMin float64, raSec float64, decDeg float64, decMin float64, decSec float64,
	gwDateDay float64, gwDateMonth int, gwDateYear int,
) float64 {
	var sunLongitudeDeg float64 = pamacro.SunLong(0, 0, 0, 0, 0, gwDateDay, gwDateMonth, gwDateYear)
	var sunRaHours float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(sunLongitudeDeg, 0, 0, 0, 0, 0, gwDateDay, gwDateMonth, gwDateYear))
	var sunDecDeg float64 = pamacro.EclipticDeclination(sunLongitudeDeg, 0, 0, 0, 0, 0, gwDateDay, gwDateMonth, gwDateYear)
	var solarElongationDeg float64 = pamacro.Angle(
		sunRaHours, 0, 0, sunDecDeg, 0, 0, raHour, raMin, raSec, decDeg, decMin, decSec, patype.AngleMeasurementType_Hours)

	return pautil.RoundTo(solarElongationDeg, 2)
}
