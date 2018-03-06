package payments

import (
    "fmt"
    paymentsDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/payments"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
    "net/http"
    "strconv"
)

func GetPaymentByID(paymentID int64) (*paymentsDomain.Payment, *apierrors.ApiError) {
    var paymentStr string
    paymentStr = fmt.Sprintf("%d", paymentID)
    if len(paymentStr) < 4 {
        return nil, &apierrors.ApiError{
            StatusCode: http.StatusBadRequest,
            Error:      "payment id too short",
        }
    }

    paymentStr = paymentStr[3:]
    orderId, err := strconv.Atoi(paymentStr)
    if err != nil {
        return nil, &apierrors.ApiError{
            StatusCode: http.StatusInternalServerError,
            Error:      err.Error(),
        }
    }

    return &paymentsDomain.Payment{
        Id:                paymentID,
        OrderId:           orderId,
        PaymentMethodId:   "account_money",
        CurrencyId:        "MXN",
        Status:            "approved",
        TransactionAmount: 100,
    }, nil
}
