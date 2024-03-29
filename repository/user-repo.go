package repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/utils"
)

type IUserRepository interface {
	Save(u *entities.User) (int, error)
	FindAll() ([]entities.User, error)
	FindById(id string) (*entities.User, error)
	Delete(id string) (*entities.User, error)
	Update(u *entities.User, uprefs *entities.UserDogPreferences) error
	CheckIfExists(id string) bool
	CheckEmail(email string) (bool, *entities.User)
}

type MySQL_U_Repo struct{}

func NewMySQLRepo() IUserRepository {
	return &MySQL_U_Repo{}
}

func (*MySQL_U_Repo) Save(u *entities.User) (int, error) {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	insertRow, err := utils.DB.Query("INSERT INTO `grpc_api_db`.`users` (`nome`,`email`,`passwd`) VALUES (?, ?, ?)", u.Name, u.Email, u.Password)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on INSERT USER query")
	}
	defer insertRow.Close()

	var userID int

	err = utils.DB.QueryRow("SELECT id FROM `grpc_api_db`.`users` WHERE email = ?", u.Email).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on SELECT from ID query")
	}

	return userID, nil
}

func (*MySQL_U_Repo) FindAll() ([]entities.User, error) {
	var users []entities.User

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := utils.DB.Query("SELECT * FROM `grpc_api_db`.`users` JOIN `grpc_api_db`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID`")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.UserPreferences.UserID,
			&user.UserPreferences.GoodWithKids,
			&user.UserPreferences.GoodWithDogs,
			&user.UserPreferences.Shedding,
			&user.UserPreferences.Grooming,
			&user.UserPreferences.Energy,
		); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return users, nil

}

func (*MySQL_U_Repo) FindById(id string) (*entities.User, error) {
	var user entities.User

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	row := utils.DB.QueryRow("SELECT * FROM `grpc_api_db`.`users` JOIN `grpc_api_db`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID` WHERE id = ?", id)
	if err := row.Scan(&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.UserPreferences.UserID,
		&user.UserPreferences.GoodWithKids,
		&user.UserPreferences.GoodWithDogs,
		&user.UserPreferences.Shedding,
		&user.UserPreferences.Grooming,
		&user.UserPreferences.Energy); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user by ID %v: no such user", id)
		}
		return &user, fmt.Errorf("user by ID %v: %v", id, err) // Checking if there is any error during the rows iteration
	}

	return &user, nil
}

/*
	DeleteUser function uses a slight similar approach to the ListUserById func. First a user
	entity is instanced to be returned by the function and to be returned by the http.ResponseWriter
	variable argument from the HandleFunc that deals with the endpoint. A quick check of the database
	conn is procceded with the Ping function, and then the deletedRow is iterated by a SELECT FROM ID statement
	and the columns of the table are mapped to the user instance by the Scan function. The errors are checked and then
	after this the delete action takes place with a query execution made by a db.Query function.
*/

func (*MySQL_U_Repo) Delete(id string) (*entities.User, error) {
	var user entities.User

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	deletedUser := utils.DB.QueryRow("SELECT * FROM `grpc_api_db`.`users` JOIN `grpc_api_db`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID` WHERE id = ?", id)
	if err := deletedUser.Scan(&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.UserPreferences.UserID,
		&user.UserPreferences.GoodWithKids,
		&user.UserPreferences.GoodWithDogs,
		&user.UserPreferences.Shedding,
		&user.UserPreferences.Grooming,
		&user.UserPreferences.Energy); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("delete user by id: %v. no such user", id)
		}
		return &user, fmt.Errorf("delete user by id: %v: %v", id, err) // Checking if there is any error during the rows iteration
	}
	_, err = utils.DB.Exec("DELETE FROM `grpc_api_db`.`user_dog_prefs` WHERE UserID = ?", id)
	if err != nil {
		log.Fatal(err.Error())
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `grpc_api_db`.`users` WHERE id = ?", id)
	if err != nil {
		return &user, fmt.Errorf(err.Error())
	}
	defer deleteAction.Close()
	return &user, nil
}

/*
	The UpdateUser func recieves a sql.DB instance and a user entity as arguments and return an value of
	int type, representing the number of rows affected by the update and an error, that in normal conditions
	if all went well will be returned as nil.
	First of all we take de user ID and store on a local variable, then is procceded a quick check of the
	db conn with de Ping function and then the db.QueryRow execs the UPDATE query, no value is returned from this
	proccedure. Some errors are checked before it returns, if no error is captured then the rows variable is incremented
	by one and the functions returns.
*/

func (*MySQL_U_Repo) Update(user *entities.User, uprefs *entities.UserDogPreferences) error {

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = utils.DB.Exec("UPDATE `grpc_api_db`.`users` SET nome = ?, email =? , passwd = ? WHERE id = ?", user.Name, user.Email, user.Password, uprefs.UserID)
	if err != nil {
		fmt.Println(err.Error(), "error during user update")
	}

	_, err = utils.DB.Exec("UPDATE `grpc_api_db`.`user_dog_prefs` SET GoodWithKids = ?, GoodWithDogs =? , Shedding = ?, Grooming = ?, Energy = ? WHERE UserID = ?",
		uprefs.GoodWithKids,
		uprefs.GoodWithDogs,
		uprefs.Shedding,
		uprefs.Grooming,
		uprefs.Energy,
		uprefs.UserID)
	if err != nil {
		fmt.Println(err.Error(), "error during user_dog_prefs update")
	}

	return nil
}

func (*MySQL_U_Repo) CheckIfExists(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		return false
	}
	var exists string
	err = utils.DB.QueryRow("SELECT id FROM `grpc_api_db`.`users` WHERE id = ?", id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such user with id: %v", id)
			return false
		}
		return false
	}
	return true
}

func (*MySQL_U_Repo) CheckEmail(email string) (bool, *entities.User) {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	var user entities.User
	checkedUser := utils.DB.QueryRow("SELECT * FROM `grpc_api_db`.`users` JOIN `grpc_api_db`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID` WHERE email = ?", email)
	if err := checkedUser.Scan(&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.UserPreferences.UserID,
		&user.UserPreferences.GoodWithKids,
		&user.UserPreferences.GoodWithDogs,
		&user.UserPreferences.Shedding,
		&user.UserPreferences.Grooming,
		&user.UserPreferences.Energy); err != nil {
		if err == sql.ErrNoRows {
			return false, &user
		}
		return false, &user // Checking if there is any error during the rows iteration
	}
	return true, &user
}
