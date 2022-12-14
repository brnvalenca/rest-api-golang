package repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type IDogRepository interface {
	Save(d *entities.Dog, id interface{}) error
	FindAll() ([]entities.Dog, error)
	FindById(id string) (*entities.Dog, error)
	Delete(id string) (*entities.Dog, error)
	Update(d *entities.Dog, id string) error
	CheckIfExists(id string) bool
}

type MySQL_D_Repo struct{}

func NewSQL_D_Repo() IDogRepository {
	return &MySQL_D_Repo{}
}

func (*MySQL_D_Repo) FindAll() ([]entities.Dog, error) {
	var dogs []entities.Dog

	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	rows, err := utils.DB.Query("SELECT * FROM `rampup`.`dogs` JOIN `rampup`.`breed_info` ON `dogs`.`BreedID` = `breed_info`.`BreedID`")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var dog entities.Dog
		if err := rows.Scan(
			&dog.KennelID,
			&dog.BreedID,
			&dog.DogID,
			&dog.DogName,
			&dog.Sex,
			&dog.Breed.ID,
			&dog.Breed.Name,
			&dog.Breed.GoodWithKids,
			&dog.Breed.GoodWithDogs,
			&dog.Breed.Shedding,
			&dog.Breed.Grooming,
			&dog.Breed.Energy,
			&dog.Breed.BreedImg,
		); err != nil {
			return nil, fmt.Errorf(err.Error(), "error during the iteration")
		}
		dogs = append(dogs, dog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(err.Error(), "error here")
	}

	return dogs, nil
}

func (*MySQL_D_Repo) FindById(id string) (*entities.Dog, error) {
	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	var dog entities.Dog
	dogRow := utils.DB.QueryRow("SELECT * FROM `rampup`.`dogs` JOIN `rampup`.`breed_info` ON `dogs`.`BreedID` = `breed_info`.`BreedID` WHERE DogID = ?", id)
	if err := dogRow.Scan(
		&dog.KennelID,
		&dog.BreedID,
		&dog.DogID,
		&dog.DogName,
		&dog.Sex,
		&dog.Breed.ID,
		&dog.Breed.Name,
		&dog.Breed.GoodWithKids,
		&dog.Breed.GoodWithDogs,
		&dog.Breed.Shedding,
		&dog.Breed.Grooming,
		&dog.Breed.Energy,
		&dog.Breed.BreedImg,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dog by ID %v: no such dog", id)
		}
		return &dog, fmt.Errorf("dog by ID %v: %v. Error during the iteration", id, err)
	}

	return &dog, nil
}

func (*MySQL_D_Repo) Save(d *entities.Dog, breedid interface{}) error {
	err := utils.DB.Ping()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`dogs` (`KennelID`, `BreedID` ,`DogName`, `Sex`) VALUES (?, ?, ?, ?)", d.KennelID, breedid, d.DogName, d.Sex)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	defer insertRow.Close()

	return nil
}

func (*MySQL_D_Repo) Delete(id string) (*entities.Dog, error) {
	var dog entities.Dog

	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	deletedRow := utils.DB.QueryRow("SELECT * FROM `rampup`.`dogs` JOIN `rampup`.`breed_info` ON `dogs`.`BreedID` = `breed_info`.`BreedID` WHERE DogID = ?", id)
	if err := deletedRow.Scan(
		&dog.KennelID,
		&dog.BreedID,
		&dog.DogID,
		&dog.DogName,
		&dog.Sex,
		&dog.Breed.ID,
		&dog.Breed.Name,
		&dog.Breed.GoodWithKids,
		&dog.Breed.GoodWithDogs,
		&dog.Breed.Shedding,
		&dog.Breed.Grooming,
		&dog.Breed.Energy,
		&dog.Breed.BreedImg,
	); err != nil {
		if err == sql.ErrNoRows {
			return &dog, fmt.Errorf("delete dog by id: %v. no such dog", id)
		}
		return &dog, fmt.Errorf("delete dog by id: %v: %v. error during the iteration", id, err)
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `rampup`.`dogs` WHERE DogID = ?", id)
	if err != nil {
		return &dog, fmt.Errorf(err.Error(), "error during the delete query")
	}

	defer deleteAction.Close()
	return &dog, nil
}

func (*MySQL_D_Repo) Update(d *entities.Dog, id string) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = utils.DB.Exec("UPDATE `rampup`.`dogs` SET KennelID = ?, BreedID = ?,  DogID = ?, DogName = ?, Sex = ? WHERE DogID = ?", d.KennelID, d.BreedID, id, d.DogName, d.Sex, id)
	if err != nil {
		log.Fatal(err.Error(), "update dog failed")
	}

	return nil
}

func (*MySQL_D_Repo) CheckIfExists(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		return false
	}

	var exists string
	err = utils.DB.QueryRow("SELECT DogID FROM `rampup`.`dogs` WHERE DogID = ?", id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such dog with id: %v", id)
			return false
		}
		return false
	}
	return true
}
