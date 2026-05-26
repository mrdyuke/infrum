package pgx

import "github.com/mrdyuke/infrum/domain"

var Paths = domain.LibraryPaths{
	"/repository/postgres/": domain.LibraryFile{"postgres.go": ""},
}

var Pgx = domain.Library{
	LibName: "Pgx",
	LibPath: Paths,
}
