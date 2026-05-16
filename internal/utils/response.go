package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, APIResponse{
		Success: true,
		Data:    data,
	})
}

func Created(c *gin.Context, data any) {
	Success(c, http.StatusCreated, data)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, APIResponse{
		Success: false,
		Message: message,
	})
}

func ValidationError(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}
