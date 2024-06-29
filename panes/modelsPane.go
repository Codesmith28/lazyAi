package panes

import (
	"fmt"

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
	ModelList.ShowSecondaryText(false).SetTitle("Models").SetBorder(true)
	Selected = &internal.Model{}

	ModelList.AddItem("Model 1", "", 0, func() {
		selectModel("Model 1")
	})
	ModelList.AddItem("Model 2", "", 0, func() {
		selectModel("Model 2")
	})

	// Enable mouse support
	ModelList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			currentItem := ModelList.GetCurrentItem()
			mainText, _ := ModelList.GetItemText(currentItem)
			selectModel(mainText)
		}
		return event
	})

	ModelsPane.AddItem(ModelList, 0, 1, true)
}

func selectModel(model string) {
	Selected.SelectedModel = model
	fmt.Printf("Selected Model: %s\n", model)
}
