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

/*
Determine if a solar eclipse is likely to occur.

Original macro name: SEOccurrence
*/
func SolarEclipseOccurrence(ds int, zc int, dy float64, mn int, yr int) patype.SolarEclipseStatus {
	var d0 float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var m0 int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var y0 int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))

	var j0 float64 = CivilDateToJulianDate(0.0, 1, float64(y0))
	var dj float64 = CivilDateToJulianDate(d0, float64(m0), float64(y0))
	var k float64 = (float64(y0) - 1900.0 + ((dj - j0) * 1.0 / 365.0)) * 12.3685
	k = Lint(k + 0.5)
	var tn float64 = k / 1236.85
	var tf float64 = (k + 0.5) / 1236.85
	var t float64 = tn
	var l6855Result1 patype.SolarEclipseOccurrence_L6855 = SolarEclipseOccurrenceL6855(t, k)
	var nb float64 = l6855Result1.F
	t = tf
	k += 0.5
	// var l6855Result2 patype.SolarEclipseOccurrence_L6855 = SolarEclipseOccurrenceL6855(t, k) // not used

	var df float64 = math.Abs(nb - 3.141592654*Lint(nb/3.141592654))

	if df > 0.37 {
		df = 3.141592654 - df
	}

	var s patype.SolarEclipseStatus = patype.SolarEclipseStatus_Certain

	if df >= 0.242600766 {
		s = patype.SolarEclipseStatus_Possible
		if df > 0.37 {
			s = patype.SolarEclipseStatus_None
		}
	}

	return s
}

/* Helper function for SolarEclipseOccurrence */
func SolarEclipseOccurrenceL6855(t float64, k float64) patype.SolarEclipseOccurrence_L6855 {
	var t2 float64 = t * t
	var e float64 = 29.53 * k
	var c float64 = 166.56 + (132.87-0.009173*t)*t
	c = pautil.DegreesToRadians(c)
	var b float64 = 0.00058868*k + (0.0001178-0.000000155*t)*t2
	b = b + 0.00033*math.Sin(c) + 0.75933
	var a float64 = k / 12.36886
	var a1 float64 = 359.2242 + 360.0*FPart(a) - (0.0000333+0.00000347*t)*t2
	var a2 float64 = 306.0253 + 360.0*FPart(k/0.9330851)
	a2 += (0.0107306 + 0.00001236*t) * t2
	a = k / 0.9214926
	var f float64 = 21.2964 + 360.0*FPart(a) - (0.0016528+0.00000239*t)*t2
	a1 = UnwindDeg(a1)
	a2 = UnwindDeg(a2)
	f = UnwindDeg(f)
	a1 = pautil.DegreesToRadians(a1)
	a2 = pautil.DegreesToRadians(a2)
	f = pautil.DegreesToRadians(f)

	var dd float64 = (0.1734-0.000393*t)*math.Sin(a1) + 0.0021*math.Sin(2.0*a1)
	dd = dd - 0.4068*math.Sin(a2) + 0.0161*math.Sin(2.0*a2) - 0.0004*math.Sin(3.0*a2)
	dd = dd + 0.0104*math.Sin(2.0*f) - 0.0051*math.Sin(a1+a2)
	dd = dd - 0.0074*math.Sin(a1-a2) + 0.0004*math.Sin(2.0*f+a1)
	dd = dd - 0.0004*math.Sin(2.0*f-a1) - 0.0006*math.Sin(2.0*f+a2) + 0.001*math.Sin(2.0*f-a2)
	dd += 0.0005 * math.Sin(a1+2.0*a2)
	var e1 float64 = math.Floor(e)
	b = b + dd + (e - e1)
	var b1 float64 = math.Floor(b)
	a = e1 + b1
	b -= b1

	return patype.SolarEclipseOccurrence_L6855{F: f, Dd: dd, E1: e1, B1: b1, A: a, B: b}
}

