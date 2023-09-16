package routes

import (
	"net/http"

	"github.com/ItsJustVaal/Holodash/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllChannels(ctx *gin.Context) {
	channelMap := controllers.HandleGetAllChannels()

	ctx.IndentedJSON(http.StatusOK, channelMap)
}
