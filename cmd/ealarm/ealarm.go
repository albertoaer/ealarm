package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/albertoaer/ealarm/audio"
	. "github.com/albertoaer/ealarm/core"
	. "github.com/albertoaer/ealarm/presets"
	"github.com/albertoaer/ealarm/ui"
)

func getProfile(presets *Presets) error {
	if flag.NArg() == 0 {
		return nil
	}
	if flag.NArg() > 1 {
		return errors.New("Expecting just one non-flag arg as profile name")
	}
	m, in := presets.Profiles[flag.Arg(0)]
	if !in {
		return fmt.Errorf("Unknown profile: %s", flag.Arg(0))
	}
	write := make(map[string]string)
	for k, v := range m {
		write[k] = v
	}
	flag.Visit(func(f *flag.Flag) {
		delete(write, f.Name)
	})
	for k, v := range write {
		if e := flag.Set(k, v); e != nil {
			return e
		}
	}
	return nil
}

func loadConfiguration(config *AlarmConfiguration, presets *Presets) (err error) {
	duration := flag.Duration("d", 0, "Duration of the wait between alarms")
	msg := flag.String("m", "ALARM!", "Message to show at the UI")
	track := flag.String("t", "", "Track file as alarm tone")
	times := flag.Int("n", -1, "Number of times to play, if it's negative it will loop infinitely")
	flag.Parse()
	if err = getProfile(presets); err != nil {
		return
	}
	if *duration < 0 {
		return errors.New("Duration must not be negative")
	}
	dur := *duration
	config.Duration = dur
	if len(*msg) == 0 {
		return errors.New("Message must not be empty")
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
	presets := LoadPresets()
	err := loadConfiguration(&config, presets)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		var optional_flags string
		flag.VisitAll(func(f *flag.Flag) {
			optional_flags += "[-" + f.Name + "] "
		})
		fmt.Printf("Usage: %s %s[profile]\n", os.Args[0], optional_flags)
		flag.PrintDefaults()
		return
	}
	cnt := NewController(&config)
	ui := ui.New()
	cnt.SetCommand(ui.NewAlarm(&config))
	cnt.SetOnQuit(func() {
		ui.Quit()
	})
	if err = cnt.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	ui.Run()
}
