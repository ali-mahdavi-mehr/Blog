package custumerror

import (
	"github.com/labstack/echo"
)

type CustomError struct {
	Status  string
	Message []string
}

type httpError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

func (e httpError) Error() string {
	return e.Key + ": " + e.Message
}

func GenerateCustomError(status string, messages []string) *CustomError {
	return &CustomError{status, messages}
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	//code := http.StatusInternalServerError
	//result := ""

	if he, ok := err.(*httpError); ok {
		c.Error(he)

	}
	c.Error(err)

}
