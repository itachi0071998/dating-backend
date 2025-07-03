package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"foodieMatch/db"
	"foodieMatch/models"
	"foodieMatch/query"
)

func CheckUserExists(tx *sql.Tx, id string) (bool, error) {
	var exists bool
	err := tx.QueryRow(query.UserExists, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("Error checking user existence: %v", err)
	}
	return exists, nil
}

func UpsertUser(user *models.User) error {
	_, err := db.DB.Exec(query.UpsertUser, user.ID, user.Mobile, user.Email, user.DOB, user.FirstName, user.LastName, user.Gender, user.Url)
	if err != nil {
		return fmt.Errorf("Error upsert user: %v", err)
	}
	return nil	
}

// func UpdateUser(tx *sql.Tx, user *models.User) error {
// 	_, err := tx.Exec(query.UpdateUser, user.ID, user.Mobile, user.Email, user.DOB, user.FirstName, user.LastName, user.Gender, user.Url, user.ID)
// 	if err != nil {
// 		return fmt.Errorf("Error updating user: %v", err)
// 	}
// 	return nil
// }

func GetUserByID(tx *sql.Tx, id string) (*models.User, error) {
	user := &models.User{}
	err := tx.QueryRow(query.GetUserById, id).Scan(
		&user.ID,
		&user.Mobile,
		&user.Email,
		&user.DOB,
		&user.FirstName,
		&user.LastName,
		&user.Gender,
		&user.Url,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