/*
Calculate time of maximum shadow for solar eclipse (UT)

Original macro name: UTMaxSolarEclipse
*/
func UtMaxSolarEclipse(dy float64, mn int, yr int, ds int, zc int, glong float64, glat float64) float64 {
	var tp float64 = 2.0 * math.Pi

	if SolarEclipseOccurrence(ds, zc, dy, mn, yr) == patype.SolarEclipseStatus_None {
		return -99.0
	}

	var dj float64 = NewMoon(ds, zc, dy, mn, yr)
	var gday float64 = JulianDateDay(dj)
	var gmonth int = JulianDateMonth(dj)
	var gyear int = JulianDateYear(dj)
	var igday float64 = math.Floor(gday)
	var xi float64 = gday - igday
	var utnm float64 = xi * 24.0
	var ut float64 = utnm - 1.0
	var ly float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var my float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var by float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hy float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	ut = utnm + 1.0
	var sb float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)) - ly
	var mz float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var bz float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hz float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))

	if sb < 0.0 {
		sb += tp
	}

	var xh float64 = utnm
	var x float64 = my
	var y float64 = by
	var tm float64 = xh - 1.0
	var hp float64 = hy
	var l7390result1 patype.UtMaxSolarEclipseL7390 = UtMaxSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	my = l7390result1.P
	by = l7390result1.Q
	x = mz
	y = bz
	tm = xh + 1.0
	hp = hz
	var l7390result2 patype.UtMaxSolarEclipseL7390 = UtMaxSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	mz = l7390result2.P
	bz = l7390result2.Q

	var x0 float64 = xh + 1.0 - (2.0 * bz / (bz - by))
	var dm float64 = mz - my

	if dm < 0.0 {
		dm += tp
	}

	var lj float64 = (dm - sb) / 2.0
	var mr float64 = my + (dm * (x0 - xh + 1.0) / 2.0)
	ut = x0 - 0.13851852
	var rr float64 = SunDist(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)
	var sr float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	sr += pautil.DegreesToRadians(NutatLong(igday, gmonth, gyear) - 0.00569)
	x = sr
	y = 0.0
	tm = ut
	hp = 0.00004263452 / rr
	var l7390result3 patype.UtMaxSolarEclipseL7390 = UtMaxSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	sr = l7390result3.P
	by -= l7390result3.Q
	bz -= l7390result3.Q
	// var p3 float64 = 0.00004263  // not used
	var zh float64 = (sr - mr) / lj
	var tc float64 = x0 + zh
	var sh float64 = (((bz - by) * (tc - xh - 1.0) / 2.0) + bz) / lj
	var s2 float64 = sh * sh
	var z2 float64 = zh * zh
	// var ps float64 = p3 / (rr * lj)  // not used
	var z1 float64 = (zh * z2 / (z2 + s2)) + x0
	var h0 float64 = (hy + hz) / (2.0 * lj)
	var rm float64 = 0.272446 * h0
	var rn float64 = 0.00465242 / (lj * rr)
	// var hd float64 = h0 * 0.99834  // not used
	// var _ru float64 = (hd - rn + ps) * 1.02  // not used
	// var _rp float64 = (hd + rn + ps) * 1.02  // not used
	// var pj float64 = math.Abs(sh * zh / math.Sqrt(s2+z2))  // not used
	var r float64 = rm + rn
	var dd float64 = z1 - x0
	dd = dd*dd - ((z2 - (r * r)) * dd / zh)

	if dd < 0.0 {
		return -99.0
	}

	// var zd float64 = math.Sqrt(dd)  // not used

	return z1
}

