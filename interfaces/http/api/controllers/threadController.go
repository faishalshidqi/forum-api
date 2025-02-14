package controllers

import (
	"forum-api/applications/security"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ThreadController struct {
	TaskUsecase  domains.ThreadUsecase
	TokenManager security.AuthnTokenManager
	Env          *bootstrap.Env
}

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
	addedThread, err := tc.TaskUsecase.Add(c, addThreadRequest, id)
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
