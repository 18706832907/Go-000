package dao

import "database/sql"

type User struct {
	UserName string
	UserId   int
}

func GetUserName(userId int) (User, error) {

	error := sql.ErrNoRows
	userInfo := User{nil, nil}

	return userInfo, error
}
