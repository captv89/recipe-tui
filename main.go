package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

/*
type item struct {
	title, author, desc, content string
}

func (i item) Title() string       { return i.title }
func (i item) Author() string      { return i.author }
func (i item) Description() string { return i.desc }
func (i item) Content() string     { return i.content }
func (i item) FilterValue() string { return i.title }

type model struct {
	list        list.Model
	viewport    viewport.Model
	showDetails bool
	currentItem item
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEscape:
			// Go back to the list
			m.showDetails = false
			// Do nothing
			return m, nil
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.list.SelectedItem() != nil {
				m.currentItem = m.list.SelectedItem().(item)
			}
			m.showDetails = !m.showDetails
			if m.showDetails && m.list.SelectedItem() != nil {
				selectedItem := m.list.SelectedItem().(item)
				m.viewport.SetContent(fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\nContent: %s",
					selectedItem.Title(),
					selectedItem.Author(),
					selectedItem.Description(),
					selectedItem.Content(),
				))
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.viewport.Width = msg.Width - h
		m.viewport.Height = msg.Height - v - 5
		if m.showDetails && m.list.SelectedItem() != nil {
			selectedItem := m.list.SelectedItem().(item)
			m.viewport.SetContent(fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\nContent: %s",
				selectedItem.Title(),
				selectedItem.Author(),
				selectedItem.Description(),
				selectedItem.Content(),
			))
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	if m.showDetails {
		m.viewport, cmd = m.viewport.Update(msg)
	}
	return m, cmd
}
*/

func main() {
	// Logging
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("Error logging to file:", err)
		os.Exit(1)
	}
	defer f.Close()

	// Create the program
	p := tea.NewProgram(MainModel{}, tea.WithAltScreen())

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
