package repository

import "rest-api/golang/exercise/domain/entities"

type IAddressRepository interface {
	SaveAddress(addr *entities.Address) error
}
