package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MealsCategory struct {
	Meals []struct {
		StrCategory string `json:"strCategory,omitempty"`
	} `json:"meals,omitempty"`
}

type MealItems struct {
	Meals []struct {
		StrMeal      string `json:"strMeal,omitempty"`
		StrMealThumb string `json:"strMealThumb,omitempty"`
		IDMeal       string `json:"idMeal,omitempty"`
	} `json:"meals,omitempty"`
}

// API placeholders
const (
	categoriesAPI = "https://www.themealdb.com/api/json/v1/1/list.php?c=list"
	mealsAPI      = "https://www.themealdb.com/api/json/v1/1/filter.php?c="
	recipeAPI     = "https://www.themealdb.com/api/json/v1/1/lookup.php?i="
)

func fetchCategories() tea.Cmd {
	return func() tea.Msg {
		// API call to get the food categories
		r, err := http.Get(categoriesAPI)
		if err != nil {
			return nil
		}
		defer r.Body.Close()

		var data MealsCategory
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return nil
		}

		var categories []Category
		for _, meal := range data.Meals {
			categories = append(categories, Category{
				Name: meal.StrCategory,
			})
		}
		return categories
	}
}

func fetchMeals(category string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s%s", mealsAPI, category)
		r, err := http.Get(url)
		if err != nil {
			return nil
		}
		defer r.Body.Close()

		var data MealItems
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return nil
		}

		var meals []Meal
		for _, meal := range data.Meals {
			meals = append(meals, Meal{
				Name:     meal.StrMeal,
				ImageURL: meal.StrMealThumb,
				ID:       meal.IDMeal,
			})
		}
		return meals
	}
}

func fetchRecipe(mealID string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s%s", recipeAPI, mealID)
		r, err := http.Get(url)
		if err != nil {
			return Recipe{}
		}
		defer r.Body.Close()

		var data Recipe
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return Recipe{}
		}
		return data
	}
}

func (m model) headerView() string {
	title := titleStyle.Render(m.Recipe.StrMeal)
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func formatRecipe(r Recipe) string {
	var ingredients []string
	for i := 1; i <= 10; i++ { // Adjusted to 10 for brevity
		ing := getIngredientField(r, i)
		measure := getMeasureField(r, i)
		if ing != "" {
			ingredients = append(ingredients, fmt.Sprintf("- %s %s", strings.TrimSpace(measure), strings.TrimSpace(ing)))
		}
	}

	return fmt.Sprintf(
		"%s\n\nCategory: %s\n\nIngredients:\n%s\n\nInstructions:\n%s",
		r.StrMeal,
		r.StrCategory,
		strings.Join(ingredients, "\n"),
		r.StrInstructions,
	)
}

func getIngredientField(r Recipe, i int) string {
	switch i {
	case 1:
		return r.StrIngredient1
	case 2:
		return r.StrIngredient2
	case 3:
		return r.StrIngredient3
	case 4:
		return r.StrIngredient4
	case 5:
		return r.StrIngredient5
	case 6:
		return r.StrIngredient6
	case 7:
		return r.StrIngredient7
	case 8:
		return r.StrIngredient8
	case 9:
		return r.StrIngredient9
	case 10:
		return r.StrIngredient10
	default:
		return ""
	}
}

func getMeasureField(r Recipe, i int) string {
	switch i {
	case 1:
		return r.StrMeasure1
	case 2:
		return r.StrMeasure2
	case 3:
		return r.StrMeasure3
	case 4:
		return r.StrMeasure4
	case 5:
		return r.StrMeasure5
	case 6:
		return r.StrMeasure6
	case 7:
		return r.StrMeasure7
	case 8:
		return r.StrMeasure8
	case 9:
		return r.StrMeasure9
	case 10:
		return r.StrMeasure10
	default:
		return ""
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
