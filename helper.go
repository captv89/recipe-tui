package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }

func loadCategories() tea.Cmd {
	return func() tea.Msg {
		categories, err := loadCategoryItems()
		if err != nil {
			return errMsg{err}
		}
		return categories
	}
}

func loadMeals(category string) tea.Cmd {
	return func() tea.Msg {
		meals, err := loadMealItems(category)
		if err != nil {
			return errMsg{err}
		}
		return meals
	}
}

func loadRecipe(mealID string) tea.Cmd {
	return func() tea.Msg {
		recipe, err := fetchRecipe(mealID)
		if err != nil {
			return errMsg{err}
		}
		return recipe
	}
}

func handleCategorySelection(m *MainModel, msg tea.Msg) tea.Cmd {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			selected := m.Category.List.SelectedItem()
			if selected != nil {
				category := selected.(Category).Name
				m.State = LoadingState
				return tea.Batch(m.Spinner.Tick, loadMeals(category))
			}
		case "ctrl+c", "q":
			return tea.Quit
		}
	}
	return nil
}

func handleMealSelection(m *MainModel, msg tea.Msg) tea.Cmd {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			selected := m.Meal.List.SelectedItem()
			if selected != nil {
				meal := selected.(Meal)
				m.State = LoadingState
				return tea.Batch(m.Spinner.Tick, loadRecipe(meal.ID))
			}
		case "esc":
			m.State = CategoriesState
			return nil
		case "ctrl+c", "q":
			return tea.Quit
		}
	}
	return nil
}
