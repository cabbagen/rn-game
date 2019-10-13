package main

import (
	"rn-game/views"
	"rn-game/routers"
	"rn-game/database"
	"github.com/gin-gonic/gin"
	"rn-game/middlewares"
)

func RegisterDatabase() database.Connector {
	databaseInstance := database.NewDatabase(
		"mysql",
		database.DefaultMysqlConfig["username"],
		database.DefaultMysqlConfig["password"],
		database.DefaultMysqlConfig["dbName"],
	)

	databaseInstance.Connect()

	return databaseInstance
}

func main() {
	engine := gin.Default()

	middleware.RegisterMiddleware(engine)

	router.RegisterRouters(engine)
	
	view.RegisterTemplate(engine)

	databaseInstance := RegisterDatabase()

	defer databaseInstance.Destroy()
	
	engine.Run(":9090")
}
