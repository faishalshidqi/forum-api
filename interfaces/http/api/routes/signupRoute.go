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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	passwordHash := security.NewBcryptPasswordHash()
	signupController := controllers.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepository, timeout),
		PasswordHash:  passwordHash,
		Env:           env,
	}
	router.POST("/users", signupController.Signup)
}
