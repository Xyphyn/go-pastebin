package routers

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xylight.dev/pastebin/common"
)

type Paste struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetPastes() *[]Paste {
	var data []Paste
	common.DB.Find(&data).Limit(50)
	return &data
}

func GetPaste(id int) (*Paste, error) {
	var data Paste
	err := common.DB.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func PastesRouter(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		data := GetPastes()

		if data != nil {
			if c.Param("short") == "true" {
				shortData := []Paste{}

				for _, paste := range *data {
					paste.Content = paste.Content[:80]
					shortData = append(shortData, paste)
				}

				c.IndentedJSON(200, shortData)
			} else {
				c.IndentedJSON(200, data)
			}
		} else {
			c.AbortWithStatus(404)
		}
	})

	r.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		data, err := GetPaste(id)
		if data != nil {
			c.IndentedJSON(200, data)
		} else {
			c.AbortWithError(404, err)
		}
	})

	r.POST("/", func(c *gin.Context) {
		var newPaste Paste

		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(data, &newPaste); err != nil {
			c.AbortWithError(400, err)
			return
		}

		common.DB.Create(&newPaste)

		c.JSON(201, newPaste)
	})
}
