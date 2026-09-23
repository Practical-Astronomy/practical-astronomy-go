package lib

type PlanetRecord struct {
	Name                        string  /* Name of planet. */
	Tp_PeriodOrbit              float64 /* Period of orbit. */
	Long_LongitudeEpoch         float64 /* Longitude at the epoch. */
	Peri_LongitudePerihelion    float64 /* Longitude of the perihelion. */
	Ecc_EccentricityOrbit       float64 /* Eccentricity of the orbit. */
	Axis_AxisOrbit              float64 /* Semi-major axis of the orbit. */
	Incl_OrbitalInclination     float64 /* Orbital inclination. */
	Node_LongitudeAscendingNode float64 /* Longitude of the ascending node. */
	Theta0_AngularDiameter      float64 /* Angular diameter at 1 AU. */
	V0_VisualMagnitude          float64 /* Visual magnitude at 1 AU. */
}

type PlanetDataPrecise struct {
	Name    string  /** Name of planet */
	Value1  float64 /** Working value 1 */
	Value2  float64 /** Working value 2 */
	Value3  float64 /** Working value 3 */
	Value4  float64 /** Working value 4 */
	Value5  float64 /** Working value 5 */
	Value6  float64 /** Working value 6 */
	Value7  float64 /** Working value 7 */
	Value8  float64 /** Working value 8 */
	Value9  float64 /** Working value 9 */
	ApValue float64 /** Working AP value */
}

func GetPlanetData(planetName string) PlanetRecord {
	var planetRecords = []PlanetRecord{
		PlanetRecord{
			Name: "Mercury", Tp_PeriodOrbit: 0.24085, Long_LongitudeEpoch: 75.5671, Peri_LongitudePerihelion: 77.612,
			Ecc_EccentricityOrbit: 0.205627, Axis_AxisOrbit: 0.387098, Incl_OrbitalInclination: 7.0051, Node_LongitudeAscendingNode: 48.449,
			Theta0_AngularDiameter: 6.74, V0_VisualMagnitude: -0.42,
		},
		PlanetRecord{
			Name: "Venus", Tp_PeriodOrbit: 0.615207, Long_LongitudeEpoch: 272.30044, Peri_LongitudePerihelion: 131.54,
			Ecc_EccentricityOrbit: 0.006812, Axis_AxisOrbit: 0.723329, Incl_OrbitalInclination: 3.3947, Node_LongitudeAscendingNode: 76.769,
			Theta0_AngularDiameter: 16.92, V0_VisualMagnitude: -4.4,
		},
		PlanetRecord{
			Name: "Earth", Tp_PeriodOrbit: 0.999996, Long_LongitudeEpoch: 99.556772, Peri_LongitudePerihelion: 103.2055,
			Ecc_EccentricityOrbit: 0.016671, Axis_AxisOrbit: 0.999985, Incl_OrbitalInclination: -99.0, Node_LongitudeAscendingNode: -99.0,
			Theta0_AngularDiameter: -99.0, V0_VisualMagnitude: -99.0,
		},
		PlanetRecord{
			Name: "Mars", Tp_PeriodOrbit: .880765, Long_LongitudeEpoch: 109.09646, Peri_LongitudePerihelion: 336.217,
			Ecc_EccentricityOrbit: 0.093348, Axis_AxisOrbit: 1.523689, Incl_OrbitalInclination: 1.8497, Node_LongitudeAscendingNode: 49.632,
			Theta0_AngularDiameter: 9.36, V0_VisualMagnitude: -1.52,
		},
		PlanetRecord{
			Name: "Jupiter", Tp_PeriodOrbit: 11.857911, Long_LongitudeEpoch: 337.917132, Peri_LongitudePerihelion: 14.6633,
			Ecc_EccentricityOrbit: 0.048907, Axis_AxisOrbit: 5.20278, Incl_OrbitalInclination: 1.3035, Node_LongitudeAscendingNode: 100.595,
			Theta0_AngularDiameter: 196.74, V0_VisualMagnitude: -9.4,
		},
		PlanetRecord{
			Name: "Saturn", Tp_PeriodOrbit: 29.310579, Long_LongitudeEpoch: 172.398316, Peri_LongitudePerihelion: 89.567,
			Ecc_EccentricityOrbit: 0.053853, Axis_AxisOrbit: 9.51134, Incl_OrbitalInclination: 2.4873, Node_LongitudeAscendingNode: 113.752,
			Theta0_AngularDiameter: 165.6, V0_VisualMagnitude: -8.88,
		},
		PlanetRecord{
			Name: "Uranus", Tp_PeriodOrbit: 84.039492, Long_LongitudeEpoch: 356.135400, Peri_LongitudePerihelion: 172.884833,
			Ecc_EccentricityOrbit: 0.046321, Axis_AxisOrbit: 19.21814, Incl_OrbitalInclination: 0.773059, Node_LongitudeAscendingNode: 73.926961,
			Theta0_AngularDiameter: 65.8, V0_VisualMagnitude: -7.19,
		},
		PlanetRecord{
			Name: "Neptune", Tp_PeriodOrbit: 165.845392, Long_LongitudeEpoch: 326.895127, Peri_LongitudePerihelion: 23.07,
			Ecc_EccentricityOrbit: 0.010483, Axis_AxisOrbit: 30.1985, Incl_OrbitalInclination: 1.7673, Node_LongitudeAscendingNode: 131.879,
			Theta0_AngularDiameter: 62.2, V0_VisualMagnitude: -6.87,
		},
	}

	for _, planetRecord := range planetRecords {
		if planetRecord.Name == planetName {
			return planetRecord
		}
	}

	return PlanetRecord{
		Name: "NOTFOUND", Tp_PeriodOrbit: -99, Long_LongitudeEpoch: -99, Peri_LongitudePerihelion: -99,
		Ecc_EccentricityOrbit: -99, Axis_AxisOrbit: -99, Incl_OrbitalInclination: -99, Node_LongitudeAscendingNode: -99,
		Theta0_AngularDiameter: -99, V0_VisualMagnitude: -99}
}
