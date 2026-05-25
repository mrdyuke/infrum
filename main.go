package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens"
)

func main() {
	generator, err := generator.NewGenerator()
	if err != nil {
		fmt.Println(err)
	}

	appNameScreen := screens.NewAppNameScreen(generator)

	program := tea.NewProgram(appNameScreen)
	_, err = program.Run()
	if err != nil {
		fmt.Println(err)
	}
}
