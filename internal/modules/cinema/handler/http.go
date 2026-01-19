package handler

import (
	"strconv"

	"github.com/geraldiaditya/ratix-backend/internal/modules/cinema/domain"
	"github.com/geraldiaditya/ratix-backend/internal/modules/cinema/service"
	"github.com/gofiber/fiber/v2"
)

type CinemaHandler struct {
	Service *service.CinemaService
}

func NewCinemaHandler(service *service.CinemaService) *CinemaHandler {
	return &CinemaHandler{Service: service}
}

func (h *CinemaHandler) RegisterRoutes(app *fiber.App) {
	cinemas := app.Group("/cinemas")
	locations := app.Group("/locations")

	locations.Get("/", h.handleGetLocations)
	cinemas.Get("/brands", h.handleGetBrands)
	cinemas.Get("/", h.handleGetCinemas)

	// Seat Selection
	app.Get("/showtimes/:id/seats", h.handleGetSeats)
}

func (h *CinemaHandler) handleGetLocations(c *fiber.Ctx) error {
	resp, err := h.Service.GetLocations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(resp)
}

func (h *CinemaHandler) handleGetBrands(c *fiber.Ctx) error {
	resp, err := h.Service.GetBrands()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(resp)
}

func (h *CinemaHandler) handleGetCinemas(c *fiber.Ctx) error {
	city := c.Query("city")
	brand := c.Query("brand")
	lat := c.QueryFloat("lat", 0)
	lon := c.QueryFloat("lon", 0)
	radius := c.QueryFloat("radius", 0)

	filter := domain.CinemaFilter{
		City:   city,
		Brand:  brand,
		Lat:    lat,
		Lon:    lon,
		Radius: radius,
	}

	resp, err := h.Service.GetCinemas(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(resp)
}

func (h *CinemaHandler) handleGetSeats(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	resp, err := h.Service.GetSeatLayout(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(resp)
}
