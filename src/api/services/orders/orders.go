package orders

import (
    ordersDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/orders"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
)

func GetOrderByID(orderID int64) (*ordersDomain.Order, *apierrors.ApiError) {
    return &ordersDomain.Order{
        OrderID: orderID,
    }, nil
}
