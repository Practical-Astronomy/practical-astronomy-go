package lib

import (
	"math"
	padata "practicalastro/lib/data"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Calculate orbital data for binary star. */
func BinaryStarOrbit(greenwichDateDay float64, greenwichDateMonth int, greenwichDateYear int, binaryName string) patype.BinaryStarOrbitalData {
	var binaryInfo padata.BinaryStarData = padata.GetBinaryStarData(binaryName)

	var yYears float64 = float64(greenwichDateYear) +
		(pamacro.CivilDateToJulianDate(greenwichDateDay, float64(greenwichDateMonth), float64(greenwichDateYear))-
			pamacro.CivilDateToJulianDate(0, 1, float64(greenwichDateYear)))/365.242191 - binaryInfo.EpochPeri
	var mDeg float64 = 360 * yYears / binaryInfo.Period
	var mRad float64 = pautil.DegreesToRadians(mDeg - 360*math.Floor(mDeg/360))
	var eccentricity float64 = binaryInfo.Ecc
	var trueAnomalyRad float64 = pamacro.TrueAnomaly(mRad, eccentricity)
	var rArcsec float64 = (1 - eccentricity*math.Cos(pamacro.EccentricAnomaly(mRad, eccentricity))) * binaryInfo.Axis
	var taPeriRad float64 = trueAnomalyRad + pautil.DegreesToRadians(binaryInfo.LongPeri)

	var y float64 = math.Sin(taPeriRad) * math.Cos(pautil.DegreesToRadians(binaryInfo.Incl))
	var x float64 = math.Cos(taPeriRad)
	var a_deg float64 = pamacro.Degrees(math.Atan2(y, x))
	var theta_deg1 float64 = a_deg + binaryInfo.PaNode
	var theta_deg2 float64 = theta_deg1 - 360*math.Floor(theta_deg1/360)
	var rho_arcsec float64 = rArcsec * math.Cos(taPeriRad) / math.Cos(pautil.DegreesToRadians((theta_deg2 - binaryInfo.PaNode)))

	var position_angle_deg float64 = pautil.RoundTo(theta_deg2, 1)
	var separation_arcsec float64 = pautil.RoundTo(rho_arcsec, 2)

	return patype.BinaryStarOrbitalData{PositionAngleDeg: position_angle_deg, SeparationArcsec: separation_arcsec}
}
