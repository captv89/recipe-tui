package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Logging
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("Error logging to file:", err)
		os.Exit(1)
	}
	defer f.Close()

	// Set the spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("202"))

	// Load Empty Category Model
	emptyCategoryModel := CategoryModel{}

	// Load Empty Meal Model
	emptyMealModel := MealModel{}

	// Load Empty Recipe Model
	emptyRecipeModel := RecipeModel{}

	// Initialise the main model
	mainModel := MainModel{
		Spinner:  s,
		Category: emptyCategoryModel,
		Meal:     emptyMealModel,
		Recipe:   emptyRecipeModel,
	}

	// Create the program
	p := tea.NewProgram(mainModel, tea.WithAltScreen())

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
