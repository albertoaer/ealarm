package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	. "github.com/albertoaer/ealarm/core"
)

type AlarmUI struct {
	w      fyne.Window
	config *AlarmConfiguration
	signal *chan bool
}

func (ui *UI) NewAlarm(config *AlarmConfiguration) *AlarmUI {
	w := ui.app.NewWindow("Alarm")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 400))
	w.SetPadded(false)

	text := canvas.NewText(config.Message, color.White)
	text.TextSize = 32

	alarm := &AlarmUI{w, config, nil}

	endfn := func(s bool) func() {
		return func() {
			w.Hide()
			config.Track.Stop()
			*alarm.signal <- s
		}
	}

	w.SetContent(container.NewCenter(container.NewVBox(container.NewCenter(text),
		widget.NewSeparator(),
		container.NewBorder(
			nil,
			nil,
			widget.NewButton("STOP", endfn(false)),
			widget.NewButton("CONTINUE", endfn(true)),
		))))

	w.SetCloseIntercept(endfn(true))

	w.CenterOnScreen()
	return alarm
}

func (alarm *AlarmUI) Show(n chan bool) {
	alarm.config.Track.PlayLoop()
	alarm.signal = &n
	alarm.w.Show()
}
