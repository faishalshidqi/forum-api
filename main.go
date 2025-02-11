package main

import (
	"forum-api/commons/bootstrap"
	"forum-api/interfaces/http/api/routes"
	"github.com/gin-gonic/gin"
	"time"
)

//	@title			Forum API
//	@version		1.0
//	@description	This is a Forum API

// @host		localhost:9000
// @BasePath	/
// @securityDefinitions.apikey
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.DB
	defer app.DB.Pool.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	router := gin.Default()
	routes.Setup(env, timeout, db, router)
	router.Run(env.ServerAddr)
}
