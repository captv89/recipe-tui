package main

import tea "github.com/charmbracelet/bubbletea"

type CategoryModel struct {
	Name string
}

func (rc CategoryModel) Title() string       { return rc.Name }
func (rc CategoryModel) Description() string { return rc.Name }
func (rc CategoryModel) FilterValue() string { return rc.Name }

// Init
func (m CategoryModel) Init() tea.Cmd {
	return nil
}

// Update
func (m CategoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View
func (m CategoryModel) View() string {
	return m.Name
}
