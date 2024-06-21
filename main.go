package main

import (
	"github.com/rivo/tview"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	mainBox = tview.NewBox().SetBorder(true).SetTitle("Cheatscript")
	app     = tview.NewApplication()
)

func main() {
	err := app.SetRoot(mainBox, true).EnableMouse(true).Run()
	checkNilErr(err)
}
