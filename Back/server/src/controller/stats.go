package controller

import (
	"github.com/labstack/echo"
	"net/http"

	"server/server/src/repository"
)

// GetStatistics Reads and returns current Covid-19 statistics for every country.
// HTTP internal server error is returned if an error happens in database operation.
func GetStatistics(c echo.Context) error {
	statistics, err := repository.FindAllStatistics()
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, statistics)
}
