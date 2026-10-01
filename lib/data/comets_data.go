package data

type CometDataElliptical struct {
	Name                          string  /* Name of comet */
	Epoch_EpochOfPerihelion       float64 /* Epoch of the perihelion */
	Peri_LongitudeOfPerihelion    float64 /* Longitude of the perihelion */
	Node_LongitudeOfAscendingNode float64 /* Longitude of the ascending node */
	Period_PeriodOfOrbit          float64 /* Period of the orbit */
	Axis_SemiMajorAxisOfOrbit     float64 /* Semi-major axis of the orbit */
	Ecc_EccentricityOfOrbit       float64 /* Eccentricity of the orbit */
	Incl_InclinationOfOrbit       float64 /* Inclination of the orbit */
}

func GetCometDataElliptical(cometName string) CometDataElliptical {
	var cometRecords = []CometDataElliptical{
		{
			Name: "Encke", Epoch_EpochOfPerihelion: 1974.32, Peri_LongitudeOfPerihelion: 160.1, Node_LongitudeOfAscendingNode: 334.2,
			Period_PeriodOfOrbit: 3.3, Axis_SemiMajorAxisOfOrbit: 2.21, Ecc_EccentricityOfOrbit: 0.85, Incl_InclinationOfOrbit: 12.0,
		},
		{
			Name: "Temple 2", Epoch_EpochOfPerihelion: 1972.87, Peri_LongitudeOfPerihelion: 310.2, Node_LongitudeOfAscendingNode: 119.3,
			Period_PeriodOfOrbit: 5.26, Axis_SemiMajorAxisOfOrbit: 3.02, Ecc_EccentricityOfOrbit: 0.55, Incl_InclinationOfOrbit: 12.5,
		},
		{
			Name: "Haneda-Campos", Epoch_EpochOfPerihelion: 1978.77, Peri_LongitudeOfPerihelion: 12.02, Node_LongitudeOfAscendingNode: 131.7,
			Period_PeriodOfOrbit: 5.37, Axis_SemiMajorAxisOfOrbit: 3.07, Ecc_EccentricityOfOrbit: 0.64, Incl_InclinationOfOrbit: 5.81,
		},
		{
			Name: "Schwassmann-Wachmann 2", Epoch_EpochOfPerihelion: 1974.7, Peri_LongitudeOfPerihelion: 123.3, Node_LongitudeOfAscendingNode: 126.0,
			Period_PeriodOfOrbit: 6.51, Axis_SemiMajorAxisOfOrbit: 3.49, Ecc_EccentricityOfOrbit: 0.39, Incl_InclinationOfOrbit: 3.7,
		},
		{
			Name: "Borrelly", Epoch_EpochOfPerihelion: 1974.36, Peri_LongitudeOfPerihelion: 67.8, Node_LongitudeOfAscendingNode: 75.1,
			Period_PeriodOfOrbit: 6.76, Axis_SemiMajorAxisOfOrbit: 3.58, Ecc_EccentricityOfOrbit: 0.63, Incl_InclinationOfOrbit: 30.2,
		},
		{
			Name: "Whipple", Epoch_EpochOfPerihelion: 1970.77, Peri_LongitudeOfPerihelion: 18.2, Node_LongitudeOfAscendingNode: 188.4,
			Period_PeriodOfOrbit: 7.47, Axis_SemiMajorAxisOfOrbit: 3.82, Ecc_EccentricityOfOrbit: 0.35, Incl_InclinationOfOrbit: 10.2,
		},
		{
			Name: "Oterma", Epoch_EpochOfPerihelion: 1958.44, Peri_LongitudeOfPerihelion: 150.0, Node_LongitudeOfAscendingNode: 155.1,
			Period_PeriodOfOrbit: 7.88, Axis_SemiMajorAxisOfOrbit: 3.96, Ecc_EccentricityOfOrbit: 0.14, Incl_InclinationOfOrbit: 4.0,
		},
		{
			Name: "Schaumasse", Epoch_EpochOfPerihelion: 1960.29, Peri_LongitudeOfPerihelion: 138.1, Node_LongitudeOfAscendingNode: 86.2,
			Period_PeriodOfOrbit: 8.18, Axis_SemiMajorAxisOfOrbit: 4.05, Ecc_EccentricityOfOrbit: 0.71, Incl_InclinationOfOrbit: 12.0,
		},
		{
			Name: "Comas Sola", Epoch_EpochOfPerihelion: 1969.83, Peri_LongitudeOfPerihelion: 102.9, Node_LongitudeOfAscendingNode: 62.8,
			Period_PeriodOfOrbit: 8.55, Axis_SemiMajorAxisOfOrbit: 4.18, Ecc_EccentricityOfOrbit: 0.58, Incl_InclinationOfOrbit: 13.4,
		},
		{
			Name: "Schwassmann-Wachmann 1", Epoch_EpochOfPerihelion: 1974.12, Peri_LongitudeOfPerihelion: 334.1, Node_LongitudeOfAscendingNode: 319.6,
			Period_PeriodOfOrbit: 15.03, Axis_SemiMajorAxisOfOrbit: 6.09, Ecc_EccentricityOfOrbit: 0.11, Incl_InclinationOfOrbit: 9.7,
		},
		{
			Name: "Neujmin 1", Epoch_EpochOfPerihelion: 1966.94, Peri_LongitudeOfPerihelion: 334.0, Node_LongitudeOfAscendingNode: 347.2,
			Period_PeriodOfOrbit: 17.93, Axis_SemiMajorAxisOfOrbit: 6.86, Ecc_EccentricityOfOrbit: 0.78, Incl_InclinationOfOrbit: 15.0,
		},
		{
			Name: "Crommelin", Epoch_EpochOfPerihelion: 1956.82, Peri_LongitudeOfPerihelion: 86.4, Node_LongitudeOfAscendingNode: 250.4,
			Period_PeriodOfOrbit: 27.89, Axis_SemiMajorAxisOfOrbit: 9.17, Ecc_EccentricityOfOrbit: 0.92, Incl_InclinationOfOrbit: 28.9,
		},
		{
			Name: "Olbers", Epoch_EpochOfPerihelion: 1956.46, Peri_LongitudeOfPerihelion: 150.0, Node_LongitudeOfAscendingNode: 85.4,
			Period_PeriodOfOrbit: 69.47, Axis_SemiMajorAxisOfOrbit: 16.84, Ecc_EccentricityOfOrbit: 0.93, Incl_InclinationOfOrbit: 44.6,
		},
		{
			Name: "Pons-Brooks", Epoch_EpochOfPerihelion: 1954.39, Peri_LongitudeOfPerihelion: 94.2, Node_LongitudeOfAscendingNode: 255.2,
			Period_PeriodOfOrbit: 70.98, Axis_SemiMajorAxisOfOrbit: 17.2, Ecc_EccentricityOfOrbit: 0.96, Incl_InclinationOfOrbit: 74.2,
		},
		{
			Name: "Halley", Epoch_EpochOfPerihelion: 1986.112, Peri_LongitudeOfPerihelion: 170.011, Node_LongitudeOfAscendingNode: 58.154,
			Period_PeriodOfOrbit: 76.0081, Axis_SemiMajorAxisOfOrbit: 17.9435, Ecc_EccentricityOfOrbit: 0.9673, Incl_InclinationOfOrbit: 162.2384,
		},
	}

	for _, cometRecord := range cometRecords {
		if cometRecord.Name == cometName {
			return cometRecord
		}
	}

	return CometDataElliptical{
		Name: "NOTFOUND", Epoch_EpochOfPerihelion: -99, Peri_LongitudeOfPerihelion: -99, Node_LongitudeOfAscendingNode: -99,
		Period_PeriodOfOrbit: -99, Axis_SemiMajorAxisOfOrbit: -99, Ecc_EccentricityOfOrbit: -99, Incl_InclinationOfOrbit: -99,
	}
}
