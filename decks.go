package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type Deck struct {
	Name  string
	Cards []Card
}

type Decks struct {
	Name      string
	CardTypes []CardType
	Decks     []Deck
}

func (d Decks) NameList() []string {
	var list []string
	for _, deck := range d.Decks {
		list = append(list, deck.Name)
	}
	return list
}

func getUserDecks() Decks {
	var decks Decks

	file, err := os.Open("json/decks.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(byteValue), &decks)
	if err != nil {
		fmt.Println(err)
	}
	return decks
}

func viewDecksScreen() fyne.CanvasObject {
	w := widget.NewVBox(
		widget.NewLabelWithStyle("Decks:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

	deckPicker := createDeckPickerUi()
	for _, deck := range deckPicker {
		w.Append(deck)
	}

	return w
}

func createDeckPickerUi() []fyne.CanvasObject {
	result := []fyne.CanvasObject{}
	decks := getUserDecks()

	for _, deck := range decks.Decks {
		result = append(result, widget.NewButton(deck.Name, func() {
			// TODO: navigate to deck viewer with that deck selected
		}))
	}

	return result
}
