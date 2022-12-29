package repository

import (
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type IAddressRepository interface {
	SaveAddress(addr *entities.Address, kennelID int) error
}

type addrRepo struct{}

func NewAddrRepo() IAddressRepository {
	return &addrRepo{}
}

func (*addrRepo) SaveAddress(addr *entities.Address, kennelID int) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	insertRow, err := utils.DB.Query("INSERT INTO `grpc_api_db`.`kennel_addr` (`ID_Kennel`, `Numero`, `Rua`, `Bairro`, `CEP`, `Cidade`) VALUES (?, ?, ?, ?, ?, ?)",
		kennelID,
		addr.Numero,
		addr.Rua,
		addr.Bairro,
		addr.CEP,
		addr.Cidade,
	)

	if err != nil {
		log.Fatal(err.Error(), "error during address query insert")
	}

	defer insertRow.Close()

	return nil
}
