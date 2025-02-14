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
	ThreadUsecase  domains.ThreadUsecase
	TokenManager   security.AuthnTokenManager
	Env            *bootstrap.Env
}

// AddComment Create Comment godoc
//
//	@Summary		Create Comment
//	@Description	Creating a new comment. Only valid users can create a comment to a valid thread
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			content			body		string	true	"content of the comment"
//	@Param			thread_id		path		string	true	"Thread ID"
//	@Success		200				{object}	domains.AddCommentResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		404				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/threads/{thread_id}/comments [post]
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
	_, err = cc.ThreadUsecase.GetById(c, threadId)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Status:  "fail",
			Message: "Thread does not exist",
		})
		return
	}
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
