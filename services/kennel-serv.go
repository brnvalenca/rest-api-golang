package services

import (
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/repository/repos"
	"rest-api/golang/exercise/services/middleware"
)

type kennelServ struct{}

var (
	kennelRepo repository.IKennelRepository
	addrRepo   repository.IAddressRepository = repos.NewAddrRepo()
)

func NewKennelService(repo repository.IKennelRepository, adrepo repository.IAddressRepository) IKennelService {
	kennelRepo = repo
	addrRepo = adrepo
	return &kennelServ{}
}

func (*kennelServ) FindAllKennels() ([]entities.Kennel, error) {
	return kennelRepo.FindAllRepo()
}

func (*kennelServ) Save(k *entities.Kennel) (int, error) {
	kennel, err := kennelRepo.SaveRepo(k)
	if err != nil {
		log.Fatal(err.Error(), "error on kennelRepo.Save()")
	}
	kennelAddr := middleware.PartitionKennelAddress(k, kennel)
	err = addrRepo.SaveAddress(kennelAddr)
	if err != nil {
		log.Fatal(err.Error(), " error with addrRepo.Save() method")
	}

	return kennel, nil
}

func (*kennelServ) FindKennelByIdServ(id string) (*entities.Kennel, error) {
	kennel, err := kennelRepo.FindByIdRepo(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return kennel, nil
}

func (*kennelServ) DeleteKennelServ(id string) (*entities.Kennel, error) {
	return kennelRepo.DeleteRepo(id)
}

func (*kennelServ) UpdateKennelServ(k *entities.Kennel, id string) error {
	return kennelRepo.UpdateRepo(k, id)
}

func (*kennelServ) CheckIfExists(id string) bool {
	return kennelRepo.CheckIfExistsRepo(id)
}
