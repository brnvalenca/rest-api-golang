package repos

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/golang/exercise/domain/entities"
	"rest-api/golang/exercise/repository"
	"rest-api/golang/exercise/utils"
)

type MySQL_U_Repo struct{}

func NewMySQLRepo() repository.IUserRepository {
	return &MySQL_U_Repo{}
}

func (*MySQL_U_Repo) Save(u *entities.User) (int, error) {
	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	insertRow, err := utils.DB.Query("INSERT INTO `rampup`.`users` (`nome`,`email`,`passwd`) VALUES (?, ?, ?)", u.Name, u.Email, u.Password)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on INSERT USER query")
	}
	defer insertRow.Close()

	var userID int

	err = utils.DB.QueryRow("SELECT id FROM `rampup`.`users` WHERE email = ?", u.Email).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf(err.Error(), "error on SELECT from ID query")
	}

	return userID, nil
}

/*
	ListUsers function recieves an instance of sql.DB and returns a slice of Users and a error.
	First is declared an slice to recieve de returned slice and the connection with the database
	is checked if is still alive.
	Then the SELECT statement is executed by the db.Query function, returning the rows or an error,
	the error is checked to see if its different from nil.
	The rows are iterated over the rows.Next function, this function prepare the row for be readed by
	the next Scan function. A instance of the user is declared to recieve the data from each
	column of the row and copy to the field structs by the Scan function. And then the user
	is appended to the users slice. During all these processes the errors are beign handled
	and at the end the function return the slice and a nil value for error.
*/

func (*MySQL_U_Repo) FindAll() ([]entities.User, error) {
	var users []entities.User

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := utils.DB.Query("SELECT * FROM `rampup`.`users` JOIN `rampup`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID`")
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

/*
	The ListUserById function recieves a *sql.DB instance and a string id as arguments and should
	return a entities.User and an error as result.
	The function uses a db.QueryRow to execute the SELECT statement to query for an user with an
	specific ID. Different from the previous function ListUsers, where we used the DB.Query function
	to execute the SELECT statement, the QueryRow function doesn't return an error. Instead, it
	arranges to return any query error from Rows.Scan later call.
	We then use the row.Scan function to copy the values from the columns into de struct fields and
	then we check for errors from Scan. In this case we check for the special sql.ErrNoRows error that
	points that the query returned no rows.
*/

func (*MySQL_U_Repo) FindById(id string) (*entities.User, error) {
	var user entities.User

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	row := utils.DB.QueryRow("SELECT * FROM `rampup`.`users` JOIN `rampup`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID` WHERE id = ?", id)
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

	deletedUser := utils.DB.QueryRow("SELECT * FROM `rampup`.`users` JOIN `rampup`.`user_dog_prefs` ON `users`.`id` = `user_dog_prefs`.`UserID` WHERE id = ?", id)
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
			return &user, fmt.Errorf("delete user by id: %v. no such user", id)
		}
		return &user, fmt.Errorf("delete user by id: %v: %v", id, err) // Checking if there is any error during the rows iteration
	}
	_, err = utils.DB.Exec("DELETE FROM `rampup`.`user_dog_prefs` WHERE UserID = ?", id)
	if err != nil {
		log.Fatal(err.Error())
	}

	deleteAction, err := utils.DB.Query("DELETE FROM `rampup`.`users` WHERE id = ?", id)
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

func (*MySQL_U_Repo) Update(u *entities.User, id string) error {

	err := utils.DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = utils.DB.Exec("UPDATE `rampup`.`users` SET nome = ?, email =? , passwd = ? WHERE id = ?", u.Name, u.Email, u.Password, id)
	if err != nil {
		fmt.Println(err.Error(), "error during user update")
	}

	_, err = utils.DB.Exec("UPDATE `rampup`.`user_dog_prefs` SET GoodWithKids = ?, GoodWithDogs =? , Shedding = ?, Grooming = ?, Energy = ? WHERE UserID = ?",
		u.UserPreferences.GoodWithKids,
		u.UserPreferences.GoodWithDogs,
		u.UserPreferences.Shedding,
		u.UserPreferences.Grooming,
		u.UserPreferences.Energy,
		id)
	if err != nil {
		fmt.Println(err.Error(), "error during user_dog_prefs update")
	}

	return nil
}

/*
	Criar uma função aqui para rodar um SELECT BY ID e checar se esse ID existe
*/

func (*MySQL_U_Repo) CheckIfExists(id string) bool {
	err := utils.DB.Ping()
	if err != nil {
		return false
	}
	var exists string
	err = utils.DB.QueryRow("SELECT id FROM `rampup`.`users` WHERE id = ?", id).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("no such user with id: %v", id)
			return false
		}
		return false
	}
	return true
}
