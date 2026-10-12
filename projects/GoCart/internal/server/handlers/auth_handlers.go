package handlers

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/utils"
	"github.com/KalyanKanuri/GoCart/internal/services"
	"github.com/gin-gonic/gin"
)

var authService *services.AuthService

func InitAuthHandlers(srv *server.Server) {
	authService = services.NewAuthService(srv.DBConn, srv.CFG, srv.Logger)
}

func tokenFingerprint(token string) string {
	if token == "" {
		return "empty"
	}
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])[:12]
}

func registerHandler(ctx *gin.Context) {
	var req *dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		authService.Log.Error().Str("handler", "register").Err(err).Msg("register request binding failed")
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := authService.RegisterUser(req)
	if err != nil {
		authService.Log.Error().Str("handler", "register").Str("email", req.Email).Err(err).Msg("user registration failed")
		utils.InternalServerError(ctx, "Registration Failed", err)
		return
	}

	authService.Log.Info().Str("handler", "register").Str("email", resp.User.Email).Uint("user_id", resp.User.ID).Msg("user registration completed successfully")
	utils.CreatedResponse(ctx, "User Registered Successfully", resp)
}

func loginHandler(ctx *gin.Context) {
	var req *dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		authService.Log.Error().Str("handler", "login").Err(err).Msg("login request binding failed")
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := authService.LoginUser(req)
	if err != nil {
		authService.Log.Error().Str("handler", "login").Str("email", req.Email).Err(err).Msg("user login failed")
		utils.InternalServerError(ctx, "Failed to Login", err)
		return
	}

	authService.Log.Info().Str("handler", "login").Str("email", req.Email).Uint("user_id", resp.User.ID).Msg("user login completed successfully")
	utils.SuccessResponse(ctx, "User Logged in Successfully", resp)
}

func refreshTokenHandler(ctx *gin.Context) {
	var req *dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		authService.Log.Error().Str("handler", "refresh_token").Err(err).Msg("refresh token request binding failed")
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := authService.RefreshToken(req)
	if err != nil {
		authService.Log.Error().Str("handler", "refresh_token").Str("token_fingerprint", tokenFingerprint(req.RefreshToken)).Err(err).Msg("refresh token generation failed")
		utils.InternalServerError(ctx, "Failed to generate refresh token", err)
		return
	}

	authService.Log.Info().Str("handler", "refresh_token").Str("email", resp.User.Email).Uint("user_id", resp.User.ID).Str("token_fingerprint", tokenFingerprint(req.RefreshToken)).Msg("refresh token generated successfully")
	utils.SuccessResponse(ctx, "Refresh token generated successfully", resp)
}

func logoutHandler(ctx *gin.Context) {
	var req *dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		authService.Log.Error().Str("handler", "logout").Err(err).Msg("logout request binding failed")
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	if err := authService.Logout(req); err != nil {
		authService.Log.Error().Str("handler", "logout").Str("token_fingerprint", tokenFingerprint(req.RefreshToken)).Err(err).Msg("user logout failed")
		utils.InternalServerError(ctx, "Unable to logout user", err)
		return
	}

	authService.Log.Info().Str("handler", "logout").Str("token_fingerprint", tokenFingerprint(req.RefreshToken)).Msg("user logout completed successfully")
	utils.SuccessResponse(ctx, "User logged out successfully")
}
