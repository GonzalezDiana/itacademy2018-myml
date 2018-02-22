package items

import (
    itemsDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/items"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
)

func GetItemByID(itemID int64) (*itemsDomain.Item, *apierrors.ApiError) {
    return &itemsDomain.Item{
        ItemID: itemID,
    }, nil
}
