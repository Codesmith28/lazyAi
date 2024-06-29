package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var KeybindingsPane = tview.NewBox().SetBorder(true).SetTitle(" Keybindings ")

// SetupGlobalKeybindings sets up global keybindings to focus on each pane.
func SetupGlobalKeybindings(app *tview.Application) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				app.SetFocus(InputPane)
				return nil
			case '2':
				app.SetFocus(ModelsPane)
				return nil
			case '3':
				app.SetFocus(OutputPane)
				return nil
			case '4':
				app.SetFocus(HistoryPane)
				return nil
			case '5':
				app.SetFocus(PromptPane)
				return nil
			case '?':
				app.SetFocus(KeybindingsPane)
				return nil
			case 's':
				saveCurrentState()
			case 'o':
				createNewState()
			}
		}
		return event
	})
}
