package transactions

import (
	"fmt"
	"net/http"
	"project-e-commerces/delivery/common"
	"project-e-commerces/entities"
	"project-e-commerces/repository/transactions"
	"strconv"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var crc coreapi.Client

type TransactionsController struct {
	Repo transactions.TransactionInterface
}

func NewTransactionsControllers(tsrep transactions.TransactionInterface) *TransactionsController {
	return &TransactionsController{Repo: tsrep}
}

func (trrep TransactionsController) PostProductsIntoTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		newPTransaction := Detail_TransactionReqeuestFormat{}
		if err := c.Bind(&newPTransaction); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		totalprice := 0
		totalQty := 0
		for i := 0; i < len(newPTransaction.Products); i++ {
			totalprice += newPTransaction.Products[i].Product_price
			totalQty += newPTransaction.Products[i].Product_qty

		}
		invoiceID := uuid.New().String()
		if res, err := trrep.Repo.GetsPaymentUrl(uint(userID), totalprice, totalQty, invoiceID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal Server Error",
				"data":    err,
			})
		} else {
			newTransaction := entities.Transaction{
				User_id:     uint(userID),
				Total_price: totalprice,
				Total_qty:   totalQty,
				Status:      "PENDING",
				Url:         res,
				Invoice:     invoiceID,
				OrderID:     "INV-" + invoiceID + "/c/" + strconv.Itoa(int(userID)),
			}

			if res, _ := trrep.Repo.InsertT(newTransaction); err != nil {
				return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
			} else {
				i := 0
				for i != len(newPTransaction.Products) {
					newDetailTransaction := entities.Detail_transaction{
						Transaction_id: res.ID,
						Product_id:     uint(newPTransaction.Products[i].ProductID),
						Product_qty:    newPTransaction.Products[i].Product_qty,
						Price:          newPTransaction.Products[i].Product_price,
					}
					if res2, err := trrep.Repo.InsertDT(newDetailTransaction); err != nil || res2.ID == 0 {
						return c.JSON(http.StatusInternalServerError, map[string]interface{}{
							"code":    500,
							"message": "Internal Server Error",
							"data":    err,
						})
					}
					i++
				}

				return c.JSON(http.StatusOK, map[string]interface{}{
					"code":    200,
					"message": "Successful Operation",
					"data":    res.Url,
				})
			}

		}

	}
}

func (trrep TransactionsController) GetStatus() echo.HandlerFunc {
	return func(c echo.Context) error {

		midtrans.ServerKey = "SB-Mid-server-WBQoXNegZ5veTRfQsX3WOGFq"
		midtrans.ClientKey = "SB-Mid-client-lbfJ_9e_8nsyvWWS"
		midtrans.Environment = midtrans.Sandbox

		crc.New(midtrans.ServerKey, midtrans.Environment)

		var notificationPayload map[string]interface{}

		if err := c.Bind(&notificationPayload); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		orderID, exists := notificationPayload["order-id"].(string)
		if !exists {
			fmt.Println("not found")
		}

		fmt.Println("notification", notificationPayload)
		fmt.Println(orderID)

		tranStatusResp, e := crc.CheckTransaction(orderID)
		if e != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			if tranStatusResp != nil {

				if res, err := trrep.Repo.Update(tranStatusResp.TransactionStatus, 1); err != nil {
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

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())

	}
}
