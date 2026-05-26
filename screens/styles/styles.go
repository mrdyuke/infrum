package styles

import (
	"charm.land/lipgloss/v2"
)

var PrimaryColor = "#7927F5"

func TextInputStyle(str string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true).
		Padding(0, 1).
		MarginTop(1).
		BorderForeground(lipgloss.Color(PrimaryColor)).
		Render(str)
}

func LogoStyle(str string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(PrimaryColor)).
		Padding(1, 2, 0, 2).
		Render(str)
}

func TitleStyle(str string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(PrimaryColor)).
		Padding(1, 3).
		Bold(true).
		MarginTop(2).
		MarginBottom(1).
		Render(str)
}
