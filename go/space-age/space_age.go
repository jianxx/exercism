package space

// Planet name
type Planet string

// Age to calculate how old someone would be on the planet
func Age(seconds float64, planet Planet) float64 {
	const SecondsEarthYear = 31557600.0
	switch planet {
	case "Mercury":
		return seconds / (SecondsEarthYear * 0.2408467)
	case "Venus":
		return seconds / (SecondsEarthYear * 0.61519726)
	case "Mars":
		return seconds / (SecondsEarthYear * 1.8808158)
	case "Jupiter":
		return seconds / (SecondsEarthYear * 11.862615)
	case "Saturn":
		return seconds / (SecondsEarthYear * 29.447498)
	case "Uranus":
		return seconds / (SecondsEarthYear * 84.016846)
	case "Neptune":
		return seconds / (SecondsEarthYear * 164.79132)
	default:
		return seconds / SecondsEarthYear
	}
}
