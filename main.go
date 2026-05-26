package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens"
	"github.com/mrdyuke/infrum/templates"
)

func main() {
	generator, err := generator.NewGenerator()
	if err != nil {
		fmt.Println(err)
	}

	appFrameworkScreen := screens.NewAppFrameworkScreen(templates.FrameworkList, generator)
	appNameScreen := screens.NewAppNameScreen(appFrameworkScreen, generator)

	program := tea.NewProgram(appNameScreen)
	_, err = program.Run()
	if err != nil {
		fmt.Println(err)
	}
}
