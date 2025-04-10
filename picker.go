package datepick

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fxtlabs/date"
)

// Default styles.
var (
	weekdayStyle           = lipgloss.NewStyle().Faint(true)
	defaultCursorTextStyle = lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("201"))
)

type interval int

const (
	day   interval = 0
	month interval = 1
	year  interval = 2
)

// incr i in direction d; bodge mod-3 indexing.
func (i interval) incr(d direction) interval {
	mod := (int(i) + int(d)) % 3
	if mod < 0 {
		return year
	}
	return interval(mod)
}

type direction int

const (
	forward  direction = 1
	backward direction = -1
)

// picker implements tea.Model for a date.Date.
type picker struct {
	date.Date
	focus interval

	promptStyle lipgloss.Style
	prompt      string
	help        help.Model

	cursorTextStyle lipgloss.Style
}

func basePicker() *picker {
	return &picker{
		Date:            date.Today(),
		focus:           day,
		prompt:          "> ",
		help:            help.New(),
		cursorTextStyle: defaultCursorTextStyle,
	}
}

func (p *picker) formatDate() string {
	raw := p.Date.Format("02 Jan 2006")
	parts := strings.Split(raw, " ")
	parts[int(p.focus)] = p.cursorTextStyle.Render(parts[int(p.focus)])
	return strings.Join(parts, " ") + " " + p.formatWeekday()
}

func (p *picker) formatWeekday() string {
	shortName := p.Date.Weekday().String()[:3]
	return weekdayStyle.Render(shortName)
}

func (p *picker) incr(d direction) {
	switch p.focus {
	case day:
		p.Date = p.Date.AddDate(0, 0, int(d))
	case month:
		p.Date = p.Date.AddDate(0, int(d), 0)
	case year:
		p.Date = p.Date.AddDate(int(d), 0, 0)
	}
}

// Init implements tea.Model.
func (p *picker) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (p *picker) Update(msg tea.Msg) (*picker, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		p.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		// "up"/"down" increment/decrement the focused component, respectively
		case key.Matches(msg, keys.Up):
			p.incr(forward)
		case key.Matches(msg, keys.Down):
			p.incr(backward)

		// increment/decrement by week
		case key.Matches(msg, keys.UpWeek):
			p.incr(direction(forward * 7))
		case key.Matches(msg, keys.DownWeek):
			p.incr(direction(backward * 7))

		// "left"/"right" cycle the focused component
		case key.Matches(msg, keys.Left):
			p.focus = p.focus.incr(backward)
		case key.Matches(msg, keys.Right):
			p.focus = p.focus.incr(forward)

		case key.Matches(msg, keys.Help):
			p.help.ShowAll = !p.help.ShowAll
		}
	}
	return p, nil
}

// View implements tea.Model.
func (p *picker) View() string {
	return p.promptStyle.Render(p.prompt) + p.formatDate() + "\n" + p.help.View(keys)
}

// Value of p.
func (p *picker) Value() date.Date {
	return p.Date
}

type keyMap struct {
	// Manipulate the date.
	Up       key.Binding
	Down     key.Binding
	UpWeek   key.Binding
	DownWeek key.Binding
	// Cycle the focused component.
	Left  key.Binding
	Right key.Binding
	// Toggle help.
	Help key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "incr"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "decr"),
	),
	UpWeek: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "next week"),
	),
	DownWeek: key.NewBinding(
		key.WithKeys("W"),
		key.WithHelp("W", "prev week"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h", "shift+tab"),
		key.WithHelp("←/h", "prev"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l", "tab"),
		key.WithHelp("→/l", "next"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Right, k.Help}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.UpWeek, k.DownWeek}, // first column
		{k.Help},                             // second column
	}
}
