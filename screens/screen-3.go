package screens

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mrdyuke/infrum/domain"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens/logo"
)

type AppDriverScreen struct {
	generator  *generator.Generator
	cursor     int
	driverList domain.LibraryList
}

func NewAppDriverScreen(list domain.LibraryList, generator *generator.Generator) *AppDriverScreen {

	return &AppDriverScreen{
		generator:  generator,
		driverList: list,
	}
}

func (s *AppDriverScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			if s.cursor < len(s.driverList)-1 {
				s.cursor++
			}

		case "enter":
			s.generator.LibList = append(s.generator.LibList, s.driverList[s.cursor])
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s *AppDriverScreen) View() tea.View {
	str := fmt.Sprintf("\n%s\n", logo.Logo)

	for k, v := range s.driverList {
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

func (s *AppDriverScreen) Init() tea.Cmd {
	return nil
}
