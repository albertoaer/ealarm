package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/albertoaer/ealarm/audio"
	. "github.com/albertoaer/ealarm/core"
	"github.com/albertoaer/ealarm/ui"
)

func loadConfiguration(config *AlarmConfiguration) (err error) {
	hours := flag.Int("h", 0, "Number of hours")
	minutes := flag.Int("m", 0, "Number of minutes")
	seconds := flag.Int("s", 0, "Number of seconds")
	msg := flag.String("d", "ALARM!", "Display message")
	track := flag.String("t", "", "Track file as alarm tone")
	times := flag.Int("n", -1, "Number of times to play, if it's negative it will loop infinitely")
	flag.Parse()
	dur := time.Duration(*hours)*time.Hour + time.Duration(*minutes)*time.Minute + time.Duration(*seconds)*time.Second
	if dur < 0 {
		err = errors.New("Elapsed time must not be negative")
		return
	}
	config.Duration = dur
	if len(*msg) == 0 {
		err = errors.New("Message must not be empty")
		return
	}
	config.Message = *msg
	if config.Track, err = audio.From(*track); err != nil {
		return
	}
	config.Times = *times
	return
}

func main() {
	config := AlarmConfiguration{}
	err := loadConfiguration(&config)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		flag.PrintDefaults()
		return
	}
	cnt := NewController(&config)
	ui := ui.New()
	cnt.SetAction(ui.NewAlarm(&config).Show)
	cnt.SetOnQuit(func() {
		ui.Quit()
	})
	if err = cnt.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	ui.Run()
}
