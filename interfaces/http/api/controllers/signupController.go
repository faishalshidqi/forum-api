package controllers

import (
	"forum-api/applications/security"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type SignupController struct {
	SignupUsecase domains.SignupUsecase
	PasswordHash  security.PasswordHash
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	request := &domains.SignupRequest{}
	if err := c.ShouldBind(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: "Invalid request body",
			Status:  "fail",
		})
		return
	}
	_, err := sc.SignupUsecase.GetUserByUsername(c, request.Username)
	if err == nil {
		c.AbortWithStatusJSON(http.StatusConflict, domains.ErrorResponse{
			Message: "User already exists",
			Status:  "fail",
		})
		return
	}
	encryptedPassword, err := sc.PasswordHash.HashPassword(request.Password)
	if err != nil {
		slog.Error("Failed to hash password")
		c.AbortWithStatusJSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
			Status:  "fail",
		})
		return
	}
	request.Password = encryptedPassword
	user, err := sc.SignupUsecase.Create(c, request)
	if err != nil {
		slog.Error("Failed to add user")
		c.AbortWithStatusJSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
			Status:  "fail",
		})
		return
	}
	c.JSON(http.StatusCreated, domains.SignupResponse{
		Message: "User created",
		Status:  "success",
		Data: domains.SignupResponseData{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
		},
	})
}
