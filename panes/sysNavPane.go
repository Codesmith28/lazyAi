package panes

import (
	_ "embed"

	"github.com/getlantern/systray"
	"github.com/rivo/tview"
)

//go:embed lazyAi.ico
var iconBytes []byte

func ApplySystemNavConfig(app *tview.Application) {

	onReady := func() {
		systray.SetIcon(iconBytes)
		systray.SetTitle("AI model is now on your clipboard!!")
		systray.SetTooltip("Started!")

		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

		go func() {
			<-mQuit.ClickedCh
			app.Stop()
			systray.Quit()
		}()
	}

	go systray.Run(onReady, onExit)
}

func onExit() {
	// clean up here
}
