package lib

import (
	"math"
	padata "practicalastro/lib/data"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Calculate approximate position of a planet. */
func ApproximatePositionOfPlanet(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int, planetName string,
) patype.PlanetPosition {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var planetInfo padata.PlanetRecord = padata.GetPlanetData(planetName)

	var gDateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var gDateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var gDateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var utHours float64 = pamacro.LocalCivilTimeToUniversalTime(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var gDays float64 = pamacro.CivilDateToJulianDate(gDateDay+(utHours/24), float64(gDateMonth), float64(gDateYear)) - pamacro.CivilDateToJulianDate(0, 1, 2010)
	var npDeg1 float64 = 360 * gDays / (365.242191 * planetInfo.Tp_PeriodOrbit)
	var npDeg2 float64 = npDeg1 - 360*math.Floor(npDeg1/360)
	var mpDeg float64 = npDeg2 + planetInfo.Long_LongitudeEpoch - planetInfo.Peri_LongitudePerihelion
	var lpDeg1 float64 = npDeg2 + (360 * planetInfo.Ecc_EccentricityOrbit * math.Sin(pautil.DegreesToRadians(mpDeg)) / math.Pi) + planetInfo.Long_LongitudeEpoch
	var lpDeg2 float64 = lpDeg1 - 360*math.Floor(lpDeg1/360)
	var planetTrueAnomalyDeg float64 = lpDeg2 - planetInfo.Peri_LongitudePerihelion
	var rAu float64 = planetInfo.Axis_AxisOrbit * (1 - math.Pow(planetInfo.Ecc_EccentricityOrbit, 2)) /
		(1 + planetInfo.Ecc_EccentricityOrbit*math.Cos(pautil.DegreesToRadians(planetTrueAnomalyDeg)))

	var earthInfo padata.PlanetRecord = padata.GetPlanetData("Earth")

	var neDeg1 float64 = 360 * gDays / (365.242191 * earthInfo.Tp_PeriodOrbit)
	var neDeg2 float64 = neDeg1 - 360*math.Floor(neDeg1/360)
	var meDeg float64 = neDeg2 + earthInfo.Long_LongitudeEpoch - earthInfo.Peri_LongitudePerihelion
	var leDeg1 float64 = neDeg2 + earthInfo.Long_LongitudeEpoch + 360*earthInfo.Ecc_EccentricityOrbit*math.Sin(pautil.DegreesToRadians(meDeg))/math.Pi
	var leDeg2 float64 = leDeg1 - 360*math.Floor(leDeg1/360)
	var earthTrueAnomalyDeg float64 = leDeg2 - earthInfo.Peri_LongitudePerihelion
	var rAu2 float64 = earthInfo.Axis_AxisOrbit * (1 - math.Pow(earthInfo.Ecc_EccentricityOrbit, 2)) /
		(1 + earthInfo.Ecc_EccentricityOrbit*math.Cos(pautil.DegreesToRadians(earthTrueAnomalyDeg)))
	var lpNodeRad float64 = pautil.DegreesToRadians(lpDeg2 - planetInfo.Node_LongitudeAscendingNode)
	var psiRad float64 = math.Asin(math.Sin(lpNodeRad) * math.Sin(pautil.DegreesToRadians(planetInfo.Incl_OrbitalInclination)))
	var y float64 = math.Sin(lpNodeRad) * math.Cos(pautil.DegreesToRadians(planetInfo.Incl_OrbitalInclination))
	var x float64 = math.Cos(lpNodeRad)
	var ldDeg float64 = pamacro.Degrees(math.Atan2(y, x)) + planetInfo.Node_LongitudeAscendingNode
	var rdAu float64 = rAu * math.Cos(psiRad)
	var leLdRad float64 = pautil.DegreesToRadians(leDeg2 - ldDeg)
	var atan2Type1 float64 = math.Atan2((rdAu * math.Sin(leLdRad)), (rAu2 - rdAu*math.Cos(leLdRad)))
	var atan2Type2 float64 = math.Atan2((rAu2 * math.Sin(-leLdRad)), (rdAu - rAu2*math.Cos(leLdRad)))

	var aRad float64 = pautil.TernaryAssign(rdAu < 1, atan2Type1, atan2Type2)

	var lamdaDeg1 float64 = pautil.TernaryAssign(
		rdAu < 1,
		180+leDeg2+pamacro.Degrees(aRad),
		pamacro.Degrees(aRad)+ldDeg,
	)

	var lamdaDeg2 float64 = lamdaDeg1 - 360*math.Floor(lamdaDeg1/360)
	var betaDeg float64 = pamacro.Degrees(math.Atan(rdAu * math.Tan(psiRad) * math.Sin(pautil.DegreesToRadians(lamdaDeg2-ldDeg)) / (rAu2 * math.Sin(-leLdRad))))
	var raHours float64 = pamacro.DecimalDegreesToDegreeHours(pamacro.EclipticRightAscension(lamdaDeg2, 0, 0, betaDeg, 0, 0, gDateDay, gDateMonth, gDateYear))
	var decDeg float64 = pamacro.EclipticDeclination(lamdaDeg2, 0, 0, betaDeg, 0, 0, gDateDay, gDateMonth, gDateYear)

	var planetRaHour int = pamacro.DecimalHoursHour(raHours)
	var planetRaMin int = pamacro.DecimalHoursMinute(raHours)
	var planetRaSec float64 = pamacro.DecimalHoursSecond(raHours)
	var planetDecDeg float64 = pamacro.DecimalDegreesDegrees(decDeg)
	var planetDecMin float64 = pamacro.DecimalDegreesMinutes(decDeg)
	var planetDecSec float64 = pamacro.DecimalDegreesSeconds(decDeg)

	return patype.PlanetPosition{
		RightAscensionHour: float64(planetRaHour), RightAscensionMinutes: float64(planetRaMin), RightAscensionSeconds: planetRaSec,
		DeclinationDegrees: planetDecDeg, DeclinationMinutes: planetDecMin, DeclinationSeconds: planetDecSec,
	}
}

/* Calculate precise position of a planet. */
func PrecisePositionOfPlanet(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int, planetName string) patype.PlanetPosition {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var coordinateResults patype.PlanetCoordinates = pamacro.PlanetCoordinates(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear, planetName)

	var planetRaHours float64 = pamacro.DecimalDegreesToDegreeHours(
		pamacro.EclipticRightAscension(coordinateResults.Longitude, 0, 0, coordinateResults.Latitude, 0, 0, localDateDay, localDateMonth, localDateYear))
	var planetDecDeg1 float64 = pamacro.EclipticDeclination(
		coordinateResults.Longitude, 0, 0, coordinateResults.Latitude, 0, 0, localDateDay, localDateMonth, localDateYear)

	var planetRaHour int = pamacro.DecimalHoursHour(planetRaHours)
	var planetRaMin int = pamacro.DecimalHoursMinute(planetRaHours)
	var planetRaSec float64 = pamacro.DecimalHoursSecond(planetRaHours)
	var planetDecDeg float64 = pamacro.DecimalDegreesDegrees(planetDecDeg1)
	var planetDecMin float64 = pamacro.DecimalDegreesMinutes(planetDecDeg1)
	var planetDecSec float64 = pamacro.DecimalDegreesSeconds(planetDecDeg1)

	return patype.PlanetPosition{
		RightAscensionHour: float64(planetRaHour), RightAscensionMinutes: float64(planetRaMin), RightAscensionSeconds: planetRaSec,
		DeclinationDegrees: planetDecDeg, DeclinationMinutes: planetDecMin, DeclinationSeconds: planetDecSec,
	}
}

/* Calculate several visual aspects of a planet. */
func VisualAspectsOfAPlanet(
	lctHour float64, lctMin float64, lctSec float64, isDaylightSaving bool, zoneCorrectionHours int,
	localDateDay float64, localDateMonth int, localDateYear int, planetName string,
) patype.PlanetVisualAspects {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var greenwichDateDay float64 = pamacro.LocalCivilTimeGreenwichDay(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var greenwichDateMonth int = int(pamacro.LocalCivilTimeGreenwichMonth(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))
	var greenwichDateYear int = int(pamacro.LocalCivilTimeGreenwichYear(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear))

	var planetCoordInfo patype.PlanetCoordinates = pamacro.PlanetCoordinates(
		lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear, planetName)

	var planetRaRad float64 = pautil.DegreesToRadians(
		pamacro.EclipticRightAscension(planetCoordInfo.Longitude, 0, 0, planetCoordInfo.Latitude, 0, 0, localDateDay, localDateMonth, localDateYear))
	var planetDecRad float64 = pautil.DegreesToRadians(
		pamacro.EclipticDeclination(planetCoordInfo.Longitude, 0, 0, planetCoordInfo.Latitude, 0, 0, localDateDay, localDateMonth, localDateYear))

	var lightTravelTimeHours float64 = planetCoordInfo.DistanceAu * 0.1386

	var planetInfo padata.PlanetRecord = padata.GetPlanetData(planetName)
	var angularDiameterArcsec float64 = planetInfo.Theta0_AngularDiameter / planetCoordInfo.DistanceAu
	var phase1 float64 = 0.5 * (1.0 + math.Cos(pautil.DegreesToRadians(planetCoordInfo.Longitude-planetCoordInfo.HLong1)))

	var sunEclLongDeg float64 = pamacro.SunLong(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var sunRaRad float64 = pautil.DegreesToRadians(
		pamacro.EclipticRightAscension(sunEclLongDeg, 0, 0, 0, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear))
	var sunDecRad float64 = pautil.DegreesToRadians(
		pamacro.EclipticDeclination(sunEclLongDeg, 0, 0, 0, 0, 0, greenwichDateDay, greenwichDateMonth, greenwichDateYear))

	var y float64 = math.Cos(sunDecRad) * math.Sin(sunRaRad-planetRaRad)
	var x float64 = math.Cos(planetDecRad)*math.Sin(sunDecRad) - math.Sin(planetDecRad)*math.Cos(sunDecRad)*math.Cos(sunRaRad-planetRaRad)

	var chiDeg float64 = pamacro.Degrees(math.Atan2(y, x))
	var radiusVectorAu float64 = planetCoordInfo.RVect
	var approximateMagnitude1 float64 = 5.0*math.Log10(radiusVectorAu*planetCoordInfo.DistanceAu/math.Sqrt(phase1)) + planetInfo.V0_VisualMagnitude

	var distanceAu float64 = pautil.RoundTo(planetCoordInfo.DistanceAu, 5)
	var angDiaArcsec float64 = pautil.RoundTo(angularDiameterArcsec, 1)
	var phase float64 = pautil.RoundTo(phase1, 2)
	var lightTimeHour int = pamacro.DecimalHoursHour(lightTravelTimeHours)
	var lightTimeMinutes int = pamacro.DecimalHoursMinute(lightTravelTimeHours)
	var lightTimeSeconds float64 = pamacro.DecimalHoursSecond(lightTravelTimeHours)
	var posAngleBrightLimbDeg float64 = pautil.RoundTo(chiDeg, 1)
	var approximateMagnitude float64 = pautil.RoundTo(approximateMagnitude1, 1)

	return patype.PlanetVisualAspects{
		DistanceAu: distanceAu, AngDiaArcsec: angDiaArcsec, Phase: phase,
		LightTimeHour: float64(lightTimeHour), LightTimeMinutes: float64(lightTimeMinutes), LightTimeSeconds: lightTimeSeconds,
		PosAngleBrightLimbDeg: posAngleBrightLimbDeg, ApproximateMagnitude: approximateMagnitude,
	}
}
