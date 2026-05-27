package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mrdyuke/infrum/generator"
	"github.com/mrdyuke/infrum/screens"
	"github.com/mrdyuke/infrum/templates"
)

func main() {
	gen, err := generator.NewGenerator()
	if err != nil {
		fmt.Printf("Error initializing generator: %v\n", err)
		os.Exit(1)
	}

	applyScreen := screens.NewApplyScreen(gen)
	appDriverScreen := screens.NewAppDriverScreen(applyScreen, templates.DriverList, gen)
	appFrameworkScreen := screens.NewAppFrameworkScreen(appDriverScreen, templates.FrameworkList, gen)
	appNameScreen := screens.NewAppNameScreen(appFrameworkScreen, gen)

	program := tea.NewProgram(appNameScreen)
	_, err = program.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}

}
