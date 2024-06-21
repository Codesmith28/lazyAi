package main

import (
	"github.com/rivo/tview"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// fmt.Println("Hello there!")
	box := tview.NewBox().SetBorder(true).SetTitle("Hello there!")
	err := tview.NewApplication().SetRoot(box, true).EnableMouse(true).Run()
	checkNilErr(err)
}
