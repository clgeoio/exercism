//Package space calculates how old someone would be on different planets
package space

import "fmt"

type Planet string

const EarthYearInSeconds = 31557600

var planetConversion = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age converts a float64 duration in seconds to respective planetary years
func Age(sec float64, planet Planet) float64 {
	earthYears := sec / EarthYearInSeconds

	convert, ok := planetConversion[planet]
	if !ok {
		fmt.Printf("Unrecognized planet %s", planet)
		return 0.0
	}
	return earthYears / convert
}
