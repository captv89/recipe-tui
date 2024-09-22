package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Meal struct {
	ID       string
	Name     string
	ImageURL string
}

func (m Meal) Title() string       { return m.Name }
func (m Meal) Description() string { return m.Name }
func (m Meal) FilterValue() string { return m.Name }

type MealModel struct {
	List list.Model
}

func (m MealModel) initializeList(category string) list.Model {
	meals, err := loadMealItems(category)
	if err != nil {
		return list.Model{}
	}
	// Convert to list items
	items := make([]list.Item, len(meals))
	for i, meal := range meals {
		items[i] = meal
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Meals"
	return l
}

// Init
func (m MealModel) Init() tea.Cmd {
	return nil
}

// Update
func (m MealModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// View
func (m MealModel) View() string {
	return listStyle.Render(m.List.View())
}
