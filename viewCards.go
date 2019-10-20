package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func viewDecksScreen() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabelWithStyle("Todo: View saved decks here :) ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
	)
}
