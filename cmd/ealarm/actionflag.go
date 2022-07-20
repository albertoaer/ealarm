package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

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
