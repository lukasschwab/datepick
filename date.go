// Package date provides a shell script interface for picking a date.
//
// The date the user selected will be sent to stdout in ISO-8601 format:
// YYYY-MM-DD.
//
// $ datepick --value 2023-11-28 > date.text
package datepick

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	header      string
	headerStyle lipgloss.Style
	picker      *picker
	quitting    bool
	aborted     bool
	hasTimeout  bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	if m.header != "" {
		header := m.headerStyle.Render(m.header)
		return lipgloss.JoinVertical(lipgloss.Left, header, m.picker.View())
	}

	return m.picker.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			m.aborted = true
			return m, tea.Quit
		case "enter":
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.picker, cmd = m.picker.Update(msg)
	return m, cmd
}
