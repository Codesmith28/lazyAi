package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal/clipboard"
	core "github.com/Codesmith28/cheatScript/internal/queue"
	"github.com/Codesmith28/cheatScript/panes"
)

var (
	historyPane     = panes.HistoryPane
	modelsPane      = panes.ModelsPane
	outputPane      = panes.OutputPane
	promptPane      = panes.PromptPane
	keybindingsPane = panes.KeybindingsPane
	inputPane       = panes.InputPane

	modelList = panes.ModelList
)

var once sync.Once

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := tview.NewApplication()

	clipboard := clipboard.NewClipboard()
	go clipboard.StartMonitoring()
	clipboard.Clear()

	// Group 2: historyPane and promptPane in a vertical layout
	group2 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(historyPane, 0, 1, false).
		AddItem(promptPane, 0, 1, false)

	// Group 4: inputPane and modelsPane in a horizontal layout
	group4 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(inputPane, 0, 1, false).
		AddItem(modelsPane, 0, 1, true) // Set focusable to true

	// intialize the queue

	// calculate how much time it takes to connect to the queue

	queue := core.NewQueue()
	var err error = queue.Connect()

	if err != nil {
		checkNilErr(err)
	}

	go func() {
		for {
			text, err := clipboard.GetClipboardText()

			if err != nil {
				fmt.Println("Error getting clipboard text:", err)
				checkNilErr(err)
			}

			app.QueueUpdateDraw(func() {
				inputPane.SetText("Prompt: " + text)
			})

			err = queue.Publish(text)

			if err != nil {
				fmt.Println("Error publishing message:", err)
				checkNilErr(err)
			}

			once.Do(func() {
				go queue.Consume(clipboard, func(msg string) {
					app.QueueUpdateDraw(func() {
						outputPane.SetText(msg)
					})

					clipboard.LastText = msg
					err := clipboard.SetClipboardText(msg)

					if err != nil {
						checkNilErr(err)
					}
				})
			})

			time.Sleep(1 * time.Second)
		}
	}()

	// Group 3: Group 4 and outputPane in a vertical layout
	group3 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group4, 0, 1, false).
		AddItem(outputPane, 0, 2, false)

	// Group 1: Group 2 and Group 3 in a horizontal layout
	group1 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(group2, 0, 1, false).
		AddItem(group3, 0, 2, false)

	// Main Flex: Group 1 and keybindingsPane in a vertical layout
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group1, 0, 4, false).
		AddItem(keybindingsPane, 3, 1, false)

	// Set up global keybindings to focus on each pane
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				app.SetFocus(inputPane)
			case '2':
				app.SetFocus(modelList)
			case '3':
				app.SetFocus(outputPane)
			case '4':
				app.SetFocus(historyPane)
			case '5':
				app.SetFocus(promptPane)
			case '?':
				app.SetFocus(keybindingsPane)
			}
		}
		return event
	})

	// Set the default focus to inputPane
	app.SetFocus(inputPane)

	// Set up the application root
	err = app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
