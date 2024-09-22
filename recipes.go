package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func loadCategoryItems() ([]Category, error) {
	// API call to get the food categories
	r, err := http.Get("https://www.themealdb.com/api/json/v1/1/list.php?c=list")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var data MealsCategory
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var categories []Category
	for _, meal := range data.Meals {
		categories = append(categories, Category{
			Name: meal.StrCategory,
		})
	}
	return categories, nil
}

func loadMealItems(category string) ([]Meal, error) {
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

	var meals []Meal
	for _, meal := range data.Meals {
		meals = append(meals, Meal{
			Name:     meal.StrMeal,
			ImageURL: meal.StrMealThumb,
			ID:       meal.IDMeal,
		})
	}
	return meals, nil
}

func fetchRecipe(mealID string) (Recipe, error) {
	url := fmt.Sprintf("https://www.themealdb.com/api/json/v1/1/lookup.php?i=%s", mealID)
	r, err := http.Get(url)
	if err != nil {
		return Recipe{}, err
	}
	defer r.Body.Close()

	var data Recipe
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return Recipe{}, err
	}
	return data, nil
}
