package lib

import "time"

func NowPtr() *time.Time {
	now := time.Now()
	return &now
}

func StrPtr(str string) *string {
	return &str
}
