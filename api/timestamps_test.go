package api

import (
	"reflect"
	"testing"
	"time"
)

func TestCalculateTimestamps(t *testing.T) {

	// Set Test Case timezone
	location, _ := time.LoadLocation("Europe/Athens")

	// Test case 1: Period "month"

	startTime := time.Date(2021, time.February, 2, 20, 46, 3, 0, time.UTC)
	endTime := time.Date(2021, time.November, 11, 12, 34, 56, 0, time.UTC)

	startTime = startTime.In(location)
	endTime = endTime.In(location)

	expected1 := []string{
		"20210228T220000Z",
		"20210331T210000Z",
		"20210430T210000Z",
		"20210531T210000Z",
		"20210630T210000Z",
		"20210731T210000Z",
		"20210831T210000Z",
		"20210930T210000Z",
		"20211031T220000Z",
	}
	result1, _ := calculateTimestamps(startTime, endTime, "1mo")
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Unexpected result for period 'month'. Expected: %v, Got: %v", expected1, result1)
	}

	// Test case 2: Period "year"

	startTime = time.Date(2021, time.February, 2, 20, 46, 3, 0, time.UTC)
	endTime = time.Date(2023, time.November, 11, 12, 34, 56, 0, time.UTC)

	startTime = startTime.In(location)
	endTime = endTime.In(location)

	expected2 := []string{
		"20211231T220000Z",
		"20221231T220000Z",
	}
	result2, _ := calculateTimestamps(startTime, endTime, "1y")
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Unexpected result for period 'year'. Expected: %v, Got: %v", expected2, result2)
	}

	// Test case 3: Period "day"

	startTime = time.Date(2021, time.February, 14, 20, 46, 3, 0, time.UTC)
	endTime = time.Date(2021, time.February, 17, 12, 34, 56, 0, time.UTC)

	startTime = startTime.In(location)
	endTime = endTime.In(location)

	expected3 := []string{
		"20210214T210000Z",
		"20210215T210000Z",
		"20210216T210000Z",
	}
	result3, _ := calculateTimestamps(startTime, endTime, "1d")
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Unexpected result for period 'day'. Expected: %v, Got: %v", expected3, result3)
	}

	// Test case 4: Period "hour"

	startTime = time.Date(2021, time.February, 14, 20, 46, 3, 0, time.UTC)
	endTime = time.Date(2021, time.February, 15, 12, 34, 56, 0, time.UTC)

	startTime = startTime.In(location)
	endTime = endTime.In(location)

	expected4 := []string{
		"20210214T210000Z",
		"20210214T220000Z",
		"20210214T230000Z",
		"20210215T000000Z",
		"20210215T010000Z",
		"20210215T020000Z",
		"20210215T030000Z",
		"20210215T040000Z",
		"20210215T050000Z",
		"20210215T060000Z",
		"20210215T070000Z",
		"20210215T080000Z",
		"20210215T090000Z",
		"20210215T100000Z",
		"20210215T110000Z",
		"20210215T120000Z",
	}
	result4, _ := calculateTimestamps(startTime, endTime, "1h")
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Unexpected result for period 'hour'. Expected: %v, Got: %v", expected4, result4)
	}
}
