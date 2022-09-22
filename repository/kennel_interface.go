package repository

import "rest-api/golang/exercise/domain/entities"

type IKennelRepository interface {
	FindAllRepo() ([]entities.Kennel, error)
	SaveRepo(u *entities.Kennel) (int, error)
	FindByIdRepo(id string) (*entities.Kennel, error)
	DeleteRepo(id string) (*entities.Kennel, error)
	UpdateRepo(u *entities.Kennel, id string) error
	CheckIfExistsRepo(id string) bool
}
