package handlers

import (
	"database/sql"
	"errors"
	"marketplace/cmd/models"
	"marketplace/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ProductErrorHandler(c echo.Context, err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	case errors.Is(err, strconv.ErrRange) || errors.Is(err, strconv.ErrSyntax):
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Operation failed"})
	}
}

func CreateProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	newProduct, err := repositories.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newProduct)
}

func HandleUpdateProduct(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product := models.Product{}
	c.Bind(&product)
	updatedProduct, err := repositories.UpdateProduct(product, idInt)
	if err != nil {
		//return c.JSON(http.StatusInternalServerError, err.Error())
		return ProductErrorHandler(c, err)
	}
	return c.JSON(http.StatusOK, updatedProduct)
}

func HandleGetProduct(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product, err := repositories.GetProduct(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve product"})

	}

	return c.JSON(http.StatusOK, product)
}

func HandleGetAllProducts(c echo.Context) error {
	product, err := repositories.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func HandleDeleteProduct(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	err = repositories.DeleteProduct(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
