package templates

import (
	"github.com/mrdyuke/infrum/domain"
	"github.com/mrdyuke/infrum/templates/drivers/gorm"
	"github.com/mrdyuke/infrum/templates/drivers/pgx"
	"github.com/mrdyuke/infrum/templates/drivers/sqlx"
	"github.com/mrdyuke/infrum/templates/frameworks/echo"
	"github.com/mrdyuke/infrum/templates/frameworks/fiber"
	"github.com/mrdyuke/infrum/templates/frameworks/gin"
	"github.com/mrdyuke/infrum/templates/frameworks/vanilla"
)

var FrameworkList = domain.LibraryList{
	vanilla.Vanilla,
	gin.Gin,
	echo.Echo,
	fiber.Fiber,
}

var DriverList = domain.LibraryList{
	pgx.Pgx,
	gorm.Gorm,
	sqlx.Sqlx,
}
