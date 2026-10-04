package tests

import (
	palib "practicalastro/lib"
	patype "practicalastro/lib/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLunarEclipseOccurrence(t *testing.T) {
	var expectedLunarEclipseOccurrence patype.LunarEclipseOccurrence = patype.LunarEclipseOccurrence{
		Status: patype.LunarEclipseStatus_Certain, EventDateDay: 4, EventDateMonth: 4, EventDateYear: 2015}
	var actualLunarEclipseOccurrence patype.LunarEclipseOccurrence = palib.LunarEclipseOccurrence(1, 4, 2015, false, 10)

	require.Equal(t, expectedLunarEclipseOccurrence, actualLunarEclipseOccurrence, "Mismatched Lunar Eclipse Occurrence")
}

func TestLunarEclipseCircumstances(t *testing.T) {
	var expectedLunarEclipseCircumstances patype.LunarEclipseCircumstances = patype.LunarEclipseCircumstances{
		CertainDateDay: 4, CertainDateMonth: 4, CertainDateYear: 2015,
		UtStartPenPhaseHour: 9, UtStartPenPhaseMinutes: 0,
		UtStartUmbralPhaseHour: 10, UtStartUmbralPhaseMinutes: 16,
		UtStartTotalPhaseHour: 11, UtStartTotalPhaseMinutes: 55,
		UtMidEclipseHour: 12, UtMidEclipseMinutes: 1,
		UtEndTotalPhaseHour: 12, UtEndTotalPhaseMinutes: 7,
		UtEndUmbralPhaseHour: 13, UtEndUmbralPhaseMinutes: 46,
		UtEndPenPhaseHour: 15, UtEndPenPhaseMinutes: 1,
		EclipseMagnitude: 1.01,
	}
	var actualLunarEclipseCircumstances patype.LunarEclipseCircumstances = palib.LunarEclipseCircumstances(1, 4, 2015, false, 10)

	require.Equal(t, expectedLunarEclipseCircumstances, actualLunarEclipseCircumstances, "Mismatched Lunar Eclipse Circumstances")
}

func TestSolarEclipseOccurrence(t *testing.T) {
	var expectedSolarEclipseOccurrence patype.SolarEclipseOccurrence = patype.SolarEclipseOccurrence{
		Status: patype.SolarEclipseStatus_Certain, EventDateDay: 20, EventDateMonth: 3, EventDateYear: 2015}
	var actualSolarEclipseOccurrence patype.SolarEclipseOccurrence = palib.SolarEclipseOccurrence(1, 4, 2015, false, 0)

	require.Equal(t, expectedSolarEclipseOccurrence, actualSolarEclipseOccurrence, "Mismatched Solar Eclipse Occurrence")
}

func TestSolarEclipseCircumstances(t *testing.T) {
	var expectedSolarEclipseCircumstances patype.SolarEclipseCircumstances = patype.SolarEclipseCircumstances{
		CertainDateDay: 20, CertainDateMonth: 3, CertainDateYear: 2015,
		UtFirstContactHour: 8, UtFirstContactMinutes: 55,
		UtMidEclipseHour: 9, UtMidEclipseMinutes: 57,
		UtLastContactHour: 10, UtLastContactMinutes: 58,
		EclipseMagnitude: 1.016,
	}
	var actualSolarEclipseCircumstances patype.SolarEclipseCircumstances = palib.SolarEclipseCircumstances(20, 3, 2015, false, 0, 0, 68.65)

	require.Equal(t, expectedSolarEclipseCircumstances, actualSolarEclipseCircumstances, "Mismatched Solar Eclipse Circumstances")
}
