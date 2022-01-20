package transactions

import (
	"net/http"
	"project-e-commerces/delivery/common"
	"project-e-commerces/entities"
	"project-e-commerces/repository/transactions"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionsController struct {
	Repo transactions.TransactionInterface
}

func NewTransactionsControllers(tsrep transactions.TransactionInterface) *TransactionsController {
	return &TransactionsController{Repo: tsrep}
}

func (trrep TransactionsController) PostProductTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		newPTransaction := ProductDetail_TransactionReqeuestFormat{}
		if err := c.Bind(&newPTransaction); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		paymentMidtrans := entities.Transaction{
			ID:          uint(userID),
			Total_price: newPTransaction.Product_price,
		}

		if res, err := trrep.Repo.GetPaymentURL(paymentMidtrans, uint(userID)); err != nil {

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": res,
			})
		} else {
			newItem := entities.Transaction{
				User_id:     uint(userID),
				Total_price: newPTransaction.Product_price,
				Total_qty:   newPTransaction.Product_qty,
				Status:      "PENDING",
				Url:         res,
			}

			if res, err := trrep.Repo.Insert(newItem); err != nil || res.ID == 0 {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    500,
					"message": "Internal Server Error",
					"data":    res,
				})

			}
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())

	}
}
