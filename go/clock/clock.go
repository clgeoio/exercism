//Package clock defines clocks and allows the manipulation of them
package clock

import (
	"fmt"
	"math"
)

type Clock struct {
	Hours, Minutes int
}

//New creates a new clock with hours and mins
func New(hour, minute int) Clock {
	h := int(math.Floor((float64(hour)*60.0+float64(minute))/60.0)) % 24
	m := minute % 60

	if h < 0 {
		h = h + 24
	}

	if m < 0 {
		m = m + 60
	}

	return Clock{Hours: h, Minutes: m}
}

//String ouputs a clock time in 24hr HH:MM time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

//Add increses a clock time by a number of minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.Hours, c.Minutes+minutes)
}

//Subtract adds negative minutes to a clock time
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
