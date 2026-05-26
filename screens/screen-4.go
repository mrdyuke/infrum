package screens

import (
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens/logo"
)

type ApplyScreen struct {
	generator *generator.Generator
	list      [2]string
	cursor    int
}

func NewApplyScreen(generator *generator.Generator) *ApplyScreen {

	return &ApplyScreen{
		generator: generator,
		list:      [2]string{"Apply", "Cancel"},
	}
}

func (s *ApplyScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			if s.cursor < len(s.list)-1 {
				s.cursor++
			}

		case "enter":
			if s.cursor == 0 {
				err := s.generator.Generate()
				if err != nil {
					log.Fatal(err)
				}
			}
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s *ApplyScreen) View() tea.View {
	str := fmt.Sprintf("\n%s\n", logo.Logo)

	for k, v := range s.list {
		pointer := ""
		if s.cursor == k {
			pointer = ">"
		}
		str += lipgloss.JoinVertical(lipgloss.Top, fmt.Sprintf("\n%s %s\n", pointer, v))
	}

	view := tea.NewView(str)
	view.AltScreen = true
	return view
}

func (s *ApplyScreen) Init() tea.Cmd {
	return nil
}
