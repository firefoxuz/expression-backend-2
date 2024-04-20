package entities

import (
	"expression-backend/internal/database"
	constErrors "expression-backend/internal/errors"
	"log"
)

type User struct {
	Id       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

func InsertUser(user *User) error {
	db, err := database.GetConnection()

	if err != nil {
		log.Println(err)
		return constErrors.CannotConnectDatabase
	}

	_, err = db.NamedExec("insert into  users (login, password) values (:login, :password)", &user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func FindUserByLogin(login string) (*User, error) {
	user := User{}

	db, err := database.GetConnection()

	if err != nil {
		log.Println(err)
		return nil, constErrors.CannotConnectDatabase
	}

	err = db.Get(&user, "select * from users where login = $1", login)

	if err != nil {
		log.Println(err)
		return nil, constErrors.CannotFindEntity
	}

	return &user, err
}
