package transactions

type Detail_TransactionReqeuestFormat struct {
	Products []ProductList `json:"productlist"`
}

type ProductList struct {
	ProductID     int `json:"product_id"`
	Product_qty   int `json:"product_qty"`
	Product_price int `json:"product_price"`
}
