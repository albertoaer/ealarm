package core

import (
	"time"

	"github.com/albertoaer/ealarm/audio"
)

type AlarmConfiguration struct {
	Message  string
	Duration time.Duration
	Track    audio.Playable
	Times    int
	Action   ReentrantCommand
}
