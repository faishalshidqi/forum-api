package routes

import (
	"forum-api/applications/usecase"
	"forum-api/commons/bootstrap"
	"forum-api/infrastructures/repository"
	"forum-api/infrastructures/security"
	"forum-api/interfaces/http/api/controllers"
	"github.com/gin-gonic/gin"
	"time"
)

func newAuthnRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	userRepository := repository.NewPostgresUserRepository(db)
	tokenManager := security.NewJwtTokenManager()
	authenticationController := controllers.AuthenticationController{
		AuthenticationUsecase: usecase.NewAuthenticationUsecase(userRepository, tokenManager, timeout),
		Env:                   env,
	}
	router.POST("/authentications", authenticationController.Login)
}
