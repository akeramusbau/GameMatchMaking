package routers

import (
	"github.com/akeramusbau/game-matchmaking/controller"
	"github.com/gin-gonic/gin"
)

type Match struct{}

func (p *Match) Route(route *gin.Engine) {
	router := route.Group("/match")
	Controller := controller.MatchController{}

	router.POST("/join", Controller.JoinMatch)
}
