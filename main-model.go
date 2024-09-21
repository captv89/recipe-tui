package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// States
type State int

const (
	CategoriesState State = iota
	MealsState
	RecipeState
)

type MainModel struct {
	State State
	Error error
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View
func (m MainModel) View() string {
	return "Main"
}
