package routers

import "github.com/gin-gonic/gin"

func APIRouter(r *gin.RouterGroup) {
	pastes := r.Group("/pastes")
	PastesRouter(pastes)
}
