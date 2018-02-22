package payments

import (
    "github.com/gin-gonic/gin"
    paymentsService "github.com/mercadolibre/itacademy-myml/src/api/services/payments"
    "github.com/mercadolibre/itacademy-myml/src/api/util/apierrors"
    "net/http"
    "strconv"
)

func GetPaymentByID(c *gin.Context) {
    paymentID, err := strconv.ParseInt(c.Param("paymentID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, apierrors.ApiError{
            StatusCode: http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

    payment, apiErr := paymentsService.GetPaymentByID(paymentID)
    if apiErr != nil {
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, payment)
}
