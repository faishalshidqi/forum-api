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

func newThreadRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	threadRepository := repository.NewPostgresThreadRepository(db)
	commentRepository := repository.NewPostgresCommentRepository(db)
	tokenManager := security.NewJwtTokenManager()
	threadController := controllers.ThreadController{
		ThreadUsecase:  usecase.NewThreadUsecase(threadRepository, timeout),
		CommentUsecase: usecase.NewCommentUsecase(commentRepository, timeout),
		TokenManager:   tokenManager,
		Env:            env,
	}
	router.POST("/threads", threadController.AddThread)
	router.GET("/threads/:thread_id", threadController.GetByThread)
}
