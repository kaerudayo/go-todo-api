package main

import (
  "app/domain/reposiroy"
  "net/http"
  "github.com/labstack/echo/v4"
//  "github.com/labstack/echo/v4/middleware"
)

func main() {
//  e := echo.New()
//  e.Use(middleware.Logger())
//  e.Use(middleware.Recover())
//  e.GET("/", hello)
//
//  // server start
//  e.Logger.Fatal(e.Start(":8080"))

  reposiroy.GetAll()
}

func hello(c echo.Context) error {
  reposiroy.GetAll()
  return c.String(http.StatusOK, "Hello, World!")
}
