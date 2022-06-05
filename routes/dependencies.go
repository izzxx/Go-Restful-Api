package routes

import (
	"github.com/allegro/bigcache/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

type Dependencies struct {
	Db     *pgxpool.Pool
	Memory *bigcache.BigCache
	App    *echo.Echo
}
