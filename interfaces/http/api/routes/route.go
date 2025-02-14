package routes

import (
	"forum-api/commons/bootstrap"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, gin *gin.Engine) {
	router := gin.Group("")
	newSignupRouter(env, timeout, db, router)
	newAuthnRouter(env, timeout, db, router)

	newThreadRouter(env, timeout, db, router)
}
