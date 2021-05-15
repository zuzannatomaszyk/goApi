package db

import (
	"github.com/zuzannatomaszyk/goApi/models"
)

func (db Database) GetUsers() (*models.UsersList, error) {
	//query execution
	list := &models.UsersList{}
	rows, err := db.Conn.Query("SELECT username FROM users")
	if err != nil {
		return list, err
	}
	//add every row returned by query to users array
	for rows.Next() {
		var user string
		err := rows.Scan(&user)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}
