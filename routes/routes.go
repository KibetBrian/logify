package routes

import (
	"github.com/gofiber/fiber/v2"
	"logify/handlers"
)

const rootUrl = "/api/v1/"

func Routes(app *fiber.App) {

	//Auth
	app.Post(rootUrl+"auth/register", handlers.Register)
	app.Post(rootUrl+"auth/login", handlers.Login)
	app.Post(rootUrl+"auth/forgot", handlers.Forgot)
	app.Post(rootUrl+"auth/reset/:token", handlers.Reset)
	app.Get(rootUrl+"auth/logout", handlers.Logout)

	//user
	app.Put(rootUrl+"user/update", handlers.UpdateInfo)
	app.Get(rootUrl+"user/all", handlers.AllUsers)

	//Product
	app.Get(rootUrl+"products/all", handlers.AllProducts)
	app.Get(rootUrl+"products/latest", handlers.GetLatestProducts)
	app.Get(rootUrl+"products/", handlers.GetBrandProducts)
	app.Post(rootUrl+"products/create", handlers.CreateProduct)
	app.Put(rootUrl+"products/update", handlers.UpdateProduct)
	app.Delete(rootUrl+"products/delete", handlers.DeleteProduct)

	//Orders
	app.Get(rootUrl+"orders/all", handlers.AllOrders)
	app.Get(rootUrl+"orders/pending", handlers.AllOrders)
	app.Get(rootUrl+"orders/transit", handlers.GetTransitOrders)
	app.Get(rootUrl+"orders/delivered", handlers.GetDeliveredOrders)
	app.Put(rootUrl+"orders/update", handlers.UpdateOrderStatus)
	app.Post(rootUrl+"orders/create", handlers.CreateOrder)
	app.Post(rootUrl+"orders/clear", handlers.ClearOrder)

}
