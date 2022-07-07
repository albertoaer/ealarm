package ui

import (
	"fyne.io/fyne/v2"
	. "github.com/albertoaer/ealarm/core"
)

type AlarmUI struct {
	w      fyne.Window
	config *AlarmConfiguration
}

func (ui *UI) NewAlarm(config *AlarmConfiguration) *AlarmUI {
	w := ui.app.NewWindow("Alarm")
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))
	return &AlarmUI{w, config}
}

func (alarm *AlarmUI) Show(n chan bool) {
	alarm.config.Track.PlayLoop()
	alarm.w.SetCloseIntercept(func() {
		alarm.w.Hide()
		alarm.config.Track.Stop()
		n <- true
	})

	alarm.w.Show()
}
