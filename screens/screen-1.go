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
	generator  *generator.Generator
	textInput  textinput.Model
	nextScreen tea.Model
}

func NewAppNameScreen(next tea.Model, generator *generator.Generator) *AppNameScreen {
	ti := textinput.New()
	ti.Focus()
	ti.SetWidth(45)
	ti.CharLimit = 45
	ti.Placeholder = "  My awesome app!"

	return &AppNameScreen{
		textInput:  ti,
		generator:  generator,
		nextScreen: next,
	}
}

func (s *AppNameScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return s, tea.Quit
		case "enter":
			if s.textInput.Value() != "" {
				s.generator.AppName = s.textInput.Value()
				return s.nextScreen, nil
			}
		}
	}

	var cmd tea.Cmd
	s.textInput, cmd = s.textInput.Update(msg)
	return s, cmd
}

func (s *AppNameScreen) View() tea.View {
	str := lipgloss.JoinVertical(lipgloss.Top, fmt.Sprintf("\n%s\n\n%s\n", logo.Logo, s.textInput.View()))
	view := tea.NewView(str)
	view.AltScreen = true
	return view
}

func (s *AppNameScreen) Init() tea.Cmd {
	return nil
}
