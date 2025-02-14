package controllers

import (
	"forum-api/applications/security"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentController struct {
	CommentUsecase domains.CommentUsecase
	TokenManager   security.AuthnTokenManager
	Env            *bootstrap.Env
}

func (cc *CommentController) AddComment(c *gin.Context) {
	token, err := cc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	id, err := cc.TokenManager.VerifyToken(token, cc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	addCommentRequest := domains.AddCommentRequest{}
	if err := c.ShouldBind(&addCommentRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	threadId := c.Param("thread_id")
	addedComment, err := cc.CommentUsecase.Add(c, addCommentRequest, id, threadId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domains.AddCommentResponse{
		Status:  "success",
		Message: "Comment added successfully",
		Data:    addedComment,
	})
}
