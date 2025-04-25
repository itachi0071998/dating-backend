package db

import (
	"database/sql"
	"fmt"
	"foodieMatch/models"
	"foodieMatch/query"
	"strings"
)


func InsertUserFoodPreference(userID string, foodPreference *models.FoodPrefrenceRequestDTO) error {
	// Start a transaction
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert main food preference record
	_, err = tx.Exec(query.InsertUserFoodPreference, userID, foodPreference.FavoriteCuisineOther, foodPreference.DietaryPreferenceOther,
		foodPreference.FoodAllergyOther, foodPreference.FavoriteDishesOther, foodPreference.CookingStyleOther, foodPreference.SpiceToleranceOther)
	if err != nil {
		return fmt.Errorf("error inserting food preference: %v", err)
	}

	var placeholders []string
	var values []interface{}
	// Insert favorite cuisines
	for _, cuisineID := range foodPreference.FavoriteCuisines {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, cuisineID)
	}
	favCusineQuery := fmt.Sprintf(query.InsertUserFavoriteCuisine, strings.Join(placeholders, ","))

	
	_, err = tx.Exec(favCusineQuery, values...)
	if err != nil {
		return fmt.Errorf("error inserting favorite cuisine: %v", err)
	}
	

	placeholders = nil
	values = nil
	// Insert dietary preferences
	for _, prefID := range foodPreference.DietaryPreferences {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, prefID)
	}
	dietaryPrefQuery := fmt.Sprintf(query.InsertUserDietaryPreference, strings.Join(placeholders, ","))
	fmt.Println("dietaryPrefQuery: ", dietaryPrefQuery)
	_, err = tx.Exec(dietaryPrefQuery, values...)
	if err != nil {
		return fmt.Errorf("error inserting dietary preference: %v", err)
	}

	placeholders = nil
	values = nil
	// Insert food allergies	
	for _, allergyID := range foodPreference.FoodAllergies {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, allergyID)
	}
	foodAllergyQuery := fmt.Sprintf(query.InsertUserFoodAllergy, strings.Join(placeholders, ","))
	fmt.Println("foodAllergyQuery: ", foodAllergyQuery)
	_, err = tx.Exec(foodAllergyQuery, values...)
	if err != nil {
		return fmt.Errorf("error inserting food allergy: %v", err)
	}
	
	placeholders = nil
	values = nil
	// Insert favorite dishes
	for _, dishID := range foodPreference.FavoriteDishes {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, dishID)
	}
	favDishQuery := fmt.Sprintf(query.InsertUserFavoriteDish, strings.Join(placeholders, ","))
	fmt.Println("favDishQuery: ", favDishQuery)
	_, err = tx.Exec(favDishQuery, values...)
	if err != nil {
		return fmt.Errorf("error inserting favorite dish: %v", err)
	}

	placeholders = nil
	values = nil
	// Insert cooking styles
	for _, styleID := range foodPreference.CookingStyles {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, styleID)
	}
	cookingStyleQuery := fmt.Sprintf(query.InsertUserCookingStyle, strings.Join(placeholders, ","))
	fmt.Println("cookingStyleQuery: ", cookingStyleQuery)
	_, err = tx.Exec(cookingStyleQuery, values...)
		if err != nil {
		return fmt.Errorf("error inserting cooking style: %v", err)
	}			

	placeholders = nil
	values = nil
	// Insert spice levels
	for _, spiceLevelID := range foodPreference.SpiceLevels {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, spiceLevelID)
	}
	spiceToleranceQuery := fmt.Sprintf(query.InsertUserSpiceTolerance, strings.Join(placeholders, ","))
	fmt.Println("spiceToleranceQuery: ", spiceToleranceQuery)
	_, err = tx.Exec(spiceToleranceQuery, values...)
		if err != nil {
		return fmt.Errorf("error inserting spice level: %v", err)
	}
	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func GetUserFoodPreference(userID string) (*models.FoodPrefrenceResponseDTO, error) {
	var foodPreference models.FoodPrefrenceResponseDTO
	var favoriteCuisines, dietaryPreferences, foodAllergies, favoriteDishes, cookingStyles, spiceLevels sql.NullString

	err := DB.QueryRow(query.GetUserFoodPreference, userID).Scan(
		&foodPreference.UserID,
		&foodPreference.FavoriteCuisineOther,
		&foodPreference.DietaryPreferenceOther,
		&foodPreference.FoodAllergyOther,
		&foodPreference.FavoriteDishesOther,
		&foodPreference.CookingStyleOther,
		&foodPreference.SpiceToleranceOther,
		&favoriteCuisines,
		&dietaryPreferences,
		&foodAllergies,
		&favoriteDishes,
		&cookingStyles,
		&spiceLevels,
	)
	if err == sql.ErrNoRows {
		fmt.Println("No food preferences found for user")
		return nil, fmt.Errorf("not found")
	}

	if err != nil {
		fmt.Println("Error getting Food Preference: ", err)
		return nil, err
	}

	if favoriteCuisines.Valid {
		foodPreference.FavoriteCuisines = strings.Split(favoriteCuisines.String, ",")
	}
	if dietaryPreferences.Valid {
		foodPreference.DietaryPreferences = strings.Split(dietaryPreferences.String, ",")
	}
	if foodAllergies.Valid {
		foodPreference.FoodAllergies = strings.Split(foodAllergies.String, ",")
	}
	if favoriteDishes.Valid {
		foodPreference.FavoriteDishes = strings.Split(favoriteDishes.String, ",")
	}
	if cookingStyles.Valid {
		foodPreference.CookingStyles = strings.Split(cookingStyles.String, ",")
	}
	if spiceLevels.Valid {
		foodPreference.SpiceLevels = strings.Split(spiceLevels.String, ",")
	}

	return &foodPreference, nil
}


