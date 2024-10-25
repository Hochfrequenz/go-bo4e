package internal

import (
	"encoding/json"
	"time"
)

type TimeOnly time.Time

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("15:04:05"))
}

func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return err
	}

	*t = TimeOnly(parsedTime)
	return nil
}
