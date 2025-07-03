package models

import (
    "errors"
)


type User struct {
	ID               string    `json:"id"`
	Mobile           string    `json:"mobile"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Gender           string    `json:"gender"`
	DOB              int64     `json:"dob"`
	Email   		 string    `json:"email"`
	Url              string    `json:"url"`
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return errors.New("name not present")
	}
	if u.Email == "" && u.Mobile == "" {
		return errors.New("email and mobile not present")
	}
	if u.Gender == "" {
		return errors.New("gender not present")
	}
	return nil
}

type UserRequestDTO struct {
	User User `json:"user"`
	FoodPreference FoodPrefrenceRequestDTO `json:"food_preference"`
}
