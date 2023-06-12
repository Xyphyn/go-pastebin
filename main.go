package main

import (
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

	api := r.Group("/api")

	routers.APIRouter(api)

	r.Run("0.0.0.0:3000")
}
