package handlers

import (
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/utils"
	"github.com/KalyanKanuri/GoCart/internal/services"
	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func InitUserHandlers(srv *server.Server) {
	userService = services.NewUserService(
		srv.DBConn,
		srv.Logger,
	)
}

func GetProfileHandler(ctx *gin.Context) {
	userID := ctx.GetUint("id")
	resp, err := userService.GetProfile(userID)
	if err != nil {
		userService.Log.Error().Str("handler", "get_profile").Uint("user_id", userID).Err(err).Msg("failed to get user profile")
		utils.NotFoundError(ctx, "user not found", err)
		return
	}

	userService.Log.Info().Str("handler", "get_profile").Uint("user_id", userID).Msg("user profile retrieved successfully")
	utils.SuccessResponse(ctx, "User Profile retrieved successfully", resp)
}

func UpdateProfileHandler(ctx *gin.Context) {
	userID := ctx.GetUint("id")
	var req *dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		userService.Log.Error().Str("handler", "update_profile").Uint("user_id", userID).Err(err).Msg("profile update request binding failed")
		utils.BadRequestError(ctx, "Invalid Request", err)
		return
	}

	resp, err := userService.UpdateProfile(userID, req)
	if err != nil {
		userService.Log.Error().Str("handler", "update_profile").Uint("user_id", userID).Err(err).Msg("failed to update user profile")
		utils.NotFoundError(ctx, "User not found", err)
		return
	}

	userService.Log.Info().Str("handler", "update_profile").Uint("user_id", userID).Msg("user profile updated successfully")
	utils.SuccessResponse(ctx, "user updated successfully", resp)
}
