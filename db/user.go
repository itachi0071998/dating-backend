package db

import (
	"database/sql"
	"errors"
	"fmt"
	"foodieMatch/models"
	"foodieMatch/query"
)

// CheckUserExists checks if a user exists by their ID
func CheckUserExists(id string) (bool, error) {
	var exists bool
	err := DB.QueryRow(query.UserExists, id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking user existence: ", err)
		return false, err
	}
	return exists, nil
}

// SaveUser saves a user to the database
func SaveUser(user *models.User) error {
	_, err := DB.Exec(query.CreateUser, user.ID, user.Mobile, user.Email, user.DOB, user.FirstName, user.LastName, user.Gender, user.Url)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return err
	}
	return nil
}

func UpdateUser(user *models.User) error {
	_, err := DB.Exec(query.UpdateUser, user.ID, user.Mobile, user.Email, user.DOB, user.FirstName, user.LastName, user.Gender, user.Url, user.ID)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		return err
	}
	return nil
}

// GetUserByID retrieves a user from the database by ID
func GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := DB.QueryRow(query.GetUserById, id).Scan(
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
