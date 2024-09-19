package main

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/glamour"
)

func loadRecipes() ([]list.Item, error) {
	// For simplicity, loading recipes from markdown files in the "recipes" directory
	files, err := os.ReadDir("./recipes")
	if err != nil {
		return nil, err
	}

	var recipes []list.Item
	for _, file := range files {
		if !file.IsDir() && file.Type().IsRegular() {
			content, err := os.ReadFile("./recipes/" + file.Name())
			if err != nil {
				return nil, err
			}
			compiler, err := glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(80),
			)
			if err != nil {
				return nil, err
			}
			parsed, err := compiler.Render(string(content))
			if err != nil {
				return nil, err
			}
			recipes = append(recipes, item{
				title:       file.Name(),
				author:      "Captain Cook",
				desc:        "You will find the recipe inside.",
				content:     parsed,
			})
		}
	}
	return recipes, nil
}
