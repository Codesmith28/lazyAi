package panes

import (
	"fmt"

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
		Selected.SelectedModel = "Model 1"
		fmt.Println("Selected Model 1")
	})
	ModelList.AddItem("Model 2", "", 0, func() {
		Selected.SelectedModel = "Model 2"
		fmt.Println("Selected Model 2")
	})

	// Add the model list to the flex layout
	ModelsPane.AddItem(ModelList, 0, 1, true)
}
