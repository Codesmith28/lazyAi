package panes

import "github.com/rivo/tview"

var InputPane = tview.NewForm()

// Create a new InputField.
var inputField = tview.NewInputField()

// Set the text view to wrap long lines.
var TextView = tview.NewTextView().SetWrap(true)
