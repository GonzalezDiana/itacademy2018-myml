package items

import (
    itemsService "github.com/emikohmann/itacademy2018-myml/src/api/services/items"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetItemByID(c *gin.Context) {
    itemID, err := strconv.ParseInt(c.Param("itemID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, apierrors.ApiError{
            StatusCode: http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

    item, apiErr := itemsService.GetItemByID(itemID)
    if apiErr != nil {
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, item)
}
