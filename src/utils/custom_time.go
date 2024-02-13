package utils

import "time"

type CustomTime time.Time

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(ct).Format("2006-01-02T15:04:05.999999Z")
	return []byte(`"` + formatted + `"`), nil
}
