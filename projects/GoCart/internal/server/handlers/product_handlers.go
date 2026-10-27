package handlers

import (
	"errors"
	"strconv"

	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/utils"
	"github.com/KalyanKanuri/GoCart/internal/services"
	"github.com/gin-gonic/gin"
)

var productService *services.ProductService

func InitProductHandlers(srv *server.Server) {
	productService = services.NewProductService(
		srv.DBConn,
		srv.Logger,
	)
}

func GetCategoriesHandler(ctx *gin.Context) {
	resp, err := productService.GetCategories()
	if err != nil {
		utils.InternalServerError(
			ctx,
			"Error fetching categories",
			err,
		)
		return
	}

	utils.SuccessResponse(ctx, "Categories fetched successfully", resp)
}

func CreateCategoryHandler(ctx *gin.Context) {
	var req *dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(ctx, "invalid request", err)
		return
	}

	resp, err := productService.CreateCategroy(req)
	if err != nil {
		utils.InternalServerError(
			ctx,
			"Error Creating Category",
			err,
		)
	}

	utils.SuccessResponse(ctx, "Category created successfully", resp)
}

func UpdateCategoryHandler(ctx *gin.Context) {
	cat_id, exists := ctx.GetQuery("cat_id")
	if !exists {
		utils.BadRequestError(
			ctx,
			"missing query param",
			errors.New("missing query param - id"),
		)
		return
	}

	categoryID, err := strconv.ParseUint(cat_id, 10, 0)
	if err != nil {
		utils.BadRequestError(ctx, "invalid category id", err)
		return
	}

	var req *dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(
			ctx,
			"Invalid Request",
			err,
		)
		return
	}

	resp, err := productService.UpdateCategory(uint(categoryID), req)
	if err != nil {
		utils.InternalServerError(
			ctx,
			"Error updating category",
			err,
		)
	}

	utils.SuccessResponse(ctx, "updated category successfully", resp)
}
