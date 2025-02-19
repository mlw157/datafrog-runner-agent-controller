package handlers

import "github.com/labstack/echo/v4"

type FrontendHandler struct {
}

func (h *InstanceHandler) ServeDashboard(c echo.Context) error {
	return c.File("templates/index.html")
}
