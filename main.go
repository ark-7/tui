package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	modal := tview.NewModal().
		SetText("Aww Snap! This TUI app is still in progress! :( Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})
	if err := app.SetRoot(modal, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
