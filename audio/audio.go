package audio

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const quality int = 1

var mx *beep.Mixer
var sampleRate beep.SampleRate

func init() {
	mx = &beep.Mixer{}
	sampleRate = beep.SampleRate(48000)
	if err := speaker.Init(sampleRate, sampleRate.N(time.Second/10)); err != nil {
		log.Fatal(err)
	}
	go speaker.Play(mx)
}

type Audio struct {
	streamer   beep.StreamSeekCloser
	sampleRate beep.SampleRate
}

func From(audio string) (*Audio, error) {
	f, err := os.Open(audio)
	if err != nil {
		return nil, err
	}

	streamer, format, err := mp3.Decode(f)

	if err != nil {
		return nil, err
	}

	return &Audio{streamer, format.SampleRate}, nil
}

func (audio *Audio) Play() {
	audio.streamer.Seek(0)
	rs := beep.Resample(quality, audio.sampleRate, sampleRate, audio.streamer)
	mx.Add(rs)
}

func (audio *Audio) PlayLoop() {
	audio.streamer.Seek(0)
	rs := beep.Resample(quality, audio.sampleRate, sampleRate, beep.Loop(-1, audio.streamer))
	mx.Add(rs)
}

func (audio *Audio) Stop() {
	audio.streamer.Seek(audio.streamer.Len())
}

func (audio *Audio) Close() {
	audio.streamer.Close()
}
