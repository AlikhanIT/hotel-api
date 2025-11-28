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

func (h *Handler) GetAll(c *gin.Context) {
	hotels, err := h.rep.GetHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get hotels"})
		return
	}
	c.JSON(http.StatusOK, hotels)
}

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

func (h *Handler) Create(c *gin.Context) {
	var input domain.Hotel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.rep.CreateHotel(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create hotel"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input domain.Hotel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	input.ID = uint(id) // обновляем ID

	if err := h.rep.UpdateHotel(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update hotel"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := h.rep.DeleteHotel(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete hotel"})
		return
	}
	c.Status(http.StatusNoContent)
}
