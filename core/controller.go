package core

import (
	"errors"
	"time"
)

type Controller struct {
	config *AlarmConfiguration
	cmd    ReentrantCommand
	quit   func()
}

func NewController(config *AlarmConfiguration) *Controller {
	return &Controller{config, nil, nil}
}

func (controller *Controller) SetCommand(cmd ReentrantCommand) {
	controller.cmd = cmd
}

func (controller *Controller) SetOnQuit(quit func()) {
	controller.quit = quit
}

func (controller *Controller) Start() error {
	n := make(chan bool, 1)
	if controller.cmd == nil {
		return errors.New("No command associated")
	}
	n <- true
	times := 0
	go func() {
		for <-n && (controller.config.Times < 0 || times < controller.config.Times) {
			time.Sleep(controller.config.Duration)
			controller.cmd.Launch(n)
			if controller.config.Times >= 0 {
				times++
			}
		}
		controller.quit()
	}()
	return nil
}
