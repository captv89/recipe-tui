package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Category struct {
	Name string
}

func (c Category) Title() string       { return c.Name }
func (c Category) Description() string { return c.Name }
func (c Category) FilterValue() string { return c.Name }

type CategoryModel struct {
	List list.Model
}

func (m CategoryModel) initializeList() []list.Item {
	var items []list.Item
	// Get the categories
	categories, err := loadCategoryItems()
	if err != nil {
		fmt.Println("Error loading categories:", err)
		return nil
	}
	for _, category := range categories {
		items = append(items, category)
	}
	return items
}

// Init
func (m CategoryModel) Init() tea.Cmd {
	return nil
}

// Update
func (m CategoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	items := m.initializeList()
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Categories"

	// Redraw the list
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
		fmt.Println("WindowSizeMsg:", msg.Width, msg.Height)
		l.SetSize(msg.Width-h, msg.Height-v)
	}

	m.List = l

	return m, cmd
}

// View
func (m CategoryModel) View() string {
	return listStyle.Render(m.List.View())
}
