package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const timestampFormat = "20060102T150405Z"

func GetTimestamps(c *gin.Context) {
	// Parse query parameters
	period := c.Query("period")
	timezone := c.Query("tz")
	startStr := c.Query("t1")
	endStr := c.Query("t2")

	// Parse timestamps
	startTime, err := time.Parse(timestampFormat, startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "desc": "Invalid start timestamp"})
		return
	}

	endTime, err := time.Parse(timestampFormat, endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "desc": "Invalid end timestamp"})
		return
	}

	// Apply timezone
	location, err := time.LoadLocation(timezone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "desc": "Invalid timezone"})
		return
	}
	startTime = startTime.In(location)
	endTime = endTime.In(location)

	// Calculate timestamps within the given range and period
	timestamps, err := calculateTimestamps(startTime, endTime, period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "desc": "Unsupported period"})
		return
	}

	// Return the matching timestamps
	c.JSON(http.StatusOK, timestamps)
	return
}

func calculateTimestamps(startTime, endTime time.Time, period string) ([]string, error) {
	timestamps := make([]string, 0)

	invocationTimestampPhase := true

	// Generate timestamps based on the period until the end time
	for current := startTime; current.Before(endTime); {

		// When we start we first set the invocationTimestamp correctly before calculating and adding results to the timestamps slice
		if invocationTimestampPhase {
			// Period day needs to add 1 hour only and not a whole day at the initial run for correct invocation period
			if period == "1d" {
				current = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour()+1, 0, 0, 0, startTime.Location())
				invocationTimestampPhase = false
				continue
			}
			invocationTimestampPhase = false
		} else {
			// Add the result to the timestamps slice in UTC format
			timestamps = append(timestamps, current.UTC().Format(timestampFormat))
		}

		// Increment the current time based on the period
		switch period {
		case "1h":
			current = time.Date(current.Year(), current.Month(), current.Day(), current.Hour()+1, 0, 0, 0, current.Location())
		case "1d":
			current = time.Date(current.Year(), current.Month(), current.Day()+1, current.Hour(), 0, 0, 0, current.Location())
		case "1mo":
			current = time.Date(current.Year(), current.Month()+1, 1, 0, 0, 0, 0, current.Location())
		case "1y":
			current = time.Date(current.Year()+1, 1, 1, 0, 0, 0, 0, current.Location())
		default:
			return nil, errors.New("invalid period")
		}
	}

	return timestamps, nil
}
