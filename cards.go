package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type Card struct {
	Name     string
	CardType CardType
}

func viewCardsScreen() fyne.CanvasObject {
	// cards := []Card{}
	// labelVal := time.Now()
	// testLabel := widget.NewLabel(labelVal.String())
	decks := getUserDecks()
	templates := getUserTemplates()
	deckSelect := widget.NewSelect(decks.NameList(), func(deckName string) { fmt.Println(deckName) })
	templateSelect := widget.NewSelect(templates.NameList(), func(string) {})
	return widget.NewForm(
		&widget.FormItem{"", widget.NewLabelWithStyle("To start generating cards please choose a deck/template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})},
		&widget.FormItem{"Selected Deck", widget.NewHBox(deckSelect, templateSelect)},
		// &widget.FormItem{"Test", testLabel}
	)
}
