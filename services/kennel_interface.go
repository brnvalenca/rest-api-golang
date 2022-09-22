package services

import "rest-api/golang/exercise/domain/entities"

type IKennelService interface {
	FindAllKennels() ([]entities.Kennel, error)
	Save(u *entities.Kennel) (int, error)
	FindKennelByIdServ(id string) (*entities.Kennel, error)
	DeleteKennelServ(id string) (*entities.Kennel, error)
	UpdateKennelServ(u *entities.Kennel, id string) error
	CheckIfExists(id string) bool
}
