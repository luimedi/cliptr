package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var Version string

func main() {
	handler := NewHandler()
	handler.AddAction("Trim", "Remove leading and trailing whitespace", Trim)
	handler.AddAction("Capitalize", "Capitalize words if text is fully uppercase", CapitalizeUppercase)

	p := tea.NewProgram(NewUIModel(handler))
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running cliptr: %v\n", err)
		os.Exit(1)
	}
}
