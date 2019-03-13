// Package space calculates age in terms of a planet's years.
package space

// Planet is the name of one of the eight planets of our solar system.
type Planet string

// OrbitalPeriod returns how long a planet takes to orbit the sun in seconds.
func (p Planet) OrbitalPeriod() float64 {
	var conversions = map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return 31557600 * conversions[p]
}

// Age calculates someone's age in terms of the years of a given planet.
func Age(seconds float64, planet Planet) float64 {
	return seconds / planet.OrbitalPeriod()
}
