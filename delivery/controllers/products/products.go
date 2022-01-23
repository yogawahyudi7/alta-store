package products

import (
	"fmt"
	"net/http"
	"project-e-commerces/delivery/common"
	"project-e-commerces/delivery/pagination"
	"project-e-commerces/entities"
	"project-e-commerces/repository/products"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Repo products.ProductInterface
}

func NewProductControllers(cc products.ProductInterface) *ProductController {
	return &ProductController{Repo: cc}
}

func (c *ProductController) GetAllProduct(e echo.Context) error {
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	page, _ := strconv.Atoi(e.QueryParam("page"))

	Pagination := pagination.ProductPagination{Limit: limit, Page: page}

	operationResult, totalPage, err := c.Repo.ProductPagination(Pagination)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	data := operationResult.(pagination.ProductPagination)

	url := e.Request().URL

	data.FirstPage = fmt.Sprint(url.Path, "?limit=", Pagination.Limit, "&page=", 0)
	data.LastPage = fmt.Sprint(url.Path, "?limit=", Pagination.Limit, "&page=", totalPage)

	if data.Page > 0 {
		data.PreviousPage = fmt.Sprint(url.Path, "?limit=", Pagination.Limit, "&page=", data.Page-1)
	}

	if data.Page < totalPage {
		data.NextPage = fmt.Sprint(url.Path, "?limit=", Pagination.Limit, "&page=", data.Page+1)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Successful Operation",
		"data":    data,
	})
}

func (c *ProductController) GetProductByID(e echo.Context) error {
	product_id, _ := strconv.Atoi(e.Param("id"))

	res, err := c.Repo.GetProductByID(product_id)

	if err != nil || res.ID == 0 {
		return e.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Successful Operation",
		"data":    res,
	})
}

func (c *ProductController) CreateProduct(e echo.Context) error {
	input := CreateProductRequest{}

	e.Bind(&input)

	product := entities.Product{}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.Category_id = uint(input.Category_id)

	res, err := c.Repo.CreateProduct(product)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Successful Operation",
		"data":    res,
	})
}

func (c *ProductController) UpdateProduct(e echo.Context) error {
	input := UpdateProductRequest{}

	product := entities.Product{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	e.Bind(&input)

	product.Name = input.Name
	product.Price = input.Price
	product.Category_id = uint(input.Category_id)

	_, err := c.Repo.UpdateProduct(product_id, product)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (c *ProductController) UpdateStockProduct(e echo.Context) error {
	input := UpdateStockProductRequest{}

	product_id, _ := strconv.Atoi(e.Param("id"))

	e.Bind(&input)

	_, err := c.Repo.UpdateStockProduct(product_id, input.Qty)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (c *ProductController) GetHistoryStockProduct(e echo.Context) error {
	product_id, _ := strconv.Atoi(e.Param("id"))

	res, err := c.Repo.GetHistoryStockProduct(product_id)

	if err != nil || res[0].Product_id == 0 {
		return e.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Successful Operation",
		"data":    res,
	})
}

func (c *ProductController) DeleteProduct(e echo.Context) error {
	product_id, _ := strconv.Atoi(e.Param("id"))

	_, err := c.Repo.DeleteProduct(product_id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (c *ProductController) ExportPDF(e echo.Context) error {
	products, err := c.Repo.GetAllProduct()

	if err != nil || len(products) == 0 {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	contents := [][]string{}

	for i := 0; i < len(products); i++ {
		temp := []string{}
		ID := strconv.Itoa(int(products[i].ID))
		Stock := strconv.Itoa(products[i].Stock)
		Price := strconv.Itoa(products[i].Price)
		CategoryID := strconv.Itoa(int(products[i].Category_id))
		temp = append(temp, ID)
		temp = append(temp, products[i].Name)
		temp = append(temp, Stock)
		temp = append(temp, Price)
		temp = append(temp, CategoryID)

		contents = append(contents, temp)
	}

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 20, 20)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("Tabel List Produk", props.Text{
					Top:    2,
					Size:   14,
					Align:  consts.Center,
					Family: consts.Arial,
				})
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	tableHeadings := []string{"ID Product", "Nama Product", "Stock", "Price", "Category ID"}

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			Style:     consts.Bold,
			GridSizes: []uint{3, 3, 2, 2, 2},
		},

		ContentProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{3, 3, 2, 2, 2},
		},
		Align:                consts.Center,
		AlternatedBackground: &color.Color{Red: 230, Blue: 230, Green: 230},
		HeaderContentSpace:   1,
		Line:                 true,
	})

	err = m.OutputFileAndClose("./hasil-export/list-product.pdf")

	if err != nil {
		return e.JSON(http.StatusBadRequest, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
