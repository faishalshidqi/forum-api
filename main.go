package main

import (
	"forum-api/commons/bootstrap"
	"forum-api/interfaces/http/api/routes"
	"github.com/gin-gonic/gin"
	"time"
)

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
