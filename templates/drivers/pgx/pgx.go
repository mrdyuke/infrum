package pgx

import (
	_ "embed"

	"github.com/mrdyuke/infrum/domain"
)

var Paths = domain.LibraryPaths{
	"/internal/repository/postgres/": domain.LibraryFile{"postgres.go": postgresTemplate},
}

//go:embed postgres.tmpl
var postgresTemplate string

var Pgx = domain.Library{
	LibName: "Pgx",
	LibPath: Paths,
}
