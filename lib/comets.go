package lib

import (
	"math"
	padata "practicalastro/lib/data"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	"practicalastro/lib/util"
	pautil "practicalastro/lib/util"
)

/* Calculate position of an elliptical comet. */
func PositionOfEllipticalComet(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int, cometName string,
) patype.EllipticalCometPosition {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var greenwichDateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var greenwichDateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var greenwichDateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var cometInfo padata.CometDataElliptical = padata.GetCometDataElliptical(cometName)

	var timeSinceEpochYears float64 = (pamacro.CivilDateToJulianDate(
		greenwichDateDay, float64(greenwichDateMonth), float64(greenwichDateYear))-
		pamacro.CivilDateToJulianDate(0.0, 1, float64(greenwichDateYear)))/365.242191 + float64(greenwichDateYear) - cometInfo.Epoch_EpochOfPerihelion
	var mcDeg float64 = 360 * timeSinceEpochYears / cometInfo.Period_PeriodOfOrbit
	var mcRad float64 = pautil.DegreesToRadians(mcDeg - 360*math.Floor(mcDeg/360))
	var eccentricity float64 = cometInfo.Ecc_EccentricityOfOrbit
	var trueAnomalyDeg float64 = pamacro.Degrees(pamacro.TrueAnomaly(mcRad, eccentricity))
	var lcDeg float64 = trueAnomalyDeg + cometInfo.Peri_LongitudeOfPerihelion
	var rAu float64 = cometInfo.Axis_SemiMajorAxisOfOrbit *
		(1 - eccentricity*eccentricity) / (1 + eccentricity*math.Cos(pautil.DegreesToRadians(trueAnomalyDeg)))
	var lcNodeRad float64 = pautil.DegreesToRadians(lcDeg - cometInfo.Node_LongitudeOfAscendingNode)
	var psiRad float64 = math.Asin(math.Sin(lcNodeRad) * math.Sin(pautil.DegreesToRadians(cometInfo.Incl_InclinationOfOrbit)))

	var y float64 = math.Sin(lcNodeRad) * math.Cos(pautil.DegreesToRadians(cometInfo.Incl_InclinationOfOrbit))
	var x float64 = math.Cos(lcNodeRad)

	var ldDeg float64 = pamacro.Degrees(math.Atan2(y, x)) + cometInfo.Node_LongitudeOfAscendingNode
	var rdAu float64 = rAu * math.Cos(psiRad)

	var earthLongitudeLeDeg float64 = pamacro.SunLong(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear) + 180.0
	var earthRadiusVectorAu float64 = pamacro.SunDist(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	var leLdRad float64 = pautil.DegreesToRadians(earthLongitudeLeDeg - ldDeg)

	var aRad float64 = util.TernaryAssign(
		rdAu < earthRadiusVectorAu,
		math.Atan2((rdAu*math.Sin(leLdRad)), (earthRadiusVectorAu-rdAu*math.Cos(leLdRad))),
		math.Atan2((earthRadiusVectorAu*math.Sin(-leLdRad)), (rdAu-earthRadiusVectorAu*math.Cos(leLdRad))),
	)

	var cometLongDeg1 float64 = pautil.TernaryAssign(
		rdAu < earthRadiusVectorAu,
		180.0+earthLongitudeLeDeg+pamacro.Degrees(aRad),
		pamacro.Degrees(aRad)+ldDeg,
	)

	var cometLongDeg float64 = cometLongDeg1 - 360*math.Floor(cometLongDeg1/360)
	var cometLatDeg float64 = pamacro.Degrees(
		math.Atan(rdAu * math.Tan(psiRad) * math.Sin(pautil.DegreesToRadians((cometLongDeg1 - ldDeg))) / (earthRadiusVectorAu * math.Sin(-leLdRad))))
	var cometRaHours1 float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(cometLongDeg, 0, 0, cometLatDeg, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear))
	var cometDecDeg1 float64 = pamacro.EclipticDeclination(cometLongDeg, 0, 0, cometLatDeg, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear)
	var cometDistanceAu float64 = math.Sqrt(
		math.Pow(earthRadiusVectorAu, 2) + math.Pow(rAu, 2) -
			2.0*earthRadiusVectorAu*rAu*math.Cos(pautil.DegreesToRadians((lcDeg-earthLongitudeLeDeg)))*math.Cos(psiRad))

	var cometRaHour int = pamacro.DecimalHoursHour(cometRaHours1 + 0.008333)
	var cometRaMin int = pamacro.DecimalHoursMinute(cometRaHours1 + 0.008333)
	var cometDecDeg float64 = pamacro.DecimalDegreesDegrees(cometDecDeg1 + 0.008333)
	var cometDecMin float64 = pamacro.DecimalDegreesMinutes(cometDecDeg1 + 0.008333)
	var cometDistEarth float64 = pautil.RoundTo(cometDistanceAu, 2)

	return patype.EllipticalCometPosition{
		RaHour: float64(cometRaHour), RaMin: float64(cometRaMin), DecDeg: cometDecDeg, DecMin: cometDecMin, DistEarth: cometDistEarth,
	}
}

/* Calculate position of a parabolic comet. */
func PositionOfParabolicComet(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int, cometName string,
) patype.ParabolicCometPosition {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var greenwichDateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var greenwichDateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var greenwichDateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var cometInfo padata.CometDataParabolic = padata.GetCometDataParabolic(cometName)

	var perihelionEpochDay float64 = cometInfo.EpochPeriDay
	var perihelionEpochMonth int = cometInfo.EpochPeriMonth
	var perihelionEpochYear int = cometInfo.EpochPeriYear
	var qAu float64 = cometInfo.PeriDist
	var inclinationDeg float64 = cometInfo.Incl
	var perihelionDeg float64 = cometInfo.ArgPeri
	var nodeDeg float64 = cometInfo.Node

	var cometLongLatDist patype.CometLongLatDist = pamacro.PCometLongLatDist(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear,
		perihelionEpochDay, perihelionEpochMonth, perihelionEpochYear, qAu, inclinationDeg, perihelionDeg, nodeDeg,
	)

	var cometRaHours float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(cometLongLatDist.LongDeg, 0, 0, cometLongLatDist.LatDeg, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear))
	var cometDecDeg1 float64 = pamacro.EclipticDeclination(
		cometLongLatDist.LongDeg, 0, 0, cometLongLatDist.LatDeg, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear)

	var cometRaHour int = pamacro.DecimalHoursHour(cometRaHours)
	var cometRaMin int = pamacro.DecimalHoursMinute(cometRaHours)
	var cometRaSec float64 = pamacro.DecimalHoursSecond(cometRaHours)
	var cometDecDeg float64 = pamacro.DecimalDegreesDegrees(cometDecDeg1)
	var cometDecMin float64 = pamacro.DecimalDegreesMinutes(cometDecDeg1)
	var cometDecSec float64 = pamacro.DecimalDegreesSeconds(cometDecDeg1)
	var cometDistEarth float64 = pautil.RoundTo(cometLongLatDist.DistAu, 2)

	return patype.ParabolicCometPosition{
		RaHour: float64(cometRaHour), RaMin: float64(cometRaMin), RaSec: cometRaSec,
		DecDeg: cometDecDeg, DecMin: cometDecMin, DecSec: cometDecSec,
		DistEarth: cometDistEarth,
	}
}
