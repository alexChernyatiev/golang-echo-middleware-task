package mv

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const ROLE_ADMIN = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hVal := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(hVal, ROLE_ADMIN) {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
