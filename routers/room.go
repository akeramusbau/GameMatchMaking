package routers

import (
	"github.com/akeramusbau/game-matchmaking/controller"
	"github.com/gin-gonic/gin"
)

type Room struct{}

func (p *Room) Route(route *gin.Engine) {
	router := route.Group("/room")
	Controller := controller.RoomController{}

	router.POST("/", Controller.CreateRoom)
}
