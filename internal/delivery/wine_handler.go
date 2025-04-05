package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wine-shop/internal/models"
	"wine-shop/internal/service"
)

type WineHandler struct {
	service *service.WineService
}

func NewWineHandler(s *service.WineService) *WineHandler {
	return &WineHandler{service: s}
}

func (h *WineHandler) GetAll(c *gin.Context) {
	wines, _ := h.service.GetAllWines()
	c.JSON(http.StatusOK, wines)
}

func (h *WineHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	wine, err := h.service.GetWineById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, wine)
}

func (h *WineHandler) Create(c *gin.Context) {
	var wineEdit models.WineEdit
	if err := c.ShouldBindJSON(&wineEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	wine, err := h.service.CreateWine(wineEdit.Name, wineEdit.Year, wineEdit.Price, wineEdit.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, wine)
}

func (h *WineHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var wineEdit models.WineEdit
	if err := c.ShouldBindJSON(&wineEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	updated, err := h.service.UpdateWine(uint(id), &wineEdit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *WineHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteWine(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
