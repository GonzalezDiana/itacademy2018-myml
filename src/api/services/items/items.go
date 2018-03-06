package items

import (
    "fmt"
    itemsDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/items"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
)

func GetItemByID(itemID int64) (*itemsDomain.Item, *apierrors.ApiError) {
    return &itemsDomain.Item{
        Id:         itemID,
        SiteId:     "MLM",
        Title:      fmt.Sprintf("Item-%d", itemID),
        CategoryId: fmt.Sprintf("%s%d", "MLM", itemID),
        Price:      100,
        CurrencyId: "MXN",
        Quantity:   20,
        Condition:  "new",
        Pictures:   []string{fmt.Sprintf("img-11%d", itemID), fmt.Sprintf("img-22%d", itemID)},
    }, nil
}
