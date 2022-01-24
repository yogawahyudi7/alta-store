package carts

type GetsDetail_CartRequestFormat struct {
	CartID uint `json:"cart_id" form:"cart_id"`
}

type AddItemIntoDetail_CartReqeuestFormat struct {
	ProductID    uint `json:"product_id" form:"product_id"`
	ProductPrice int  `json:"product_price" form:"product_price"`
	Qty          int  `json:"qty" form:"qty"`
}
type DeleteItemIntoDetail_CartReqeuestFormat struct {
	ProductID uint `json:"product_id" form:"product_id"`
}
