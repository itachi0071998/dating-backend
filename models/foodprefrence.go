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
	FavDrinkOther          string `json:"fav_drink_other"`
	DiningStyleOther       string `json:"dining_style_other"`
	FoodHabitOther         string `json:"food_habit_other"`
	CookingLevelOther      string `json:"cooking_level_other"`
}

func (u *UserFoodPreference) Validate() error {
	if u.UserID == "" {
		return errors.New("user_id not present")
	}
	return nil
}

type PrefrenceTableIds struct {
	Add    []int `json:"add"`
	Remove []int `json:"remove"`
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


type UserFavDrink struct {
	UserID    string `json:"user_id"`
	FavDrinkID int    `json:"fav_drink_id"`
}

type UserDiningStyle struct {
	UserID    string `json:"user_id"`
	DiningStyleID int    `json:"dining_style_id"`
}

type UserFoodHabit struct {
	UserID    string `json:"user_id"`
	FoodHabitID int    `json:"food_habit_id"`
}

type UserCookingLevel struct {
	UserID    string `json:"user_id"`
	CookingLevelID int    `json:"cooking_level_id"`
}

type FoodPrefrenceRequestDTO struct {
	FavoriteCuisineOther     []string   `json:"favorite_cuisine_other"`
	DietaryPreferenceOther   []string   `json:"dietary_preference_other"`
	FoodAllergyOther         []string   `json:"food_allergy_other"`
	FavoriteDishesOther      []string   `json:"favorite_dishes_other"`
	CookingStyleOther        []string   `json:"cooking_style_other"`
	SpiceToleranceOther      []string   `json:"spice_tolerance_other"`
	FavDrinkOther            []string   `json:"fav_drink_other"`
	FoodDiningStyleOther     []string   `json:"food_dining_style_other"`
	FoodHabitOther           []string   `json:"food_habit_other"`
	CookingLevelOther       []string   `json:"cooking_level_other"`
	FavoriteCuisines         []int      `json:"favorite_cuisines"`
	DietaryPreferences       []int      `json:"dietary_preferences"`
	FoodAllergies            []int      `json:"food_allergies"`
	FavoriteDishes           []int      `json:"favorite_dishes"`
	CookingStyles            []int      `json:"cooking_styles"`
	SpiceLevels              []int      `json:"spice_levels"`
	CookingLevels            []int      `json:"cooking_levels"`
	FavDrinks                []int      `json:"fav_drinks"`
	DiningStyles             []int      `json:"dining_styles"`
	FoodHabits               []int      `json:"food_habits"`
}

type FoodPrefrenceUpdateRequestDTO struct {
	FavoriteCuisine     RequestDTO   `json:"favorite_cuisine"`
	DietaryPreference   RequestDTO   `json:"dietary_preference"`
	FoodAllergy         RequestDTO   `json:"food_allergy"`
	FavoriteDishes      RequestDTO   `json:"favorite_dishes"`
	CookingStyle        RequestDTO   `json:"cooking_style"`
	SpiceTolerance      RequestDTO   `json:"spice_tolerance"`
	FavDrink            RequestDTO   `json:"fav_drink"`
	FoodDiningStyle     RequestDTO   `json:"food_dining_style"`
	FoodHabit           RequestDTO   `json:"food_habit"`
	CookingLevel        RequestDTO   `json:"cooking_level"`
	DiningStyle         RequestDTO   `json:"dining_style"`
}

type FoodPrefrenceResponseDTO struct {
	UserID                   string   `json:"user_id"`
	FavoriteCuisineOther     []string   `json:"favorite_cuisine_other"`
	DietaryPreferenceOther   []string   `json:"dietary_preference_other"`
	FoodAllergyOther         []string   `json:"food_allergy_other"`
	FavoriteDishesOther      []string   `json:"favorite_dishes_other"`
	CookingStyleOther        []string   `json:"cooking_style_other"`
	SpiceToleranceOther      []string   `json:"spice_tolerance_other"`
	CookingLevelOther        []string   `json:"cooking_level_other"`
	FavDrinkOther            []string   `json:"fav_drink_other"`
	DiningStyleOther     []string   `json:"dining_style_other"`
	FoodHabitOther           []string   `json:"food_habit_other"`
	FavoriteCuisines         []int `json:"favorite_cuisines"`
	DietaryPreferences       []int `json:"dietary_preferences"`
	FoodAllergies            []int `json:"food_allergies"`
	FavoriteDishes           []int `json:"favorite_dishes"`
	CookingStyles            []int `json:"cooking_styles"`
	SpiceLevels              []int `json:"spice_levels"`
	CookingLevels            []int `json:"cooking_levels"`
	FavDrinks                []int `json:"fav_drinks"`
	DiningStyles             []int `json:"dining_styles"`
	FoodHabits               []int `json:"food_habits"`
}




