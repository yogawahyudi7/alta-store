package transactions

type ProductDetail_TransactionReqeuestFormat struct {
	ProductID     uint `json:"product_id" form:"product_id"`
	Product_qty   int  `json:"product_qty" form:"product_qty"`
	Product_price int  `json:"product_price" form:"product_price"`
}
