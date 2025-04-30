package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wine-shop/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, _ := h.service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteUserByID(id); err != nil { // <-- исправлено
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
