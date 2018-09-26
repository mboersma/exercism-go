// Package gigasecond calculates the moment 10^9 seconds from a given time.
package gigasecond

import "time"

// AddGigasecond returns a time with a gigasecond (10^9 seconds) added to it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
