package timeparse

import (
	"fmt"
	"testing"
)

func TestTimeParse(t *testing.T) {
	time := "14:07:33"
	val, err := ParseTime(time)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(val)
	}
}
func TestLeftPad(t *testing.T) {
	time := "14:7:33"
	val, err := ParseTime(time)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(val)
	}
}
func TestTooManyComponents(t *testing.T) {
	time := "14:07:33:30"
	val, err := ParseTime(time)
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("component should only accept time strings with 3 components. Got: %v", val)
	}
}
func TestMissingComponents(t *testing.T) {
	time := "14:07"
	val, err := ParseTime(time)
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("component should only accept time strings with 3 components. Got: %v", val)
	}
}
func TestTooLongComponents(t *testing.T) {
	time := "14:072:33"
	val, err := ParseTime(time)
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("component should be less than 2 runes long. Got: %v", val)
	}
}
func TestComponentOutOfBounds(t *testing.T) {
	time := "14:72:33"
	val, err := ParseTime(time)
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("component should not exceed the maximum limit for a minute. Got: %v", val)
	}
}
func TestComponentIncludesLetters(t *testing.T) {
	time := "14:w2:33"
	val, err := ParseTime(time)
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("component should not exceed the maximum limit for a minute. Got: %v", val)
	}
}
