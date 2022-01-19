package products

import (
	"net/http"
	"project-e-commerces/entities"
	"project-e-commerces/repository/products"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Repo products.ProductInterface
}

func NewProductControllers(cc products.ProductInterface) *ProductController {
	return &ProductController{Repo: cc}
}

func (c *ProductController) GetAllProduct(e echo.Context) error {
	res := GetProductResponseFormat{}

	products, err := c.Repo.GetAllProduct()

	if err != nil || len(products) == 0 {
		res.Message = "no products data found"
		res.Data = products

		return e.JSON(http.StatusNotFound, res)
	}

	res.Message = "success get all product"
	res.Data = append(res.Data, products...)

	return e.JSON(http.StatusOK, res)
}

func (c *ProductController) GetProductByID(e echo.Context) error {
	res := GetProductResponseFormat{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	// product_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	product, err := c.Repo.GetProductByID(product_id)

	if err != nil || product.ID == 0 {
		res.Message = "product not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	res.Message = "success get product"
	res.Data = append(res.Data, product)

	return e.JSON(http.StatusOK, res)
}

func (c *ProductController) CreateProduct(e echo.Context) error {
	res := GetProductResponseFormat{}

	input := CreateProductRequest{}

	e.Bind(&input)

	// err := e.Bind(&input)

	// if err != nil {
	// 	res.Message = "can't get the input"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	// err = e.Validate(&input)

	// if err != nil {
	// 	res.Message = "there is data that has not been filled in"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	product := entities.Product{}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.Category_id = uint(input.Category_id)

	newProduct, err := c.Repo.CreateProduct(product)

	if err != nil {
		res.Data = nil
		res.Message = "can't create product"

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success create product"
	res.Data = append(res.Data, newProduct)

	return e.JSON(http.StatusOK, res)
}

func (c *ProductController) UpdateProduct(e echo.Context) error {
	res := GetProductResponseFormat{}

	input := UpdateProductRequest{}

	product := entities.Product{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	e.Bind(&input)
	// product_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	// err = e.Validate(&input)

	// if err != nil {
	// 	res.Message = "there is data that has not been filled in"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	product.Name = input.Name
	product.Price = input.Price
	product.Category_id = uint(input.Category_id)

	updatedProduct, err := c.Repo.UpdateProduct(product_id, product)

	if err != nil {
		res.Message = "can't update product"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success update product"
	res.Data = append(res.Data, updatedProduct)

	return e.JSON(http.StatusOK, res)
}

func (c *ProductController) UpdateStockProduct(e echo.Context) error {
	res := GetProductResponseFormat{}

	input := UpdateStockProductRequest{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	e.Bind(&input)

	// product_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	// err = e.Bind(&input)

	// if err != nil {
	// 	res.Message = "can't get the data"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	// err = e.Validate(&input)

	// if err != nil {
	// 	res.Message = "there is data that has not been filled in"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	updatedStockProduct, err := c.Repo.UpdateStockProduct(product_id, input.Qty)

	if err != nil {
		res.Message = "can't update stock product"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success update stock product"
	res.Data = append(res.Data, updatedStockProduct)

	return e.JSON(http.StatusOK, res)
}

func (c *ProductController) GetHistoryStockProduct(e echo.Context) error {
	res := GetStockProductResponseFormat{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	product, err := c.Repo.GetHistoryStockProduct(product_id)

	if err != nil || product[0].Product_id == 0 {
		res.Message = "history product not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	res.Message = "success get history product"
	res.Data = append(res.Data, product...)

	return e.JSON(http.StatusOK, res)
}

func (c *ProductController) DeleteProduct(e echo.Context) error {
	res := GetProductResponseFormat{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	// product_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	_, err := c.Repo.DeleteProduct(product_id)

	if err != nil {
		res.Message = "can't delete product"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success delete product"
	res.Data = nil

	return e.JSON(http.StatusOK, res)
}

// func (c *ProductController) Pagination(e echo.Context) error {
// 	input := ProductPagination{}

// 	e.Bind(&input)

// 	limit, _ := strconv.Atoi(e.QueryParam("limit"))

// 	return e.JSON(http.StatusOK, map[string]interface{}{"Data": limit})
// }
