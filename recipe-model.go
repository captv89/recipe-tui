package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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

type RecipeModel struct {
	Recipe   Recipe
	Viewport viewport.Model
}

// Init
func (m RecipeModel) Init() tea.Cmd {
	return nil
}

// Update
func (m RecipeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View
func (m RecipeModel) View() string {
	body := fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
	return viewportStyle.Render(body)
}

func (m RecipeModel) headerView() string {
	title := titleStyle.Render("Recipe Name Here")
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m RecipeModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
