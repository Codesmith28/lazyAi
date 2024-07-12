package panes

import (
	_ "embed"

	"os"
	"os/signal"

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
			if app != nil {
				app.Stop()
			}
			systray.Quit()
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				if app != nil {
					app.Stop()
				}
				systray.Quit()
				os.Exit(0)
			}
		}()
	}

	go systray.Run(onReady, onExit)
}

func onExit() {
	// clean up here
}
