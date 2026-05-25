package screens

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens/logo"
)

type AppNameScreen struct {
	generator *generator.Generator
	textinput textinput.Model
}

func NewAppNameScreen(generator *generator.Generator) *AppNameScreen {
	ti := textinput.New()
	ti.Focus()

	return &AppNameScreen{
		textinput: ti,
		generator: generator,
	}
}

func (s *AppNameScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return s, tea.Quit
		case "enter":
			if s.textinput.Value() != "" {
				s.generator.AppName = s.textinput.Value()
				return s, nil
			}
		}
	}

	var cmd tea.Cmd
	s.textinput, cmd = s.textinput.Update(msg)
	return s, cmd
}

func (s *AppNameScreen) View() tea.View {
	str := lipgloss.JoinVertical(lipgloss.Top, fmt.Sprintf("\n%s\n%s\n", logo.Logo, s.textinput.View()))
	view := tea.NewView(str)
	view.AltScreen = true
	return view
}

func (s *AppNameScreen) Init() tea.Cmd {
	return nil
}
