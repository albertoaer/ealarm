package controller

import (
	"errors"
	"time"

	. "github.com/albertoaer/ealarm/core"
)

type Controller struct {
	config *AlarmConfiguration
	action func(chan bool)
}

func New(config *AlarmConfiguration) *Controller {
	return &Controller{config, nil}
}

func (controller *Controller) SetAction(action func(chan bool)) {
	controller.action = action
}

func (controller *Controller) Start() error {
	n := make(chan bool, 2)
	if controller.action == nil {
		return errors.New("No action associated")
	}
	n <- true
	times := 0
	go func() {
		for times < controller.config.Times && <-n {
			time.Sleep(controller.config.Duration)
			controller.action(n)
			times++
		}
	}()
	return nil
}
