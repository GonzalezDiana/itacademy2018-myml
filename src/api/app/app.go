package app

import (
    "github.com/gin-gonic/gin"
    addressesController "github.com/mercadolibre/itacademy-myml/src/api/controllers/addresses"
    itemsController "github.com/mercadolibre/itacademy-myml/src/api/controllers/items"
    ordersController "github.com/mercadolibre/itacademy-myml/src/api/controllers/orders"
    paymentsController "github.com/mercadolibre/itacademy-myml/src/api/controllers/payments"
    pingController "github.com/mercadolibre/itacademy-myml/src/api/controllers/ping"
)

func StartApp() {
    router := gin.Default()
    router.GET("/ping", pingController.Ping)
    router.GET("/orders/:orderID", ordersController.GetOrderByID)
    router.GET("/items/:itemID", itemsController.GetItemByID)
    router.GET("/payments/:paymentID", paymentsController.GetPaymentByID)
    router.GET("/addresses/:addressID", addressesController.GetAddressByID)
    router.Run(":8080")
}
