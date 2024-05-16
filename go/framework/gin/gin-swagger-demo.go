package main

import (
	"net/http"

	docs "api/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} resp
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	// g.JSON(http.StatusOK, "helloworld")
	g.JSON(http.StatusOK, render.JSON{
		Data: resp{
			Code:    0,
			Message: "helloworld",
			Data:    nil,
		},
	})
}

type resp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
