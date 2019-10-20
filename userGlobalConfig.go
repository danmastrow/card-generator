package main

import (
	"fyne.io/fyne"
)

type UserGlobalConfig struct {
	cardWidth  float64
	cardHeight float64
}

func showSetGlobalConfiguration(app fyne.App) {
	w := app.NewWindow("Card Generator - First Time Setup")
	w.CenterOnScreen()
	w.ShowAndRun()
}
