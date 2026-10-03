package macros

import (
	"math"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

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
