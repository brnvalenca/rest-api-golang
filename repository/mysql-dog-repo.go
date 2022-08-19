package repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type MySQL_D_Repo struct{}

func NewSQL_D_Repo() DogRepositoryI {
	return &MySQL_D_Repo{}
}

func (*MySQL_D_Repo) FindAll() ([]entities.Dog, error) {
	var dogs []entities.Dog

	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	rows, err := utils.DB.Query("SELECT * FROM `rampup`.`dogs`")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var dog entities.Dog
		if err := rows.Scan(&dog.ID, &dog.Name, &dog.Sex, &dog.Breed, &dog.Age); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		dogs = append(dogs, dog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return dogs, nil
}

func (*MySQL_D_Repo) FindById(id string) (*entities.Dog, error) {
	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	var dog entities.Dog
	dogRow := utils.DB.QueryRow("SELECT * FROM `rampup`.`dogs` WHERE id = ?", id)
	if err := dogRow.Scan(&dog.ID, &dog.Name, &dog.Sex, &dog.Breed, &dog.Age); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dog by ID %v: no such dog", id)
		}
		return &dog, fmt.Errorf("dog by ID %v: %v. Error during the iteration", id, err)
	}

	return &dog, nil
}

func (*MySQL_D_Repo) Save(d *entities.Dog) (*entities.Dog, error) {
	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`dogs` (`dogname`, `sex`, `breed`, `age`) VALUES (?, ?, ?, ?)", d.Name, d.Sex, d.Breed, d.Age)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer insertRow.Close()
	var dog entities.Dog

	if err := insertRow.Scan(&dog.ID, &dog.Name, &dog.Sex, &dog.Breed, &dog.Age); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &dog, nil
}

func (*MySQL_D_Repo) Delete(id string) (*entities.Dog, error) {
	var dog entities.Dog

	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	deletedRow := utils.DB.QueryRow("SELECT * FROM `rampup`.`dogs` WHERE id = ?", id)
	if err := deletedRow.Scan(&dog.ID, &dog.Name, &dog.Sex, &dog.Breed, &dog.Age); err != nil {
		if err == sql.ErrNoRows {
			return &dog, fmt.Errorf("delete dog by id: %v. no such dog", id)
		}
		return &dog, fmt.Errorf("delete dog by id: %v: %v. error during the iteration", id, err)
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `rampup`.`dogs` WHERE id = ?", id)
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
	_, err = utils.DB.Exec("UPDATE `rampup`.`dogs` SET dogname = ?, sex =? , breed = ?, age = ? WHERE id = ?", d.Name, d.Sex, d.Breed, d.Age, id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (*MySQL_D_Repo) CheckIfExists(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		return false
	}

	var exists string
	err = utils.DB.QueryRow("SELECT id FROM `rampup`.`dogs` WHERE id = ?", id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such user with id: %v", id)
			return false
		}
		return false
	}
	return true
}
