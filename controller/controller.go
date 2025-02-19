package controller

import (
	"crud_fire/model"
	"crud_fire/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type productController struct {
	productService service.ProducService
}

func NewProductController(product service.ProducService) *productController {
	return &productController{productService: product}

}

func (c *productController) GetAllProducts(ctx echo.Context) error {
	// get data
	products, err := c.productService.GetAllProducts(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, products)
}

func (c *productController) InsertProduct(ctx echo.Context) error {
	var product model.Product

	// binding data
	err := ctx.Bind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// validasi input
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// insert data
	err = c.productService.InsertProduct(ctx.Request().Context(), product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "success insert data")
}

func (c *productController) EditProduct(ctx echo.Context) error {
	var product model.Product

	// binding data
	err := ctx.Bind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// validasi input
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// edit data
	err = c.productService.EditProduct(ctx.Request().Context(), product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "success edit data")
}

func (c *productController) DeleteProduct(ctx echo.Context) error {
	var product model.Product

	// binding data
	err := ctx.Bind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// validasi input
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// delete data
	err = c.productService.DeleteProduct(ctx.Request().Context(), product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success delete data")
}
