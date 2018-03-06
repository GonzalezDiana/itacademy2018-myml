package addresses

import (
    addressesDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/addresses"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
)

func GetAddressByID(addressID int64) (*addressesDomain.Address, *apierrors.ApiError) {
    return &addressesDomain.Address{
        Id:           addressID,
        StreetName:   "Calle Falsa",
        StreetNumber: 123,
        City:         "Springfield",
    }, nil
}
