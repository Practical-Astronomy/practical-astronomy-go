package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryStar(t *testing.T) {
	var expectedBinaryStar patype.BinaryStarOrbitalData = patype.BinaryStarOrbitalData{PositionAngleDeg: 318.5, SeparationArcsec: 0.41}

	var actualBinaryStar patype.BinaryStarOrbitalData = palib.BinaryStarOrbit(1, 1, 1980, "eta-Cor")

	require.Equal(t, expectedBinaryStar, actualBinaryStar, "Mismatched Binary Star Orbital Data")
}
