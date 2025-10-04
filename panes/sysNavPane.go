package panes

import (
	_ "embed"

	"os"
	"os/signal"

	"github.com/Codesmith28/lazyAi/internal/clipboard"
	"github.com/getlantern/systray"
	"github.com/rivo/tview"
)

//go:embed lazyAi.ico
var iconBytes []byte

func ApplySystemNavConfig(app *tview.Application, clipboard *clipboard.Clipboard) {
	onReady := func() {
		systray.SetIcon(iconBytes)
		systray.SetTitle("AI model is now on your clipboard!!")
		systray.SetTooltip("Started!")

		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
		stopMonitoring := systray.AddMenuItem("Pause", "Pause the clipboard monitoring")

		go func() {
			<-mQuit.ClickedCh
			if app != nil {
				app.Stop()
			}
			systray.Quit()
		}()

		go func() {
			for {
				<-stopMonitoring.ClickedCh

				if clipboard.Disable {
					clipboard.Disable = false
					stopMonitoring.SetTitle("Pause")
					stopMonitoring.SetTooltip("Pause the clipboard monitoring")
				} else {
					clipboard.Disable = true
					stopMonitoring.SetTitle("Resume")
					stopMonitoring.SetTooltip("Resume the clipboard monitoring")
				}
			}
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
}
