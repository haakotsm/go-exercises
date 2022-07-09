//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeComponent struct {
	Hours, Minutes, Seconds int
}

func checkComponent(component string) (int, error) {
	if len(component) > 2 {
		return 0, fmt.Errorf("length of component %v cannot exceed 2", component)
	}
	if len(component) < 2 {
		component = fmt.Sprintf("%02s", component)
	}
	convertedValue, err := strconv.Atoi(component)
	return convertedValue, err

}
func ParseTimeWithDelimiter(timeString string, delimiter string) (TimeComponent, error) {
	timeString = strings.ReplaceAll(timeString, delimiter, ":")
	timeComponent, err := ParseTime(timeString)
	return timeComponent, err
}

func ParseTime(timeString string) (TimeComponent, error) {
	split := strings.Split(timeString, ":")
	timeCompnent := TimeComponent{}
	if len(split) != 3 {
		return timeCompnent, fmt.Errorf("not enough components in supplied string. Got: %v, expected 3", len(split))
	}
	hours, err := checkComponent(split[0])
	if err != nil {
		return TimeComponent{}, fmt.Errorf("error converting hours component to int: %v", err)
	} else {
		if hours > 23 {
			return TimeComponent{}, fmt.Errorf("hours cannot be greater than 23. Got: %v", hours)
		}
		timeCompnent.Hours = hours
	}
	minutes, err := checkComponent(split[1])
	if err != nil {
		return TimeComponent{}, fmt.Errorf("error converting mintes component to int: %v", err)
	} else {
		if minutes > 59 {
			return TimeComponent{}, fmt.Errorf("minutes cannot be greater than 59. Got: %v", minutes)
		}
		timeCompnent.Minutes = minutes
	}
	seconds, err := checkComponent(split[2])
	if err != nil {
		return TimeComponent{}, fmt.Errorf("error converting seconds component to int: %v", err)
	} else {
		if seconds > 59 {
			return TimeComponent{}, fmt.Errorf("seconds cannot be greater than 59. Got: %v", seconds)
		}
		timeCompnent.Seconds = seconds
	}
	return timeCompnent, nil
}
