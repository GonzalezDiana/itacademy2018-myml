package payments

import (
    paymentsDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/payments"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
)

func GetPaymentByID(paymentID int64) (*paymentsDomain.Payment, *apierrors.ApiError) {
    return &paymentsDomain.Payment{
        PaymentID: paymentID,
    }, nil
}
