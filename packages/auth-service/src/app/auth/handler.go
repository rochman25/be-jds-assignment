package auth

import (
	"auth-service/src/dto"
	"auth-service/src/factory"
	"auth-service/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: NewService(f),
	}
}

func (h *Handler) Register(c *gin.Context) {
	var input dto.RegisterRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := util.GenerateCustomMessages(err)
		c.JSON(http.StatusBadRequest, gin.H{"error_message": errorMessages})
		return
	}

	data, err := h.service.RegisterService(c, input)
	if err != nil {
		errCode := util.GetErrorCode(err)
		statusCode := http.StatusInternalServerError
		if errCode > 0 {
			statusCode = errCode
		}
		c.JSON(statusCode, util.ApiErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, data)
	return
}
