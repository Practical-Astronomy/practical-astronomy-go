package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPositionOfEllipticalComet(t *testing.T) {
	var expectedPositionOfEllipticalComet patype.EllipticalCometPosition = patype.EllipticalCometPosition{RaHour: 6, RaMin: 29, DecDeg: 10, DecMin: 13, DistEarth: 8.13}

	var actualPositionOfEllipticalComet patype.EllipticalCometPosition = palib.PositionOfEllipticalComet(0, 0, 0, false, 0, 1, 1, 1984, "Halley")

	require.Equal(t, expectedPositionOfEllipticalComet, actualPositionOfEllipticalComet, "Mismatched Position of Elliptical Comet")
}
