package handlers

import (
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/utils"
	"github.com/KalyanKanuri/GoCart/internal/services"
	"github.com/gin-gonic/gin"
)

var authService *services.AuthService

func InitAuthHandlers(srv *server.Server) {
	authService = services.NewAuthService(srv.DBConn, srv.CFG)
}

func registerHandler(ctx *gin.Context) {
	var req *dto.RegisterRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := authService.RegisterUser(req)
	if err != nil {
		utils.InternalServerError(ctx, "Registration Failed", err)
		return
	}

	utils.CreatedResponse(ctx, "User Registered Successfully", resp)
}

func loginHandler(ctx *gin.Context) {
	var req *dto.LoginRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := authService.LoginUser(req)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to Login", err)
		return
	}

	utils.SuccessResponse(ctx, "User Logged in Successfully", resp)
}

func refreshTokenHandler(ctx *gin.Context) {
	var req *dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := authService.RefreshToken(req)
	if err != nil {
		utils.InternalServerError(ctx, "Failed to generate refresh token", err)
		return
	}

	utils.SuccessResponse(ctx, "Refresh token generated successfully", resp)
}

func logoutHandler(ctx *gin.Context) {
	var req *dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	if err := authService.Logout(req); err != nil {
		utils.InternalServerError(ctx, "Unable to logout user", err)
		return
	}

	utils.SuccessResponse(ctx, "User logged out successfully")
}
