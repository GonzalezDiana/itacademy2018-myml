package addresses

import (
    addressesDomain "github.com/mercadolibre/itacademy-myml/src/api/domain/addresses"
    "github.com/mercadolibre/itacademy-myml/src/api/util/apierrors"
)

func GetAddressByID(addressID int64) (*addressesDomain.Address, *apierrors.ApiError) {
    return &addressesDomain.Address{
        AddressID: addressID,
    }, nil
}