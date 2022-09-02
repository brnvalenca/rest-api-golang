package repository

import "rest-api/golang/exercise/domain/entities"

type IAddressRepository interface {
	Save(addr *entities.Address) error
}
