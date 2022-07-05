package core

import "time"

type AlarmConfiguration struct {
	Message  string
	Duration time.Duration
	Track    string
}
