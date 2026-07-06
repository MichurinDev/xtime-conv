package timeutil

import "time"

// Displaying the unixTimestamp timestamp relative to the tz time zone
func InLocalTimestamp(unixTimestamp time.Time, tz int) time.Time {
	return unixTimestamp.In(time.FixedZone("Other", tz*3600))
}
