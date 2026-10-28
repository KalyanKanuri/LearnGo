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
		productService.Log.Error().Str("handler", "get_categories").Err(err).Msg("failed to fetch categories")
		utils.InternalServerError(
			ctx,
			"Error fetching categories",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "get_categories").Int("count", len(resp)).Msg("categories fetched successfully")
	utils.SuccessResponse(ctx, "Categories fetched successfully", resp)
}

func CreateCategoryHandler(ctx *gin.Context) {
	var req *dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		productService.Log.Error().Str("handler", "create_category").Err(err).Msg("create category request binding failed")
		utils.BadRequestError(ctx, "invalid request", err)
		return
	}

	resp, err := productService.CreateCategory(req)
	if err != nil {
		productService.Log.Error().Str("handler", "create_category").Str("name", req.Name).Err(err).Msg("failed to create category")
		utils.InternalServerError(
			ctx,
			"Error Creating Category",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "create_category").Uint("category_id", resp.ID).Msg("category created successfully")
	utils.SuccessResponse(ctx, "Category created successfully", resp)
}

func UpdateCategoryHandler(ctx *gin.Context) {
	cat_id, exists := ctx.GetQuery("cat_id")
	if !exists {
		productService.Log.Error().Str("handler", "update_category").Msg("missing cat_id query param")
		utils.BadRequestError(
			ctx,
			"missing query param",
			errors.New("missing query param - id"),
		)
		return
	}

	categoryID, err := strconv.ParseUint(cat_id, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "update_category").Str("cat_id", cat_id).Err(err).Msg("invalid category id")
		utils.BadRequestError(ctx, "invalid category id", err)
		return
	}

	var req *dto.UpdateCategoryRequest
	if bindErr := ctx.ShouldBindJSON(&req); bindErr != nil {
		productService.Log.Error().Str("handler", "update_category").Uint("category_id", uint(categoryID)).Err(bindErr).Msg("update category request binding failed")
		utils.BadRequestError(
			ctx,
			"Invalid Request",
			bindErr,
		)
		return
	}

	resp, err := productService.UpdateCategory(uint(categoryID), req)
	if err != nil {
		productService.Log.Error().Str("handler", "update_category").Uint("category_id", uint(categoryID)).Err(err).Msg("failed to update category")
		utils.InternalServerError(
			ctx,
			"Error updating category",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "update_category").Uint("category_id", uint(categoryID)).Msg("category updated successfully")
	utils.SuccessResponse(ctx, "updated category successfully", resp)
}

func DeleteCategoryHandler(ctx *gin.Context) {
	cat_id, exists := ctx.GetQuery("cat_id")
	if !exists {
		productService.Log.Error().Str("handler", "delete_category").Msg("missing cat_id query param")
		utils.BadRequestError(
			ctx,
			"missing category id in request params",
			errors.New("missing category id in request params"),
		)
		return
	}

	categoryID, err := strconv.ParseUint(cat_id, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "delete_category").Str("cat_id", cat_id).Err(err).Msg("invalid category id")
		utils.InternalServerError(
			ctx,
			"error converting string to uint",
			err,
		)
		return
	}

	isDeleted, err := productService.DeleteCategory(uint(categoryID))
	if err != nil {
		productService.Log.Error().Str("handler", "delete_category").Uint("category_id", uint(categoryID)).Err(err).Msg("failed to delete category")
		utils.InternalServerError(
			ctx,
			"error deleting category",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "delete_category").Uint("category_id", uint(categoryID)).Msg("category deleted successfully")
	utils.SuccessResponse(ctx, "Category deleted successfully", isDeleted)
}

func AddProductHandler(ctx *gin.Context) {
	var req *dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		productService.Log.Error().Str("handler", "add_product").Err(err).Msg("add product request binding failed")
		utils.BadRequestError(
			ctx,
			"Invalid Request",
			err,
		)
		return
	}

	resp, err := productService.AddProduct(req)
	if err != nil {
		productService.Log.Error().Str("handler", "add_product").Str("sku", req.SKU).Err(err).Msg("failed to create product")
		utils.InternalServerError(
			ctx,
			"error creating product",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "add_product").Uint("product_id", resp.ID).Msg("product created successfully")
	utils.SuccessResponse(ctx, "Product created successfully", resp)
}

