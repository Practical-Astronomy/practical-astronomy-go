package macros

import (
	"math"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

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
Calculate distance from the Earth to the Moon (km)

Original macro name: MoonDist
*/
func MoonDist(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var hp float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lh, lm, ls, ds, zc, dy, mn, yr))
	var r float64 = 6378.14 / math.Sin(hp)

	return r
}

/*
Calculate the Moon's angular diameter (degrees)

Original macro name: MoonSize
*/
func MoonSize(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var hp float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lh, lm, ls, ds, zc, dy, mn, yr))
	var r float64 = 6378.14 / math.Sin(hp)
	var th float64 = 384401.0 * 0.5181 / r

	return th
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
Longitude, latitude, and horizontal parallax of the Moon.

Original macro names: MoonLong, MoonLat, MoonHP
*/
func MoonLongLatHp(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) patype.MoonLongLatHP {
	var ut float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var t float64 = ((CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020.0) / 36525.0) + (ut / 876600.0)
	var t2 float64 = t * t

	var m1 float64 = 27.32158213
	var m2 float64 = 365.2596407
	var m3 float64 = 27.55455094
	var m4 float64 = 29.53058868
	var m5 float64 = 27.21222039
	var m6 float64 = 6798.363307
	var q float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020.0 + (ut / 24.0)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360.0 * (m1 - math.Floor(m1))
	m2 = 360.0 * (m2 - math.Floor(m2))
	m3 = 360.0 * (m3 - math.Floor(m3))
	m4 = 360.0 * (m4 - math.Floor(m4))
	m5 = 360.0 * (m5 - math.Floor(m5))
	m6 = 360.0 * (m6 - math.Floor(m6))

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
	ms -= 0.001778 * s1
	md = md + 0.000817*s1 + s3 + 0.002541*s2
	mf = mf + s3 - 0.024691*s2 - 0.004328*s4
	me1 = me1 + 0.002011*s1 + s3 + 0.001964*s2
	var e float64 = 1.0 - (0.002495+0.00000752*t)*t
	var e2 float64 = e * e
	ml = pautil.DegreesToRadians(ml)
	ms = pautil.DegreesToRadians(ms)
	na = pautil.DegreesToRadians(na)
	me1 = pautil.DegreesToRadians(me1)
	mf = pautil.DegreesToRadians(mf)
	md = pautil.DegreesToRadians(md)

	// Longitude-specific
	var l float64 = 6.28875*math.Sin(md) + 1.274018*math.Sin(2.0*me1-md)
	l = l + 0.658309*math.Sin(2.0*me1) + 0.213616*math.Sin(2.0*md)
	l = l - e*0.185596*math.Sin(ms) - 0.114336*math.Sin(2.0*mf)
	l += 0.058793 * math.Sin(2.0*(me1-md))
	l = l + 0.057212*e*math.Sin(2.0*me1-ms-md) + 0.05332*math.Sin(2.0*me1+md)
	l = l + 0.045874*e*math.Sin(2.0*me1-ms) + 0.041024*e*math.Sin(md-ms)
	l = l - 0.034718*math.Sin(me1) - e*0.030465*math.Sin(ms+md)
	l = l + 0.015326*math.Sin(2.0*(me1-mf)) - 0.012528*math.Sin(2.0*mf+md)
	l = l - 0.01098*math.Sin(2.0*mf-md) + 0.010674*math.Sin(4.0*me1-md)
	l = l + 0.010034*math.Sin(3.0*md) + 0.008548*math.Sin(4.0*me1-2.0*md)
	l = l - e*0.00791*math.Sin(ms-md+2.0*me1) - e*0.006783*math.Sin(2.0*me1+ms)
	l = l + 0.005162*math.Sin(md-me1) + e*0.005*math.Sin(ms+me1)
	l = l + 0.003862*math.Sin(4.0*me1) + e*0.004049*math.Sin(md-ms+2.0*me1)
	l = l + 0.003996*math.Sin(2.0*(md+me1)) + 0.003665*math.Sin(2.0*me1-3.0*md)
	l = l + e*0.002695*math.Sin(2.0*md-ms) + 0.002602*math.Sin(md-2.0*(mf+me1))
	l = l + e*0.002396*math.Sin(2.0*(me1-md)-ms) - 0.002349*math.Sin(md+me1)
	l = l + e2*0.002249*math.Sin(2.0*(me1-ms)) - e*0.002125*math.Sin(2.0*md+ms)
	l = l - e2*0.002079*math.Sin(2.0*ms) + e2*0.002059*math.Sin(2.0*(me1-ms)-md)
	l = l - 0.001773*math.Sin(md+2.0*(me1-mf)) - 0.001595*math.Sin(2.0*(mf+me1))
	l = l + e*0.00122*math.Sin(4.0*me1-ms-md) - 0.00111*math.Sin(2.0*(md+mf))
	l = l + 0.000892*math.Sin(md-3.0*me1) - e*0.000811*math.Sin(ms+md+2.0*me1)
	l += e * 0.000761 * math.Sin(4.0*me1-ms-2.0*md)
	l += e2 * 0.000704 * math.Sin(md-2.0*(ms+me1))
	l += e * 0.000693 * math.Sin(ms-2.0*(md-me1))
	l += e * 0.000598 * math.Sin(2.0*(me1-mf)-ms)
	l = l + 0.00055*math.Sin(md+4.0*me1) + 0.000538*math.Sin(4.0*md)
	l = l + e*0.000521*math.Sin(4.0*me1-ms) + 0.000486*math.Sin(2.0*md-me1)
	l += e2 * 0.000717 * math.Sin(md-2.0*ms)
	var mm float64 = Unwind(ml + pautil.DegreesToRadians(l))

	// Latitude-specific
	var g float64 = 5.128189*math.Sin(mf) + 0.280606*math.Sin(md+mf)
	g = g + 0.277693*math.Sin(md-mf) + 0.173238*math.Sin(2.0*me1-mf)
	g = g + 0.055413*math.Sin(2.0*me1+mf-md) + 0.046272*math.Sin(2.0*me1-mf-md)
	g = g + 0.032573*math.Sin(2.0*me1+mf) + 0.017198*math.Sin(2.0*md+mf)
	g = g + 0.009267*math.Sin(2.0*me1+md-mf) + 0.008823*math.Sin(2.0*md-mf)
	g = g + e*0.008247*math.Sin(2.0*me1-ms-mf) + 0.004323*math.Sin(2.0*(me1-md)-mf)
	g = g + 0.0042*math.Sin(2.0*me1+mf+md) + e*0.003372*math.Sin(mf-ms-2.0*me1)
	g += e * 0.002472 * math.Sin(2.0*me1+mf-ms-md)
	g += e * 0.002222 * math.Sin(2.0*me1+mf-ms)
	g += e * 0.002072 * math.Sin(2.0*me1-mf-ms-md)
	g = g + e*0.001877*math.Sin(mf-ms+md) + 0.001828*math.Sin(4.0*me1-mf-md)
	g = g - e*0.001803*math.Sin(mf+ms) - 0.00175*math.Sin(3.0*mf)
	g = g + e*0.00157*math.Sin(md-ms-mf) - 0.001487*math.Sin(mf+me1)
	g = g - e*0.001481*math.Sin(mf+ms+md) + e*0.001417*math.Sin(mf-ms-md)
	g = g + e*0.00135*math.Sin(mf-ms) + 0.00133*math.Sin(mf-me1)
	g = g + 0.001106*math.Sin(mf+3.0*md) + 0.00102*math.Sin(4.0*me1-mf)
	g = g + 0.000833*math.Sin(mf+4.0*me1-md) + 0.000781*math.Sin(md-3.0*mf)
	g = g + 0.00067*math.Sin(mf+4.0*me1-2.0*md) + 0.000606*math.Sin(2.0*me1-3.0*mf)
	g += 0.000597 * math.Sin(2.0*(me1+md)-mf)
	g = g + e*0.000492*math.Sin(2.0*me1+md-ms-mf) + 0.00045*math.Sin(2.0*(md-me1)-mf)
	g = g + 0.000439*math.Sin(3.0*md-mf) + 0.000423*math.Sin(mf+2.0*(me1+md))
	g = g + 0.000422*math.Sin(2.0*me1-mf-3.0*md) - e*0.000367*math.Sin(ms+mf+2.0*me1-md)
	g = g - e*0.000353*math.Sin(ms+mf+2.0*me1) + 0.000331*math.Sin(mf+4.0*me1)
	g += e * 0.000317 * math.Sin(2.0*me1+mf-ms+md)
	g = g + e2*0.000306*math.Sin(2.0*(me1-ms)-mf) - 0.000283*math.Sin(md+3.0*mf)
	var w1 float64 = 0.0004664 * math.Cos(na)
	var w2 float64 = 0.0000754 * math.Cos(c)
	var bm float64 = pautil.DegreesToRadians(g) * (1.0 - w1 - w2)

	// Horizontal parallax-specific
	var pm float64 = 0.950724 + 0.051818*math.Cos(md) + 0.009531*math.Cos(2.0*me1-md)
	pm = pm + 0.007843*math.Cos(2.0*me1) + 0.002824*math.Cos(2.0*md)
	pm = pm + 0.000857*math.Cos(2.0*me1+md) + e*0.000533*math.Cos(2.0*me1-ms)
	pm += e * 0.000401 * math.Cos(2.0*me1-md-ms)
	pm = pm + e*0.00032*math.Cos(md-ms) - 0.000271*math.Cos(me1)
	pm = pm - e*0.000264*math.Cos(ms+md) - 0.000198*math.Cos(2.0*mf-md)
	pm = pm + 0.000173*math.Cos(3.0*md) + 0.000167*math.Cos(4.0*me1-md)
	pm = pm - e*0.000111*math.Cos(ms) + 0.000103*math.Cos(4.0*me1-2.0*md)
	pm = pm - 0.000084*math.Cos(2.0*md-2.0*me1) - e*0.000083*math.Cos(2.0*me1+ms)
	pm = pm + 0.000079*math.Cos(2.0*me1+2.0*md) + 0.000072*math.Cos(4.0*me1)
	pm = pm + e*0.000064*math.Cos(2.0*me1-ms+md) - e*0.000063*math.Cos(2.0*me1+ms-md)
	pm = pm + e*0.000041*math.Cos(ms+me1) + e*0.000035*math.Cos(2.0*md-ms)
	pm = pm - 0.000033*math.Cos(3.0*md-2.0*me1) - 0.00003*math.Cos(md+me1)
	pm = pm - 0.000029*math.Cos(2.0*(mf-me1)) - e*0.000029*math.Cos(2.0*md+ms)
	pm = pm + e2*0.000026*math.Cos(2.0*(me1-ms)) - 0.000023*math.Cos(2.0*(mf-me1)+md)
	pm += e * 0.000019 * math.Cos(4.0*me1-ms-md)

	var moonLongDeg float64 = Degrees(mm)
	var moonLatDeg float64 = Degrees(bm)
	var moonHorPara float64 = pm

	return patype.MoonLongLatHP{LongDeg: moonLongDeg, LatDeg: moonLatDeg, HorPara: moonHorPara}
}

