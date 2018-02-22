package orders

import (
    ordersService "github.com/emikohmann/itacademy2018-myml/src/api/services/orders"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetOrderByID(c *gin.Context) {
    orderID, err := strconv.ParseInt(c.Param("orderID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, apierrors.ApiError{
            StatusCode: http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

    order, apiErr := ordersService.GetOrderByID(orderID)
    if apiErr != nil {
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, order)
}
