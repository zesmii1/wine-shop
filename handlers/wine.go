package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"wine-shop/models"
)

var wines = []models.Wine{
	{ID: 1, Name: "Cabernet Sauvignon", Origin: "France", Year: 2018, Price: 29.99},
	{ID: 2, Name: "Merlot", Origin: "Italy", Year: 2020, Price: 19.99},
}

func GetWines(c *gin.Context) {
	c.JSON(http.StatusOK, wines)
}

func GetWineByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine ID"})
		return
	}

	for _, wine := range wines {
		if wine.ID == uint(id) {
			c.JSON(http.StatusOK, wine)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}

func CreateWine(c *gin.Context) {
	var newWine models.Wine
	if err := c.ShouldBindJSON(&newWine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newWine.ID = uint(len(wines) + 1)
	wines = append(wines, newWine)
	c.JSON(http.StatusCreated, newWine)
}

func UpdateWine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine ID"})
		return
	}

	var updatedWine models.Wine
	if err := c.ShouldBindJSON(&updatedWine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, wine := range wines {
		if wine.ID == uint(id) {
			wines[i] = updatedWine
			wines[i].ID = uint(id)
			c.JSON(http.StatusOK, wines[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}

func DeleteWine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wine ID"})
		return
	}

	for i, wine := range wines {
		if wine.ID == uint(id) {
			wines = append(wines[:i], wines[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Wine deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
}
