package gommon

import (
	"fmt"
	"strconv"
	"time"
)

const (
	TimezoneTehran   = "Asia/Tehran"
	TimezoneDubai    = "Asia/Dubai"
	TimezoneIstanbul = "Asia/Istanbul"
	TimezoneKabul    = "Asia/Kabul"
	TimezoneKuwait   = "Asia/Kuwait"
)

// PersianDateSplitter splits a number into x, y, z parts based on specific rules.
// for example : '13690516' => 1369 , 05 , 16
func PersianDateSplitter(input string) (x, y, z int, err error) {
	// Convert the input number to a string
	//str := fmt.Sprintf("%08d", input) // Ensures it's padded to 8 digits if needed
	str := input // Ensures it's padded to 8 digits if needed

	// Ensure the string length is exactly 8
	if len(str) != 8 {
		return 0, 0, 0, fmt.Errorf("input must be an 8-digit number")
	}

	// Parse the parts
	x, err = strconv.Atoi(str[:4]) // First 4 digits
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse x: %v", err)
	}
	y, err = strconv.Atoi(str[4:6]) // Next 2 digits
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse y: %v", err)
	}
	z, err = strconv.Atoi(str[6:]) // Last 2 digits
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse z: %v", err)
	}

	return x, y, z, nil
}

// DurationUntilInTimezone returns remaining duration until a target time
func DurationUntilInTimezone(hour, minute int, timezone string) (time.Duration, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, fmt.Errorf("failed to load timezone: %w", err)
	}

	now := time.Now().In(loc)
	targetTime := time.Date(
		now.Year(), now.Month(), now.Day(),
		hour, minute, 0, 0, loc,
	)

	if now.After(targetTime) {
		targetTime = targetTime.Add(24 * time.Hour)
	}

	return time.Until(targetTime), nil
}

// DurationUntilNextHour returns remaining duration until next hour
func DurationUntilNextHour(timezone string) (time.Duration, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, fmt.Errorf("failed to load timezone %s: %w", timezone, err)
	}

	now := time.Now().In(loc)
	nextHour := time.Date(
		now.Year(), now.Month(), now.Day(),
		now.Hour()+1, 0, 0, 0, loc,
	)

	return time.Until(nextHour), nil
}

// DurationUntilNextInterval returns the duration until the next rounded interval (e.g., 5, 10, 15 mins) in the given timezone.
func DurationUntilNextInterval(timezone string, intervalMinutes int) (time.Duration, error) {
	if intervalMinutes <= 0 || intervalMinutes > 60 {
		return 0, fmt.Errorf("invalid interval: must be between 1 and 60 minutes")
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, fmt.Errorf("failed to load timezone %s: %w", timezone, err)
	}

	now := time.Now().In(loc)

	// Calculate how many minutes to add to reach the next interval
	minutes := now.Minute()
	roundedUpMin := ((minutes / intervalMinutes) + 1) * intervalMinutes

	// Handle hour overflow
	addHours := 0
	if roundedUpMin >= 60 {
		roundedUpMin = 0
		addHours = 1
	}

	next := time.Date(
		now.Year(), now.Month(), now.Day(),
		now.Hour()+addHours, roundedUpMin, 0, 0, loc,
	)

	return time.Until(next), nil
}

func ConvertFloatToTime(t float64) time.Time {
	seconds := int64(t)
	nanoseconds := int64((t - float64(seconds)) * 1e9)
	return time.Unix(seconds, nanoseconds)
}
