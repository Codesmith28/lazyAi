package panes

import (
	"github.com/rivo/tview"
)

var (
	// ModelsPane is a flex layout containing the title and list of models.
	ModelsPane = tview.NewFlex()
	modelList  = tview.NewList()
)

func init() {
	// Configure the model list
	modelList.ShowSecondaryText(false).SetTitle("Models").SetBorder(true)
	modelList.AddItem("Model 1", "", 0, nil)
	modelList.AddItem("Model 2", "", 0, nil)

	// Add the model list to the flex layout
	ModelsPane.AddItem(modelList, 0, 1, true)
}
