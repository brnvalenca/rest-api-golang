package repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type MySQL_K_Repo struct{}

func NewKennelRepository() IKennelInterface {
	return &MySQL_K_Repo{}
}

func (*MySQL_K_Repo) FindAll() ([]entities.Kennel, error) {
	var kennels []entities.Kennel

	err := utils.DB.Ping()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	rows, err := utils.DB.Query("SELECT * FROM `rampup`.`kennels`")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var kennel entities.Kennel
		if err := rows.Scan(&kennel.ID, &kennel.Name); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		kennels = append(kennels, kennel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return kennels, nil
}

func (*MySQL_K_Repo) Save(k *entities.Kennel) (int, error) {
	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`kennels` (`KennelName`) VALUES (?)", k.Name)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on INSERT KENNEL query")
	}
	defer insertRow.Close()

	var kennelID int

	err = utils.DB.QueryRow("SELECT KennelID from `rampup`.`kennels` WHERE KennelName = ?", k.Name).Scan(&kennelID)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on SELECT from ID query")
	}

	return kennelID, nil
}

func (*MySQL_K_Repo) FindById(id string) (*entities.Kennel, error) {
	var kennel entities.Kennel

	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	row := utils.DB.QueryRow("SELECT * FROM `rampup`.`kennels` WHERE KennelID = ?", id)
	if err := row.Scan(
		&kennel.ID,
		&kennel.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("kennel by ID %v: no such kennel", err)
		}
		return &kennel, fmt.Errorf("kennel by ID %v: %v", id, err)
	}

	return &kennel, nil
}

func (*MySQL_K_Repo) Delete(id string) (*entities.Kennel, error) {
	var kennel entities.Kennel

	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	deletedKennel := utils.DB.QueryRow("SELECT * FROM `rampup`.`kennels` WHERE KennelID = ?", id)
	if err := deletedKennel.Scan(&kennel.ID, &kennel.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("delete kennel error: no such kennel with id: %v", id)
		}
		return nil, fmt.Errorf("error with iteration: %v: %v", id, err)
	}

	_, err = utils.DB.Exec("DELETE FROM `rampup`.`kennels` WHERE KennelID = ?", id)
	if err != nil {
		log.Fatal(err.Error(), "error during the DELETE query exec")
	}

	return &kennel, nil
}

func (*MySQL_K_Repo) Update(k *entities.Kennel, id string) error {
	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}

	_, err = utils.DB.Exec("UPDATE `rampup`.`kennels` SET KennelName = ? WHERE KennelID = ?", k.Name, id)
	if err != nil {
		log.Fatal(err.Error(), "error during kennel update in db")
	}

	return nil
}

func (*MySQL_K_Repo) CheckIfExists(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		log.Fatal(err.Error(), "db conn error")
	}
	var exists string
	err = utils.DB.QueryRow("SELECT KennelID FROM `rampup`.`kennels` WHERE KennelID = ?", id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such kennel with id: %v\n", id)
			return false
		}
		return false
	}
	return true
}