func UpdateProductHandler(ctx *gin.Context) {
	prod_id, exists := ctx.GetQuery("prod_id")
	if !exists {
		productService.Log.Error().Str("handler", "update_product").Msg("missing prod_id query param")
		utils.BadRequestError(
			ctx,
			"Invalid Request",
			errors.New("product_id missing in query params"),
		)
		return
	}

	productID, err := strconv.ParseUint(prod_id, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "update_product").Str("prod_id", prod_id).Err(err).Msg("invalid product id")
		utils.InternalServerError(
			ctx,
			"error converting query params",
			err,
		)
		return
	}

	var req *dto.UpdateProductRequest
	if bindErr := ctx.ShouldBindJSON(&req); bindErr != nil {
		productService.Log.Error().Str("handler", "update_product").Uint("product_id", uint(productID)).Err(bindErr).Msg("update product request binding failed")
		utils.BadRequestError(
			ctx,
			"invalid request",
			bindErr,
		)
		return
	}

	resp, err := productService.UpdateProduct(uint(productID), req)
	if err != nil {
		productService.Log.Error().Str("handler", "update_product").Uint("product_id", uint(productID)).Err(err).Msg("failed to update product")
		utils.InternalServerError(
			ctx,
			"error updating product",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "update_product").Uint("product_id", uint(productID)).Msg("product updated successfully")
	utils.SuccessResponse(ctx, "updated product successfully", resp)
}

func GetProductHandler(ctx *gin.Context) {
	prod_id, exists := ctx.GetQuery("prod_id")
	if !exists {
		productService.Log.Error().Str("handler", "get_product").Msg("missing prod_id query param")
		utils.BadRequestError(
			ctx,
			"Invalid Request",
			errors.New("product_id missing in query params"),
		)
		return
	}

	productID, err := strconv.ParseUint(prod_id, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "get_product").Str("prod_id", prod_id).Err(err).Msg("invalid product id")
		utils.InternalServerError(
			ctx,
			"error converting query params",
			err,
		)
		return
	}

	resp, err := productService.GetProduct(uint(productID))
	if err != nil {
		productService.Log.Error().Str("handler", "get_product").Uint("product_id", uint(productID)).Err(err).Msg("failed to fetch product")
		utils.InternalServerError(
			ctx,
			"error fetching product",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "get_product").Uint("product_id", uint(productID)).Msg("product fetched successfully")
	utils.SuccessResponse(ctx, "product fetched successfully", resp)
}

func GetProductsHandler(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageNum, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "get_products").Str("page", page).Err(err).Msg("invalid page query param")
		utils.InternalServerError(
			ctx,
			"error converting query params",
			err,
		)
		return
	}

	queryLimit, err := strconv.ParseInt(limit, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "get_products").Str("limit", limit).Err(err).Msg("invalid limit query param")
		utils.InternalServerError(
			ctx,
			"error converting query params",
			err,
		)
		return
	}

	resp, paginationInfo, err := productService.GetProducts(int(pageNum), int(queryLimit))
	if err != nil {
		productService.Log.Error().Str("handler", "get_products").Int64("page", pageNum).Int64("limit", queryLimit).Err(err).Msg("failed to fetch products")
		utils.InternalServerError(
			ctx,
			"error fetching products",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "get_products").Int64("page", pageNum).Int64("limit", queryLimit).Int("count", len(resp)).Msg("products fetched successfully")
	utils.PaginatedResponse(ctx, "products fetched successfully", resp, *paginationInfo)
}

func DeleteProductHandler(ctx *gin.Context) {
	prod_id, exists := ctx.GetQuery("prod_id")
	if !exists {
		productService.Log.Error().Str("handler", "delete_product").Msg("missing prod_id query param")
		utils.BadRequestError(
			ctx,
			"Invalid Request",
			errors.New("product_id missing in query params"),
		)
		return
	}

	productID, err := strconv.ParseUint(prod_id, 10, 0)
	if err != nil {
		productService.Log.Error().Str("handler", "delete_product").Str("prod_id", prod_id).Err(err).Msg("invalid product id")
		utils.InternalServerError(
			ctx,
			"error converting query params",
			err,
		)
		return
	}

	resp, err := productService.DeleteProduct(uint(productID))
	if err != nil {
		productService.Log.Error().Str("handler", "delete_product").Uint("product_id", uint(productID)).Err(err).Msg("failed to delete product")
		utils.InternalServerError(
			ctx,
			"error fetching product",
			err,
		)
		return
	}

	productService.Log.Info().Str("handler", "delete_product").Uint("product_id", uint(productID)).Msg("product deleted successfully")
	utils.SuccessResponse(ctx, "product deletedS successfully", resp)
}
