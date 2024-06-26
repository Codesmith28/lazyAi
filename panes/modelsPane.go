package panes

import (
	"fmt"

	"github.com/rivo/tview"
)

var (
	// ModelsPane is a flex layout containing the title and list of models.
	ModelsPane    = tview.NewFlex()
	ModelList     = tview.NewList() // Export ModelList to access it
	selectedModel string
)

func init() {
	// Configure the model list
	ModelList.ShowSecondaryText(false).SetTitle("Models").SetBorder(true)

	// Add models to the list
	ModelList.AddItem("Model 1", "", 0, func() {
		selectedModel = "Model 1"
		fmt.Println("Selected Model 1")
	})
	ModelList.AddItem("Model 2", "", 0, func() {
		selectedModel = "Model 2"
		fmt.Println("Selected Model 2")
	})

	// Add the model list to the flex layout
	ModelsPane.AddItem(ModelList, 0, 1, true)
}
