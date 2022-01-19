package carts

type AddItemIntoDetail_CartReqeuestFormat struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Qty       int  `json:"qty" form:"qty"`
}
type DeleteItemIntoDetail_CartReqeuestFormat struct {
	ProductID uint `json:"product_id" form:"product_id"`
}
