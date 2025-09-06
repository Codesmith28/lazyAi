package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/lazyAi/internal"
)

var (
	ModelList = tview.NewList()
	Selected  *internal.Model
)

var AvailableModels = map[string]*internal.Model{
	"Gemini 2.0 Flash":      {SelectedModel: "gemini-2.0-flash"},
	"Gemini 2.0 Flash-Lite": {SelectedModel: "gemini-2.0-flash-lite"},
	"Gemini Pro 1.5":        {SelectedModel: "gemini-1.5-pro"},
	"Gemini Flash":          {SelectedModel: "gemini-1.5-flash"},
}

func init() {
	// Configure the model list
	ModelList.ShowSecondaryText(false).SetTitle(" Models ").SetBorder(true)
	Selected = &internal.Model{}

	SelectModel(AvailableModels["Gemini 2.0 Flash"].SelectedModel)

	// Add models to the list
	for key, model := range AvailableModels {
		currentModel := model
		ModelList.AddItem(key, "", 0, func() {
			SelectModel(currentModel.SelectedModel)
		})
	}

	// Enable mouse support
	ModelList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			currentItem := ModelList.GetCurrentItem()
			mainText, _ := ModelList.GetItemText(currentItem)
			SelectModel(AvailableModels[mainText].SelectedModel)
		}
		return event
	})
}
