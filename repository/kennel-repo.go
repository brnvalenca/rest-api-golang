package repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type IKennelRepository interface {
	FindAllRepo() ([]entities.Kennel, error)
	SaveRepo(u *entities.Kennel) (int, error)
	FindByIdRepo(id string) (*entities.Kennel, error)
	DeleteRepo(id string) (*entities.Kennel, error)
	UpdateRepo(u *entities.Kennel, addr *entities.Address, id string) error
	CheckIfExistsRepo(id string) bool
}

type MySQL_K_Repo struct{}

var (
	findAllQuery            string = "SELECT * FROM `rampup`.`kennels` JOIN `rampup`.`kennel_addr` ON `kennels`.`KennelID` = `kennel_addr`.`ID_Kennel`"
	insertQuery             string = "INSERT INTO `rampup`.`kennels` (`KennelName`, `ContactNumber`) VALUES (?, ?)"
	findByIdQuery           string = "SELECT * FROM `rampup`.`kennels` JOIN `rampup`.`kennel_addr` ON `kennels`.`KennelID` = `kennel_addr`.`ID_Kennel` WHERE KennelID = ?"
	deleteAddrQuery         string = "DELETE FROM `rampup`.`kennel_addr` WHERE ID_Kennel = ?"
	deleteKennelQuery       string = "DELETE FROM `rampup`.`kennels` WHERE KennelID = ?"
	deleteDogsInKennelQuery string = "DELETE FROM `rampup`.`dogs` WHERE KennelID = ?"
	updateKennelQuery       string = "UPDATE `rampup`.`kennels` SET KennelName = ?, ContactNumber = ? WHERE KennelID = ?"
	updateKennelAddrQuery   string = "UPDATE `rampup`.`kennel_addr` SET Numero = ?, Rua = ?, Bairro = ?, CEP = ?, Cidade = ? WHERE ID_Kennel = ?"
	CheckIfExistsQuery      string = "SELECT KennelID FROM `rampup`.`kennels` WHERE KennelID = ?"
	dogByKennelQuery        string = "SELECT * from `rampup`.`kennels` JOIN `rampup`.`dogs` ON `kennels`.`KennelID` = `dogs`.`KennelID` WHERE KennelID = ?"
)

func NewKennelRepository() IKennelRepository {
	return &MySQL_K_Repo{}
}

func MatchDogWithOneKennel(dogs []entities.Dog, kennel entities.Kennel) entities.Kennel {

	for i := 0; i < len(dogs); i++ {
		if dogs[i].KennelID == kennel.ID {
			kennel.Dogs = append(kennel.Dogs, dogs[i])
		}
	}
	return kennel
}

func MatchDogsWithKennels(dogs []entities.Dog, kennels []entities.Kennel) []entities.Kennel {

	for i := 0; i < len(kennels); i++ {
		for j := 0; j < len(dogs); j++ {
			if dogs[j].KennelID == kennels[i].ID {
				kennels[i].Dogs = append(kennels[i].Dogs, dogs[j])
			}
		}
	}
	return kennels
}

func ReturnDogsArr(dogs []entities.Dog) ([]entities.Dog, error) {
	dogQuery := "SELECT * FROM `rampup`.`dogs`"
	breedQuery := "SELECT * FROM `rampup`.`breed_info`"
	dogRows, err := utils.DB.Query(dogQuery)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	breedRows, err := utils.DB.Query(breedQuery)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var breeds []entities.DogBreed
	for breedRows.Next() {
		var breed entities.DogBreed
		if err := breedRows.Scan(
			&breed.ID,
			&breed.Name,
			&breed.GoodWithKids,
			&breed.GoodWithDogs,
			&breed.Shedding,
			&breed.Grooming,
			&breed.Energy,
			&breed.BreedImg,
		); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		breeds = append(breeds, breed)
	}

	for dogRows.Next() {
		var dog entities.Dog
		if err := dogRows.Scan(
			&dog.KennelID,
			&dog.BreedID,
			&dog.DogID,
			&dog.DogName,
			&dog.Sex,
		); err != nil {
			return nil, fmt.Errorf(err.Error(), "error during dog iteration")
		}
		index := dog.BreedID
		dog.Breed = breeds[index-1]
		dogs = append(dogs, dog)
	}

	defer dogRows.Close()
	return dogs, nil
}

