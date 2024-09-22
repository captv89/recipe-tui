package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// States
type State int

const (
	CategoriesState State = iota
	MealsState
	RecipeState
	LoadingState
)

type MainModel struct {
	State    State
	Category CategoryModel
	Meal     MealModel
	Recipe   RecipeModel
	Spinner  spinner.Model
	Error    error
	Current  string // Current category, meal, or recipe
}

func (m MainModel) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Handle updates from keys
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}

	switch m.State {
	case CategoriesState:
		var cmdCategory tea.Cmd
		updatedModel, cmdCategory := m.Category.Update(msg)
		var ok bool
		m.Category, ok = updatedModel.(CategoryModel)
		if !ok {
			m.Error = fmt.Errorf("failed to update category model")
			return m, cmdCategory
		}
		return m, cmdCategory

	case MealsState:
		var cmdMeal tea.Cmd
		updatedModel, cmdMeal := m.Meal.Update(msg)
		m.Meal, _ = updatedModel.(MealModel)
		return m, cmdMeal

	case RecipeState:
		var cmdRecipe tea.Cmd
		updatedModel, cmdRecipe := m.Recipe.Update(msg)
		m.Recipe, _ = updatedModel.(RecipeModel)
		return m, cmdRecipe

	case LoadingState:
		return m, m.Spinner.Tick
	}

	return m, cmd
}

// View
func (m MainModel) View() string {
	if m.Error != nil {
		return fmt.Sprintf("Error: %v", m.Error)
	}
	switch m.State {
	case CategoriesState:
		return m.Category.View()
	case MealsState:
		return m.Meal.View()
	case RecipeState:
		return m.Recipe.View()
	case LoadingState:
		return m.Spinner.View() + "Loading..."
	default:
		return "Unknown State"
	}
}
