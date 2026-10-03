package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApproximatePositionOfMoon(t *testing.T) {
	var expectedMoonPosition patype.MoonApproximatePosition = patype.MoonApproximatePosition{
		RaHour: 14, RaMin: 12, RaSec: 42.31, DecDeg: -11, DecMin: 31, DecSec: 38.27}
	var actualMoonPosition patype.MoonApproximatePosition = palib.ApproximatePositionOfMoon(0, 0, 0, false, 0, 1, 9, 2003)

	require.Equal(t, expectedMoonPosition, actualMoonPosition, "Mismatched Approximate Position of Moon")
}

func TestPrecisePositionOfMoon(t *testing.T) {
	var expectedMoonPosition patype.MoonPrecisePosition = patype.MoonPrecisePosition{
		RaHour: 14, RaMin: 12, RaSec: 10.21, DecDeg: -11, DecMin: 34, DecSec: 57.83, EarthMoonDistKm: 367964, HorParallaxDeg: 0.993191}
	var actualMoonPosition patype.MoonPrecisePosition = palib.PrecisePositionOfMoon(0, 0, 0, false, 0, 1, 9, 2003)

	require.Equal(t, expectedMoonPosition, actualMoonPosition, "Mismatched Precise Position of Moon")
}

func TestMoonPhase(t *testing.T) {
	var expectedMoonPhase patype.MoonPhase = patype.MoonPhase{Phase: 0.22, BrightLimbDeg: -71.58}
	var actualMoonPhase patype.MoonPhase = palib.MoonPhase(0, 0, 0, false, 0, 1, 9, 2003, patype.AccuracyLevel_Approximate)

	require.Equal(t, expectedMoonPhase, actualMoonPhase, "Mismatched Moon Phase")
}

func TestNewMoonFullMoon(t *testing.T) {
	var expectedNewMoonFullMoon patype.MoonNewFull = patype.MoonNewFull{
		NewLocalTimeHour: 17, NewLocalTimeMin: 27, NewLocalDateDay: 27, NewLocalDateMonth: 8, NewLocalDateYear: 2003,
		FullLocalTimeHour: 16, FullLocalTimeMin: 36, FullLocalDateDay: 10, FullLocalDateMonth: 9, FullLocalDateYear: 2003,
	}
	var actualNewMoonFullMoon patype.MoonNewFull = palib.TimesOfNewMoonAndFullMoon(false, 0, 1, 9, 2003)

	require.Equal(t, expectedNewMoonFullMoon, actualNewMoonFullMoon, "Mismatched New Moon/Full Moon")
}

func TestMoonDistAngDiamHorParallax(t *testing.T) {
	var expectedMoonDist patype.MoonDistDiameterHorParallax = patype.MoonDistDiameterHorParallax{
		EarthMoonDist: 367964, AngDiameterDeg: 0, AngDiameterMin: 32, HorParallaxDeg: 0, HorParallaxMin: 59, HorParallaxSec: 35.49}
	var actualMoonDist patype.MoonDistDiameterHorParallax = palib.MoonDistAngDiamHorParallax(0, 0, 0, false, 0, 1, 9, 2003)

	require.Equal(t, expectedMoonDist, actualMoonDist, "Mismatched Moon Distance/Angular Diameter/Horizontal Parallax")
}

func TestMoonRiseSet(t *testing.T) {
	var expectedMoonRiseSet patype.MoonRiseSet = patype.MoonRiseSet{
		RiseLocalTimeHour: 4, RiseLocalTimeMin: 21,
		RiseLocalDateDay: 6, RiseLocalDateMonth: 3, RiseLocalDateYear: 1986,
		RiseAzimuthDeg:   127.34,
		SetLocalTimeHour: 13, SetLocalTimeMin: 8,
		SetLocalDateDay: 6, SetLocalDateMonth: 3, SetLocalDateYear: 1986,
		SetAzimuthDeg: 234.05,
	}
	var actualMoonRiseSet patype.MoonRiseSet = palib.MoonriseAndMoonset(6, 3, 1986, false, -5, -71.05, 42.3667)

	require.Equal(t, expectedMoonRiseSet, actualMoonRiseSet, "Mismatched Moon Rise/Set")
}
