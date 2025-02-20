package handlers

import "github.com/labstack/echo/v4"

type FrontendHandler struct {
}

func (h *FrontendHandler) ServeDashboard(c echo.Context) error {
	return c.File("../../templates/pages/index.html")
}
