package services

import (
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
)

type kennelServ struct{}

var (
	kennelRepo repository.IKennelInterface
)

func NewKennelService(repo repository.IKennelInterface) IKennelService {
	kennelRepo = repo
	return &kennelServ{}
}

func (*kennelServ) FindAll() ([]entities.Kennel, error) {
	return kennelRepo.FindAll()
}

func (*kennelServ) Save(k *entities.Kennel) (int, error) {
	return kennelRepo.Save(k)
}

func (*kennelServ) FindById(id string) (*entities.Kennel, error) {
	return kennelRepo.FindById(id)
}

func (*kennelServ) Delete(id string) (*entities.Kennel, error) {
	return kennelRepo.Delete(id)
}

func (*kennelServ) Update(k *entities.Kennel, id string) error {
	return kennelRepo.Update(k, id)
}

func (*kennelServ) CheckIfExists(id string) bool {
	return kennelRepo.CheckIfExists(id)
}
