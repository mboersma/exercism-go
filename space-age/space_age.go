// Package space calculates age in terms of a planet's years.
package space

// Planet is the name of one of the eight planets of our solar system.
type Planet string

// OrbitalPeriod returns how long a planet takes to orbit the sun in seconds.
func (p Planet) OrbitalPeriod() float64 {
	const earth = 31557600.0
	switch p {
	case "Mercury":
		return earth * 0.2408467
	case "Venus":
		return earth * 0.61519726
	case "Mars":
		return earth * 1.8808158
	case "Jupiter":
		return earth * 11.862615
	case "Saturn":
		return earth * 29.447498
	case "Uranus":
		return earth * 84.016846
	case "Neptune":
		return earth * 164.79132
	default:
		return earth
	}
}

// Age calculates someone's age in terms of the years of a given planet.
func Age(seconds float64, planet Planet) float64 {
	return seconds / planet.OrbitalPeriod()
}
