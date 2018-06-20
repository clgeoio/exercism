// Package gigasecond provides a function to add a gigasecond (10^9 seconds) to a time
package gigasecond

import "time"

const GigaSec = time.Duration(1E9) * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(GigaSec)
}
