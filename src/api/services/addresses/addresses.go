package addresses

import (
    addressesDomain "github.com/emikohmann/itacademy2018-myml/src/api/domain/addresses"
    "github.com/emikohmann/itacademy2018-myml/src/api/util/apierrors"
)

func GetAddressByID(addressID int64) (*addressesDomain.Address, *apierrors.ApiError) {
    return &addressesDomain.Address{
        AddressID: addressID,
    }, nil
}
