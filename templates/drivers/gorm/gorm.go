package gorm

import (
	_ "embed"

	"github.com/mrdyuke/infrum/domain"
)

var Paths = domain.LibraryPaths{
	"/internal/repository/postgres/": domain.LibraryFile{"postgres.go": postgresTemplate},
}

//go:embed postgres.tmpl
var postgresTemplate string

var Gorm = domain.Library{
	LibName: "Gorm",
	LibPath: Paths,
}
