package controllers

import (
	"fmt"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationController struct {
	AuthenticationUsecase  domains.AuthenticationUsecase
	RefreshTokenRepository domains.RefreshTokenRepository
	Env                    *bootstrap.Env
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
//	@Success		201			{object}	domains.LoginResponse
//	@Failure		400			{object}	domains.ErrorResponse
//	@Failure		401			{object}	domains.ErrorResponse
//	@Failure		500			{object}	domains.ErrorResponse
//	@Router			/authentications [post]
func (ac *AuthenticationController) Login(c *gin.Context) {
	loginRequest := domains.LoginRequest{}
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid request body",
		})
		return
	}
	user, err := ac.AuthenticationUsecase.GetUserByUsername(c, loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid username or password",
		})
		return
	}
	err = ac.AuthenticationUsecase.CheckPasswordHash(loginRequest.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid username or password",
		})
		return
	}
	accessToken, err := ac.AuthenticationUsecase.CreateAccessToken(user, ac.Env.AccessTokenKey, ac.Env.AccessTokenAge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	refreshToken, err := ac.AuthenticationUsecase.CreateRefreshToken(user, ac.Env.RefreshTokenKey, ac.Env.RefreshTokenAge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	err = ac.RefreshTokenRepository.Add(c, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domains.LoginResponse{
		Status: "success",
		Data: domains.LoginResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

// RefreshToken Refresh Authentication, generating new access and refresh token godoc
//
//	@Summary		Refresh Authentication
//	@Description	Generating new access token using a refresh token. Only valid refresh token will generate new
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		string	true	"refresh token possessed by the user"
//	@Success		200				{object}	domains.RefreshResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/authentications [put]
func (ac *AuthenticationController) RefreshToken(c *gin.Context) {
	refreshRequest := domains.RefreshRequest{}
	if err := c.ShouldBind(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid request body",
		})
		return
	}
	id, err := ac.AuthenticationUsecase.ValidateToken(refreshRequest.RefreshToken, ac.Env.RefreshTokenKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: fmt.Sprintf("Invalid refresh token: %s", err.Error()),
		})
		return
	}
	_, err = ac.RefreshTokenRepository.Fetch(c, refreshRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid refresh token",
		})
		return
	}
	user, err := ac.AuthenticationUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Status:  "fail",
			Message: "user not found",
		})
		return
	}
	accessToken, err := ac.AuthenticationUsecase.CreateAccessToken(user, ac.Env.AccessTokenKey, ac.Env.AccessTokenAge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domains.RefreshResponse{
		Status: "success",
		Data: domains.RefreshResponseData{
			AccessToken: accessToken,
		},
	})
}

// Logout Sign out godoc
//
//	@Summary		Sign Out
//	@Description	Signing User Out. Requires refresh token
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		string	true	"refresh token possessed by the user"
//	@Success		200				{object}	domains.SuccessResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/authentications [delete]
func (ac *AuthenticationController) Logout(c *gin.Context) {
	refreshRequest := domains.RefreshRequest{}
	if err := c.ShouldBind(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid request body",
		})
		return
	}
	_, err := ac.RefreshTokenRepository.Fetch(c, refreshRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Status:  "fail",
			Message: "Invalid refresh token",
		})
		return
	}
	err = ac.RefreshTokenRepository.Delete(c, refreshRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domains.SuccessResponse{
		Status:  "success",
		Message: "Successfully logged out",
	})
}
