package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func showSummary() string   { return "showSummary" }
func showProcesses() string { return "showProcesses" }
func showDownloads() string { return "showDownloads" }
func showSettings() string  { return "showSettings" }

type Model struct {
	tabs         []string // items on the to-do list
	cursor       int      // which to-do list item our cursor is pointing at
	page         int
	pageMappings []func() string
}

func NewModel() Model {
	return Model{
		// Our to-do list is a grocery list
		tabs: []string{"Summary", "Processes", "Downloads", "Settings"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		pageMappings: []func() string{
			showSummary,
			showProcesses,
			showDownloads,
			showSettings,
		},
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "left", "a":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.tabs) - 1
			}

		// The "down" and "j" keys move the cursor down
		case "right", "d":
			if m.cursor < len(m.tabs)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			m.page = m.cursor
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	screen := ""

	tabRow := ""
	for i, tab := range m.tabs {
		active := i == m.page
		selected := i == m.cursor

		sTab := tab
		if active {
			sTab = fmt.Sprintf("(%s)", sTab)
		} else {
			sTab = fmt.Sprintf(" %s ", sTab)
		}

		if selected {
			sTab = fmt.Sprintf("<%s>", sTab)
		} else {
			sTab = fmt.Sprintf(" %s ", sTab)
		}

		tabRow += sTab + " | "
	}
	screen += tabRow + "\n"

	screen += m.pageMappings[m.page]() + "\n"

	screen += "[Quit]\n"

	return screen
}
