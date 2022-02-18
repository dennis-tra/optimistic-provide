package util

import "time"

func NowPtr() *time.Time {
	now := time.Now()
	return &now
}

func StrPtr(str string) *string {
	return &str
}

func TimeToStr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	if t.IsZero() {
		return nil
	}
	return StrPtr(t.Format(time.RFC3339))
}
