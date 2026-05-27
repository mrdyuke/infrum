package screens

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mrdyuke/infrum/domain"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens/logo"
	"github.com/mrdyuke/infrum/screens/styles"
)

type AppDriverScreen struct {
	generator  *generator.Generator
	cursor     int
	driverList domain.LibraryList
	nextScreen tea.Model
}

func NewAppDriverScreen(next tea.Model, list domain.LibraryList, generator *generator.Generator) *AppDriverScreen {

	return &AppDriverScreen{
		generator:  generator,
		driverList: list,
		nextScreen: next,
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
			return s.nextScreen, nil
		}
	}

	return s, nil
}

func (s *AppDriverScreen) View() tea.View {
	str := styles.LogoStyle(logo.Logo)
	str += styles.TitleStyle("Select driver ↓")
	for k, v := range s.driverList {
		pointer := ""
		if s.cursor == k {
			pointer = styles.CursorPointerStyle()
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
