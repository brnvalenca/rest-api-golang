package repos

import (
	"fmt"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/utils"
)

type PrefsRepo struct{}

func NewPrefsRepo() repository.IPrefsRepository {
	return &PrefsRepo{}
}

func (*PrefsRepo) Save(u *entities.UserDogPreferences) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`user_dog_prefs` (`UserID`, `GoodWithKids`, `GoodWithDogs`, `Shedding`, `Grooming`, `Energy`) VALUES (?, ?, ?, ?, ?, ?)", u.UserID, u.GoodWithKids, u.GoodWithDogs, u.Shedding, u.Grooming, u.Energy)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer insertRow.Close()

	return nil

}

func (*PrefsRepo) Update(u *entities.UserDogPreferences, id string) error {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = utils.DB.Exec("UPDATE `rampup`.`user_dog_prefs` SET GoodWithKids = ?, GoodWithDogs = ?, Shedding = ?, Grooming = ?, Energy = ?",
		u.GoodWithKids, u.GoodWithDogs, u.Shedding, u.Grooming, u.Energy)
	if err != nil {
		fmt.Println(err.Error(), "update userdogprefs error")
	}
	return nil
}

func (*PrefsRepo) Delete(id string) error {

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `rampup`.`user_dog_prefs` WHERE UserID = ?", id)
	if err != nil {
		return err
	}
	defer deleteAction.Close()
	return nil
}