/*
Calculate current phase of Moon.

Original macro name: MoonPhase
*/
func MoonPhase(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var moonResult patype.MoonLongLatHP = MoonLongLatHp(lh, lm, ls, ds, zc, dy, mn, yr)

	var cd float64 = math.Cos(pautil.DegreesToRadians(moonResult.LongDeg-SunLong(lh, lm, ls, ds, zc, dy, mn, yr))) *
		math.Cos(pautil.DegreesToRadians(moonResult.LatDeg))
	var d float64 = math.Acos(cd)
	var sd float64 = math.Sin(d)
	var i float64 = 0.1468 * sd * (1.0 - 0.0549*math.Sin(MoonMeanAnomaly(lh, lm, ls, ds, zc, dy, mn, yr)))
	i /= (1.0 - 0.0167*math.Sin(SunMeanAnomaly(lh, lm, ls, ds, zc, dy, mn, yr)))
	i = 3.141592654 - d - pautil.DegreesToRadians(i)
	var k float64 = (1.0 + math.Cos(i)) / 2.0

	return pautil.RoundTo(k, 2)
}

/*
Calculate the Moon's mean anomaly.

Original macro name: MoonMeanAnomaly
*/
func MoonMeanAnomaly(lh float64, lm float64, ls float64, ds int, zc int, dy float64, mn int, yr int) float64 {
	var ut float64 = LocalCivilTimeToUniversalTime(lh, lm, ls, ds, zc, dy, mn, yr)
	var gd float64 = LocalCivilTimeGreenwichDay(lh, lm, ls, ds, zc, dy, mn, yr)
	var gm int = int(LocalCivilTimeGreenwichMonth(lh, lm, ls, ds, zc, dy, mn, yr))
	var gy int = int(LocalCivilTimeGreenwichYear(lh, lm, ls, ds, zc, dy, mn, yr))
	var t float64 = ((CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020.0) / 36525.0) + (ut / 876600.0)
	var t2 float64 = t * t

	var m1 float64 = 27.32158213
	var m2 float64 = 365.2596407
	var m3 float64 = 27.55455094
	var m4 float64 = 29.53058868
	var m5 float64 = 27.21222039
	var m6 float64 = 6798.363307
	var q float64 = CivilDateToJulianDate(gd, float64(gm), float64(gy)) - 2415020.0 + (ut / 24.0)
	m1 = q / m1
	m2 = q / m2
	m3 = q / m3
	m4 = q / m4
	m5 = q / m5
	m6 = q / m6
	m1 = 360.0 * (m1 - math.Floor(m1))
	m2 = 360.0 * (m2 - math.Floor(m2))
	m3 = 360.0 * (m3 - math.Floor(m3))
	m4 = 360.0 * (m4 - math.Floor(m4))
	m5 = 360.0 * (m5 - math.Floor(m5))
	m6 = 360.0 * (m6 - math.Floor(m6))

	// var ml float64 = 270.434164 + m1 - (0.001133-0.0000019*t)*t2 // not used
	// var ms float64 = 358.475833 + m2 - (0.00015+0.0000033*t)*t2  // not used
	var md float64 = 296.104608 + m3 + (0.009192+0.0000144*t)*t2
	var na float64 = 259.183275 - m6 + (0.002078+0.0000022*t)*t2
	var a float64 = pautil.DegreesToRadians(51.2 + 20.2*t)
	var s1 float64 = math.Sin(a)
	var s2 float64 = math.Sin(pautil.DegreesToRadians(na))
	var b float64 = 346.56 + (132.87-0.0091731*t)*t
	var s3 float64 = 0.003964 * math.Sin(pautil.DegreesToRadians(b))
	// var c float64 = pautil.DegreesToRadians(na + 275.05 - 2.3*t) // not used
	md = md + 0.000817*s1 + s3 + 0.002541*s2

	return pautil.DegreesToRadians(md)
}

