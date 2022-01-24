package routes

import (
	"project-e-commerces/delivery/controllers/carts"
	"project-e-commerces/delivery/controllers/categorys"
	"project-e-commerces/delivery/controllers/products"
	"project-e-commerces/delivery/controllers/transactions"

	"project-e-commerces/constants"
	controllers "project-e-commerces/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserPath(e *echo.Echo, uc controllers.UserController) {
	// e.Pre(middleware.RemoveTrailingSlash())
	// auth := e.Group("")
	// auth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	//REGISTER & LOGIN
	e.POST("/users/register", uc.Register)
	e.POST("/users/login", uc.Login)

	//RUD USER
	e.GET("/users/profile", uc.Get(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.DELETE("/users/delete", uc.Delete(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.PUT("/users/update", uc.Update(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

}

func RegisterPath(e *echo.Echo, crCtrl *carts.CartsController, tsCtrl *transactions.TransactionsController, cc *categorys.CategoryController, pc *products.ProductController) {

	// ---------------------------------------------------------------------
	// CRUD Carts
	// ---------------------------------------------------------------------
	e.PUT("/carts/additem", crCtrl.PutItemIntoDetail_CartCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.GET("/carts", crCtrl.Gets(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.DELETE("/carts/delitem", crCtrl.DeleteItemFromDetail_CartCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.POST("/transactions/live", tsCtrl.PostProductsIntoTransactionCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.POST("/transactions/status", tsCtrl.GetStatus())
	// e.POST("/transactions/cart", tsCtrl.PostCartIntoTransactionCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	e.GET("/categorys", cc.GetAllCategory)
	e.GET("/categorys/:id", cc.GetCategoryByID)
	e.POST("/categorys", cc.CreateCategory)
	e.PUT("/categorys/:id", cc.UpdateCategory)
	e.DELETE("/categorys/:id", cc.DeleteCategory)

	e.GET("/products", pc.GetAllProduct)
	e.GET("/products/:id", pc.GetProductByID)
	e.GET("/products/stocks/:id", pc.GetHistoryStockProduct)
	e.POST("/products", pc.CreateProduct)
	e.POST("/products/stocks/:id", pc.UpdateStockProduct)
	e.PUT("/products/:id", pc.UpdateProduct)
	e.DELETE("/products/:id", pc.DeleteProduct)
}
