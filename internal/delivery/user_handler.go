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
