package ui

import (
	"fyne.io/fyne/v2"
	fynapp "fyne.io/fyne/v2/app"
)

type UI struct {
	app fyne.App
}

func New() *UI {
	return &UI{fynapp.New()}
}

func (ui *UI) Run() {
	ui.app.Run()
}
