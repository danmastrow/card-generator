package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

func generateCardsScreen() fyne.CanvasObject {
	decks := getUserDecks()
	deckSelect := widget.NewSelect(decks.NameList(), func(string) {})

	fileInput := widget.NewEntry()
	browseButton := widget.NewButtonWithIcon("Browse", theme.ContentCopyIcon(), func() {})
	browseButton.OnTapped = func() { openFileBrowser(fileInput, browseButton) }
	return widget.NewForm(
		&widget.FormItem{"", widget.NewLabelWithStyle("To start generating cards please choose a template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})},
		&widget.FormItem{"Selected Deck", widget.NewHBox(deckSelect)},
		&widget.FormItem{"Template Image", widget.NewHBox(browseButton, fileInput)},
	)
}

func openFileBrowser(input *widget.Entry, button *widget.Button) {
	button.Disable()
	file, err := dialog.File().Title("Select image template").Filter("*.png", "png").Load()
	if err != nil && err.Error() != "Cancelled" {
		fmt.Println(err)

		dialog.Message("Please select a valid file").Title("Error selecting image").Error()
	} else {
		input.SetText(file)
	}
	button.Enable()
}