/*
Calculate Julian date of New Moon.

Original macro name: NewMoon
*/
func NewMoon(ds int, zc int, dy float64, mn int, yr int) float64 {
	var d0 float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var m0 int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var y0 int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))

	var j0 float64 = CivilDateToJulianDate(0.0, 1, float64(y0)) - 2415020.0
	var dj float64 = CivilDateToJulianDate(d0, float64(m0), float64(y0)) - 2415020.0
	var k float64 = Lint(((float64(y0) - 1900.0 + ((dj - j0) / 365.0)) * 12.3685) + 0.5)
	var tn float64 = k / 1236.85
	var tf float64 = (k + 0.5) / 1236.85
	var t float64 = tn
	var nmfmResult1 patype.NewMoonFullMoonL6855 = NewMoonFullMoonL6855(k, t)
	var ni float64 = nmfmResult1.A
	var nf float64 = nmfmResult1.B
	t = tf
	k += 0.5
	// var nmfmResult2 patype.NewMoonFullMoonL6855 = NewMoonFullMoonL6855(k, t) // unused

	return ni + 2415020.0 + nf
}

/*
Calculate Julian date of Full Moon.

Original macro name: FullMoon
*/
func FullMoon(ds int, zc int, dy float64, mn int, yr int) float64 {
	var d0 float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var m0 int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var y0 int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))

	var j0 float64 = CivilDateToJulianDate(0.0, 1, float64(y0)) - 2415020.0
	var dj float64 = CivilDateToJulianDate(d0, float64(m0), float64(y0)) - 2415020.0
	var k float64 = Lint(((float64(y0) - 1900.0 + ((dj - j0) / 365.0)) * 12.3685) + 0.5)
	var tn float64 = k / 1236.85
	var tf float64 = (k + 0.5) / 1236.85
	var t float64 = tn
	// var nmfmResult1 patype.NewMoonFullMoonL6855 = NewMoonFullMoonL6855(k, t) // not used
	t = tf
	k += 0.5
	var nmfmResult2 patype.NewMoonFullMoonL6855 = NewMoonFullMoonL6855(k, t)
	var fi float64 = nmfmResult2.A
	var ff float64 = nmfmResult2.B

	return fi + 2415020.0 + ff
}

