package controller

import (
	"fmt"
	"net/http"

	"github.com/akeramusbau/game-matchmaking/services"
	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

type createPlayerInput struct {
	Username string `json:"username" binding:"required"`
}

func (controller PlayerController) CreatePlayer(c *gin.Context) {
	var input createPlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var response = ErrorResponse{
			Msg: "Validation Error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	player, err := services.PlayerService.CreateOne(input.Username)
	if err != nil {
		var response = ErrorResponse{
			Msg: "Failed to Create Player",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	var responseString = fmt.Sprintf("Successfully created a new player with username : %s", player.Username)
	var response = SuccessResponse{
		Status: true,
		Msg:    responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}
