package audio

import (
	"os"

	"github.com/faiface/beep/mp3"
)

type Playable interface {
	Play()
	PlayLoop()
	Stop()
	Close()
}

func From(audio string) (Playable, error) {
	f, err := os.Open(audio)
	if err != nil {
		return nil, err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return nil, err
	}

	return &Audio{streamer, format.SampleRate, false, false}, nil
}

func Silence() Playable {
	return &noaudio{}
}

type noaudio struct{}

func (*noaudio) Play()     {}
func (*noaudio) PlayLoop() {}
func (*noaudio) Stop()     {}
func (*noaudio) Close()    {}
