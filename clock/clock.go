// Package clock handles times without dates.
package clock

import (
	"fmt"
)

const hourMinutes = 60
const dayMinutes = 24 * hourMinutes

// Clock is a simple representation of a time with hour and minute.
type Clock struct {
	hour   int
	minute int
}

// New returns a clock at a certain time. It does not consider dates.
func New(hour, minute int) Clock {
	// This comment contains a slightly different implementation that does better on
	// BenchmarkCreateClocks but worse on the add and subtract benchmarks.
	// minutes := (hour*hourMinutes + minute) % dayMinutes
	// if minutes < 0 {
	//	minutes += dayMinutes
	// }
	minutes := ((hour*hourMinutes+minute)%(dayMinutes) + (dayMinutes)) % (dayMinutes)
	return Clock{minutes / hourMinutes, minutes % hourMinutes}
}

// String returns the clock's 24-hour time formatted as a string, such as "09:22" or "15:59".
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add returns a clock that is a number of minutes ahead of this one.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract returns a clock that is a number of minutes behind this one.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
