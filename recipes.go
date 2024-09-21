package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/bubbles/list"
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

func loadFoodCategories() ([]list.Item, error) {
	// API call to get the food categories
	r, err := http.Get("https://www.themealdb.com/api/json/v1/1/list.php?c=list")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// Variable must be of type []map[string]interface{}
	var data MealsCategory
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var categories []list.Item
	for _, meal := range data.Meals {
		categories = append(categories, RecipeCategory{
			Name: meal.StrCategory,
		})
	}
	return categories, nil
}

func loadMealItems(category string) ([]list.Item, error) {
	url := fmt.Sprintf("https://www.themealdb.com/api/json/v1/1/filter.php?c=%s", category)
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var data MealItems
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var meals []list.Item
	for _, meal := range data.Meals {
		meals = append(meals, Meal{
			Name:     meal.StrMeal,
			ImageURL: meal.StrMealThumb,
			ID:       meal.IDMeal,
		})
	}
	return meals, nil
}
