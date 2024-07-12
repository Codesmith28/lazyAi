package panes

import (
	"strings"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/getlantern/systray"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/rivo/tview"
)

var (
	OutputPane *tview.TextView
	OutputText = &internal.Output{}
)

func init() {
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

	styledContent := markdownToTview(content)

	if app != nil {
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

func markdownToTview(md string) string {
	// Parse the markdown
	extensions := parser.NoExtensions
	p := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML([]byte(md), p, nil)

	// Convert HTML to tview format
	return htmlToTview(string(html))
}

func htmlToTview(html string) string {
	html = strings.ReplaceAll(html, "<h1>", "[#FFFF00::b]")
	html = strings.ReplaceAll(html, "</h1>", "[-::-]")
	html = strings.ReplaceAll(html, "<h2>", "[#00FF00::b]")
	html = strings.ReplaceAll(html, "</h2>", "[-::-]")
	html = strings.ReplaceAll(html, "<p>", "")
	html = strings.ReplaceAll(html, "</p>", "")
	html = strings.ReplaceAll(html, "<strong>", "[::b]")
	html = strings.ReplaceAll(html, "</strong>", "[::-]")
	html = strings.ReplaceAll(html, "<em>", "[::i]")
	html = strings.ReplaceAll(html, "</em>", "[::-]")
	html = strings.ReplaceAll(html, "<code>", "[#4CAF50]")
	html = strings.ReplaceAll(html, "</code>", "[-]")
	html = strings.ReplaceAll(html, "<pre>", "[#4CAF50]")
	html = strings.ReplaceAll(html, "</pre>", "[-]")
	html = strings.ReplaceAll(html, "<ul>", "")
	html = strings.ReplaceAll(html, "</ul>", "")
	html = strings.ReplaceAll(html, "<li>", "• ")
	html = strings.ReplaceAll(html, "</li>", "")
	html = strings.ReplaceAll(html, "<ol>", "")
	html = strings.ReplaceAll(html, "</ol>", "")
	html = strings.ReplaceAll(html, "&rsquo;", "'")

	return html
}
