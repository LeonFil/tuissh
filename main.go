package main

import (
	"os"
	"os/signal"

	"github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)

var app *tv.Application
var serverView *tv.TextView
var errView *tv.TextView
var screen tcell.Screen

func main() {
	signal.Ignore(os.Interrupt)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	serverTree, err := createServerTree(homeDir + "/.ssh/config")
	if err != nil {
		panic(err)
	}

	app = tv.NewApplication()

	ui := tv.NewFlex()
	serverView = tv.NewTextView() // right-up
	errView = tv.NewTextView()    // right-down
	ui.AddItem(serverTree, 0, 1, true)
	ui.AddItem(tv.NewFlex().SetDirection(tv.FlexRow).
		AddItem(serverView, 0, 1, true).
		AddItem(errView, 0, 1, false),
		0, 2, false)

	// right panel
	serverView.SetBorder(true).SetTitle("server info")
	errView.SetBorder(true).SetTitle("message")

	screen, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	// main loop
	if err := app.SetScreen(screen).SetRoot(ui, true).Run(); err != nil {
		panic(err)
	}
}
