package generator

import (
	"os"

	"github.com/mrdyuke/infrum/domain"
)

var FrameworkList = domain.LibraryList{}

var DriverList = domain.LibraryList{}

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
