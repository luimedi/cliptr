package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.design/x/clipboard"
)

var Version string

func main() {
	// Initialize clipboard
	if err := clipboard.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to initialize clipboard: %v\n", err)
		os.Exit(1)
	}

	handler := NewHandler()
	handler.AddAction("Trim", "Remove leading and trailing whitespace", Trim)
	handler.AddAction("Capitalize", "Capitalize words if text is fully uppercase", CapitalizeUppercase)

	// Listen to clipboard changes in the background
	go handler.Listen()

	p := tea.NewProgram(NewUIModel(handler))
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running cliptr: %v\n", err)
		os.Exit(1)
	}
}
