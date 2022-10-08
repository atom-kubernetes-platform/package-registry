package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type packageInfo struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/packages", func(c *gin.Context) {
		c.JSON(http.StatusOK, []packageInfo{
			packageInfo{Name: "MyPackage"},
		})
	})
	_ = r.Run()
}
