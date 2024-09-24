package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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

type MealRecipetype struct {
	Meals []Recipe `json:"meals"`
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

		var data MealRecipetype
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return Recipe{}
		}

		// Reflect the data into a Recipe struct
		recipe := data.Meals[0]
		log.Printf("Recipe: %+v", recipe)

		return recipe
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		"\nCategory: %s\n\nArea: %s\n\nIngredients:\n%s\n\nInstructions:\n%s",
		r.StrCategory,
		r.StrArea,
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
	case 11:
		return r.StrIngredient11
	case 12:
		return r.StrIngredient12
	case 13:
		return r.StrIngredient13
	case 14:
		return r.StrIngredient14
	case 15:
		return r.StrIngredient15
	case 16:
		return r.StrIngredient16
	case 17:
		return r.StrIngredient17
	case 18:
		return r.StrIngredient18
	case 19:
		return r.StrIngredient19
	case 20:
		return r.StrIngredient20
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
	case 11:
		return r.StrMeasure11
	case 12:
		return r.StrMeasure12
	case 13:
		return r.StrMeasure13
	case 14:
		return r.StrMeasure14
	case 15:
		return r.StrMeasure15
	case 16:
		return r.StrMeasure16
	case 17:
		return r.StrMeasure17
	case 18:
		return r.StrMeasure18
	case 19:
		return r.StrMeasure19
	case 20:
		return r.StrMeasure20
	default:
		return ""
	}
}
