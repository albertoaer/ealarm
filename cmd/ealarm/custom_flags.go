package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/albertoaer/ealarm/audio"
	. "github.com/albertoaer/ealarm/core"
	. "github.com/albertoaer/ealarm/presets"
)

type ActionFlag struct {
	repr      string
	action    ReentrantCommandBuilder
	actionmap map[string]ReentrantCommandBuilder
}

func (c *ActionFlag) String() string {
	return c.repr
}

func (c *ActionFlag) Set(r string) error {
	total := make([]ReentrantCommandBuilder, 0)
	for _, e := range strings.Split(r, "&") {
		command, in := c.actionmap[strings.Trim(e, " \t")]
		if !in {
			return fmt.Errorf("Unknown action: %s", e)
		}
		total = append(total, command)
	}
	if len(total) == 1 {
		c.action = total[0]
	} else {
		c.action = ManyCommandsBuilder(total...)
	}
	c.repr = r
	return nil
}

func getProfile(presets *Presets) (Profile, error) {
	if flag.NArg() == 0 {
		d := presets.Profiles["default"]
		return d, nil //If the profile is not inside profiles the profile will be null
	}
	if flag.NArg() > 1 {
		return nil, errors.New("Expecting just one non-flag arg as profile name")
	}
	m, in := presets.Profiles[flag.Arg(0)]
	if !in {
		return nil, fmt.Errorf("Unknown profile: %s", flag.Arg(0))
	}
	return m, nil
}

func applyProfileTo(presets *Presets) error {
	profile, err := getProfile(presets)
	if err != nil {
		return err
	}
	if profile == nil {
		return nil
	}
	write := make(map[string]string)
	for k, v := range profile {
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

type TrackFlag struct {
	src   string
	track audio.Playable
}

func NewTrackFlag() *TrackFlag {
	return &TrackFlag{"SILENCE", audio.Silence()}
}

func (c *TrackFlag) String() string {
	return c.src
}

func (c *TrackFlag) Set(r string) (err error) {
	c.src = r
	c.track, err = audio.From(r)
	return
}
