package http

import (
	"net/http"
	"strconv"

	"github.com/AlikhanIT/hotel-api/internal/domain/hotel"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	rep domain.Repository
}

// NewHandler создает новый обработчик
func NewHandler(r domain.Repository) Handler {
	return Handler{rep: r}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/hotels", h.GetAll)
	r.GET("/hotels/:id", h.GetByID)
	r.POST("/hotels", h.Create)
	r.PUT("/hotels/:id", h.Update)
	r.DELETE("/hotels/:id", h.Delete)
}

// GetAll godoc
// @Summary      Get list of hotels
// @Description  Returns all hotels
// @Tags         hotels
// @Produce      json
// @Success      200  {array}   domain.Hotel
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/hotels [get]
func (h *Handler) GetAll(c *gin.Context) {
	hotels, err := h.rep.GetHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get hotels"})
		return
	}
	c.JSON(http.StatusOK, hotels)
}

// GetByID godoc
// @Summary      Get hotel by ID
// @Description  Returns hotel by its ID
// @Tags         hotels
// @Produce      json
// @Param        id   path      int  true  "Hotel ID"
// @Success      200  {object}  domain.Hotel
// @Failure      404  {object}  map[string]string
// @Router       /api/v1/hotels/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	hotel, err := h.rep.GetHotelByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "hotel not found"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

// Create godoc
// @Summary      Create hotel
// @Description  Adds a new hotel
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        hotel  body     HotelDTO  true  "Hotel"
// @Success      201    {object} domain.Hotel
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /api/v1/hotels [post]
func (h *Handler) Create(c *gin.Context) {
	var input HotelDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.rep.CreateHotel(input.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create hotel"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// Update godoc
// @Summary      Update hotel
// @Description  Updates hotel's data
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        id     path     int          true  "Hotel ID"
// @Param        hotel  body     domain.Hotel true  "Hotel"
// @Success      200    {object} domain.Hotel
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /api/v1/hotels/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input domain.Hotel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	input.ID = uint(id)

	if err := h.rep.UpdateHotel(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update hotel"})
		return
	}

	c.JSON(http.StatusOK, input)
}

// Delete godoc
// @Summary      Delete hotel
// @Description  Removes hotel by ID
// @Tags         hotels
// @Param        id   path      int  true  "Hotel ID"
// @Success      204  "No Content"
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/hotels/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := h.rep.DeleteHotel(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete hotel"})
		return
	}
	c.Status(http.StatusNoContent)
}
