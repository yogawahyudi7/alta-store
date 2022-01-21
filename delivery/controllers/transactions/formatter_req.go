package transactions

type ProductDetail_TransactionReqeuestFormat struct {
	ProductID     uint `json:"product_id" form:"product_id"`
	Product_qty   int  `json:"product_qty" form:"product_qty"`
	Product_price int  `json:"product_price" form:"product_price"`
}

type CartDetail_TransactionReqeuestFormat struct {
	Products []ProductList `json:"productlist"`
}

type ProductList struct {
	ProductID     interface{} `json:"product_id"`
	Product_qty   interface{} `json:"product_qty"`
	Product_price interface{} `json:"product_price"`
}
