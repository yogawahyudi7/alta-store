package carts

import (
	"net/http"
	"project-e-commerces/delivery/common"
	"project-e-commerces/entities"
	"project-e-commerces/repository/carts"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartsController struct {
	Repo carts.CartInterface
}

func NewCartsControllers(crrep carts.CartInterface) *CartsController {
	return &CartsController{Repo: crrep}
}

func (crrep CartsController) PutItemIntoDetail_CartCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		cartID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		newItemReq := AddItemIntoDetail_CartReqeuestFormat{}

		if err := c.Bind(&newItemReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		newItem := entities.Detail_cart{
			CartID:    uint(cartID),
			ProductID: uint(newItemReq.ProductID),
			Qty:       newItemReq.Qty,
		}
		if res, err := crrep.Repo.InsertProduct(uint(cartID), newItem); err != nil || res.ID == 0 {
			if res2, err2 := crrep.Repo.UpdateProduct(newItem.CartID, newItem.ProductID, newItemReq.Qty); err2 != nil || res2.ID == 0 {
				return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
			}
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())

	}
}

func (crrep CartsController) DeleteItemFromDetail_CartCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		cartID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		delItemReq := DeleteItemIntoDetail_CartReqeuestFormat{}
		if err := c.Bind(&delItemReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		delItem := entities.Detail_cart{
			ProductID: delItemReq.ProductID,
		}

		if res, err := crrep.Repo.DeleteProduct(uint(cartID), delItem.ProductID); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}
	}
}