func (*MySQL_K_Repo) FindAllRepo() ([]entities.Kennel, error) {
	var kennels []entities.Kennel
	var dogs []entities.Dog

	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	rows, err := utils.DB.Query(findAllQuery)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var kennel entities.Kennel
		if err := rows.Scan(
			&kennel.ID,
			&kennel.Name,
			&kennel.ContactNumber,
			&kennel.Address.ID_Kennel,
			&kennel.Address.Numero,
			&kennel.Address.Rua,
			&kennel.Address.Bairro,
			&kennel.Address.CEP,
			&kennel.Address.Cidade,
		); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		kennels = append(kennels, kennel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	dogs, err = ReturnDogsArr(dogs)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	kennels = MatchDogsWithKennels(dogs, kennels)
	return kennels, nil
}

func (*MySQL_K_Repo) SaveRepo(k *entities.Kennel) (int, error) {
	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	insertRow, err := utils.DB.Query(insertQuery, k.Name, k.ContactNumber)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on insert kennel query")
	}
	defer insertRow.Close()

	var kennelID int
	// the line above takes the kennelID to be used as FK in Address table
	err = utils.DB.QueryRow("SELECT KennelID from `rampup`.`kennels` WHERE KennelName = ?", k.Name).Scan(&kennelID)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on SELECT from ID query")
	}

	return kennelID, nil
}

func (*MySQL_K_Repo) FindByIdRepo(id string) (*entities.Kennel, error) {
	var kennel entities.Kennel
	var dogs []entities.Dog

	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	row := utils.DB.QueryRow(findByIdQuery, id)
	if err := row.Scan(
		&kennel.ID,
		&kennel.Name,
		&kennel.ContactNumber,
		&kennel.Address.ID_Kennel,
		&kennel.Address.Numero,
		&kennel.Address.Rua,
		&kennel.Address.Bairro,
		&kennel.Address.CEP,
		&kennel.Address.Cidade); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("kennel by ID %v: no such kennel", id)
		}
		return &kennel, fmt.Errorf("kennel by ID %v: %v", id, err)
	}

	dogs, err = ReturnDogsArr(dogs)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	kennel = MatchDogWithOneKennel(dogs, kennel)
	return &kennel, nil
}

func (*MySQL_K_Repo) DeleteRepo(id string) (*entities.Kennel, error) {
	var kennel entities.Kennel

	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	row := utils.DB.QueryRow(findByIdQuery, id)
	if err := row.Scan(
		&kennel.ID,
		&kennel.Name,
		&kennel.ContactNumber,
		&kennel.Address.ID_Kennel,
		&kennel.Address.Numero,
		&kennel.Address.Rua,
		&kennel.Address.Bairro,
		&kennel.Address.CEP,
		&kennel.Address.Cidade,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("delete kennel error: no such kennel with id: %v", id)
		}
		return nil, fmt.Errorf("error with iteration: %v: %v", id, err)
	}
	_, err = utils.DB.Exec(deleteAddrQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error during deleting kennel address query %w", err)
	}

	_, err = utils.DB.Exec(deleteKennelQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error during deleting kennel query %w", err)
	}

	_, err = utils.DB.Exec(deleteDogsInKennelQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error during deleting dogs in kennel query %w", err)
	}
	return &kennel, nil
}

func (*MySQL_K_Repo) UpdateRepo(k *entities.Kennel, addr *entities.Address, id string) error {
	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	_, err = utils.DB.Exec(updateKennelQuery, k.Name, k.ContactNumber, id)
	if err != nil {
		return fmt.Errorf(err.Error(), "error during kennel update in db")
	}

	_, err = utils.DB.Exec(updateKennelAddrQuery,
		addr.Numero,
		addr.Rua,
		addr.Bairro,
		addr.CEP,
		addr.Cidade,
		id,
	)
	if err != nil {
		return fmt.Errorf(err.Error(), "error during the update address query")
	}

	return nil
}

func (*MySQL_K_Repo) CheckIfExistsRepo(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}
	var exists string
	err = utils.DB.QueryRow(CheckIfExistsQuery, id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such kennel with id: %v\n", id)
			return false
		}
		return false
	}
	return true
}
