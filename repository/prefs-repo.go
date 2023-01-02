package repository

import (
	"fmt"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type IPrefsRepository interface {
	SavePrefs(u *entities.UserDogPreferences, userID int) error
	DeletePrefs(id string) error
	UpdatePrefs(u *entities.UserDogPreferences, id string) error
}

type PrefsRepo struct{}

func NewPrefsRepo() IPrefsRepository {
	return &PrefsRepo{}
}

func (*PrefsRepo) SavePrefs(u *entities.UserDogPreferences, userID int) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	insertRow, err := utils.DB.Query("INSERT INTO `grpc_api_db`.`user_dog_prefs` (`UserID`, `GoodWithKids`, `GoodWithDogs`, `Shedding`, `Grooming`, `Energy`) VALUES (?, ?, ?, ?, ?, ?)", userID, u.GoodWithKids, u.GoodWithDogs, u.Shedding, u.Grooming, u.Energy)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer insertRow.Close()

	return nil

}

func (*PrefsRepo) UpdatePrefs(u *entities.UserDogPreferences, id string) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = utils.DB.Exec("UPDATE `grpc_api_db`.`user_dog_prefs` SET GoodWithKids = ?, GoodWithDogs = ?, Shedding = ?, Grooming = ?, Energy = ?",
		u.GoodWithKids, u.GoodWithDogs, u.Shedding, u.Grooming, u.Energy)
	if err != nil {
		fmt.Println(err.Error(), "update userdogprefs error")
	}
	return nil
}

func (*PrefsRepo) DeletePrefs(id string) error {

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `grpc_api_db`.`user_dog_prefs` WHERE UserID = ?", id)
	if err != nil {
		return err
	}
	defer deleteAction.Close()
	return nil
}
