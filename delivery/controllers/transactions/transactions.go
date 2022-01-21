package transactions

import (
	"fmt"
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

func (trrep TransactionsController) PostProductIntoTransactionCtrl() echo.HandlerFunc {
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
			newTransaction := entities.Transaction{
				User_id:     uint(userID),
				Total_price: newPTransaction.Product_price,
				Total_qty:   newPTransaction.Product_qty,
				Status:      "PENDING",
				Url:         res,
			}

			if res, err := trrep.Repo.InsertT(newTransaction); err != nil || res.ID == 0 {

				newDetailTransaction := entities.Detail_transaction{
					Transaction_id: res.ID,
					Product_id:     newPTransaction.ProductID,
					Product_qty:    newPTransaction.Product_qty,
					Price:          newPTransaction.Product_price,
				}

				if res2, _ := trrep.Repo.InsertDT(newDetailTransaction); err != nil || res2.ID == 0 {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"code":    500,
						"message": "Internal Server Error",
					})

				}

				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    500,
					"message": "Internal Server Error",
				})

			} else {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"code":    200,
					"message": "Successful Operation",
					"data":    res,
				})
			}
		}

	}
}
func (trrep TransactionsController) PostCartIntoTransactionCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		cartID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		newCTransaction := CartDetail_TransactionReqeuestFormat{}

		if err := c.Bind(&newCTransaction); err != nil {
			fmt.Println("anu", err)
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		fmt.Println("cartID", cartID)
		fmt.Println("productlist", newCTransaction)
		Totalprice := 0
		TotalQty := 0
		for i := 0; i < len(newCTransaction.Products); i++ {
			Totalprice += newCTransaction.Products[i].Product_price
			TotalQty += newCTransaction.Products[i].Product_qty
		}

		fmt.Println("totalprice", Totalprice)
		fmt.Println("totalqty", TotalQty)

		// paymentMidtrans := entities.Transaction{
		// 	ID:          uint(userID),
		// 	Total_price: newPTransaction.Product_price,
		// }

		// if res, err := trrep.Repo.GetPaymentURL(paymentMidtrans, uint(userID)); err != nil {

		// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 		"message": res,
		// 	})
		// } else {
		// 	newTransaction := entities.Transaction{
		// 		User_id:     uint(userID),
		// 		Total_price: newPTransaction.Product_price,
		// 		Total_qty:   newPTransaction.Product_qty,
		// 		Status:      "PENDING",
		// 		Url:         res,
		// 	}

		// 	if res, err := trrep.Repo.InsertT(newTransaction); err != nil || res.ID == 0 {

		// 		newDetailTransaction := entities.Detail_transaction{
		// 			Transaction_id: res.ID,
		// 			Product_id:     newPTransaction.ProductID,
		// 			Product_qty:    newPTransaction.Product_qty,
		// 			Price:          newPTransaction.Product_price,
		// 		}

		// 		if res2, _ := trrep.Repo.InsertDT(newDetailTransaction); err != nil || res2.ID == 0 {
		// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 				"code":    500,
		// 				"message": "Internal Server Error",
		// 			})

		// 		}

		// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 			"code":    500,
		// 			"message": "Internal Server Error",
		// 		})

		// 	} else {
		// 		return c.JSON(http.StatusOK, map[string]interface{}{
		// 			"code":    200,
		// 			"message": "Successful Operation",
		// 			"data":    res,
		// 		})
		// 	}
		// }

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Successful Operation",
			// "data":    res,
		})
	}
}
