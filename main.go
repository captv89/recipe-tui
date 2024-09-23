package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var viewportStyle = lipgloss.NewStyle().Margin(0, 2)

type item struct {
	title, author, desc, content string
}

func (i item) Title() string       { return i.title }
func (i item) Author() string      { return i.author }
func (i item) Description() string { return i.desc }
func (i item) Content() string     { return i.content }
func (i item) FilterValue() string { return i.title }

type model struct {
	list        list.Model
	viewport    viewport.Model
	showDetails bool
	currentItem item
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "enter":
			if m.list.SelectedItem() != nil {
				m.currentItem = m.list.SelectedItem().(item)
			}
			m.showDetails = !m.showDetails
			if m.showDetails && m.list.SelectedItem() != nil {
				selectedItem := m.list.SelectedItem().(item)
				m.viewport.SetContent(fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\nContent: %s",
					selectedItem.Title(),
					selectedItem.Author(),
					selectedItem.Description(),
					selectedItem.Content(),
				))
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.viewport.Width = msg.Width - h
		m.viewport.Height = msg.Height - v - 5
		// if m.showDetails && m.list.SelectedItem() != nil {
		// 	selectedItem := m.list.SelectedItem().(item)
		// 	m.viewport.SetContent(fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\nContent: %s",
		// 		selectedItem.Title(),
		// 		selectedItem.Author(),
		// 		selectedItem.Description(),
		// 		selectedItem.Content(),
		// 	))
		// }
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	if m.showDetails {
		m.viewport, cmd = m.viewport.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.showDetails {
		body := fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
		return viewportStyle.Render(body)
	}

	return docStyle.Render(m.list.View())
}

func (m model) headerView() string {
	title := titleStyle.Render(m.currentItem.Title())
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func main() {
	// items := []list.Item{
	// 	item{title: "Raspberry Pi's", author: "Raspberry Pi's", desc: "I have 'em all over my house"},
	// 	item{title: "Nutella", author: "Nutella", desc: "It's good on toast"},
	// 	item{title: "Bitter melon", author: "Bitter melon", desc: "It cools you down"},
	// 	item{title: "Nice socks", author: "Nice socks", desc: "And by that I mean socks without holes"},
	// 	item{title: "Eight hours of sleep", author: "Eight hours of sleep", desc: "I had this once"},
	// 	item{title: "Cats", author: "Cats", desc: "Usually"},
	// 	item{title: "Plantasia, the album", author: "Plantasia, the album", desc: "My plants love it too"},
	// 	item{title: "Pour over coffee", author: "Pour over coffee", desc: "It takes forever to make though"},
	// 	item{title: "VR", author: "VR", desc: "Virtual reality...what is there to say?"},
	// 	item{title: "Noguchi Lamps", author: "Noguchi Lamps", desc: "Such pleasing organic forms"},
	// 	item{title: "Linux", author: "Linux", desc: "Pretty much the best OS"},
	// 	item{title: "Business school", author: "Business school", desc: "Just kidding"},
	// 	item{title: "Pottery", author: "Pottery", desc: "Wet clay is a great feeling"},
	// 	item{title: "Shampoo", author: "Shampoo", desc: "Nothing like clean hair"},
	// 	item{title: "Table tennis", author: "Table tennis", desc: "It's surprisingly exhausting"},
	// 	item{title: "Milk crates", author: "Milk crates", desc: "Great for packing in your extra stuff"},
	// 	item{title: "Afternoon tea", author: "Afternoon tea", desc: "Especially the tea sandwich part"},
	// 	item{title: "Stickers", author: "Stickers", desc: "The thicker the vinyl the better"},
	// 	item{title: "20° Weather", author: "20° Weather", desc: "Celsius, not Fahrenheit"},
	// 	item{title: "Warm light", author: "Warm light", desc: "Like around 2700 Kelvin"},
	// 	item{title: "The vernal equinox", author: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
	// 	item{title: "Gaffer's tape", author: "Gaffer's tape", desc: "Basically sticky fabric"},
	// 	item{title: "Terrycloth", author: "Terrycloth", desc: "In other words, towel fabric"},
	// }

	items, err := loadRecipes()
	if err != nil {
		fmt.Println("Error loading recipes:", err)
		os.Exit(1)
	}

	l := list.New(items, list.NewDefaultDelegate(), 20, 10)
	l.Title = "My Fave Things"

	// Initialize the viewport with default size; it will be set in Update
	vp := viewport.New(20, 10)

	m := model{
		list:     l,
		viewport: vp,
	}

	p := tea.NewProgram(&m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
