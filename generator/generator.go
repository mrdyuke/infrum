package generator

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mrdyuke/infrum/domain"
)

type Generator struct {
	AppName   string
	TargetDir string
	LibList   domain.LibraryList
}

func NewGenerator() (*Generator, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return &Generator{
		AppName:   "MyApp",
		TargetDir: currDir,
		LibList:   make(domain.LibraryList, 0, 2),
	}, nil
}

func (g *Generator) Generate() error {
	g.TargetDir = filepath.Join(g.TargetDir, g.AppName)

	for _, lib := range g.LibList {
		for path, file := range lib.LibPath {
			err := os.MkdirAll(filepath.Join(g.TargetDir, path), 0755)
			if err != nil {
				return err
			}
			for name, data := range file {
				err := os.WriteFile(filepath.Join(g.TargetDir, path, name), []byte(data), 0644)
				if err != nil {
					return err
				}
			}
		}
	}

	err := os.Chdir(g.TargetDir)
	if err != nil {
		return err
	}

	err = exec.Command("go", "mod", "init", g.AppName).Run()
	if err != nil {
		return err
	}

	err = exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		return err
	}

	return nil
}
