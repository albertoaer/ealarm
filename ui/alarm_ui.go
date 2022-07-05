package ui

import (
	"fyne.io/fyne/v2"
)

type AlarmUI struct {
	w fyne.Window
}

func (ui *UI) NewAlarm() *AlarmUI {
	w := ui.app.NewWindow("Alarm")
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))
	return &AlarmUI{w}
}

func (alarm *AlarmUI) Show(n chan bool) {
	alarm.w.SetCloseIntercept(func() {
		n <- true
		alarm.w.Hide()
	})
	alarm.w.Show()
}
