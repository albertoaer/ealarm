package audio

import (
	"log"
	"time"

	"github.com/faiface/beep"
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
	looping    bool
	playing    bool
}

func (audio *Audio) Stream(samples [][2]float64) (n int, ok bool) {
	if audio.streamer.Err() != nil {
		audio.playing = false
		return 0, false
	}
	for len(samples) > 0 {
		sn, sok := audio.streamer.Stream(samples)
		if !sok {
			if !audio.looping {
				audio.playing = false
				break
			}
			if err := audio.streamer.Seek(0); err != nil {
				audio.playing = false
				return n, true //Return last streamed on error seeking 0
			}
			continue
		}
		samples = samples[sn:]
		n += sn
	}
	return n, true
}

func (audio *Audio) Err() error {
	return audio.streamer.Err()
}

func (audio *Audio) Play() {
	audio.streamer.Seek(0)
	audio.looping = false
	if !audio.playing {
		rs := beep.Resample(quality, audio.sampleRate, sampleRate, audio)
		mx.Add(rs)
		audio.playing = true
	}
}

func (audio *Audio) PlayLoop() {
	audio.streamer.Seek(0)
	audio.looping = true
	if !audio.playing {
		rs := beep.Resample(quality, audio.sampleRate, sampleRate, audio)
		mx.Add(rs)
		audio.playing = true
	}
}

func (audio *Audio) Stop() {
	audio.looping = false
	audio.streamer.Seek(audio.streamer.Len())
}

func (audio *Audio) Close() {
	audio.streamer.Close()
}
