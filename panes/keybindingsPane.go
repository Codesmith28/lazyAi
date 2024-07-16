package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/lazyAi/internal"
)

var KeybindingsPane *tview.TextView

func init() {
	KeybindingsPane = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true)

	KeybindingsPane.SetText(
		`M-1: Prompt | M-2: Input | M-3: Output | M-4: History | M-5: Models | M-S: Save current and create new state | M-O: Create new state | d: Delete history item`,
	).SetTextColor(Theme.TitleCol)
}

// SetupGlobalKeybindings sets up global keybindings to focus on each pane.
func SetupGlobalKeybindings(app *tview.Application) {
	HistoryLocation := internal.GetHistoryLocation()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Modifiers()&tcell.ModAlt != 0 {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case '1':
					app.SetFocus(PromptPane)
					return nil
				case '2':
					app.SetFocus(InputPane)
					return nil
				case '3':
					app.SetFocus(OutputPane)
					return nil
				case '4':
					app.SetFocus(HistoryPane)
					return nil
				case '5':
					app.SetFocus(ModelList)
					return nil
				case 'S', 's':
					saveCurrentState(HistoryLocation)
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
