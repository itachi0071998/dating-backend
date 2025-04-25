package models

type FoodPrefrenceResponseDTO struct {
	UserID                   string   `json:"user_id"`
	FavoriteCuisineOther     string   `json:"favorite_cuisine_other"`
	DietaryPreferenceOther   string   `json:"dietary_preference_other"`
	FoodAllergyOther         string   `json:"food_allergy_other"`
	FavoriteDishesOther      string   `json:"favorite_dishes_other"`
	CookingStyleOther        string   `json:"cooking_style_other"`
	SpiceToleranceOther      string   `json:"spice_tolerance_other"`
	FavoriteCuisines         []string `json:"favorite_cuisines"`
	DietaryPreferences       []string `json:"dietary_preferences"`
	FoodAllergies            []string `json:"food_allergies"`
	FavoriteDishes           []string `json:"favorite_dishes"`
	CookingStyles            []string `json:"cooking_styles"`
	SpiceLevels              []string `json:"spice_levels"`
}


