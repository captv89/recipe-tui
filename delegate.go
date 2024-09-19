package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type customDelegate struct {
	list.DefaultDelegate
}

func (d customDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(item)
	if !ok {
		return
	}

	// Define styles
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	authorStyle := lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("238"))
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	// Render the item
	fmt.Fprintf(w, "%s\n%s\n%s\n\n",
		titleStyle.Render(item.Title()),
		authorStyle.Render("By: "+item.Author()),
		descStyle.Render(item.Description()),
	)
}
