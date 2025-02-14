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

func newCommentRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, router *gin.RouterGroup) {
	commentRepository := repository.NewPostgresCommentRepository(db)
	threadRepository := repository.NewPostgresThreadRepository(db)
	tokenManager := security.NewJwtTokenManager()
	commentController := controllers.CommentController{
		CommentUsecase: usecase.NewCommentUsecase(commentRepository, timeout),
		ThreadUsecase:  usecase.NewThreadUsecase(threadRepository, timeout),
		TokenManager:   tokenManager,
		Env:            env,
	}
	router.POST("/threads/:thread_id/comments", commentController.AddComment)
}
