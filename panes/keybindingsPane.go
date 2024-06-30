package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var KeybindingsPane *tview.TextView

func init() {
	KeybindingsPane = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true)

	KeybindingsPane.SetText(
		`M-1: Input | M-2: Models | M-3: Output | M-4: History | M-5: Prompt | M-?: Keybindings | M-S: Save current state | M-O: Create new state`,
	)
}

// SetupGlobalKeybindings sets up global keybindings to focus on each pane.
func SetupGlobalKeybindings(app *tview.Application) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Modifiers()&tcell.ModAlt != 0 {
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
				case 'S', 's':
					saveCurrentState()
					return nil
				case 'O', 'o':
					createNewState()
					return nil
				}
			}
		}
		return event
	})
}