/**
 * Helper function for new_moon() and full_moon() """
 */
func NewMoonFullMoonL6855(k float64, t float64) patype.NewMoonFullMoonL6855 {
	var t2 float64 = t * t
	var e float64 = 29.53 * k
	var c float64 = 166.56 + (132.87-0.009173*t)*t
	c = pautil.DegreesToRadians(c)
	var b float64 = 0.00058868*k + (0.0001178-0.000000155*t)*t2
	b = b + 0.00033*math.Sin(c) + 0.75933
	var a float64 = k / 12.36886
	var a1 float64 = 359.2242 + 360.0*Fract(a) - (0.0000333+0.00000347*t)*t2
	var a2 float64 = 306.0253 + 360.0*Fract(k/0.9330851)
	a2 += (0.0107306 + 0.00001236*t) * t2
	a = k / 0.9214926
	var f float64 = 21.2964 + 360.0*Fract(a) - (0.0016528+0.00000239*t)*t2
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

	return patype.NewMoonFullMoonL6855{A: a, B: b, F: f}
}

/*
Local time of moonrise.

Original macro name: MoonRiseLCT
*/
func MoonRiseLct(dy float64, mn int, yr int, ds int, zc int, gLong float64, gLat float64) float64 {
	var gdy float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var gmn int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var gyr int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var lct float64 = 12.0
	var dy1 float64 = dy
	var mn1 int = mn
	var yr1 int = yr

	var lct6700Result1 patype.MoonRiseLCTL6700 = MoonRiseLctL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
	var lu float64 = lct6700Result1.Lu
	lct = lct6700Result1.Lct

	if lct == -99.0 {
		return lct
	}

	var la float64 = lu

	var x float64
	var ut float64
	var g1 float64 = 0.0
	var gu float64 = 0.0

	for k := 1; k < 9; k++ {
		x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
		ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

		if k == 1 {
			g1 = ut
		} else {
			g1 = gu
		}

		gu = ut
		ut = gu

		var lct6680Result patype.MoonRiseLCTL6680 = MoonRiseLctL6680(x, ds, zc, gdy, gmn, gyr, g1, ut)
		lct = lct6680Result.Lct
		dy1 = lct6680Result.Dy1
		mn1 = lct6680Result.Mn1
		yr1 = lct6680Result.Yr1
		gdy = lct6680Result.Gdy
		gmn = lct6680Result.Gmn
		gyr = lct6680Result.Gyr

		var lct6700Result2 patype.MoonRiseLCTL6700 = MoonRiseLctL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
		lu = lct6700Result2.Lu
		lct = lct6700Result2.Lct

		if lct == -99.0 {
			return lct
		}

		la = lu
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
	ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	lct = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)

	return lct
}

