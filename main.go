package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Types for categories, meals, and recipes
type Category struct {
	Name string `json:"name"`
}

type Meal struct {
	ID       string
	Name     string
	ImageURL string
}

type Recipe struct {
	IDMeal                      string `json:"idMeal"`
	StrMeal                     string `json:"strMeal"`
	StrDrinkAlternate           string `json:"strDrinkAlternate"`
	StrCategory                 string `json:"strCategory"`
	StrArea                     string `json:"strArea"`
	StrInstructions             string `json:"strInstructions"`
	StrMealThumb                string `json:"strMealThumb"`
	StrTags                     string `json:"strTags"`
	StrYoutube                  string `json:"strYoutube"`
	StrIngredient1              string `json:"strIngredient1"`
	StrIngredient2              string `json:"strIngredient2"`
	StrIngredient3              string `json:"strIngredient3"`
	StrIngredient4              string `json:"strIngredient4"`
	StrIngredient5              string `json:"strIngredient5"`
	StrIngredient6              string `json:"strIngredient6"`
	StrIngredient7              string `json:"strIngredient7"`
	StrIngredient8              string `json:"strIngredient8"`
	StrIngredient9              string `json:"strIngredient9"`
	StrIngredient10             string `json:"strIngredient10"`
	StrIngredient11             string `json:"strIngredient11"`
	StrIngredient12             string `json:"strIngredient12"`
	StrIngredient13             string `json:"strIngredient13"`
	StrIngredient14             string `json:"strIngredient14"`
	StrIngredient15             string `json:"strIngredient15"`
	StrIngredient16             string `json:"strIngredient16"`
	StrIngredient17             string `json:"strIngredient17"`
	StrIngredient18             string `json:"strIngredient18"`
	StrIngredient19             string `json:"strIngredient19"`
	StrIngredient20             string `json:"strIngredient20"`
	StrMeasure1                 string `json:"strMeasure1"`
	StrMeasure2                 string `json:"strMeasure2"`
	StrMeasure3                 string `json:"strMeasure3"`
	StrMeasure4                 string `json:"strMeasure4"`
	StrMeasure5                 string `json:"strMeasure5"`
	StrMeasure6                 string `json:"strMeasure6"`
	StrMeasure7                 string `json:"strMeasure7"`
	StrMeasure8                 string `json:"strMeasure8"`
	StrMeasure9                 string `json:"strMeasure9"`
	StrMeasure10                string `json:"strMeasure10"`
	StrMeasure11                string `json:"strMeasure11"`
	StrMeasure12                string `json:"strMeasure12"`
	StrMeasure13                string `json:"strMeasure13"`
	StrMeasure14                string `json:"strMeasure14"`
	StrMeasure15                string `json:"strMeasure15"`
	StrMeasure16                string `json:"strMeasure16"`
	StrMeasure17                string `json:"strMeasure17"`
	StrMeasure18                string `json:"strMeasure18"`
	StrMeasure19                string `json:"strMeasure19"`
	StrMeasure20                string `json:"strMeasure20"`
	StrSource                   string `json:"strSource"`
	StrImageSource              string `json:"strImageSource"`
	StrCreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
	DateModified                string `json:"dateModified"`
}

// App states
type state int

const (
	categoryList state = iota
	mealList
	recipeDetail
)

// API placeholders
const (
	categoriesAPI = "https://www.themealdb.com/api/json/v1/1/list.php?c=list"
	mealsAPI      = "https://www.themealdb.com/api/json/v1/1/filter.php?c="
	recipeAPI     = "https://www.themealdb.com/api/json/v1/1/lookup.php?i="
)

// Main model struct
type model struct {
	state         state
	categories    []Category
	meals         []Meal
	recipe        Recipe
	selectedIndex int
	list          list.Model
}

// Initialize app with categories
func (m model) Init() tea.Cmd {
	return fetchCategories()
}

// // Fetch categories from API
// func fetchCategories() tea.Cmd {
// 	return func() tea.Msg {
// 		resp, err := http.Get(categoriesAPI)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer resp.Body.Close()
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var categories []Category
// 		json.Unmarshal(body, &categories)
// 		return categories
// 	}
// }

// // Fetch meals based on selected category
// func fetchMeals(categoryID string) tea.Cmd {
// 	return func() tea.Msg {
// 		resp, err := http.Get(mealsAPI + categoryID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer resp.Body.Close()
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var meals []Meal
// 		json.Unmarshal(body, &meals)
// 		return meals
// 	}
// }

// // Fetch recipe based on selected meal
// func fetchRecipe(mealID string) tea.Cmd {
// 	return func() tea.Msg {
// 		resp, err := http.Get(recipeAPI + mealID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer resp.Body.Close()
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		var recipe Recipe
// 		json.Unmarshal(body, &recipe)
// 		return recipe
// 	}
// }

// Update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "enter":
			switch m.state {
			case categoryList:
				// Move to meal list
				selectedCategory := m.categories[m.selectedIndex]
				m.state = mealList
				return m, fetchMeals(selectedCategory.Name)

			case mealList:
				// Move to recipe detail
				selectedMeal := m.meals[m.selectedIndex]
				m.state = recipeDetail
				return m, fetchRecipe(selectedMeal.ID)

			case recipeDetail:
				// No further action
				return m, nil
			}

		case "up":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case "down":
			if m.state == categoryList && m.selectedIndex < len(m.categories)-1 {
				m.selectedIndex++
			}
			if m.state == mealList && m.selectedIndex < len(m.meals)-1 {
				m.selectedIndex++
			}

		case "backspace":
			// Handle going back
			if m.state == mealList {
				m.state = categoryList
			} else if m.state == recipeDetail {
				m.state = mealList
			}
		}

	case []Category:
		m.categories = msg
		m.state = categoryList
		m.selectedIndex = 0
		return m, nil

	case []Meal:
		m.meals = msg
		m.state = mealList
		m.selectedIndex = 0
		return m, nil

	case Recipe:
		m.recipe = msg
		m.state = recipeDetail
		return m, nil
	}

	return m, nil
}

// View function
func (m model) View() string {
	switch m.state {
	case categoryList:
		return m.viewCategories()
	case mealList:
		return m.viewMeals()
	case recipeDetail:
		return m.viewRecipe()
	}
	return ""
}

// View categories list
func (m model) viewCategories() string {
	s := "Select a category:\n\n"
	for i, category := range m.categories {
		cursor := " " // no cursor
		if m.selectedIndex == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf("%s %s\n", cursor, category.Name)
	}
	return s + "\nPress q to quit."
}

// View meals list
func (m model) viewMeals() string {
	s := "Select a meal:\n\n"
	for i, meal := range m.meals {
		cursor := " " // no cursor
		if m.selectedIndex == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf("%s %s\n", cursor, meal.Name)
	}
	return s + "\nPress backspace to go back."
}

// View recipe details
func (m model) viewRecipe() string {
	return fmt.Sprintf("Recipe details:\n\n%s\n\nPress backspace to go back.", m.recipe.StrMeal)
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
