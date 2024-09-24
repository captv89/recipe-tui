package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// Main model struct
type model struct {
	state        state
	categories   []Category
	meals        []Meal
	recipe       Recipe
	categoryList list.Model
	mealList     list.Model
	recipeView   viewport.Model
	width        int
	loading      bool
}

// listItem is a helper for managing the list of categories and meals
type listItem struct {
	title string
	desc  string
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return i.desc }
func (i listItem) FilterValue() string { return i.title }

// Initialize app with categories
func (m model) Init() tea.Cmd {
	return fetchCategories()
}

// Update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "enter":
			switch m.state {
			case categoryList:
				// Move to meal list
				selectedCategory := m.categories[m.categoryList.Index()]
				m.state = mealList
				m.loading = true
				return m, fetchMeals(selectedCategory.Name)

			case mealList:
				// Move to recipe detail
				selectedMeal := m.meals[m.mealList.Index()]
				m.state = recipeDetail
				m.loading = true
				return m, fetchRecipe(selectedMeal.ID)

			case recipeDetail:
				// No further action
				return m, nil
			}

		case "backspace":
			// Handle going back
			if m.state == mealList {
				m.state = categoryList
			} else if m.state == recipeDetail {
				m.state = mealList
			}

		case "up", "k":
			if m.state == recipeDetail {
				if m.recipeView.YOffset > 0 {
					m.recipeView.YOffset -= 1
				}
				log.Println("up", m.recipeView.YOffset)
			}
		case "down", "j":
			if m.state == recipeDetail {
				if m.recipeView.YOffset < m.recipeView.TotalLineCount() {
					m.recipeView.YOffset += 1
				}
				log.Println("At Bottom", m.recipeView.AtBottom())
				log.Println("down", m.recipeView.YOffset)
			}
		}

	case []Category:
		items := make([]list.Item, len(msg))
		for i, category := range msg {
			items[i] = listItem{title: category.Name, desc: "Category"}
		}
		m.categories = msg
		m.categoryList.SetItems(items)
		m.categoryList.Title = "Meal Categories"
		m.state = categoryList
		m.loading = false
		return m, tea.Batch(nil)

	case []Meal:
		items := make([]list.Item, len(msg))
		for i, meal := range msg {
			items[i] = listItem{title: meal.Name, desc: "Meal"}
		}
		m.meals = msg
		m.mealList.SetItems(items)
		m.mealList.Title = "Meals"
		m.state = mealList
		m.loading = false
		return m, tea.Batch(nil)

	case Recipe:
		m.recipe = msg
		content := formatRecipe(msg)
		log.Println("Formated Recipe: ", content)
		m.recipeView.SetContent(content)
		log.Printf("Viewport height: %d, Content height: %d", m.recipeView.Height, len(strings.Split(content, "\n")))
		log.Println("Width", m.recipeView.Width, "Rendered Width", m.width)
		log.Printf("Viewport Scroll Y: %d", m.recipeView.YOffset)
		log.Printf("Viewport Y Position: %d, Visible Lines: %d", m.recipeView.YPosition, m.recipeView.VisibleLineCount())
		log.Printf("Viewport Total Lines: %d", m.recipeView.TotalLineCount())
		m.state = recipeDetail
		m.loading = false
		return m, tea.Batch(nil)

	}

	// Update Viewport height
	switch msg := msg.(type) {
	// Use the full screen for the list and viewport
	case tea.WindowSizeMsg:
		// Get the size of the list style
		h, v := listStyle.GetFrameSize()
		m.categoryList.SetSize(msg.Width-h, msg.Height-v)
		m.mealList.SetSize(msg.Width-h, msg.Height-v)

		// Get the size of the viewport style
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight
		log.Println("Vertical Height", verticalMarginHeight)
		hv, wv := viewportStyle.GetFrameSize()
		log.Println("x", hv, "y", wv)
		m.recipeView.Height = msg.Height - verticalMarginHeight
		viewportStyle.Width(msg.Width - hv - 2)
		m.width = msg.Width - hv - 2
		m.recipeView.Width = msg.Width - hv - 2
	}

	// Update lists and viewport based on state
	// These are updated based on the user's interaction
	switch m.state {
	case categoryList:
		var cmd tea.Cmd
		m.categoryList, cmd = m.categoryList.Update(msg)
		return m, cmd
	case mealList:
		var cmd tea.Cmd
		m.mealList, cmd = m.mealList.Update(msg)
		return m, cmd
	case recipeDetail:
		var cmd tea.Cmd
		m.recipeView, cmd = m.recipeView.Update(msg)
		return m, cmd
	}

	return m, nil
}

// View function
func (m model) View() string {
	if m.loading {
		return "Loading..."
	}

	switch m.state {
	case categoryList:
		return listStyle.Render(m.categoryList.View())
	case mealList:
		return listStyle.Render(m.mealList.View())
	case recipeDetail:
		log.Println(m.recipeView.Height, m.recipeView.YPosition, m.recipeView.VisibleLineCount(), m.recipeView.TotalLineCount(), m.recipeView.YOffset)
		body := fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.recipeView.View(), m.footerView())
		return viewportStyle.Render(body)
	}
	return ""
}

func (m model) headerView() string {
	title := titleStyle.Render(m.recipe.StrMeal)
	line := strings.Repeat("─", max(0, m.width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.recipeView.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func main() {
	// Log
	// Open or create the log file
	logFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close()

	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Set the log output to the file
	log.SetOutput(logFile)

	m := model{}

	// Initialise the Model
	m.categoryList = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	m.categoryList.Title = "Meal Categories"
	m.categoryList.SetSize(20, 10) // Ensure the size is set

	m.mealList = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	m.mealList.Title = "Meals"
	m.mealList.SetSize(20, 10) // Ensure the size is set

	m.recipeView = viewport.New(0, 0)
	m.recipeView.SetContent("")

	m.loading = true

	// Start the Bubble Tea program
	log.Println("Starting Recipe TUI program")
	p := tea.NewProgram(m,
		tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
