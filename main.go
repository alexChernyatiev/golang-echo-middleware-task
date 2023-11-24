package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server is running")

	s := echo.New()

	s.GET("/status", Handler)

	s.Use(MV)

	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	s := fmt.Sprintf("Count of days: %d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

func MV(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hVal := ctx.Request().Header.Get("User-Role")

		if hVal == "admin" {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
