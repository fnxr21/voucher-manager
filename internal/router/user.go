package router

import (
	"github.com/fnxr21/voucher-manager/internal/handler"
	repositories "github.com/fnxr21/voucher-manager/internal/repository"
	"github.com/fnxr21/voucher-manager/pkg/middleware"
	"github.com/fnxr21/voucher-manager/pkg/mysql"
	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Group) {
	repo := repositories.Repository(mysql.DB)
	h := handler.HandlerUser(repo)
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)
	e.GET("/reauth", middleware.Auth(h.Reauth))
}
