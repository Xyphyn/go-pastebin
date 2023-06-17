package main

import (
	"io"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xylight.dev/pastebin/common"
	"xylight.dev/pastebin/routers"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&routers.Paste{})
}

func main() {
	db := common.Init()
	Migrate(db)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.GET("/", func(c *gin.Context) {
		data := routers.GetPastes()
		c.HTML(200, "index.go.html", gin.H{
			"Pastes": data,
		})
	})

	r.POST("/", func(c *gin.Context) {
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}

		common.DB.Create(&routers.Paste{
			Title:   "No title",
			Content: string(data),
		})
		c.JSON(201, gin.H{
			"message": "Created",
		})
	})

	api := r.Group("/api")

	routers.APIRouter(api)

	r.Run("0.0.0.0:3000")
}
