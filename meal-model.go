package main

import tea "github.com/charmbracelet/bubbletea"

type MealModel struct {
	ID       string
	Name     string
	ImageURL string
}

func (m MealModel) Title() string       { return m.Name }
func (m MealModel) Description() string { return m.Name }
func (m MealModel) FilterValue() string { return m.Name }

// Init
func (m MealModel) Init() tea.Cmd {
	return nil
}

// Update
func (m MealModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View
func (m MealModel) View() string {
	return ""
}
