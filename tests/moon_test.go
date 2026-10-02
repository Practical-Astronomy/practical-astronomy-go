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
