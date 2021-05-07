package v1

import (
	"github.com/EgorMizerov/kindergarten/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initAuthRoutes(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
		auth.POST("/refresh", h.refresh)

		authorized := auth.Group("", h.authentication())
		authorized.POST("/logout", h.logout)
	}
}

type inputSignUp struct {
	Email       string `binding:"required,email"`
	Password    string `binding:"required"`
	FirstName   string `binding:"required"`
	LastName    string `binding:"required"`
	MiddleName  string `binding:"required"`
	DateOfBirth int64  `binding:"required"`
}

func (h *Handler) singUp(ctx *gin.Context) {
	var input inputSignUp

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Auth.SignUp(ctx, domain.User{
		Email:        input.Email,
		PasswordHash: input.Password,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		MiddleName:   input.MiddleName,
		DateOfBirth:  input.DateOfBirth})
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type inputSignIn struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (h *Handler) singIn(ctx *gin.Context) {
	var input inputSignIn

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	accessToken, refreshToken, err := h.service.Auth.SignIn(ctx, domain.User{
		Email:        input.Email,
		PasswordHash: input.Password})
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

type inputRefreshToken struct {
	ID           string `binding:"required,len=24"`
	RefreshToken string `binding:"required,uuid"`
}

func (h *Handler) refresh(ctx *gin.Context) {
	var input inputRefreshToken

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	accessToken, refreshToken, err := h.service.Auth.Refresh(ctx, input.ID, input.RefreshToken)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
func (h *Handler) logout(ctx *gin.Context) {
	id := ctx.GetString("sub")
	err := h.service.Auth.Logout(ctx, id)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "logout",
	})
}
