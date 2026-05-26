package gin

import "github.com/mrdyuke/infrum/domain"

var Paths = domain.LibraryPaths{
	"/cmd/api/":           domain.LibraryFile{"main.go": ""},
	"/config/":            domain.LibraryFile{"config.go": ""},
	"/internal/domain/":   domain.LibraryFile{"domain.go": ""},
	"/internal/routes/":   domain.LibraryFile{"routes.go": ""},
	"/internal/usecases/": domain.LibraryFile{"usecases.go": ""},
}

var Gin = domain.Library{
	LibName: "Gin",
	LibPath: Paths,
}
