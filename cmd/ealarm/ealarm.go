package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	. "github.com/albertoaer/ealarm/core"
	. "github.com/albertoaer/ealarm/presets"
	"github.com/albertoaer/ealarm/ui"
)

func prepareActionFlag(ui *ui.UI, config *AlarmConfiguration, presets *Presets) *ActionFlag {
	actionmap := make(map[string]ReentrantCommandBuilder)
	for n, ac := range presets.Actions {
		actionmap[n] = ReentrantCommandBuilder(func() ReentrantCommand { return ac })
	}
	actionmap["ShowUI"] = ReentrantCommandBuilder(func() ReentrantCommand { return ui.NewAlarm(config) })
	actionflag := &ActionFlag{"ShowUI", actionmap["ShowUI"], actionmap}
	flag.Var(actionflag, "a", "Action to be executed, Syntax: CMD1[&CMD2[&CMDn]]")
	return actionflag
}

func loadConfiguration(config *AlarmConfiguration, presets *Presets) (err error) {
	duration := flag.Duration("d", 0, "Duration of the wait between alarms")
	msg := flag.String("m", "ALARM!", "Message to show at the UI")
	trackflag := NewTrackFlag()
	flag.Var(trackflag, "t", "Track file as alarm tone")
	times := flag.Int("n", -1, "Number of times to play, if it's negative it will loop infinitely")
	flag.Parse()
	if err = applyProfileTo(presets); err != nil {
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
	config.Track = trackflag.track
	config.Times = *times
	return
}

func main() {
	config := AlarmConfiguration{}
	presets := LoadPresets()
	ui := ui.New()
	actionflag := prepareActionFlag(ui, &config, presets)
	err := loadConfiguration(&config, presets)
	config.Action = actionflag.action()
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
	cnt.SetCommand(config.Action)
	cnt.SetOnQuit(func() {
		ui.Quit()
	})
	if err = cnt.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	ui.Run()
}
