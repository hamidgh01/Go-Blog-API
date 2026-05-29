package handlers

import (
	"Go-Blog-API/internal/application/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Create(c *gin.Context) {

}

func (h *UserHandler) UpdateUsername(c *gin.Context) {

}

func (h *UserHandler) UpdateEmail(c *gin.Context) {

}

func (h *UserHandler) UpdateBio(c *gin.Context) {

}

func (h *UserHandler) UpdatePassword(c *gin.Context) {

}

func (h *UserHandler) UpdateEnabled(c *gin.Context) {

}

func (h *UserHandler) Delete(c *gin.Context) {

}

// func (h *UserHandler) GetList(c *gin.Context) {}

func (h *UserHandler) GetByID(c *gin.Context) {

}

func (h *UserHandler) GetByUsername(c *gin.Context) {

}

func (h *UserHandler) GetByEmail(c *gin.Context) {

}
