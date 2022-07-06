package audio

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Audio struct {
	streamer beep.StreamSeekCloser
}

func Play(audio string) (*Audio, error) {
	f, err := os.Open(audio)
	if err != nil {
		return nil, err
	}

	streamer, format, err := mp3.Decode(f)

	if err != nil {
		return nil, err
	}

	if err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return nil, err
	}

	go speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		streamer.Close()
	})))

	return &Audio{streamer}, nil
}

func (audio *Audio) Stop() {
	audio.streamer.Close()
}
