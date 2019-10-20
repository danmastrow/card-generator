package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type Card struct {
	Name     string
	CardType CardType
}

func viewCardsScreen() fyne.CanvasObject {
	decks := getUserDecks()
	templates := getUserTemplates()
	deckSelect := widget.NewSelect(decks.NameList(), func(string) {})
	templateSelect := widget.NewSelect(templates.NameList(), func(string) {})
	fileInput := widget.NewEntry()
	browseButton := widget.NewButtonWithIcon("Browse", theme.ContentCopyIcon(), func() {})
	browseButton.OnTapped = func() { openImageFileBrowser(fileInput, browseButton) }
	return widget.NewForm(
		&widget.FormItem{"", widget.NewLabelWithStyle("To start generating cards please choose a deck/template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})},
		&widget.FormItem{"Selected Deck", widget.NewHBox(deckSelect, templateSelect)},
	)
}
