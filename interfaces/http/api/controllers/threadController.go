package controllers

import (
	"forum-api/applications/security"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ThreadController struct {
	ThreadUsecase  domains.ThreadUsecase
	CommentUsecase domains.CommentUsecase
	TokenManager   security.AuthnTokenManager
	Env            *bootstrap.Env
}

// AddThread Create A New Thread godoc
//
//	@Summary		Create Thread
//	@Description	Creating a new thread. Only valid users can create a thread
//	@Tags			threads
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer Token"
//	@Param			title			body		string	true	"title of the thread"
//	@Param			body			body		string	true	"body of the thread"
//	@Success		201				{object}	domains.AddThreadResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/threads [post]
func (tc *ThreadController) AddThread(c *gin.Context) {
	token, err := tc.TokenManager.GetBearerToken(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	id, err := tc.TokenManager.VerifyToken(token, tc.Env.AccessTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	addThreadRequest := domains.AddThreadRequest{}
	if err := c.ShouldBind(&addThreadRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	addedThread, err := tc.ThreadUsecase.Add(c, addThreadRequest, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domains.AddThreadResponse{
		Message: "successfully added thread",
		Status:  "success",
		Data:    addedThread,
	})
}

func (tc *ThreadController) GetByThread(c *gin.Context) {
	threadId := c.Param("thread_id")
	thread, err := tc.ThreadUsecase.GetById(c, threadId)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Status:  "fail",
			Message: "Thread does not exist",
		})
		return
	}
	comments, _ := tc.CommentUsecase.GetByThread(c, threadId)
	thread.Comments = comments
	c.JSON(http.StatusOK, domains.GetCommentsByThreadResponse{
		Status:  "success",
		Message: "Thread fetched successfully",
		Data: struct {
			Thread domains.GetThreadByIDResponseData `json:"thread"`
		}{Thread: thread},
	})
}
