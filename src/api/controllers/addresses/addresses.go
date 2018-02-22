package addresses

import (
    addressesService "github.com/emikohmann/itacademy2018-myml/src/api/services/addresses"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetAddressByID(c *gin.Context) {
    addressID, err := strconv.ParseInt(c.Param("addressID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, apierrors.ApiError{
            StatusCode: http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

    address, apiErr := addressesService.GetAddressByID(addressID)
    if apiErr != nil {
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, address)
}
