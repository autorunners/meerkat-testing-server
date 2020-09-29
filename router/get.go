package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetJson(c *gin.Context) {
	log.Println("")
	c.JSON(200, gin.H{"code": 200, "message": "pongv2.0 ping"})
}

func GetNormal(c *gin.Context) {
	log.Println("")
	c.Status(200)
}

func GetNormalString(c *gin.Context) {
	log.Println("")
	c.String(200, "%v", "hello world")
}
