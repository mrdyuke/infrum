package screens

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mrdyuke/infrum/domain"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens/logo"
)

type AppFrameworkScreen struct {
	generator     *generator.Generator
	cursor        int
	frameworkList domain.LibraryList
}

func NewAppFrameworkScreen(list domain.LibraryList, generator *generator.Generator) *AppFrameworkScreen {

	return &AppFrameworkScreen{
		generator:     generator,
		frameworkList: list,
	}
}

func (s *AppFrameworkScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return s, tea.Quit

		case "up":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down":
			if s.cursor < len(s.frameworkList)-1 {
				s.cursor++
			}

		case "enter":
			return s, nil
		}
	}

	return s, nil
}

func (s *AppFrameworkScreen) View() tea.View {
	str := fmt.Sprintf("\n%s\n", logo.Logo)

	for k, v := range s.frameworkList {
		pointer := ""
		if s.cursor == k {
			pointer = ">"
		}
		str += lipgloss.JoinVertical(lipgloss.Top, fmt.Sprintf("\n%s %s\n", pointer, v.LibName))
	}

	view := tea.NewView(str)
	view.AltScreen = true
	return view
}

func (s *AppFrameworkScreen) Init() tea.Cmd {
	return nil
}
