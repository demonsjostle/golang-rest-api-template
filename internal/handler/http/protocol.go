package protocol

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPProtocol struct{}

func NewHTTPProtocol() *HTTPProtocol {
	return &HTTPProtocol{}
}

func (p *HTTPProtocol) RegisterRoutes(e *echo.Echo) {
	// Root route template
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to API")
	})
}
