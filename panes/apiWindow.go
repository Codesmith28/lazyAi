package panes

import (
	"github.com/rivo/tview"
)

func CreateCredentialModal(
	app *tview.Application,
	onSubmit func(password string),
) tview.Primitive {
	apiKey := ""

	form := tview.NewForm().
		AddPasswordField("Api key", "", 43, '*', func(text string) {
			apiKey = text
		}).
		AddButton("Submit", func() {
			onSubmit(apiKey)
		}).
		AddButton("Cancel", func() {
			app.Stop()
		}).
		SetButtonsAlign(tview.AlignCenter)

	// Set form attributes
	form.SetBorder(true).
		SetTitle(" API Missing ").
		SetTitleAlign(tview.AlignCenter)

	// Create a flex to center the form horizontally
	formContainer := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(form, 55, 1, true).
		AddItem(nil, 0, 1, false)

	// Create a flex to center the form vertically
	centeredFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 2, false).
		AddItem(formContainer, 0, 1, true).
		AddItem(nil, 0, 2, false)

	return tview.NewPages().
		AddPage("form", centeredFlex, true, true)
}
