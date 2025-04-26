.PHONY: build

deps:
	@echo "Installing dependencies for getlantern/systray (required for tray functionality on Linux) and xclip (again, required for linux to interact with clipboard)...."
	sudo apt-get update
	sudo apt-get install -y gcc libgtk-3-dev libayatana-appindicator3-dev xclip

build:
	go build -o lazyAi

run:
	./lazyAi
