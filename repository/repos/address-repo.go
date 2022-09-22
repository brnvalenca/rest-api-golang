package repos

import (
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/utils"
)

type addrRepo struct{}

func NewAddrRepo() repository.IAddressRepository {
	return &addrRepo{}
}

func (*addrRepo) SaveAddress(addr *entities.Address) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`kennel_addr` (`ID_Kennel`, `Numero`, `Rua`, `Bairro`, `CEP`, `Cidade`) VALUES (?, ?, ?, ?, ?, ?)",
		addr.ID_Kennel,
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