/* Helper function for ut_max_solar_eclipse */
func UtMaxSolarEclipse_L7390(x float64, y float64, igday float64, gmonth int, gyear int, tm float64, glong float64, glat float64, hp float64) patype.UtMaxSolarEclipseL7390 {
	var paa float64 = EclipticRightAscension(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var qaa float64 = EclipticDeclination(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var xaa float64 = RightAscensionToHourAngle(DecimalDegreesToDegreeHours(paa), 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var pbb float64 = ParallaxHa(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var qbb float64 = ParallaxDec(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var xbb float64 = HourAngleToRightAscension(pbb, 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var p float64 = pautil.DegreesToRadians(EqeLong(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))
	var q float64 = pautil.DegreesToRadians(EqeLat(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))

	return patype.UtMaxSolarEclipseL7390{Paa: paa, Qaa: qaa, Xaa: xaa, Pbb: pbb, Qbb: qbb, Xbb: xbb, P: p, Q: q}
}

/*
Calculate time of first contact for solar eclipse (UT)

Original macro name: UTFirstContactSolarEclipse
*/
func UtFirstContactSolarEclipse(dy float64, mn int, yr int, ds int, zc int, glong float64, glat float64) float64 {
	var tp float64 = 2.0 * math.Pi

	if SolarEclipseOccurrence(ds, zc, dy, mn, yr) == patype.SolarEclipseStatus_None {
		return -99.0
	}

	var dj float64 = NewMoon(ds, zc, dy, mn, yr)
	var gday float64 = JulianDateDay(dj)
	var gmonth int = JulianDateMonth(dj)
	var gyear int = JulianDateYear(dj)
	var igday float64 = math.Floor(gday)
	var xi float64 = gday - igday
	var utnm float64 = xi * 24.0
	var ut float64 = utnm - 1.0
	var ly float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var my float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var by float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hy float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	ut = utnm + 1.0
	var sb float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)) - ly
	var mz float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var bz float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hz float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))

	if sb < 0.0 {
		sb += tp
	}

	var xh float64 = utnm
	var x float64 = my
	var y float64 = by
	var tm float64 = xh - 1.0
	var hp float64 = hy
	var l7390result1 patype.UtFirstContactSolarEclipseL7390 = UtFirstContactSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	my = l7390result1.P
	by = l7390result1.Q
	x = mz
	y = bz
	tm = xh + 1.0
	hp = hz
	var l7390result2 patype.UtFirstContactSolarEclipseL7390 = UtFirstContactSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	mz = l7390result2.P
	bz = l7390result2.Q

	var x0 float64 = xh + 1.0 - (2.0 * bz / (bz - by))
	var dm float64 = mz - my

	if dm < 0.0 {
		dm += tp
	}

	var lj float64 = (dm - sb) / 2.0
	var mr float64 = my + (dm * (x0 - xh + 1.0) / 2.0)
	ut = x0 - 0.13851852
	var rr float64 = SunDist(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)
	var sr float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	sr += pautil.DegreesToRadians(NutatLong(igday, gmonth, gyear) - 0.00569)
	x = sr
	y = 0.0
	tm = ut
	hp = 0.00004263452 / rr
	var l7390result3 patype.UtFirstContactSolarEclipseL7390 = UtFirstContactSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	sr = l7390result3.P
	by -= l7390result3.Q
	bz -= l7390result3.Q
	// var p3 float64 = 0.00004263 // not used
	var zh float64 = (sr - mr) / lj
	var tc float64 = x0 + zh
	var sh float64 = (((bz - by) * (tc - xh - 1.0) / 2.0) + bz) / lj
	var s2 float64 = sh * sh
	var z2 float64 = zh * zh
	// var ps float64 = p3 / (rr * lj) // not used
	var z1 float64 = (zh * z2 / (z2 + s2)) + x0
	var h0 float64 = (hy + hz) / (2.0 * lj)
	var rm float64 = 0.272446 * h0
	var rn float64 = 0.00465242 / (lj * rr)
	// var hd float64 = h0 * 0.99834  // not used
	// var _ru float64 = (hd - rn + ps) * 1.02  // not used
	// var _rp float64 = (hd + rn + ps) * 1.02  // not used
	// var pj float64 = math.Abs(sh * zh / math.Sqrt(s2+z2))  // not used
	var r float64 = rm + rn
	var dd float64 = z1 - x0
	dd = dd*dd - ((z2 - (r * r)) * dd / zh)

	if dd < 0.0 {
		return -99.0
	}

	var zd float64 = math.Sqrt(dd)
	var z6 float64 = z1 - zd

	if z6 < 0.0 {
		z6 += 24.0
	}

	return z6
}

/* Helper function for UTFirstContactSolarEclipse */
func UtFirstContactSolarEclipse_L7390(
	x float64, y float64, igday float64, gmonth int, gyear int, tm float64, glong float64, glat float64, hp float64,
) patype.UtFirstContactSolarEclipseL7390 {
	var paa float64 = EclipticRightAscension(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var qaa float64 = EclipticDeclination(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var xaa float64 = RightAscensionToHourAngle(DecimalDegreesToDegreeHours(paa), 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var pbb float64 = ParallaxHa(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var qbb float64 = ParallaxDec(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var xbb float64 = HourAngleToRightAscension(pbb, 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var p float64 = pautil.DegreesToRadians(EqeLong(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))
	var q float64 = pautil.DegreesToRadians(EqeLat(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))

	return patype.UtFirstContactSolarEclipseL7390{Paa: paa, Qaa: qaa, Xaa: xaa, Pbb: pbb, Qbb: qbb, Xbb: xbb, P: p, Q: q}
}

/*
Calculate time of last contact for solar eclipse (UT)

Original macro name: UTLastContactSolarEclipse
*/
func UtLastContactSolarEclipse(dy float64, mn int, yr int, ds int, zc int, glong float64, glat float64) float64 {
	var tp float64 = 2.0 * math.Pi

	if SolarEclipseOccurrence(ds, zc, dy, mn, yr) == patype.SolarEclipseStatus_None {
		return -99.0
	}

	var dj float64 = NewMoon(ds, zc, dy, mn, yr)
	var gday float64 = JulianDateDay(dj)
	var gmonth int = JulianDateMonth(dj)
	var gyear int = JulianDateYear(dj)
	var igday float64 = math.Floor(gday)
	var xi float64 = gday - igday
	var utnm float64 = xi * 24.0
	var ut float64 = utnm - 1.0
	var ly float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var my float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var by float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hy float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	ut = utnm + 1.0
	var sb float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)) - ly
	var mz float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var bz float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hz float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))

	if sb < 0.0 {
		sb += tp
	}

	var xh float64 = utnm
	var x float64 = my
	var y float64 = by
	var tm float64 = xh - 1.0
	var hp float64 = hy
	var l7390result1 patype.UtLastContactSolarEclipseL7390 = UtLastContactSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	my = l7390result1.P
	by = l7390result1.Q
	x = mz
	y = bz
	tm = xh + 1.0
	hp = hz
	var l7390result2 patype.UtLastContactSolarEclipseL7390 = UtLastContactSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	mz = l7390result2.P
	bz = l7390result2.Q

	var x0 float64 = xh + 1.0 - (2.0 * bz / (bz - by))
	var dm float64 = mz - my

	if dm < 0.0 {
		dm += tp
	}

	var lj float64 = (dm - sb) / 2.0
	var mr float64 = my + (dm * (x0 - xh + 1.0) / 2.0)
	ut = x0 - 0.13851852
	var rr float64 = SunDist(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)
	var sr float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	sr += pautil.DegreesToRadians(NutatLong(igday, gmonth, gyear) - 0.00569)
	x = sr
	y = 0.0
	tm = ut
	hp = 0.00004263452 / rr
	var l7390result3 patype.UtLastContactSolarEclipseL7390 = UtLastContactSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	sr = l7390result3.P
	by -= l7390result3.Q
	bz -= l7390result3.Q
	// var p3 float64 = 0.00004263  // not used
	var zh float64 = (sr - mr) / lj
	var tc float64 = x0 + zh
	var sh float64 = (((bz - by) * (tc - xh - 1.0) / 2.0) + bz) / lj
	var s2 float64 = sh * sh
	var z2 float64 = zh * zh
	// var ps float64 = p3 / (rr * lj)  // not used
	var z1 float64 = (zh * z2 / (z2 + s2)) + x0
	var h0 float64 = (hy + hz) / (2.0 * lj)
	var rm float64 = 0.272446 * h0
	var rn float64 = 0.00465242 / (lj * rr)
	// var hd float64 = h0 * 0.99834  // not used
	// var _ru float64 = (hd - rn + ps) * 1.02  // not used
	// var _rp float64 = (hd + rn + ps) * 1.02  // not used
	// var pj float64 = math.Abs(sh * zh / math.Sqrt(s2+z2))  // not used
	var r float64 = rm + rn
	var dd float64 = z1 - x0
	dd = dd*dd - ((z2 - (r * r)) * dd / zh)

	if dd < 0.0 {
		return -99.0
	}

	var zd float64 = math.Sqrt(dd)
	var z7 float64 = z1 + zd - Lint((z1+zd)/24.0)*24.0

	return z7
}

/* Helper function for ut_last_contact_solar_eclipse */
func UtLastContactSolarEclipse_L7390(
	x float64, y float64, igday float64, gmonth int, gyear int, tm float64, glong float64, glat float64, hp float64,
) patype.UtLastContactSolarEclipseL7390 {
	var paa float64 = EclipticRightAscension(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var qaa float64 = EclipticDeclination(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var xaa float64 = RightAscensionToHourAngle(DecimalDegreesToDegreeHours(paa), 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var pbb float64 = ParallaxHa(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var qbb float64 = ParallaxDec(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var xbb float64 = HourAngleToRightAscension(pbb, 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var p float64 = pautil.DegreesToRadians(EqeLong(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))
	var q float64 = pautil.DegreesToRadians(EqeLat(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))

	return patype.UtLastContactSolarEclipseL7390{Paa: paa, Qaa: qaa, Xaa: xaa, Pbb: pbb, Qbb: qbb, Xbb: xbb, P: p, Q: q}
}

/*
Calculate magnitude of solar eclipse.

Original macro name: MagSolarEclipse
*/
func MagSolarEclipse(dy float64, mn int, yr int, ds int, zc int, glong float64, glat float64) float64 {
	var tp float64 = 2.0 * math.Pi

	if SolarEclipseOccurrence(ds, zc, dy, mn, yr) == patype.SolarEclipseStatus_None {
		return -99.0
	}

	var dj float64 = NewMoon(ds, zc, dy, mn, yr)
	var gday float64 = JulianDateDay(dj)
	var gmonth int = JulianDateMonth(dj)
	var gyear int = JulianDateYear(dj)
	var igday float64 = math.Floor(gday)
	var xi float64 = gday - igday
	var utnm float64 = xi * 24.0
	var ut float64 = utnm - 1.0
	var ly float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var my float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var by float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hy float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	ut = utnm + 1.0
	var sb float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)) - ly
	var mz float64 = pautil.DegreesToRadians(MoonLongitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var bz float64 = pautil.DegreesToRadians(MoonLatitude(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	var hz float64 = pautil.DegreesToRadians(MoonHorizontalParallax(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))

	if sb < 0.0 {
		sb += tp
	}

	var xh float64 = utnm
	var x float64 = my
	var y float64 = by
	var tm float64 = xh - 1.0
	var hp float64 = hy
	var l7390result1 patype.MagSolarEclipseL7390 = MagSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	my = l7390result1.P
	by = l7390result1.Q
	x = mz
	y = bz
	tm = xh + 1.0
	hp = hz
	var l7390result2 patype.MagSolarEclipseL7390 = MagSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	mz = l7390result2.P
	bz = l7390result2.Q

	var x0 float64 = xh + 1.0 - (2.0 * bz / (bz - by))
	var dm float64 = mz - my

	if dm < 0.0 {
		dm += tp
	}

	var lj float64 = (dm - sb) / 2.0
	var mr float64 = my + (dm * (x0 - xh + 1.0) / 2.0)
	ut = x0 - 0.13851852
	var rr float64 = SunDist(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear)
	var sr float64 = pautil.DegreesToRadians(SunLong(ut, 0.0, 0.0, 0, 0, igday, gmonth, gyear))
	sr += pautil.DegreesToRadians(NutatLong(igday, gmonth, gyear) - 0.00569)
	x = sr
	y = 0.0
	tm = ut
	hp = 0.00004263452 / rr
	var l7390result3 patype.MagSolarEclipseL7390 = MagSolarEclipse_L7390(x, y, igday, gmonth, gyear, tm, glong, glat, hp)
	sr = l7390result3.P
	by -= l7390result3.Q
	bz -= l7390result3.Q
	// var p3 float64 = 0.00004263  // not used
	var zh float64 = (sr - mr) / lj
	var tc float64 = x0 + zh
	var sh float64 = (((bz - by) * (tc - xh - 1.0) / 2.0) + bz) / lj
	var s2 float64 = sh * sh
	var z2 float64 = zh * zh
	// var ps float64 = p3 / (rr * lj)  // not used
	var z1 float64 = (zh * z2 / (z2 + s2)) + x0
	var h0 float64 = (hy + hz) / (2.0 * lj)
	var rm float64 = 0.272446 * h0
	var rn float64 = 0.00465242 / (lj * rr)
	// var hd float64 = h0 * 0.99834  // not used
	// var _ru float64 = (hd - rn + ps) * 1.02  // not used
	// var _rp float64 = (hd + rn + ps) * 1.02  // not used
	var pj float64 = math.Abs(sh * zh / math.Sqrt(s2+z2))
	var r float64 = rm + rn
	var dd float64 = z1 - x0
	dd = dd*dd - ((z2 - (r * r)) * dd / zh)

	if dd < 0.0 {
		return -99.0
	}

	// var zd float64 = math.Sqrt(dd)  // not used

	var mg float64 = (rm + rn - pj) / (2.0 * rn)

	return mg
}

/* Helper function for mag_solar_eclipse */
func MagSolarEclipse_L7390(
	x float64, y float64, igday float64, gmonth int, gyear int, tm float64, glong float64, glat float64, hp float64,
) patype.MagSolarEclipseL7390 {
	var paa float64 = EclipticRightAscension(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var qaa float64 = EclipticDeclination(Degrees(x), 0.0, 0.0, Degrees(y), 0.0, 0.0, igday, gmonth, gyear)
	var xaa float64 = RightAscensionToHourAngle(DecimalDegreesToDegreeHours(paa), 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var pbb float64 = ParallaxHa(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var qbb float64 = ParallaxDec(xaa, 0.0, 0.0, qaa, 0.0, 0.0, patype.CoordinateType_Actual, glat, 0.0, Degrees(hp))
	var xbb float64 = HourAngleToRightAscension(pbb, 0.0, 0.0, tm, 0.0, 0.0, 0, 0, igday, gmonth, gyear, glong)
	var p float64 = pautil.DegreesToRadians(EqeLong(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))
	var q float64 = pautil.DegreesToRadians(EqeLat(xbb, 0.0, 0.0, qbb, 0.0, 0.0, igday, gmonth, gyear))

	return patype.MagSolarEclipseL7390{Paa: paa, Qaa: qaa, Xaa: xaa, Pbb: pbb, Qbb: qbb, Xbb: xbb, P: p, Q: q}
}
