# infrum

A CLI tool that generates Go project scaffolding with your choice of web framework and database driver.

<img width="553" height="487" alt="image" src="https://github.com/user-attachments/assets/f1a4cbb8-f206-4ecc-9cb2-f3f3005c13ac" />

## Install

```bash
go install github.com/mrdyuke/infrum@latest
```

## Usage

Run the TUI wizard:

```bash
infrum
```

Follow the steps to set your project name, pick a framework, a database driver, and generate the project.

### Frameworks

- [Gin](https://github.com/gin-gonic/gin)
- [Echo](https://github.com/labstack/echo)
- [Fiber](https://github.com/gofiber/fiber)
- Vanilla (net/http)

### Database Drivers

- [pgx](https://github.com/jackc/pgx)
- [GORM](https://gorm.io)
- [sqlx](https://github.com/jmoiron/sqlx)

## Output

Each generated project includes:

- Clean architecture layout (`internal/domain`, `internal/usecases`, `internal/routes`, `config`)
- CRUD for a `User` resource
- Dockerfile
- Makefile
- `.env` and `.gitignore`

## Requirements

- Go 1.23+
- `go` binary in `$PATH`

## License

[MIT](https://github.com/mrdyuke/infrum/blob/main/LICENSE)
