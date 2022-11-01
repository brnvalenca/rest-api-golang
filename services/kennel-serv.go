package services

import (
	"errors"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/domain/entities/dtos"
	"rest-api/golang/exercise/middleware"
	"rest-api/golang/exercise/repository"
)

type kennelServ struct{}

type IKennelService interface {
	FindAllKennels() ([]dtos.KennelDTO, error)
	SaveKennel(u *dtos.KennelDTO) (int, error)
	FindKennelByIdServ(id string) (*dtos.KennelDTO, error)
	DeleteKennelServ(id string) (*dtos.KennelDTO, error)
	UpdateKennelServ(u *dtos.KennelDTO, id string) error
	CheckIfExists(id string) bool
}

var (
	kennelRepo repository.IKennelRepository
	addrRepo   repository.IAddressRepository = repository.NewAddrRepo()
)

func NewKennelService(repo repository.IKennelRepository, adrepo repository.IAddressRepository) IKennelService {
	kennelRepo = repo
	addrRepo = adrepo
	return &kennelServ{}
}

func ValidateKennel(k *entities.Kennel) error {
	if k == nil {
		err := errors.New("the kennel is empty")
		return err
	}

	if k.Address.Rua == "" {
		err := errors.New("the kennel street name is empty")
		return err
	}

	if k.Address.CEP == "" {
		err := errors.New("the kennel postal code is empty")
		return err
	}

	if k.Address.Cidade == "" {
		err := errors.New("the kennel city is empty")
		return err
	}

	if k.Address.Bairro == "" {
		err := errors.New("the kennel neighborhood is empty")
		return err
	}

	if k.Name == "" {
		err := errors.New("the kennel name is empty")
		return err
	}

	if k.ContactNumber == "" {
		err := errors.New("the kennel contact number is empty")
		return err
	}

	return nil
}

func (*kennelServ) FindAllKennels() ([]dtos.KennelDTO, error) {
	kennels, err := kennelRepo.FindAllRepo()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	var kennelDTO []dtos.KennelDTO
	kbuilder := dtos.NewKennelBuilderDTO()

	for i := 0; i < len(kennels); i++ {
		kbuilder.Has().
			ID(kennels[i].ID).
			ContactNumber(kennels[i].ContactNumber).
			Name(kennels[i].Name).
			Numero(kennels[i].Address.Numero).
			Rua(kennels[i].Address.Rua).
			Bairro(kennels[i].Address.Bairro).
			CEP(kennels[i].Address.CEP).
			Cidade(kennels[i].Address.Cidade)
		kennel := kbuilder.BuildKennel()
		kennelDTO = append(kennelDTO, *kennel)
	}

	return kennelDTO, nil
}

func (*kennelServ) SaveKennel(k *dtos.KennelDTO) (int, error) {
	kennelAddr, kennelInfo := middleware.PartitionKennelDTO(k)
	err := ValidateKennel(kennelInfo)
	if err != nil {
		return 0, err
	}

	kennelID, err := kennelRepo.SaveRepo(kennelInfo)
	if err != nil {
		log.Fatal(err.Error(), "error on kennelRepo.Save()")
	}

	err = addrRepo.SaveAddress(kennelAddr, kennelID)
	if err != nil {
		log.Fatal(err.Error(), " error with addrRepo.Save() method")
	}
	return kennelID, nil
}

func (*kennelServ) FindKennelByIdServ(id string) (*dtos.KennelDTO, error) {
	kennel, err := kennelRepo.FindByIdRepo(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	kbuilder := dtos.NewKennelBuilderDTO()
	kbuilder.Has().
		ID(kennel.ID).
		ContactNumber(kennel.ContactNumber).
		Name(kennel.Name).
		Numero(kennel.Address.Numero).
		Rua(kennel.Address.Rua).
		Bairro(kennel.Address.Bairro).
		CEP(kennel.Address.CEP).
		Cidade(kennel.Address.Cidade)
	kennelDTO := kbuilder.BuildKennel()
	return kennelDTO, nil
}

func (*kennelServ) DeleteKennelServ(id string) (*dtos.KennelDTO, error) {
	kennel, err := kennelRepo.DeleteRepo(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	kbuilder := dtos.NewKennelBuilderDTO()
	kbuilder.Has().
		ID(kennel.ID).
		ContactNumber(kennel.ContactNumber).
		Name(kennel.Name).
		Numero(kennel.Address.Numero).
		Rua(kennel.Address.Rua).
		Bairro(kennel.Address.Bairro).
		CEP(kennel.Address.CEP).
		Cidade(kennel.Address.Cidade)
	kennelDTO := kbuilder.BuildKennel()
	return kennelDTO, nil
}

// TODO : Fazer uma validação dos campos de endereço, Por exemplo: CEP só pode ser aceito se tiver 9 caracteres contando com o -
func (*kennelServ) UpdateKennelServ(k *dtos.KennelDTO, id string) error {
	kennelAddr, kennelInfo := middleware.PartitionKennelDTO(k)
	return kennelRepo.UpdateRepo(kennelInfo, kennelAddr, id)
}

func (*kennelServ) CheckIfExists(id string) bool {
	return kennelRepo.CheckIfExistsRepo(id)
}
