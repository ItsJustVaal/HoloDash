package routes

import (
	"github.com/ItsJustVaal/Holodash/controllers"
	"github.com/gin-gonic/gin"
)

var tempch controllers.Channel

func AddChannelToDB(ctx *gin.Context) {

	tempch.Id = "testid"
	tempch.Name = "testname"
	tempch.Region = "testregion"
	controllers.HandleAddChannelToDB(tempch)
}
