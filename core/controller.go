package core

import (
	"errors"
	"time"
)

type Controller struct {
	config *AlarmConfiguration
	action func(chan bool)
	quit   func()
}

func NewController(config *AlarmConfiguration) *Controller {
	return &Controller{config, nil, nil}
}

func (controller *Controller) SetAction(action func(chan bool)) {
	controller.action = action
}

func (controller *Controller) SetOnQuit(quit func()) {
	controller.quit = quit
}

func (controller *Controller) Start() error {
	n := make(chan bool, 1)
	if controller.action == nil {
		return errors.New("No action associated")
	}
	n <- true
	times := 0
	go func() {
		for <-n && (controller.config.Times < 0 || times < controller.config.Times) {
			time.Sleep(controller.config.Duration)
			controller.action(n)
			if controller.config.Times >= 0 {
				times++
			}
		}
		controller.quit()
	}()
	return nil
}
