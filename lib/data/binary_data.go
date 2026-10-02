package data

type BinaryStarData struct {
	Name      string  /* Name of binary system.  */
	Period    float64 /* Period of the orbit. */
	EpochPeri float64 /* Epoch of the perihelion. */
	LongPeri  float64 /* Longitude of the perihelion. */
	Ecc       float64 /* Eccentricity of the orbit. */
	Axis      float64 /* Semi-major axis of the orbit. */
	Incl      float64 /* Orbital inclination. */
	PaNode    float64 /* Position angle of the ascending node. */
}

func GetBinaryStarData(binaryStarName string) BinaryStarData {
	var binaryRecords = []BinaryStarData{
		{Name: "eta-Cor", Period: 41.623, EpochPeri: 1934.008, LongPeri: 219.907, Ecc: 0.2763, Axis: 0.907, Incl: 59.025, PaNode: 23.717},
		{Name: "gamma-Vir", Period: 171.37, EpochPeri: 1836.433, LongPeri: 252.88, Ecc: 0.8808, Axis: 3.746, Incl: 146.05, PaNode: 31.78},
		{Name: "eta-Cas", Period: 480.0, EpochPeri: 1889.6, LongPeri: 268.59, Ecc: 0.497, Axis: 11.9939, Incl: 34.76, PaNode: 278.42},
		{Name: "zeta-Ori", Period: 1508.6, EpochPeri: 2070.6, LongPeri: 47.3, Ecc: 0.07, Axis: 2.728, Incl: 72.0, PaNode: 155.5},
		{Name: "alpha-CMa", Period: 50.09, EpochPeri: 1894.13, LongPeri: 147.27, Ecc: 0.5923, Axis: 7.5, Incl: 136.53, PaNode: 44.57},
		{Name: "delta-Gem", Period: 1200.0, EpochPeri: 1437.0, LongPeri: 57.19, Ecc: 0.11, Axis: 6.9753, Incl: 63.28, PaNode: 18.38},
		{Name: "alpha-Gem", Period: 420.07, EpochPeri: 1965.3, LongPeri: 261.43, Ecc: 0.33, Axis: 6.295, Incl: 115.94, PaNode: 40.47},
		{Name: "aplah-CMi", Period: 40.65, EpochPeri: 1927.6, LongPeri: 269.8, Ecc: 0.4, Axis: 4.548, Incl: 35.7, PaNode: 284.3},
		{Name: "alpha-Cen", Period: 79.92, EpochPeri: 1955.56, LongPeri: 231.56, Ecc: 0.516, Axis: 17.583, Incl: 79.24, PaNode: 204.868},
		{Name: "alpha Sco", Period: 900.0, EpochPeri: 1889.0, LongPeri: 0.0, Ecc: 0.0, Axis: 3.21, Incl: 86.3, PaNode: 273.0},
	}

	for _, binaryRecord := range binaryRecords {
		if binaryRecord.Name == binaryStarName {
			return binaryRecord
		}
	}

	return BinaryStarData{Name: "NOTFOUND", Period: -99, EpochPeri: -99, LongPeri: -99, Ecc: -99, Axis: -99, Incl: -99, PaNode: -99}
}
