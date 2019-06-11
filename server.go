package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	"log"
	"os"
	"fmt"
)

type Data struct {
	Value    string `json:"value"`
}


func PrintPostForm(c *gin.Context) {
	fmt.Println(c.PostForm("value"))
	c.JSON(200, gin.H{"value": 1})
}

func PrintPostForm2(c *gin.Context) {
	var data = Data{}
	c.JSON(200, gin.H{"value": data.Value})
}

func main() {
	log.SetOutput(os.Stdout)

	r := gin.New()
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"authorization", "cache-control", "client-id", "client-type", "content-type"},
		AllowCredentials: true,
		MaxAge:           30 * time.Minute,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"project": "getFormExample"})
	})

	r.POST("/abc", PrintPostForm)

	r.Run(":3000")
}