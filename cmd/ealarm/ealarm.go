package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/albertoaer/ealarm/controller"
	. "github.com/albertoaer/ealarm/core"
	"github.com/albertoaer/ealarm/ui"
)

func getInput(config *AlarmConfiguration) (err error) {
	hours := flag.Int("h", 0, "Number of hours")
	minutes := flag.Int("m", 0, "Number of minutes")
	seconds := flag.Int("s", 0, "Number of seconds")
	flag.Parse()
	dur := time.Duration(*hours)*time.Hour + time.Duration(*minutes)*time.Minute + time.Duration(*seconds)*time.Second
	if dur == 0 {
		err = errors.New("Trying to elapse time 0")
	} else {
		config.Duration = dur
	}
	return
}

func main() {
	config := AlarmConfiguration{}
	err := getInput(&config)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		flag.PrintDefaults()
		return
	}
	cnt := controller.New(config.Duration)
	ui := ui.New()
	cnt.SetAction(ui.NewAlarm(&config).Show)
	if err = cnt.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	ui.Run()
}
