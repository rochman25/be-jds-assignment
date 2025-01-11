package auth

import (
	"auth-service/src/dto"
	"auth-service/src/factory"
	"auth-service/util"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
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
		if err == io.EOF {
			errorMessages = []string{"Payload is required"}
		}
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

func (h *Handler) Login(c *gin.Context) {
	var input dto.LoginRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := util.GenerateCustomMessages(err)
		if err == io.EOF {
			errorMessages = []string{"Payload is required"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error_message": errorMessages})
		return
	}

	data, err := h.service.LoginService(c, input)
	if err != nil {
		errCode := util.GetErrorCode(err)
		statusCode := http.StatusInternalServerError
		if errCode > 0 {
			statusCode = errCode
		}
		c.JSON(statusCode, util.ApiErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func (h *Handler) DebugToken(c *gin.Context) {
	header := c.Request.Header["Authorization"]
	if len(header) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error_message": "token required"})
		return
	}
	token := strings.Replace(header[0], "Bearer ", "", -1)
	data, err := util.ParseAccessToken(token)
	if err != nil {
		errCode := util.GetErrorCode(err)
		statusCode := http.StatusInternalServerError
		if errCode > 0 {
			statusCode = errCode
		}
		c.JSON(statusCode, util.ApiErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
