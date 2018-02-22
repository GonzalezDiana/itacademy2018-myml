package items

import (
    itemsDomain "github.com/mercadolibre/itacademy-myml/src/api/domain/items"
    "github.com/mercadolibre/itacademy-myml/src/api/util/apierrors"
)

func GetItemByID(itemID int64) (*itemsDomain.Item, *apierrors.ApiError) {
    return &itemsDomain.Item{
        ItemID: itemID,
    }, nil
}