/* Helper function for MoonRiseLCT */
func MoonRiseLctL6680(x float64, ds int, zc int, gdy float64, gmn int, gyr int, g1 float64, ut float64) patype.MoonRiseLCTL6680 {
	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	var lct float64 = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var dy1 float64 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var mn1 int = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var yr1 int = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	gdy = LocalCivilTimeGreenwichDay(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	gmn = int(LocalCivilTimeGreenwichMonth(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	gyr = int(LocalCivilTimeGreenwichYear(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	ut -= 24.0 * math.Floor(ut/24.0)

	return patype.MoonRiseLCTL6680{Ut: ut, Lct: lct, Dy1: dy1, Mn1: mn1, Yr1: yr1, Gdy: gdy, Gmn: gmn, Gyr: gyr}
}

/* Helper function for MoonRiseLCT */
func MoonRiseLctL6700(lct float64, ds int, zc int, dy1 float64, mn1 int, yr1 int, gdy float64, gmn int, gyr int, gLat float64) patype.MoonRiseLCTL6700 {
	var mm float64 = MoonLongitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var bm float64 = MoonLatitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var pm float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	var dp float64 = NutatLong(gdy, gmn, gyr)
	var th float64 = 0.27249 * math.Sin(pm)
	var di float64 = th + 0.0098902 - pm
	var p float64 = DecimalDegreesToDegreeHours(EclipticRightAscension(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr))
	var q float64 = EclipticDeclination(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr)
	var lu float64 = RiseSetLocalSiderealTimeRise(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)

	if ERiseSet(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat) != patype.RiseSetStatus_OK {
		lct = -99.0
	}

	return patype.MoonRiseLCTL6700{Mm: mm, Bm: bm, Pm: pm, Dp: dp, Th: th, Di: di, P: p, Q: q, Lu: lu, Lct: lct}
}

/*
Local date of moonrise.

Original macro names: MoonRiseLcDay, MoonRiseLcMonth, MoonRiseLcYear
*/
func MoonRiseLcDmy(dy float64, mn int, yr int, ds int, zc int, gLong float64, gLat float64) patype.FullDatePrecise {
	var gdy float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var gmn int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var gyr int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var lct float64 = 12.0
	var dy1 float64 = dy
	var mn1 int = mn
	var yr1 int = yr

	var lct6700Result1 patype.MoonRiseLcDMYL6700 = MoonRiseLcDmyL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
	var lu float64 = lct6700Result1.Lu
	lct = lct6700Result1.Lct

	if lct == -99.0 {
		return patype.FullDatePrecise{Month: -99, Day: -99, Year: -99}
	}

	var la float64 = lu

	var x float64
	var ut float64
	var g1 float64 = 0.0
	var gu float64 = 0.0
	for k := 1; k < 9; k++ {
		x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
		ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

		if k == 1 {
			g1 = ut
		} else {
			g1 = gu
		}

		gu = ut
		ut = gu

		var lct6680Result1 patype.MoonRiseLcDMYL6680 = MoonRiseLcDmyL6680(x, ds, zc, gdy, gmn, gyr, g1, ut)
		lct = lct6680Result1.Lct
		dy1 = lct6680Result1.Dy1
		mn1 = lct6680Result1.Mn1
		yr1 = lct6680Result1.Yr1
		gdy = lct6680Result1.Gdy
		gmn = lct6680Result1.Gmn
		gyr = lct6680Result1.Gyr

		var lct6700Result2 patype.MoonRiseLcDMYL6700 = MoonRiseLcDmyL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)

		lu = lct6700Result2.Lu
		lct = lct6700Result2.Lct

		if lct == -99.0 {
			return patype.FullDatePrecise{Month: -99, Day: -99, Year: -99}
		}

		la = lu
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
	ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	dy1 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	mn1 = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	yr1 = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)

	return patype.FullDatePrecise{Month: mn1, Day: dy1, Year: yr1}
}

/* Helper function for MoonRiseLcDMY */
func MoonRiseLcDmyL6680(x float64, ds int, zc int, gdy float64, gmn int, gyr int, g1 float64, ut float64) patype.MoonRiseLcDMYL6680 {
	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	var lct float64 = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var dy1 float64 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var mn1 int = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var yr1 int = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	gdy = LocalCivilTimeGreenwichDay(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	gmn = int(LocalCivilTimeGreenwichMonth(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	gyr = int(LocalCivilTimeGreenwichYear(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	ut -= 24.0 * math.Floor(ut/24.0)

	return patype.MoonRiseLcDMYL6680{Ut: ut, Lct: lct, Dy1: dy1, Mn1: mn1, Yr1: yr1, Gdy: gdy, Gmn: gmn, Gyr: gyr}
}

/* Helper function for MoonRiseLcDMY */
func MoonRiseLcDmyL6700(
	lct float64, ds int, zc int, dy1 float64, mn1 int, yr1 int, gdy float64, gmn int, gyr int, gLat float64,
) patype.MoonRiseLcDMYL6700 {
	var mm float64 = MoonLongitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var bm float64 = MoonLatitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var pm float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	var dp float64 = NutatLong(gdy, gmn, gyr)
	var th float64 = 0.27249 * math.Sin(pm)
	var di float64 = th + 0.0098902 - pm
	var p float64 = DecimalDegreesToDegreeHours(EclipticRightAscension(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr))
	var q float64 = EclipticDeclination(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr)
	var lu float64 = RiseSetLocalSiderealTimeRise(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)

	return patype.MoonRiseLcDMYL6700{Mm: mm, Bm: bm, Pm: pm, Dp: dp, Th: th, Di: di, P: p, Q: q, Lu: lu, Lct: lct}
}

/*
Local azimuth of moonrise.

Original macro name: MoonRiseAz
*/
func MoonRiseAz(dy float64, mn int, yr int, ds int, zc int, gLong float64, gLat float64) float64 {
	var gdy float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var gmn int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var gyr int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var lct float64 = 12.0
	var dy1 float64 = dy
	var mn1 int = mn
	var yr1 int = yr

	var az6700Result1 patype.MoonRiseAzL6700 = MoonRiseAzL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
	var lu float64 = az6700Result1.Lu
	lct = az6700Result1.Lct
	var au float64

	if lct == -99.0 {
		return lct
	}

	var la float64 = lu

	var x float64
	var ut float64
	var g1 float64
	var gu float64 = 0.0
	var aa float64 = 0.0
	for k := 1; k < 9; k++ {
		x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
		ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

		if k == 1 {
			g1 = ut
		} else {
			g1 = gu
		}

		gu = ut
		ut = gu

		var az6680Result1 patype.MoonRiseAzL6680 = MoonRiseAzL6680(x, ds, zc, gdy, gmn, gyr, g1, ut)
		lct = az6680Result1.Lct
		dy1 = az6680Result1.Dy1
		mn1 = az6680Result1.Mn1
		yr1 = az6680Result1.Yr1
		gdy = az6680Result1.Gdy
		gmn = az6680Result1.Gmn
		gyr = az6680Result1.Gyr

		var az6700Result2 patype.MoonRiseAzL6700 = MoonRiseAzL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
		lu = az6700Result2.Lu
		lct = az6700Result2.Lct
		au = az6700Result2.Au

		if lct == -99.0 {
			return lct
		}

		la = lu
		aa = au
	}

	au = aa

	return au
}

/* Helper function for MoonRiseAz */
func MoonRiseAzL6680(x float64, ds int, zc int, gdy float64, gmn int, gyr int, g1 float64, ut float64) patype.MoonRiseAzL6680 {
	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	var lct float64 = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var dy1 float64 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var mn1 int = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var yr1 int = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	gdy = LocalCivilTimeGreenwichDay(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	gmn = int(LocalCivilTimeGreenwichMonth(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	gyr = int(LocalCivilTimeGreenwichYear(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	ut -= 24.0 * math.Floor(ut/24.0)

	return patype.MoonRiseAzL6680{Ut: ut, Lct: lct, Dy1: dy1, Mn1: mn1, Yr1: yr1, Gdy: gdy, Gmn: gmn, Gyr: gyr}
}

/* Helper function for MoonRiseAz */
func MoonRiseAzL6700(lct float64, ds int, zc int, dy1 float64, mn1 int, yr1 int, gdy float64, gmn int, gyr int, gLat float64) patype.MoonRiseAzL6700 {
	var mm float64 = MoonLongitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var bm float64 = MoonLatitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var pm float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	var dp float64 = NutatLong(gdy, gmn, gyr)
	var th float64 = 0.27249 * math.Sin(pm)
	var di float64 = th + 0.0098902 - pm
	var p float64 = DecimalDegreesToDegreeHours(EclipticRightAscension(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr))
	var q float64 = EclipticDeclination(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr)
	var lu float64 = RiseSetLocalSiderealTimeRise(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)
	var au float64 = RiseSetAzimuthRise(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)

	return patype.MoonRiseAzL6700{Mm: mm, Bm: bm, Pm: pm, Dp: dp, Th: th, Di: di, P: p, Q: q, Lu: lu, Lct: lct, Au: au}
}

/*
Local time of moonset.

Original macro name: MoonSetLCT
*/
func MoonSetLct(dy float64, mn int, yr int, ds int, zc int, gLong float64, gLat float64) float64 {
	var gdy float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var gmn int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var gyr int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var lct float64 = 12.0
	var dy1 float64 = dy
	var mn1 int = mn
	var yr1 int = yr

	var lct6700Result1 patype.MoonSetLCTL6700 = MoonSetLctL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
	var lu float64 = lct6700Result1.Lu
	lct = lct6700Result1.Lct

	if lct == -99.0 {
		return lct
	}

	var la float64 = lu

	var x float64
	var ut float64
	var g1 float64 = 0.0
	var gu float64 = 0.0
	for k := 1; k < 9; k++ {
		x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
		ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

		if k == 1 {
			g1 = ut
		} else {
			g1 = gu
		}

		gu = ut
		ut = gu

		var lct6680Result1 patype.MoonSetLCTL6680 = MoonSetLctL6680(x, ds, zc, gdy, gmn, gyr, g1, ut)
		lct = lct6680Result1.Lct
		dy1 = lct6680Result1.Dy1
		mn1 = lct6680Result1.Mn1
		yr1 = lct6680Result1.Yr1
		gdy = lct6680Result1.Gdy
		gmn = lct6680Result1.Gmn
		gyr = lct6680Result1.Gyr

		var lct6700Result2 patype.MoonSetLCTL6700 = MoonSetLctL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
		lu = lct6700Result2.Lu
		lct = lct6700Result2.Lct

		if lct == -99.0 {
			return lct
		}

		la = lu
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
	ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	lct = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)

	return lct
}

/* Helper function for MoonSetLCT */
func MoonSetLctL6680(x float64, ds int, zc int, gdy float64, gmn int, gyr int, g1 float64, ut float64) patype.MoonSetLCTL6680 {
	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	var lct float64 = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var dy1 float64 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var mn1 int = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var yr1 int = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	gdy = LocalCivilTimeGreenwichDay(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	gmn = int(LocalCivilTimeGreenwichMonth(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	gyr = int(LocalCivilTimeGreenwichYear(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	ut -= 24.0 * math.Floor(ut/24.0)

	return patype.MoonSetLCTL6680{Ut: ut, Lct: lct, Dy1: dy1, Mn1: mn1, Yr1: yr1, Gdy: gdy, Gmn: gmn, Gyr: gyr}
}

/* Helper function for MoonSetLCT */
func MoonSetLctL6700(lct float64, ds int, zc int, dy1 float64, mn1 int, yr1 int, gdy float64, gmn int, gyr int, gLat float64) patype.MoonSetLCTL6700 {
	var mm float64 = MoonLongitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var bm float64 = MoonLatitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var pm float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	var dp float64 = NutatLong(gdy, gmn, gyr)
	var th float64 = 0.27249 * math.Sin(pm)
	var di float64 = th + 0.0098902 - pm
	var p float64 = DecimalDegreesToDegreeHours(EclipticRightAscension(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr))
	var q float64 = EclipticDeclination(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr)
	var lu float64 = RiseSetLocalSiderealTimeSet(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)

	if ERiseSet(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat) != patype.RiseSetStatus_OK {
		lct = -99.0
	}

	return patype.MoonSetLCTL6700{Mm: mm, Bm: bm, Pm: pm, Dp: dp, Th: th, Di: di, P: p, Q: q, Lu: lu, Lct: lct}
}

/*
Local date of moonset.

Original macro names: MoonSetLcDay, MoonSetLcMonth, MoonSetLcYear
*/
func MoonSetLcDmy(dy float64, mn int, yr int, ds int, zc int, gLong float64, gLat float64) patype.FullDatePrecise {
	var gdy float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var gmn int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var gyr int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var lct float64 = 12.0
	var dy1 float64 = dy
	var mn1 int = mn
	var yr1 int = yr

	var dmy6700_result1 patype.MoonSetLcDMYL6700 = MoonSetLcDmyL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
	var lu float64 = dmy6700_result1.Lu
	lct = dmy6700_result1.Lct

	if lct == -99.0 {
		return patype.FullDatePrecise{Month: int(lct), Day: lct, Year: int(lct)}
	}

	var la float64 = lu

	var x float64
	var ut float64
	var g1 float64 = 0.0
	var gu float64 = 0.0
	for k := 1; k < 9; k++ {
		x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
		ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

		if k == 1 {
			g1 = ut
		} else {
			g1 = gu
		}

		gu = ut
		ut = gu

		var dmy6680_result1 patype.MoonSetLcDMYL6680 = MoonSetLcDmyL6680(x, ds, zc, gdy, gmn, gyr, g1, ut)
		lct = dmy6680_result1.Lct
		dy1 = dmy6680_result1.Dy1
		mn1 = dmy6680_result1.Mn1
		yr1 = dmy6680_result1.Yr1
		gdy = dmy6680_result1.Gdy
		gmn = dmy6680_result1.Gmn
		gyr = dmy6680_result1.Gyr

		var dmy6700_result2 patype.MoonSetLcDMYL6700 = MoonSetLcDmyL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
		lu = dmy6700_result2.Lu
		lct = dmy6700_result2.Lct

		if lct == -99.0 {
			return patype.FullDatePrecise{Month: int(lct), Day: lct, Year: int(lct)}
		}

		la = lu
	}

	x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
	ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	dy1 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	mn1 = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	yr1 = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)

	return patype.FullDatePrecise{Month: mn1, Day: dy1, Year: yr1}
}

/* Helper function for MoonSetLcDMY */
func MoonSetLcDmyL6680(x float64, ds int, zc int, gdy float64, gmn int, gyr int, g1 float64, ut float64) patype.MoonSetLcDMYL6680 {
	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	var lct float64 = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var dy1 float64 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var mn1 int = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var yr1 int = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	gdy = LocalCivilTimeGreenwichDay(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	gmn = int(LocalCivilTimeGreenwichMonth(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	gyr = int(LocalCivilTimeGreenwichYear(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	ut -= 24.0 * math.Floor(ut/24.0)

	return patype.MoonRiseLcDMYL6680{Ut: ut, Lct: lct, Dy1: dy1, Mn1: mn1, Yr1: yr1, Gdy: gdy, Gmn: gmn, Gyr: gyr}
}

/* Helper function for MoonSetLcDMY */
func MoonSetLcDmyL6700(lct float64, ds int, zc int, dy1 float64, mn1 int, yr1 int, gdy float64, gmn int, gyr int, gLat float64) patype.MoonSetLcDMYL6700 {
	var mm float64 = MoonLongitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var bm float64 = MoonLatitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var pm float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	var dp float64 = NutatLong(gdy, gmn, gyr)
	var th float64 = 0.27249 * math.Sin(pm)
	var di float64 = th + 0.0098902 - pm
	var p float64 = DecimalDegreesToDegreeHours(EclipticRightAscension(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr))
	var q float64 = EclipticDeclination(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr)
	var lu float64 = RiseSetLocalSiderealTimeSet(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)

	return patype.MoonSetLcDMYL6700{Mm: mm, Bm: bm, Pm: pm, Dp: dp, Th: th, Di: di, P: p, Q: q, Lu: lu, Lct: lct}
}

/*
Local azimuth of moonset.

Original macro name: MoonSetAz
*/
func MoonSetAz(dy float64, mn int, yr int, ds int, zc int, gLong float64, gLat float64) float64 {
	var gdy float64 = LocalCivilTimeGreenwichDay(12.0, 0.0, 0.0, ds, zc, dy, mn, yr)
	var gmn int = int(LocalCivilTimeGreenwichMonth(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var gyr int = int(LocalCivilTimeGreenwichYear(12.0, 0.0, 0.0, ds, zc, dy, mn, yr))
	var lct float64 = 12.0
	var dy1 float64 = dy
	var mn1 int = mn
	var yr1 int = yr

	var az6700Result1 patype.MoonSetAzL6700 = MoonSetAzL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
	var lu float64 = az6700Result1.Lu
	lct = az6700Result1.Lct

	var au float64

	if lct == -99.0 {
		return lct
	}

	var la float64 = lu

	var x float64
	var ut float64
	var g1 float64
	var gu float64 = 0.0
	var aa float64 = 0.0
	for k := 1; k < 9; k++ {
		x = LocalSiderealTimeToGreenwichSiderealTime(la, 0.0, 0.0, gLong)
		ut = GreenwichSiderealTimeToUniversalTime(x, 0.0, 0.0, gdy, gmn, gyr)

		if k == 1 {
			g1 = ut
		} else {
			g1 = gu
		}

		gu = ut
		ut = gu

		var az6680Result1 patype.MoonSetAzL6680 = MoonSetAzL6680(x, ds, zc, gdy, gmn, gyr, g1, ut)
		lct = az6680Result1.Lct
		dy1 = az6680Result1.Dy1
		mn1 = az6680Result1.Mn1
		yr1 = az6680Result1.Yr1
		gdy = az6680Result1.Gdy
		gmn = az6680Result1.Gmn
		gyr = az6680Result1.Gyr

		var az6700Result2 patype.MoonSetAzL6700 = MoonSetAzL6700(lct, ds, zc, dy1, mn1, yr1, gdy, gmn, gyr, gLat)
		lu = az6700Result2.Lu
		lct = az6700Result2.Lct
		au = az6700Result2.Au

		if lct == -99.0 {
			return lct
		}

		la = lu
		aa = au
	}

	au = aa

	return au
}

/* Helper function for MoonSetAz */
func MoonSetAzL6680(x float64, ds int, zc int, gdy float64, gmn int, gyr int, g1 float64, ut float64) patype.MoonSetAzL6680 {
	if EGstUt(x, 0.0, 0.0, gdy, gmn, gyr) != patype.WarningFlag_OK {
		if math.Abs(g1-ut) > 0.5 {
			ut += 23.93447
		}
	}

	ut = UtDayAdjust(ut, g1)
	var lct float64 = UniversalTimeToLocalCivilTime(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var dy1 float64 = UniversalTimeLocalCivilDay(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var mn1 int = UniversalTimeLocalCivilMonth(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	var yr1 int = UniversalTimeLocalCivilYear(ut, 0.0, 0.0, ds, zc, gdy, gmn, gyr)
	gdy = LocalCivilTimeGreenwichDay(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	gmn = int(LocalCivilTimeGreenwichMonth(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	gyr = int(LocalCivilTimeGreenwichYear(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	ut -= 24.0 * math.Floor(ut/24.0)

	return patype.MoonSetAzL6680{Ut: ut, Lct: lct, Dy1: dy1, Mn1: mn1, Yr1: yr1, Gdy: gdy, Gmn: gmn, Gyr: gyr}
}

/* Helper function for MoonSetAz */
func MoonSetAzL6700(lct float64, ds int, zc int, dy1 float64, mn1 int, yr1 int, gdy float64, gmn int, gyr int, gLat float64) patype.MoonSetAzL6700 {
	var mm float64 = MoonLongitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var bm float64 = MoonLatitude(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1)
	var pm float64 = pautil.DegreesToRadians(MoonHorizontalParallax(lct, 0.0, 0.0, ds, zc, dy1, mn1, yr1))
	var dp float64 = NutatLong(gdy, gmn, gyr)
	var th float64 = 0.27249 * math.Sin(pm)
	var di float64 = th + 0.0098902 - pm
	var p float64 = DecimalDegreesToDegreeHours(EclipticRightAscension(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr))
	var q float64 = EclipticDeclination(mm+dp, 0.0, 0.0, bm, 0.0, 0.0, gdy, gmn, gyr)
	var lu float64 = RiseSetLocalSiderealTimeSet(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)
	var au float64 = RiseSetAzimuthSet(p, 0.0, 0.0, q, 0.0, 0.0, Degrees(di), gLat)

	return patype.MoonSetAzL6700{Mm: mm, Bm: bm, Pm: pm, Dp: dp, Th: th, Di: di, P: p, Q: q, Lu: lu, Lct: lct, Au: au}
}
