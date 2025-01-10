package auth

import "github.com/gin-gonic/gin"

func (h *Handler) Router(g *gin.RouterGroup) {
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
}
