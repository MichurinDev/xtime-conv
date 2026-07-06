package parser

import "time"

// Timestamp unification in different farms
// Supported: seconds, milliseconds, microseconds, nanoseconds
func ParseFlexibleUnix(timestamp int64) time.Time {
	// 10^11 is the boundary: everything that is less than exactly seconds (before 2286)
	// 10^14 is the boundary for milliseconds
	// 10^17 is the boundary for microseconds (in case they appear)

	switch {
	case timestamp < 100000000000: // ~11 digits: seconds
		return time.Unix(timestamp, 0)

	case timestamp < 100000000000000: // ~14 digits: milliseconds
		sec := timestamp / 1000
		nsec := (timestamp % 1000) * 1000000
		return time.Unix(sec, nsec)

	case timestamp < 100000000000000000: // ~17 digits: microseconds
		sec := timestamp / 1000000
		nsec := (timestamp % 1000000) * 1000
		return time.Unix(sec, nsec)

	default: // ~19 digits: nanoseconds
		sec := timestamp / 1000000000
		nsec := timestamp % 1000000000
		return time.Unix(sec, nsec)
	}
}
