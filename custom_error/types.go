package custumerror

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo"
	"net/http"
)

var (
	duplicateError string = "23505"
)

func FindDBError(inputErr error, modelName string) *echo.HTTPError {
	err := inputErr.(*pgconn.PgError)
	switch {
	case err.Code == duplicateError:
		return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("%s already exists!", modelName))
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, inputErr.Error())

	}
}
