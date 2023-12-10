package user

import (
	"context"
	"time"

	"github.com/imzoloft/gonetmaster/api/config"
	"github.com/imzoloft/gonetmaster/api/database"
	"github.com/imzoloft/gonetmaster/logger"
)

func InsertUser(uuid string, key []byte) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	insertUserQuery := `
	INSERT INTO client (ID, KEY) VALUES ($1, $2)
	`

	_, err := database.Db.ExecContext(ctx, insertUserQuery, uuid, key)

	if err != nil {
		logger.Log.Warn("Error inserting user: ", err)
		return err
	}
	return nil
}

func RetrieveUsers() (*[]User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	retrieveUsersQuery := `
	SELECT ID, KEY FROM client
	`

	rows, err := database.Db.QueryContext(ctx, retrieveUsersQuery)
	defer rows.Close()

	if err != nil {
		logger.Log.Warn("Error retrieving users: ", err)
		return nil, err
	}

	var users []User

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Key)

		if err != nil {
			logger.Log.Warn("Error scanning user: ", err)
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}
