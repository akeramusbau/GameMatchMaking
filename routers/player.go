package routers

import (
	"github.com/akeramusbau/game-matchmaking/controller"
	"github.com/gin-gonic/gin"
)

type Player struct{}

func (p *Player) Route(route *gin.Engine) {
	router := route.Group("/player")
	Controller := controller.PlayerController{}

	router.POST("/", Controller.CreatePlayer)
}
