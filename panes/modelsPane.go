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

var availableModels = map[string]*internal.Model{
	"Gemini 2.5 Pro":        {SelectedModel: "gemini-2.5-pro"},
	"Gemini 2.5 Flash":      {SelectedModel: "gemini-2.5-flash"},
	"Gemini 2.5 Flash-Lite": {SelectedModel: "gemini-2.5-flash-lite"},
}

func init() {
	// Configure the model list
	ModelList.ShowSecondaryText(false).SetTitle(" Models ").SetBorder(true)
	Selected = &internal.Model{}

	SelectModel(availableModels["Gemini 2.5 Flash"].SelectedModel)

	// Add models to the list
	for key, model := range availableModels {
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
			SelectModel(availableModels[mainText].SelectedModel)
		}
		return event
	})
}
