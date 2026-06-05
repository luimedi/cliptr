package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UIModel struct {
	handler  *Handler
	cursor   int
	spinner  spinner.Model
	quitting bool
}

func NewUIModel(h *Handler) UIModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F87"))

	return UIModel{
		handler: h,
		spinner: s,
	}
}

func (m UIModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < m.handler.Len()-1 {
				m.cursor++
			}
		case " ", "enter":
			m.handler.ToggleAction(m.cursor)
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF06B7")).
			Bold(true)

	activeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00F5D4")).
			Bold(true)

	inactiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#767676"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5C5C5C")).
			Italic(true).
			MarginTop(1)

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2)
)

func (m UIModel) View() string {
	if m.quitting {
		return "Exiting cliptr...\n"
	}

	var s strings.Builder

	// Header / Title
	s.WriteString(titleStyle.Render("Cliptr - Clipboard Transformer"))
	s.WriteString("\n\n")

	// Handlers list
	s.WriteString(lipgloss.NewStyle().Bold(true).Render("Active Handlers:"))
	s.WriteByte('\n')
	actions := m.handler.GetActionsInfo()
	for i, action := range actions {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("❯ ")
		}

		checkbox := "[ ]"
		var textStyle lipgloss.Style
		if action.IsActive {
			checkbox = activeStyle.Render("[x]")
			textStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFFFF"))
		} else {
			checkbox = inactiveStyle.Render("[ ]")
			textStyle = inactiveStyle
		}

		desc := ""
		if action.Tooltip != "" {
			desc = inactiveStyle.Render(" - " + action.Tooltip)
		}

		s.WriteString(fmt.Sprintf("%s%s %s%s\n", cursor, checkbox, textStyle.Render(action.Title), desc))
	}

	s.WriteString("\n")

	// Spinner & status
	spinnerView := m.spinner.View()
	statusText := lipgloss.NewStyle().Foreground(lipgloss.Color("#5C5C5C")).Render("Listening to clipboard...")
	s.WriteString(fmt.Sprintf("%s %s\n", spinnerView, statusText))

	// Help / footer
	s.WriteString(helpStyle.Render("Space/Enter: toggle • Up/Down: navigate • Q: quit"))

	// Wrap everything in a nice border
	return borderStyle.Render(s.String()) + "\n"
}
