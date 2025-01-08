package http

import (
	"auth-service/src/factory"
	"auth-service/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Here we define route function for user Handlers that accepts gin.Engine and factory parameters
func NewHttp(g *gin.Engine, f *factory.Factory) {
	g.Use(middleware.CORSMiddleware())

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	//override response when route not found
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, nil)
	})
}
