package gorm

import "github.com/mrdyuke/infrum/domain"

var Paths = domain.LibraryPaths{
	"/repository/postgres/": domain.LibraryFile{"postgres.go": ""},
}

var Gorm = domain.Library{
	LibName: "Gorm",
	LibPath: Paths,
}
