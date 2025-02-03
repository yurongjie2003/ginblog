package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yurongjie2003/ginblog/config"
	"net/http"
)

func Init() error {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	err := r.Run(config.HttpPort)
	return err
}
