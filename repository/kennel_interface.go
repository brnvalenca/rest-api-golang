package repository

import "rest-api/golang/exercise/domain/entities"

type IKennelRepository interface {
	FindAll() ([]entities.Kennel, error)
	Save(u *entities.Kennel) (int, error)
	FindById(id string) (*entities.Kennel, error)
	Delete(id string) (*entities.Kennel, error)
	Update(u *entities.Kennel, id string) error
	CheckIfExists(id string) bool
}
