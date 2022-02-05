package route

import (
	"echo-sample/handler"
	"echo-sample/repository"
	"echo-sample/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Musicians handler.Musicians
}

func NeweHandler() Handler {
	musicianRepository := repository.NewMusician()
	musicianService := service.NewMusician(musicianRepository)

	return Handler{
		Musicians: handler.NewMusicians(musicianService),
	}
}

func Route(h Handler, e *echo.Echo) error {
	musicians := e.Group("/musicians")
	musicians.GET("", h.Musicians.List)

	return nil
}
