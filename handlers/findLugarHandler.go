package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/julioolivares90/tiendaCercaApi/scraper"
)

func GetLugar(c echo.Context) error {
	lugar := c.QueryParam("lugar")
	idLugar := scraper.GetIDMunicipio(lugar)

	datos := scraper.FindLugarByID(idLugar)

	return c.JSON(http.StatusOK,datos)
}