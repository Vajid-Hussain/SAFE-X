package cmd

import (
	"time"
)

func GetTimeInTimeZone(timezone string) (string, error) {
	localtion, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	currentTime := time.Now().In(localtion)
	return currentTime.Format(time.RFC1123), nil
}
