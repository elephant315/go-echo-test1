package http

import (
	"github.com/elephant315/go-echo-test1/internal/app/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RegisterHandlers registers the HTTP handlers
func RegisterHandlers(e *echo.Echo) {
	e.POST("/itinerary", handleItinerary)
}

// handleItinerary processes the POST request to reconstruct the itinerary
// @Summary Reconstruct flight itinerary
// @Description Reconstruct the complete flight itinerary from given tickets
// @Accept  json
// @Produce  json
// @Param   tickets  body  [][][]string  true  "Array of ticket pairs"
// @Success 200 {array} string "Complete itinerary"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Failed to construct itinerary"
// @Router /itinerary [post]
func handleItinerary(c echo.Context) error {
	var tickets [][]string
	if err := c.Bind(&tickets); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	itinerary, err := services.ReconstructItinerary(tickets)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to construct itinerary"})
	}

	return c.JSON(http.StatusOK, itinerary)
}
