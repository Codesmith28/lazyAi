package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
)

var (
	ModelsPane = tview.NewFlex()
	ModelList  = tview.NewList()
	Selected   *internal.Model
)

func init() {
	// Configure the model list
	ModelList.ShowSecondaryText(false).SetTitle("Models").SetBorder(true)
	Selected = &internal.Model{}

	// Add models to the list
	ModelList.AddItem("Model 1", "", 0, func() {
		SelectModel("Model 1")
	})
	ModelList.AddItem("Model 2", "", 0, func() {
		SelectModel("Model 2")
	})

	// Enable mouse support
	ModelList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			currentItem := ModelList.GetCurrentItem()
			mainText, _ := ModelList.GetItemText(currentItem)
			SelectModel(mainText)
		}
		return event
	})

	ModelsPane.AddItem(ModelList, 0, 1, true)
}
