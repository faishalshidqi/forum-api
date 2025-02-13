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

// Signup AddUser godoc
//
//	@Summary		Register A User
//	@Description	New user must have a unique email address
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			email		body		string	true	"email address of the new user, must be unique"
//	@Param			password	body		string	true	"password of the new user"
//	@Param			name		body		string	true	"name of the new user"
//	@Success		201			{object}	domains.SignupResponse
//	@Failure		400			{object}	domains.ErrorResponse
//	@Failure		409			{object}	domains.ErrorResponse
//	@Failure		500			{object}	domains.ErrorResponse
//	@Router			/users [post]
func (sc *SignupController) Signup(c *gin.Context) {
	request := &domains.SignupRequest{}
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: "Invalid request body",
			Status:  "fail",
		})
		return
	}
	_, err := sc.SignupUsecase.GetUserByUsername(c, request.Username)
	if err == nil {
		c.JSON(http.StatusConflict, domains.ErrorResponse{
			Message: "User already exists",
			Status:  "fail",
		})
		return
	}
	encryptedPassword, err := sc.PasswordHash.HashPassword(request.Password)
	if err != nil {
		slog.Error("Failed to hash password")
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
			Status:  "fail",
		})
		return
	}
	request.Password = encryptedPassword
	user, err := sc.SignupUsecase.Create(c, request)
	if err != nil {
		slog.Error("Failed to add user")
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
			Status:  "fail",
		})
		return
	}
	c.JSON(http.StatusCreated, domains.SignupResponse{
		Message: "User created",
		Status:  "success",
		Data:    user,
	})
}
