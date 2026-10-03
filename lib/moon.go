package lib

import (
	"math"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Calculate approximate position of the Moon. */
func ApproximatePositionOfMoon(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int,
) patype.MoonApproximatePosition {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var l0 float64 = 91.9293359879052
	var p0 float64 = 130.143076320618
	var n0 float64 = 291.682546643194
	var i float64 = 5.145396

	var gdateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var gdateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var gdateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var utHours float64 = pamacro.LocalCivilTimeToUniversalTime(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var dDays float64 = pamacro.CivilDateToJulianDate(
		gdateDay, float64(gdateMonth), float64(gdateYear)) - pamacro.CivilDateToJulianDate(0.0, 1, 2010) + utHours/24
	var sunLongDeg float64 = pamacro.SunLong(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var sunMeanAnomalyRad float64 = pamacro.SunMeanAnomaly(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var lmDeg float64 = pamacro.UnwindDeg(13.1763966*dDays + l0)
	var mmDeg float64 = pamacro.UnwindDeg(lmDeg - 0.1114041*dDays - p0)
	var nDeg float64 = pamacro.UnwindDeg(n0 - (0.0529539 * dDays))
	var evDeg float64 = 1.2739 * math.Sin(pautil.DegreesToRadians(2.0*(lmDeg-sunLongDeg)-mmDeg))
	var aeDeg float64 = 0.1858 * math.Sin(sunMeanAnomalyRad)
	var a3Deg float64 = 0.37 * math.Sin(sunMeanAnomalyRad)
	var mmdDeg float64 = mmDeg + evDeg - aeDeg - a3Deg
	var ecDeg float64 = 6.2886 * math.Sin(pautil.DegreesToRadians(mmdDeg))
	var a4Deg float64 = 0.214 * math.Sin(2.0*pautil.DegreesToRadians(mmdDeg))
	var ldDeg float64 = lmDeg + evDeg + ecDeg - aeDeg + a4Deg
	var vDeg float64 = 0.6583 * math.Sin(2.0*pautil.DegreesToRadians(ldDeg-sunLongDeg))
	var lddDeg float64 = ldDeg + vDeg
	var ndDeg float64 = nDeg - 0.16*math.Sin(sunMeanAnomalyRad)
	var y float64 = math.Sin(pautil.DegreesToRadians(lddDeg-ndDeg)) * math.Cos(pautil.DegreesToRadians(i))
	var x float64 = math.Cos(pautil.DegreesToRadians(lddDeg - ndDeg))

	var moonLongDeg float64 = pamacro.UnwindDeg(pamacro.Degrees(math.Atan2(y, x)) + ndDeg)
	var moonLatDeg float64 = pamacro.Degrees(math.Asin(math.Sin(pautil.DegreesToRadians(lddDeg-ndDeg)) * math.Sin(pautil.DegreesToRadians(i))))
	var moonRaHours1 float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(moonLongDeg, 0, 0, moonLatDeg, 0, 0, gdateDay, gdateMonth, gdateYear))
	var moonDecDeg1 float64 = pamacro.EclipticDeclination(moonLongDeg, 0, 0, moonLatDeg, 0, 0, gdateDay, gdateMonth, gdateYear)

	var moonRaHour int = pamacro.DecimalHoursHour(moonRaHours1)
	var moonRaMin int = pamacro.DecimalHoursMinute(moonRaHours1)
	var moonRaSec float64 = pamacro.DecimalHoursSecond(moonRaHours1)
	var moonDecDeg float64 = pamacro.DecimalDegreesDegrees(moonDecDeg1)
	var moonDecMin float64 = pamacro.DecimalDegreesMinutes(moonDecDeg1)
	var moonDecSec float64 = pamacro.DecimalDegreesSeconds(moonDecDeg1)

	return patype.MoonApproximatePosition{
		RaHour: float64(moonRaHour), RaMin: float64(moonRaMin), RaSec: moonRaSec,
		DecDeg: moonDecDeg, DecMin: moonDecMin, DecSec: moonDecSec,
	}
}

/* Calculate precise position of the Moon. */
func PrecisePositionOfMoon(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int,
) patype.MoonPrecisePosition {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var gdateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var gdateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var gdateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var moonResult patype.MoonLongLatHP = pamacro.MoonLongLatHp(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	var nutationInLongitudeDeg float64 = pamacro.NutatLong(gdateDay, gdateMonth, gdateYear)
	var correctedLongDeg float64 = moonResult.LongDeg + nutationInLongitudeDeg
	var earthMoonDistanceKm float64 = 6378.14 / math.Sin(pautil.DegreesToRadians(moonResult.HorPara))
	var moonRaHours1 float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(correctedLongDeg, 0, 0, moonResult.LatDeg, 0, 0, gdateDay, gdateMonth, gdateYear))
	var moonDecDeg1 float64 = pamacro.EclipticDeclination(correctedLongDeg, 0, 0, moonResult.LatDeg, 0, 0, gdateDay, gdateMonth, gdateYear)

	var moonRaHour int = pamacro.DecimalHoursHour(moonRaHours1)
	var moonRaMin int = pamacro.DecimalHoursMinute(moonRaHours1)
	var moonRaSec float64 = pamacro.DecimalHoursSecond(moonRaHours1)
	var moonDecDeg float64 = pamacro.DecimalDegreesDegrees(moonDecDeg1)
	var moonDecMin float64 = pamacro.DecimalDegreesMinutes(moonDecDeg1)
	var moonDecSec float64 = pamacro.DecimalDegreesSeconds(moonDecDeg1)
	var earthMoonDistKm float64 = pautil.RoundTo(earthMoonDistanceKm, 0)
	var moonHorParallaxDeg float64 = pautil.RoundTo(moonResult.HorPara, 6)

	return patype.MoonPrecisePosition{
		RaHour: float64(moonRaHour), RaMin: float64(moonRaMin), RaSec: moonRaSec,
		DecDeg: moonDecDeg, DecMin: moonDecMin, DecSec: moonDecSec,
		EarthMoonDistKm: earthMoonDistKm, HorParallaxDeg: moonHorParallaxDeg,
	}
}

/* Calculate Moon phase and position angle of bright limb. */
func MoonPhase(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int, accuracyLevel patype.AccuracyLevel,
) patype.MoonPhase {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var gdateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var gdateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var gdateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var sunLongDeg float64 = pamacro.SunLong(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var moonResult patype.MoonLongLatHP = pamacro.MoonLongLatHp(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var dRad float64 = pautil.DegreesToRadians(moonResult.LongDeg - sunLongDeg)

	var moonPhase1 float64
	if accuracyLevel == patype.AccuracyLevel_Precise {
		moonPhase1 = pamacro.MoonPhase(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	} else {
		moonPhase1 = (1.0 - math.Cos(dRad)) / 2.0
	}

	var sunRaRad float64 = pautil.DegreesToRadians(pamacro.EclipticRightAscension(sunLongDeg, 0, 0, 0, 0, 0, gdateDay, gdateMonth, gdateYear))
	var moonRaRad float64 = pautil.DegreesToRadians(pamacro.EclipticRightAscension(
		moonResult.LongDeg, 0, 0, moonResult.LatDeg, 0, 0, gdateDay, gdateMonth, gdateYear))
	var sunDecRad float64 = pautil.DegreesToRadians(pamacro.EclipticDeclination(sunLongDeg, 0, 0, 0, 0, 0, gdateDay, gdateMonth, gdateYear))
	var moonDecRad float64 = pautil.DegreesToRadians(pamacro.EclipticDeclination(
		moonResult.LongDeg, 0, 0, moonResult.LatDeg, 0, 0, gdateDay, gdateMonth, gdateYear))

	var y float64 = math.Cos(sunDecRad) * math.Sin(sunRaRad-moonRaRad)
	var x float64 = math.Cos(moonDecRad)*math.Sin(sunDecRad) - math.Sin(moonDecRad)*math.Cos(sunDecRad)*math.Cos(sunRaRad-moonRaRad)

	var chiDeg float64 = pamacro.Degrees(math.Atan2(y, x))

	var moonPhase float64 = pautil.RoundTo(moonPhase1, 2)
	var brightLimbDeg float64 = pautil.RoundTo(chiDeg, 2)

	return patype.MoonPhase{Phase: moonPhase, BrightLimbDeg: brightLimbDeg}
}

/* Calculate new moon and full moon instances. */
func TimesOfNewMoonAndFullMoon(
	isDaylightSaving bool, zoneCorrectionHours int, localDateDay float64, localDateMonth int, localDateYear int,
) patype.MoonNewFull {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var jdOfNewMoonDays float64 = pamacro.NewMoon(daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var jdOfFullMoonDays float64 = pamacro.FullMoon(3, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	var gDateOfNewMoonDay float64 = pamacro.JulianDateDay(jdOfNewMoonDays)
	var integerDay1 float64 = math.Floor(gDateOfNewMoonDay)
	var gDateOfNewMoonMonth int = pamacro.JulianDateMonth(jdOfNewMoonDays)
	var gDateOfNewMoonYear int = pamacro.JulianDateYear(jdOfNewMoonDays)

	var gDateOfFullMoonDay float64 = pamacro.JulianDateDay(jdOfFullMoonDays)
	var integerDay2 float64 = math.Floor(gDateOfFullMoonDay)
	var gDateOfFullMoonMonth int = pamacro.JulianDateMonth(jdOfFullMoonDays)
	var gDateOfFullMoonYear int = pamacro.JulianDateYear(jdOfFullMoonDays)

	var utOfNewMoonHours float64 = 24.0 * (gDateOfNewMoonDay - integerDay1)
	var utOfFullMoonHours float64 = 24.0 * (gDateOfFullMoonDay - integerDay2)
	var lctOfNewMoonHours float64 = pamacro.UniversalTimeToLocalCivilTime(
		utOfNewMoonHours+0.008333, 0, 0, daylightSaving, zoneCorrectionHours, integerDay1, gDateOfNewMoonMonth, gDateOfNewMoonYear)
	var lctOfFullMoonHours float64 = pamacro.UniversalTimeToLocalCivilTime(
		utOfFullMoonHours+0.008333, 0, 0, daylightSaving, zoneCorrectionHours, integerDay2, gDateOfFullMoonMonth, gDateOfFullMoonYear)

	var nmLocalTimeHour int = pamacro.DecimalHoursHour(lctOfNewMoonHours)
	var nmLocalTimeMin int = pamacro.DecimalHoursMinute(lctOfNewMoonHours)
	var nmLocalDateDay float64 = pamacro.UniversalTimeLocalCivilDay(
		utOfNewMoonHours, 0, 0, daylightSaving, zoneCorrectionHours, integerDay1, gDateOfNewMoonMonth, gDateOfNewMoonYear)
	var nmLocalDateMonth int = pamacro.UniversalTimeLocalCivilMonth(
		utOfNewMoonHours, 0, 0, daylightSaving, zoneCorrectionHours, integerDay1, gDateOfNewMoonMonth, gDateOfNewMoonYear)
	var nmLocalDateYear int = pamacro.UniversalTimeLocalCivilYear(
		utOfNewMoonHours, 0, 0, daylightSaving, zoneCorrectionHours, integerDay1, gDateOfNewMoonMonth, gDateOfNewMoonYear)
	var fmLocalTimeHour int = pamacro.DecimalHoursHour(lctOfFullMoonHours)
	var fmLocalTimeMin int = pamacro.DecimalHoursMinute(lctOfFullMoonHours)
	var fmLocalDateDay float64 = pamacro.UniversalTimeLocalCivilDay(
		utOfFullMoonHours, 0, 0, daylightSaving, zoneCorrectionHours, integerDay2, gDateOfFullMoonMonth, gDateOfFullMoonYear)
	var fmLocalDateMonth int = pamacro.UniversalTimeLocalCivilMonth(
		utOfFullMoonHours, 0, 0, daylightSaving, zoneCorrectionHours, integerDay2, gDateOfFullMoonMonth, gDateOfFullMoonYear)
	var fmLocalDateYear int = pamacro.UniversalTimeLocalCivilYear(
		utOfFullMoonHours, 0, 0, daylightSaving, zoneCorrectionHours, integerDay2, gDateOfFullMoonMonth, gDateOfFullMoonYear)

	return patype.MoonNewFull{
		NewLocalTimeHour: float64(nmLocalTimeHour), NewLocalTimeMin: float64(nmLocalTimeMin),
		NewLocalDateDay: nmLocalDateDay, NewLocalDateMonth: nmLocalDateMonth, NewLocalDateYear: nmLocalDateYear,
		FullLocalTimeHour: float64(fmLocalTimeHour), FullLocalTimeMin: float64(fmLocalTimeMin),
		FullLocalDateDay: fmLocalDateDay, FullLocalDateMonth: fmLocalDateMonth, FullLocalDateYear: fmLocalDateYear,
	}
}

/* Calculate Moon's distance, angular diameter, and horizontal parallax. */
func MoonDistAngDiamHorParallax(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int,
) patype.MoonDistDiameterHorParallax {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var moonDistance float64 = pamacro.MoonDist(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var moonAngularDiameter float64 = pamacro.MoonSize(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var moonHorizontalParallax float64 = pamacro.MoonHorizontalParallax(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	var earthMoonDist float64 = pautil.RoundTo(moonDistance, 0)
	var angDiameterDeg float64 = pamacro.DecimalDegreesDegrees(moonAngularDiameter + 0.008333)
	var angDiameterMin float64 = pamacro.DecimalDegreesMinutes(moonAngularDiameter + 0.008333)
	var horParallaxDeg float64 = pamacro.DecimalDegreesDegrees(moonHorizontalParallax)
	var horParallaxMin float64 = pamacro.DecimalDegreesMinutes(moonHorizontalParallax)
	var horParallaxSec float64 = pamacro.DecimalDegreesSeconds(moonHorizontalParallax)

	return patype.MoonDistDiameterHorParallax{
		EarthMoonDist: earthMoonDist, AngDiameterDeg: angDiameterDeg, AngDiameterMin: angDiameterMin,
		HorParallaxDeg: horParallaxDeg, HorParallaxMin: horParallaxMin, HorParallaxSec: horParallaxSec,
	}
}

/* Calculate date/time of local moonrise and moonset. */
func MoonriseAndMoonset(
	localDateDay float64, localDateMonth int, localDateYear int, isDaylightSaving bool, zoneCorrectionHours int,
	geogLongDeg float64, geogLatDeg float64,
) patype.MoonRiseSet {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var localTimeOfMoonriseHours float64 = pamacro.MoonRiseLct(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours, geogLongDeg, geogLatDeg)
	var moonRiseLcResult patype.FullDatePrecise = pamacro.MoonRiseLcDmy(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours, geogLongDeg, geogLatDeg)
	var localAzimuthDeg1 float64 = pamacro.MoonRiseAz(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours, geogLongDeg, geogLatDeg)

	var localTimeOfMoonsetHours float64 = pamacro.MoonSetLct(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours, geogLongDeg, geogLatDeg)
	var moonSetLcResult patype.FullDatePrecise = pamacro.MoonSetLcDmy(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours, geogLongDeg, geogLatDeg)
	var localAzimuthDeg2 float64 = pamacro.MoonSetAz(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours, geogLongDeg, geogLatDeg)

	var mrLtHour int = pamacro.DecimalHoursHour(localTimeOfMoonriseHours + 0.008333)
	var mrLtMin int = pamacro.DecimalHoursMinute(localTimeOfMoonriseHours + 0.008333)
	var mrLocalDateDay float64 = moonRiseLcResult.Day
	var mrLocalDateMonth int = moonRiseLcResult.Month
	var mrLocalDateYear int = moonRiseLcResult.Year
	var mrAzimuthDeg float64 = pautil.RoundTo(localAzimuthDeg1, 2)
	var msLtHour int = pamacro.DecimalHoursHour(localTimeOfMoonsetHours + 0.008333)
	var msLtMin int = pamacro.DecimalHoursMinute(localTimeOfMoonsetHours + 0.008333)
	var msLocalDateDay float64 = moonSetLcResult.Day
	var msLocalDateMonth int = moonSetLcResult.Month
	var msLocalDateYear int = moonSetLcResult.Year
	var msAzimuthDeg float64 = pautil.RoundTo(localAzimuthDeg2, 2)

	return patype.MoonRiseSet{
		RiseLocalTimeHour: float64(mrLtHour), RiseLocalTimeMin: float64(mrLtMin),
		RiseLocalDateDay: mrLocalDateDay, RiseLocalDateMonth: mrLocalDateMonth, RiseLocalDateYear: mrLocalDateYear,
		RiseAzimuthDeg:   mrAzimuthDeg,
		SetLocalTimeHour: float64(msLtHour), SetLocalTimeMin: float64(msLtMin),
		SetLocalDateDay: msLocalDateDay, SetLocalDateMonth: msLocalDateMonth, SetLocalDateYear: msLocalDateYear,
		SetAzimuthDeg: msAzimuthDeg,
	}
}
