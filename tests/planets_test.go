package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApproximatePositionOfPlanet(t *testing.T) {
	var lctHour float64 = 0
	var lctMin float64 = 0
	var lctSec float64 = 0
	var isDaylightSaving bool = false
	var zoneCorrectionHours int = 0
	var localDateDay float64 = 22
	var localDateMonth int = 11
	var localDateYear int = 2003
	var planetName string = "Jupiter"

	var expectedApproximatePositionofPlanet patype.PlanetPosition = patype.PlanetPosition{
		RightAscensionHour: 11, RightAscensionMinutes: 11, RightAscensionSeconds: 13.8,
		DeclinationDegrees: 6, DeclinationMinutes: 21, DeclinationSeconds: 25.1,
	}
	var actualApproximatePositionofPlanet patype.PlanetPosition = palib.ApproximatePositionOfPlanet(
		lctHour, lctMin, lctSec, isDaylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear, planetName)

	require.Equal(t, expectedApproximatePositionofPlanet, actualApproximatePositionofPlanet, "Mismatch in Approximate Position of Planet")
}
