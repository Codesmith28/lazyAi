package panes

import (
	"github.com/Codesmith28/lazyAi/api"
	"github.com/Codesmith28/lazyAi/internal"
	"github.com/Codesmith28/lazyAi/internal/clipboard"
	"github.com/charmbracelet/glamour"
	"github.com/gdamore/tcell/v2"
	"github.com/getlantern/systray"
	"github.com/rivo/tview"
)

var (
	OutputPane   *tview.TextView
	OutputText   = &internal.Output{}
	HelpCommands string
)

func init() {

	var commands string = "\n## Commands:\n\n" +
		"- run the application: `lazyAi`\n\n" +
		"- run the application in detached mode: `lazyAi -d`\n" +
		"- run the application with a default prompt: `lazyAi -p \"your prompt here\"`\n"

	helpcmd := true
	HelpCommands = markdownToTview(commands, &helpcmd)

	OutputPane = tview.NewTextView()

	OutputText = &internal.Output{
		OutputString: "",
	}

	OutputPane.SetText(OutputText.OutputString)

	OutputPane.
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetScrollable(true).
		SetBorder(true)

	OutputPane.
		SetTitle(" Output ").
		SetBorderPadding(1, 1, 2, 2)

	OutputPane.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyUp {
			row, _ := OutputPane.GetScrollOffset()
			if row > 0 {
				OutputPane.ScrollTo(row-1, 0)
			}
		} else if event.Key() == tcell.KeyDown {
			row, _ := OutputPane.GetScrollOffset()
			OutputPane.ScrollTo(row+1, 0)
		}
		return event
	})
}

func HandlePromptChange(
	query *internal.Query,
	clipboard *clipboard.Clipboard,
	app *tview.Application,
) {
	content, err := api.SendPrompt(
		query.PromptString,
		query.SelectedModel,
		query.InputString,
	)
	if err != nil {
		panic(err)
	}

	if app != nil {
		styledContent := markdownToTview(content, nil)
		app.QueueUpdateDraw(func() {
			OutputPane.SetText(styledContent)
		})
	}

	clipboard.Mu.Lock()
	OutputText.OutputString = content
	clipboard.OutputText = content
	err = clipboard.SetClipboardText(content)

	if err != nil {
		panic(err)
	}

	clipboard.Mu.Unlock()

	systray.SetTooltip("Ready!!")
}

func markdownToTview(md string, helpCommand *bool) string {

	formattedOutput, err := glamour.Render(md, "dark")

	if err != nil {
		panic(err)
	}

	if helpCommand != nil && *helpCommand {
		return formattedOutput
	}

	finalOutput := tview.TranslateANSI(formattedOutput)

	return finalOutput
}
