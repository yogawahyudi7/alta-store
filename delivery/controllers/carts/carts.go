package carts

import (
	"net/http"
	"project-e-commerces/delivery/common"
	"project-e-commerces/entities"
	"project-e-commerces/repository/carts"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CartsController struct {
	Repo carts.CartInterface
}

func NewCartsControllers(crrep carts.CartInterface) *CartsController {
	return &CartsController{Repo: crrep}
}

func (crrep CartsController) Gets() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		cartID := int(claims["userid"].(float64))

		if res, err := crrep.Repo.Get(uint(cartID)); err != nil || len(res) == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code":    200,
				"message": "Successful Operation",
				"data":    res,
			})
		}
	}
}

func (crrep CartsController) PutItemIntoDetail_CartCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		newItemReq := AddItemIntoDetail_CartReqeuestFormat{}

		if err := c.Bind(&newItemReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		cartID := int(claims["userid"].(float64))

		newItem := entities.Detail_cart{
			CartID:     uint(cartID),
			ProductID:  uint(newItemReq.ProductID),
			Qty:        newItemReq.Qty,
			Price:      newItemReq.ProductPrice,
			TotalPrice: newItemReq.ProductPrice * newItemReq.Qty,
		}

		if res, err := crrep.Repo.InsertProduct(newItem); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}

	}
}

func (crrep CartsController) DeleteItemFromDetail_CartCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		delItemReq := DeleteItemIntoDetail_CartReqeuestFormat{}
		if err := c.Bind(&delItemReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		cartID := int(claims["userid"].(float64))

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
