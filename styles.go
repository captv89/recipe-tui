package main

import "github.com/charmbracelet/lipgloss"

var (
	listStyle = lipgloss.NewStyle().Margin(1, 2)

	viewportStyle = lipgloss.NewStyle().Margin(1, 1)

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)
