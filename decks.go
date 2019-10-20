package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Card struct {
	Name string
}

type Deck struct {
	Name  string
	Cards []Card
}

type Decks struct {
	Name  string
	Decks []Deck
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

	file, err := os.Open("decks.json")
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
	fmt.Println(decks)
	return decks
}
