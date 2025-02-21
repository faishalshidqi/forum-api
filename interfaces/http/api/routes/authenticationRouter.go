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
	passwordHash := security.NewBcryptPasswordHash()
	refreshTokenRepository := repository.NewPostgresRefreshTokenRepository(db)
	authenticationController := controllers.AuthenticationController{
		AuthenticationUsecase:  usecase.NewAuthenticationUsecase(userRepository, tokenManager, passwordHash, timeout),
		RefreshTokenRepository: refreshTokenRepository,
		Env:                    env,
	}
	router.POST("/authentications", authenticationController.Login)
	router.PUT("/authentications", authenticationController.RefreshToken)
	router.DELETE("/authentications", authenticationController.Logout)
}
