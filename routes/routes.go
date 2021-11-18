package routes

import (
	"github.com/gin-gonic/gin"
	"webSocket_dev/controller"
)

func Setup(app *gin.Engine)  {
	app.GET("/",controller.GetWerSocket)
}


