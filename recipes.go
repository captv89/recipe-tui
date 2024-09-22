package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
