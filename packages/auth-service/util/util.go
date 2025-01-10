package util

import "github.com/gin-gonic/gin"

func Index(g *gin.Engine, name string, version string) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{
			Name:    name,
			Version: version,
		})
	})
}
