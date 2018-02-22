package orders

import (
    ordersDomain "github.com/mercadolibre/itacademy-myml/src/api/domain/orders"
    "github.com/mercadolibre/itacademy-myml/src/api/util/apierrors"
)

func GetOrderByID(orderID int64) (*ordersDomain.Order, *apierrors.ApiError) {
    return &ordersDomain.Order{
        OrderID: orderID,
    }, nil
}
