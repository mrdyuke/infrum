package gin

import (
	_ "embed"

	"github.com/mrdyuke/infrum/domain"
)

var Paths = domain.LibraryPaths{
	"/":                  domain.LibraryFile{".env": envTemplate, ".gitignore": gitignoreTemplate, "Dockerfile": dockerfileTemplate, "Makefile": makefileTemplate},
	"/cmd/api/":          domain.LibraryFile{"main.go": mainTemplate},
	"/config/":           domain.LibraryFile{"config.go": configTemplate},
	"/internal/domain/":  domain.LibraryFile{"domain.go": domainTemplate},
	"/internal/routes/":  domain.LibraryFile{"routes.go": routesTemplate},
	"/internal/usecases/": domain.LibraryFile{"usecases.go": usecasesTemplate},
}

//go:embed main.tmpl
var mainTemplate string

//go:embed config.tmpl
var configTemplate string

//go:embed domain.tmpl
var domainTemplate string

//go:embed routes.tmpl
var routesTemplate string

//go:embed usecases.tmpl
var usecasesTemplate string

//go:embed env.tmpl
var envTemplate string

//go:embed gitignore.tmpl
var gitignoreTemplate string

//go:embed Dockerfile.tmpl
var dockerfileTemplate string

//go:embed Makefile.tmpl
var makefileTemplate string

var Gin = domain.Library{
	LibName: "Gin",
	LibPath: Paths,
}
