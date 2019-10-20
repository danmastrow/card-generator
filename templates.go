package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

type Template struct {
	Name          string
	DefaultWidth  float64
	DefaultHeight float64
}

type Templates struct {
	Templates []Template
}

func (t Templates) NameList() []string {
	var list []string
	for _, deck := range t.Templates {
		list = append(list, deck.Name)
	}
	return list
}
func viewTemplatesScreen() fyne.CanvasObject {
	w := widget.NewVBox(
		widget.NewLabelWithStyle("Templates:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

	templatePicker := createTemplatePickerUi()
	for _, deck := range templatePicker {
		w.Append(deck)
	}

	return w
}

func getUserTemplates() Templates {
	var templates Templates

	file, err := os.Open("json/templates.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(byteValue), &templates)
	if err != nil {
		fmt.Println(err)
	}
	return templates
}

func createTemplatePickerUi() []fyne.CanvasObject {
	result := []fyne.CanvasObject{}
	templates := getUserTemplates()

	for _, template := range templates.Templates {
		result = append(result, widget.NewButton(template.Name, func() {
			// TODO: navigate to template viewer
		}))
	}

	return result
}

func addTemplateScreen() {
	// fileInput := widget.NewEntry()
	// browseButton := widget.NewButtonWithIcon("Browse", theme.ContentCopyIcon(), func() {})
	// browseButton.OnTapped = func() { openImageFileBrowser(fileInput, browseButton) }
	// return widget.NewForm(
	// 	&widget.FormItem{"Template Image", widget.NewHBox(browseButton, fileInput)},
	// )
}

func openImageFileBrowser(input *widget.Entry, button *widget.Button) {
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
