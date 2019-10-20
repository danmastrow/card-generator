package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	file, newlyCreated := getUserGlobalConfig("config.json")
	if newlyCreated {

	}
	fmt.Println(file)

	showMainWindow(app)
}

func showMainWindow(app fyne.App) {
	w := app.NewWindow("Card generator")
	w.CenterOnScreen()

	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
		// a quit item will be appended to the first menu
	), fyne.NewMenu("Edit",
		fyne.NewMenuItem("Settings", func() { fmt.Println("Menu Settings") }),
	)))
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("View Decks", theme.HomeIcon(), viewDecksScreen()),
		widget.NewTabItemWithIcon("Generate Cards", theme.ContentCopyIcon(), generateCardsScreen()))
	tabs.SetTabLocation(widget.TabLocationLeading)

	w.SetContent(tabs)

	w.ShowAndRun()
}
