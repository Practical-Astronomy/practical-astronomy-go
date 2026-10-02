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
