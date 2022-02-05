package main

import (
	"echo-sample/config"
	"echo-sample/route"
	"fmt"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func main() {
	conf := (&config.Config{}).Build()
	err := run(conf)
	if err != nil {
		log.Fatal("fail.run", err)
	}
}

func run(conf *config.Config) error {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := route.NeweHandler()

	route.Route(h, e)

	err := e.Start(fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		err = errors.Wrapf(err, "fail.start")
		return err
	}
	return nil
}
