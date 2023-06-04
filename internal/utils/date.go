package utils

import "time"

// ParseStringToTime convert date with string type to time.Time type
func ParseStringToTime(dob string) (parsedTime time.Time) {
	parsedTime, _ = time.Parse("02/01/2006", dob)
	return parsedTime
}

// ParseTimeToString convert date with time.Time type to string type
func ParseTimeToString(dob time.Time) (parsedTime string) {
	parsedTime = dob.Format("02/01/2006")
	return parsedTime
}
