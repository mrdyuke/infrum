package generator

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

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
	if len(strings.TrimSpace(g.AppName)) == 0 {
		return errors.New("empty app name")
	}
	if len(g.LibList) == 0 {
		return errors.New("no libraries selected")
	}

	g.TargetDir = filepath.Join(g.TargetDir, g.AppName)

	for _, lib := range g.LibList {
		for path, file := range lib.LibPath {
			err := os.MkdirAll(filepath.Join(g.TargetDir, path), 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", path, err)
			}
			for name, data := range file {
				filePath := filepath.Join(g.TargetDir, path, name)
				f, err := os.Create(filePath)
				if err != nil {
					return fmt.Errorf("failed to create file %s: %w", filePath, err)
				}

				tmpl, err := template.New(name).Parse(data)
				if err != nil {
					_ = f.Close()
					return fmt.Errorf("failed to parse template %s: %w", name, err)
				}

				err = tmpl.Execute(f, g)
				closeErr := f.Close()
				if err != nil {
					return fmt.Errorf("failed to execute template %s: %w", name, err)
				}
				if closeErr != nil {
					return fmt.Errorf("failed to close file %s: %w", filePath, closeErr)
				}
			}
		}
	}

	if err := os.Chdir(g.TargetDir); err != nil {
		return fmt.Errorf("failed to change to target directory: %w", err)
	}

	if err := exec.Command("go", "mod", "init", g.AppName).Run(); err != nil {
		return fmt.Errorf("go mod init failed: %w", err)
	}

	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return fmt.Errorf("go mod tidy failed: %w", err)
	}

	return nil
}
