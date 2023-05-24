package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/derailed/tcell/v2"
	"github.com/derailed/tview"
)

func Throw(err any) {
	if err != nil {
		fmt.Println("error:", err)
		fmt.Printf("%s\n", debug.Stack())
		os.Exit(-1)
	}
}

func main() {
	var text = tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(a) to add a new contact \n(q) to quit")

	pages := tview.NewPages()

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).AddItem(text, 0, 1, false)
	//
	//flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//	switch event.Rune() {
	//	case 'q':
	//	case 'a':
	//		pages.SwitchToPage("addNewName")
	//	}
	//	return event
	//})

	form := tview.NewForm()

	var nameList []string

	var newName = ""
	form.AddInputField("Name", "", 20, nil, func(text string) {
		newName = text
	})

	form.AddButton("Save", func() {
		nameList = append(nameList, newName)
		pages.SwitchToPage("Menu")
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("addNewName", form, true, false)
	app := tview.NewApplication()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
		case 'a':
			//form.Clear(true)
			pages.SwitchToPage("addNewName")
		}
		return event
	})
	app.SetRoot(pages, true).EnableMouse(true)
	Throw(app.Run())
}
