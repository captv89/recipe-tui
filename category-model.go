package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type CategoryModel struct {
	List list.Model
}

type Category struct {
	Name string
}

func (c Category) Title() string       { return c.Name }
func (c Category) Description() string { return c.Name }
func (c Category) FilterValue() string { return c.Name }

// Init
func (m CategoryModel) Init() tea.Cmd {
	// Get the category list
	categories, _ := loadFoodCategories()

	// Create the list
	l := list.New(categories, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Categories"

	m.List = l

	return nil
}

// Update
func (m CategoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View
func (m CategoryModel) View() string {
	return listStyle.Render(m.List.View())
}
