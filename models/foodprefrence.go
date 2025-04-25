package models

import (
	"errors"
)

type UserFoodPreference struct {
	UserID                 string `json:"user_id"`
	FavoriteCuisineOther   string `json:"favorite_cuisine_other"`
	DietaryPreferenceOther string `json:"dietary_preference_other"`
	FoodAllergyOther       string `json:"food_allergy_other"`
	FavoriteDishesOther    string `json:"favorite_dishes_other"`
	CookingStyleOther      string `json:"cooking_style_other"`
	SpiceToleranceOther    string `json:"spice_tolerance_other"`
}

func (u *UserFoodPreference) Validate() error {
	if u.UserID == "" {
		return errors.New("user_id not present")
	}
	return nil
}

type UserFavoriteCuisine struct {
	UserID    string `json:"user_id"`
	CuisineID int    `json:"cuisine_id"`
}

type UserDietaryPreference struct {
	UserID    string `json:"user_id"`
	DietaryPreferenceID int    `json:"dietary_preference_id"`
}

type UserFoodAllergy struct {
	UserID    string `json:"user_id"`
	FoodAllergyID int    `json:"food_allergy_id"`
}

type UserFavoriteDish struct {
	UserID    string `json:"user_id"`
	FavoriteDishID int    `json:"favorite_dish_id"`
}

type UserCookingStyle struct {
	UserID    string `json:"user_id"`
	CookingStyleID int    `json:"cooking_style_id"`
}

type UserSpiceLevel struct {
	UserID    string `json:"user_id"`
	SpiceLevelID int    `json:"spice_level_id"`
}








