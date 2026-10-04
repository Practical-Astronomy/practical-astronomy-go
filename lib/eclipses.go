package lib

import (
	"math"
	pamacro "practicalastro/lib/macros"
	patype "practicalastro/lib/types"
	pautil "practicalastro/lib/util"
)

/* Determine if a lunar eclipse is likely to occur. */
func LunarEclipseOccurrence(
	localDateDay float64, localDateMonth int, localDateYear int, isDaylightSaving bool, zoneCorrectionHours int,
) patype.LunarEclipseOccurrence {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var julianDateOfFullMoon float64 = pamacro.FullMoon(daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	var gDateOfFullMoonDay float64 = pamacro.JulianDateDay(julianDateOfFullMoon)
	var integerDay float64 = math.Floor(gDateOfFullMoonDay)
	var gDateOfFullMoonMonth int = pamacro.JulianDateMonth(julianDateOfFullMoon)
	var gDateOfFullMoonYear int = pamacro.JulianDateYear(julianDateOfFullMoon)
	var utOfFullMoonHours float64 = gDateOfFullMoonDay - integerDay

	var localCivilDateDay float64 = pamacro.UniversalTimeLocalCivilDay(
		utOfFullMoonHours, 0.0, 0.0, daylightSaving, zoneCorrectionHours, integerDay, gDateOfFullMoonMonth, gDateOfFullMoonYear)
	var localCivilDateMonth int = pamacro.UniversalTimeLocalCivilMonth(
		utOfFullMoonHours, 0.0, 0.0, daylightSaving, zoneCorrectionHours, integerDay, gDateOfFullMoonMonth, gDateOfFullMoonYear)
	var localCivilDateYear int = pamacro.UniversalTimeLocalCivilYear(
		utOfFullMoonHours, 0.0, 0.0, daylightSaving, zoneCorrectionHours, integerDay, gDateOfFullMoonMonth, gDateOfFullMoonYear)

	var eclipseOccurrence patype.LunarEclipseStatus = pamacro.LunarEclipseOccurrence(
		daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	var status patype.LunarEclipseStatus = eclipseOccurrence
	var eventDateDay float64 = localCivilDateDay
	var eventDateMonth int = localCivilDateMonth
	var eventDateYear int = localCivilDateYear

	return patype.LunarEclipseOccurrence{Status: status, EventDateDay: eventDateDay, EventDateMonth: eventDateMonth, EventDateYear: eventDateYear}
}

/* Calculate the circumstances of a lunar eclipse. */
func LunarEclipseCircumstances(
	localDateDay float64, localDateMonth int, localDateYear int, isDaylightSaving bool, zoneCorrectionHours int,
) patype.LunarEclipseCircumstances {
	var daylightSaving int = pautil.BoolToInt(isDaylightSaving)

	var julianDateOfFullMoon float64 = pamacro.FullMoon(daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	var gDateOfFullMoonDay float64 = pamacro.JulianDateDay(julianDateOfFullMoon)
	var integerDay float64 = math.Floor(gDateOfFullMoonDay)
	var gDateOfFullMoonMonth int = pamacro.JulianDateMonth(julianDateOfFullMoon)
	var gDateOfFullMoonYear int = pamacro.JulianDateYear(julianDateOfFullMoon)
	var utOfFullMoonHours float64 = gDateOfFullMoonDay - integerDay

	var localCivilDateDay float64 = pamacro.UniversalTimeLocalCivilDay(
		utOfFullMoonHours, 0.0, 0.0, daylightSaving, zoneCorrectionHours, integerDay, gDateOfFullMoonMonth, gDateOfFullMoonYear)
	var localCivilDateMonth int = pamacro.UniversalTimeLocalCivilMonth(
		utOfFullMoonHours, 0.0, 0.0, daylightSaving, zoneCorrectionHours, integerDay, gDateOfFullMoonMonth, gDateOfFullMoonYear)
	var localCivilDateYear int = pamacro.UniversalTimeLocalCivilYear(
		utOfFullMoonHours, 0.0, 0.0, daylightSaving, zoneCorrectionHours, integerDay, gDateOfFullMoonMonth, gDateOfFullMoonYear)

	var utMaxEclipse float64 = pamacro.UtMaxLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)
	var utFirstContact float64 = pamacro.UtFirstContactLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)
	var utLastContact float64 = pamacro.UtLastContactLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)
	var utStartUmbralPhase float64 = pamacro.UtStartUmbraLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)
	var utEndUmbralPhase float64 = pamacro.UtEndUmbraLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)
	var utStartTotalPhase float64 = pamacro.UtStartTotalLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)
	var utEndTotalPhase float64 = pamacro.UtEndTotalLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)

	var eclipseMagnitude1 float64 = pamacro.MagLunarEclipse(
		localDateDay, localDateMonth, localDateYear, daylightSaving, zoneCorrectionHours)

	var lunarEclipseCertainDateDay float64 = localCivilDateDay
	var lunarEclipseCertainDateMonth int = localCivilDateMonth
	var lunarEclipseCertainDateYear int = localCivilDateYear

	var utStartPenPhaseHour float64 = pautil.TernaryAssignFloat64(
		utFirstContact == -99.0, -99, float64(pamacro.DecimalHoursHour(utFirstContact+0.008333)))
	var utStartPenPhaseMinutes float64 = pautil.TernaryAssignFloat64(
		utFirstContact == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utFirstContact+0.008333)))

	var utStartUmbralPhaseHour float64 = pautil.TernaryAssignFloat64(
		utStartUmbralPhase == -99.0, -99.0, float64(pamacro.DecimalHoursHour(utStartUmbralPhase+0.008333)))
	var utStartUmbralPhaseMinutes float64 = pautil.TernaryAssignFloat64(
		utStartUmbralPhase == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utStartUmbralPhase+0.008333)))

	var utStartTotalPhaseHour float64 = pautil.TernaryAssignFloat64(
		utStartTotalPhase == -99.0, -99.0, float64(pamacro.DecimalHoursHour(utStartTotalPhase+0.008333)))
	var utStartTotalPhaseMinutes float64 = pautil.TernaryAssignFloat64(
		utStartTotalPhase == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utStartTotalPhase+0.008333)))

	var utMidEclipseHour float64 = pautil.TernaryAssignFloat64(
		utMaxEclipse == -99.0, -99.0, float64(pamacro.DecimalHoursHour(utMaxEclipse+0.008333)))
	var utMidEclipseMinutes float64 = pautil.TernaryAssignFloat64(
		utMaxEclipse == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utMaxEclipse+0.008333)))

	var utEndTotalPhaseHour float64 = pautil.TernaryAssignFloat64(
		utEndTotalPhase == -99.0, -99.0, float64(pamacro.DecimalHoursHour(utEndTotalPhase+0.008333)))
	var utEndTotalPhaseMinutes float64 = pautil.TernaryAssignFloat64(
		utEndTotalPhase == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utEndTotalPhase+0.008333)))

	var utEndUmbralPhaseHour float64 = pautil.TernaryAssignFloat64(
		utEndUmbralPhase == -99.0, -99.0, float64(pamacro.DecimalHoursHour(utEndUmbralPhase+0.008333)))
	var utEndUmbralPhaseMinutes float64 = pautil.TernaryAssignFloat64(
		utEndUmbralPhase == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utEndUmbralPhase+0.008333)))

	var utEndPenPhaseHour float64 = pautil.TernaryAssignFloat64(
		utLastContact == -99.0, -99.0, float64(pamacro.DecimalHoursHour(utLastContact+0.008333)))
	var utEndPenPhaseMinutes float64 = pautil.TernaryAssignFloat64(
		utLastContact == -99.0, -99.0, float64(pamacro.DecimalHoursMinute(utLastContact+0.008333)))

	var eclipseMagnitude float64 = pautil.TernaryAssignFloat64(eclipseMagnitude1 == -99.0, -99.0, pautil.RoundTo(eclipseMagnitude1, 2))

	return patype.LunarEclipseCircumstances{
		CertainDateDay:      lunarEclipseCertainDateDay,
		CertainDateMonth:    float64(lunarEclipseCertainDateMonth),
		CertainDateYear:     float64(lunarEclipseCertainDateYear),
		UtStartPenPhaseHour: utStartPenPhaseHour, UtStartPenPhaseMinutes: utStartPenPhaseMinutes,
		UtStartUmbralPhaseHour: utStartUmbralPhaseHour, UtStartUmbralPhaseMinutes: utStartUmbralPhaseMinutes,
		UtStartTotalPhaseHour: utStartTotalPhaseHour, UtStartTotalPhaseMinutes: utStartTotalPhaseMinutes,
		UtMidEclipseHour: utMidEclipseHour, UtMidEclipseMinutes: utMidEclipseMinutes,
		UtEndTotalPhaseHour: utEndTotalPhaseHour, UtEndTotalPhaseMinutes: utEndTotalPhaseMinutes,
		UtEndUmbralPhaseHour: utEndUmbralPhaseHour, UtEndUmbralPhaseMinutes: utEndUmbralPhaseMinutes,
		UtEndPenPhaseHour: utEndPenPhaseHour, UtEndPenPhaseMinutes: utEndPenPhaseMinutes,
		EclipseMagnitude: eclipseMagnitude,
	}
}