func UpdateUserFoodPreference(userID string, foodPreference *models.UserFoodPreference) error {
	_, err := DB.Exec(query.UpdateUserFoodPreference, foodPreference.FavoriteCuisineOther, foodPreference.DietaryPreferenceOther,
		foodPreference.FoodAllergyOther, foodPreference.FavoriteDishesOther, foodPreference.CookingStyleOther, foodPreference.SpiceToleranceOther, userID)
	if err != nil {
		fmt.Println("Error updating Food Preference: ", err)
		return err
	}	
	return nil
}

func UpdateUserFavoriteCuisine(userID string, updateCuisine *models.FavoriteCuisineRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string	
	var values []interface{}

	for _, cuisineID := range updateCuisine.AddFavoriteCuisine {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, cuisineID)
	}
	if len(placeholders) > 0 {
		favCusineQuery := fmt.Sprintf(query.InsertUserFavoriteCuisine, strings.Join(placeholders, ","))
		fmt.Println("favCusineQuery: ", favCusineQuery)
		_, err = tx.Exec(favCusineQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting favorite cuisine: %v", err)
		}
	}

	for _, cuisineID := range updateCuisine.RemoveFavoriteCuisine {
		_, err = tx.Exec(query.DeleteUserFavoriteCuisine, userID, cuisineID)
		if err != nil {
			return fmt.Errorf("error deleting favorite cuisine: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func UpdateUserDietaryPreference(userID string, updateDietaryPreference *models.DietaryPreferenceRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string		
	var values []interface{}

	for _, prefID := range updateDietaryPreference.AddDietaryPreference {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, prefID)
	}
	if len(placeholders) > 0 {
		dietaryPrefQuery := fmt.Sprintf(query.InsertUserDietaryPreference, strings.Join(placeholders, ","))
		fmt.Println("dietaryPrefQuery: ", dietaryPrefQuery)
		_, err = tx.Exec(dietaryPrefQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting dietary preference: %v", err)
		}
	}

	for _, prefID := range updateDietaryPreference.RemoveDietaryPreference {
		_, err = tx.Exec(query.DeleteUserDietaryPreference, userID, prefID)
		if err != nil {
			return fmt.Errorf("error deleting dietary preference: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}	

func UpdateUserFoodAllergy(userID string, updateFoodAllergy *models.FoodAllergyRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}	
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string		
	var values []interface{}

	for _, allergyID := range updateFoodAllergy.AddFoodAllergy {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, allergyID)
	}
	if len(placeholders) > 0 {
		foodAllergyQuery := fmt.Sprintf(query.InsertUserFoodAllergy, strings.Join(placeholders, ","))
		fmt.Println("foodAllergyQuery: ", foodAllergyQuery)
		_, err = tx.Exec(foodAllergyQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting food allergy: %v", err)
		}
	}	

	for _, allergyID := range updateFoodAllergy.RemoveFoodAllergy {
		_, err = tx.Exec(query.DeleteUserFoodAllergy, userID, allergyID)
		if err != nil {
			return fmt.Errorf("error deleting food allergy: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func UpdateUserFavoriteDish(userID string, updateFavoriteDish *models.FavoriteDishRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}	
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string		
	var values []interface{}

	for _, dishID := range updateFavoriteDish.AddFavoriteDish {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, dishID)
	}
	if len(placeholders) > 0 {
		favDishQuery := fmt.Sprintf(query.InsertUserFavoriteDish, strings.Join(placeholders, ","))
		fmt.Println("favDishQuery: ", favDishQuery)
		_, err = tx.Exec(favDishQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting favorite dish: %v", err)
		}
	}

	for _, dishID := range updateFavoriteDish.RemoveFavoriteDish {
		_, err = tx.Exec(query.DeleteUserFavoriteDish, userID, dishID)
		if err != nil {
			return fmt.Errorf("error deleting favorite dish: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}	
	return nil
}

func UpdateUserCookingStyle(userID string, updateCookingStyle *models.CookingStyleRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}	
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string		
	var values []interface{}

	for _, styleID := range updateCookingStyle.AddCookingStyle {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, styleID)
	}
	if len(placeholders) > 0 {
		cookingStyleQuery := fmt.Sprintf(query.InsertUserCookingStyle, strings.Join(placeholders, ","))
		fmt.Println("cookingStyleQuery: ", cookingStyleQuery)
		_, err = tx.Exec(cookingStyleQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting cooking style: %v", err)
		}
	}

	for _, styleID := range updateCookingStyle.RemoveCookingStyle {
		_, err = tx.Exec(query.DeleteUserCookingStyle, userID, styleID)
		if err != nil {
			return fmt.Errorf("error deleting cooking style: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func UpdateUserSpiceTolerance(userID string, updateSpiceTolerance *models.SpiceToleranceRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}	
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string		
	var values []interface{}

	for _, spiceLevelID := range updateSpiceTolerance.AddSpiceTolerance {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, spiceLevelID)
	}
	if len(placeholders) > 0 {
		spiceToleranceQuery := fmt.Sprintf(query.InsertUserSpiceTolerance, strings.Join(placeholders, ","))
		fmt.Println("spiceToleranceQuery: ", spiceToleranceQuery)
		_, err = tx.Exec(spiceToleranceQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting spice tolerance: %v", err)
		}
	}

	for _, spiceLevelID := range updateSpiceTolerance.RemoveSpiceTolerance {
		_, err = tx.Exec(query.DeleteUserSpiceLevel, userID, spiceLevelID)
		if err != nil {
			return fmt.Errorf("error deleting spice tolerance: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

									