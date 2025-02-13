package controllers

import (
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationController struct {
	AuthenticationUsecase domains.AuthenticationUsecase
	Env                   *bootstrap.Env
}

// Login Log In godoc
//
//	@Summary		Login with Username & Password
//	@Description	authenticate user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			username	body		string	true	"username address of the user"
//	@Param			password	body		string	true	"password of the user"
//	@Success		201			{object}	domains.AuthenticationResponse
//	@Failure		400			{object}	domains.ErrorResponse
//	@Failure		401			{object}	domains.ErrorResponse
//	@Failure		500			{object}	domains.ErrorResponse
//	@Router			/authentications [post]
func (au *AuthenticationController) Login(c *gin.Context) {
	loginRequest := domains.LoginRequest{}
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid request body",
		})
		return
	}
	user, err := au.AuthenticationUsecase.GetUserByUsername(c, loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid username or password",
		})
		return
	}
	err = au.AuthenticationUsecase.CheckPasswordHash(loginRequest.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid username or password",
		})
		return
	}
	accessToken, err := au.AuthenticationUsecase.CreateAccessToken(user, au.Env.AccessTokenKey, au.Env.AccessTokenAge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	refreshToken, err := au.AuthenticationUsecase.CreateRefreshToken(user, au.Env.RefreshTokenKey, au.Env.RefreshTokenAge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domains.AuthenticationResponse{
		Status: "success",
		Data: domains.AuthnResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
