package orders

import (
    ordersDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/orders"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
    "fmt"
)

func GetOrderByID(orderID int64) (*ordersDomain.Order, *apierrors.ApiError) {
    return &ordersDomain.Order{
        Id: orderID,
        DateCreated: "2018-03-06T15:00:00.000-04:00",
        Items: []string{fmt.Sprintf("111%d", orderID)},
        Address: fmt.Sprintf("222%d", orderID),
        TotalAmount: 100,
        Status: "paid",
        CurrencyId: "MXN",
        Tags: []string{"delivered","paid"},
        Buyer: fmt.Sprintf("BUYER-%d", orderID),
        Seller: fmt.Sprintf("SELLER-%d", orderID),
        Payments: []string{fmt.Sprintf("333%d", orderID)},
    }, nil
}
