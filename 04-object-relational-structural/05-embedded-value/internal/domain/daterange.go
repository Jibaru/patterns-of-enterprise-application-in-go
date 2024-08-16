package domain

import "time"

// DateRange represents a value object for a range of dates.
type DateRange struct {
	Start time.Time // Start date of the employment
	End   time.Time // End date of the employment (can be nil)
}
