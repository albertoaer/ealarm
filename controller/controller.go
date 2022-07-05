package controller

import (
	"errors"
	"time"
)

type Controller struct {
	lapse  time.Duration
	action func(chan bool)
}

func New(lapse time.Duration) *Controller {
	return &Controller{lapse, nil}
}

func (controller *Controller) SetAction(action func(chan bool)) {
	controller.action = action
}

func (controller *Controller) Start() error {
	n := make(chan bool, 1)
	if controller.action == nil {
		return errors.New("No action associated")
	}
	n <- true
	go func() {
		for <-n {
			time.Sleep(controller.lapse)
			controller.action(n)
		}
	}()
	return nil
}
