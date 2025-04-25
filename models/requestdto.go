package models


type FoodPrefrenceRequestDTO struct {
	FavoriteCuisineOther     string   `json:"favorite_cuisine_other"`
	DietaryPreferenceOther   string   `json:"dietary_preference_other"`
	FoodAllergyOther         string   `json:"food_allergy_other"`
	FavoriteDishesOther      string   `json:"favorite_dishes_other"`
	CookingStyleOther        string   `json:"cooking_style_other"`
	SpiceToleranceOther      string   `json:"spice_tolerance_other"`
	FavoriteCuisines         []int    `json:"favorite_cuisines"`
	DietaryPreferences       []int    `json:"dietary_preferences"`
	FoodAllergies            []int    `json:"food_allergies"`
	FavoriteDishes           []int    `json:"favorite_dishes"`
	CookingStyles            []int    `json:"cooking_styles"`
	SpiceLevels              []int    `json:"spice_levels"`
}

type DietaryPreferenceRequestDTO struct {
	AddDietaryPreference    []int    `json:"add_dietary_preference"`
	RemoveDietaryPreference []int    `json:"remove_dietary_preference"`
}

type FoodAllergyRequestDTO struct {
	AddFoodAllergy    []int    `json:"add_food_allergy"`
	RemoveFoodAllergy []int    `json:"remove_food_allergy"`
}	

type FavoriteDishRequestDTO struct {
	AddFavoriteDish    []int    `json:"add_favorite_dish"`
	RemoveFavoriteDish []int    `json:"remove_favorite_dish"`
}	

type CookingStyleRequestDTO struct {
	AddCookingStyle    []int    `json:"add_cooking_style"`
	RemoveCookingStyle []int    `json:"remove_cooking_style"`
}			

type SpiceToleranceRequestDTO struct {
	AddSpiceTolerance    []int    `json:"add_spice_tolerance"`
	RemoveSpiceTolerance []int    `json:"remove_spice_tolerance"`
}	

type FavoriteCuisineRequestDTO struct {
	AddFavoriteCuisine    []int    `json:"add_favorite_cuisine"`
	RemoveFavoriteCuisine []int    `json:"remove_favorite_cuisine"`
}		







