package repos

import (
	"database/sql"
	"fmt"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/utils"
)

type breedRepo struct{}

func NewBreedRepository() repository.IBreedRepository {
	return &breedRepo{}
}

/*
	Function recieves a DogBreed object and returns the BreedID saved as int if execution is succeded
	and a nil error. It checks errors on some points. Starts checking the DB connection than makes an
	insert query to database. The BreedID is extracted and scanned to an int variable, to be used as
	argument to dog.Save() function and used as FK on dog table.
*/

func (*breedRepo) Save(d *entities.DogBreed) (int, error) {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`breed_info` (`BreedName`, `GoodWithKids`, `GoodWithDogs`, `Shedding`, `Grooming`, `Energy`, `BreedImg`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		d.Name,
		d.GoodWithKids,
		d.GoodWithDogs,
		d.Shedding,
		d.Grooming,
		d.Energy,
		d.BreedImg,
	)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on INSERT BREED query")
	}
	defer insertRow.Close()

	var breedID int

	err = utils.DB.QueryRow("SELECT BreedID FROM `rampup`.`breed_info` WHERE BreedName = ?", d.Name).Scan(&breedID)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on SELECT from ID query")
	}

	return breedID, nil
}

/*
	Select breed with specified BreedID.
*/

func (*breedRepo) FindById(id string) (*entities.DogBreed, error) {
	var breed entities.DogBreed

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	row := utils.DB.QueryRow("SELECT * FROM `rampup`.`breed_info` WHERE BreedID = ?", id)
	if err := row.Scan(&breed.ID,
		&breed.Name,
		&breed.GoodWithKids,
		&breed.GoodWithDogs,
		&breed.Shedding,
		&breed.Grooming,
		&breed.Energy,
		&breed.BreedImg,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("breed by ID %v: no such breed", id)
		}
		return &breed, fmt.Errorf("breed by ID %v: %v", id, err) // Checking if there is any error during the rows iteration
	}

	return &breed, nil
}

/*
	Return a list with all breeds registered
*/

func (*breedRepo) FindAll() ([]entities.DogBreed, error) {
	var breeds []entities.DogBreed

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := utils.DB.Query("SELECT * FROM `rampup`.`breed_info`")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var breed entities.DogBreed
		if err := rows.Scan(
			&breed.ID,
			&breed.Name,
			&breed.GoodWithKids,
			&breed.GoodWithDogs,
			&breed.Shedding,
			&breed.Grooming,
			&breed.Energy,
			&breed.BreedImg,
		); err != nil {
			return nil, fmt.Errorf(err.Error(), "error during scan")
		}
		breeds = append(breeds, breed)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return breeds, nil

}

func (*breedRepo) Update(d *entities.DogBreed) error {

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = utils.DB.Exec("UPDATE `rampup`.`breed_info` SET BreedName = ?, GoodWithKids = ?, GoodWithDogs = ?, Shedding = ?, Grooming = ?, Energy = ?, BreedImg = ? WHERE BreedID = ?",
		d.Name,
		d.GoodWithKids,
		d.GoodWithDogs,
		d.Shedding,
		d.Grooming,
		d.Energy,
		d.BreedImg,
		d.ID,
	)
	if err != nil {
		fmt.Println(err.Error(), "error during breed update")
	}

	return nil
}

func (*breedRepo) Delete(id string) (*entities.DogBreed, error) {
	var breed entities.DogBreed

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error(), "error with db conn")
	}

	deletedBreed := utils.DB.QueryRow("SELECT * FROM `rampup`.`breed_info` WHERE BreedID = ?", id)
	if err := deletedBreed.Scan(&breed.ID,
		&breed.GoodWithKids,
		&breed.GoodWithDogs,
		&breed.Shedding,
		&breed.Grooming,
		&breed.Energy,
		&breed.Name,
		&breed.BreedImg,
	); err != nil {
		if err == sql.ErrNoRows {
			return &breed, fmt.Errorf("delete breed by id: %v. no such user", id)
		}
		return &breed, fmt.Errorf("delete breed by id: %v: %v", id, err) // Checking if there is any error during the rows iteration
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `rampup`.`breed_info` WHERE BreedID = ?", id)
	if err != nil {
		return &breed, fmt.Errorf(err.Error(), "error with the delete breed query")
	}
	defer deleteAction.Close()
	return &breed, nil
}

func (*breedRepo) CheckIfExists(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		return false
	}
	var exists string
	err = utils.DB.QueryRow("SELECT BreedID FROM `rampup`.`breed_info` WHERE BreedID = ?", id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such breed with id: %v", id)
			return false
		}
		return false
	}
	return true
}
