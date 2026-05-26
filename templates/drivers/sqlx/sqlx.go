package sqlx

import "github.com/mrdyuke/infrum/domain"

var Paths = domain.LibraryPaths{
	"/repository/postgres/": domain.LibraryFile{"postgres.go": ""},
}

var Sqlx = domain.Library{
	LibName: "Sqlx",
	LibPath: Paths,
}
