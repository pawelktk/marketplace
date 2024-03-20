package front

import (
	"marketplace/cmd/models"
	"marketplace/cmd/repositories"
	"strconv"

	"github.com/labstack/echo/v4"

	"net/http"
)

func MainHandler(c echo.Context) error {
	products, _ := repositories.GetAllProducts()
	return c.Render(http.StatusOK, "index.html", products)
}

func HandleProductAdd(c echo.Context) error {
	newProduct := models.Product{}
	c.Bind(&newProduct)
	repositories.CreateProduct(newProduct)
	products, _ := repositories.GetAllProducts()

	return c.Render(http.StatusOK, "products.html", products)
}
func HandleProductUpdate(c echo.Context) error {
	newProduct := models.Product{}
	id, _ := strconv.Atoi(c.Param("id"))

	c.Bind(&newProduct)
	newProduct, _ = repositories.UpdateProduct(newProduct, id)
	return c.Render(http.StatusOK, "row.html", newProduct)
}

func AddProduct(c echo.Context) error {
	products, _ := repositories.GetAllProducts()

	return c.Render(http.StatusOK, "product-add.html", products)
}

func EditProduct(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	selectedProduct, _ := repositories.GetProduct(id)
	return c.Render(http.StatusOK, "row-edit.html", selectedProduct)
}

func CancelAdd(c echo.Context) error {
	products, _ := repositories.GetAllProducts()
	return c.Render(http.StatusOK, "products.html", products)
}

func CancelEdit(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	row, _ := repositories.GetProduct(id)

	return c.Render(http.StatusOK, "row.html", row)
}
