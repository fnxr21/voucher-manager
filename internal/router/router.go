package router

import (
	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	Auth(e)

}
