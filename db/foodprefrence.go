package db

// import (
// 	"database/sql"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"foodieMatch/models"
// 	"foodieMatch/query"
// )

// func InsertUserFoodPreference(userID string, foodPreference *models.FoodPrefrenceRequestDTO) error {
// 	// Start a transaction
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// Insert main food preference record
// 	_, err = tx.Exec(query.InsertUserFoodPreference, userID, strings.Join(foodPreference.FavoriteCuisineOther, ","), strings.Join(foodPreference.DietaryPreferenceOther, ","),
// 		strings.Join(foodPreference.FoodAllergyOther, ","), strings.Join(foodPreference.FavoriteDishesOther, ","), strings.Join(foodPreference.CookingStyleOther, ","),
// 		strings.Join(foodPreference.SpiceToleranceOther, ","), strings.Join(foodPreference.CookingLevelOther, ","), strings.Join(foodPreference.FavDrinkOther, ","), strings.Join(foodPreference.FoodDiningStyleOther, ","), strings.Join(foodPreference.FoodHabitOther, ","))

// 	if err != nil {
// 		return fmt.Errorf("error inserting food preference: %v", err)
// 	}

// 	var placeholders []string
// 	var values []interface{}
// 	// Insert favorite cuisines
// 	for _, cuisineID := range foodPreference.FavoriteCuisines {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, cuisineID)
// 	}
// 	if len(placeholders) > 0 {
// 		favCusineQuery := fmt.Sprintf(query.InsertUserFavoriteCuisine, strings.Join(placeholders, ","))
// 		fmt.Println("favCusineQuery: ", favCusineQuery)
// 		_, err = tx.Exec(favCusineQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting favorite cuisine: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	// Insert dietary preferences
// 	for _, prefID := range foodPreference.DietaryPreferences {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, prefID)
// 	}
// 	if len(placeholders) > 0 {
// 		dietaryPrefQuery := fmt.Sprintf(query.InsertUserDietaryPreference, strings.Join(placeholders, ","))
// 		fmt.Println("dietaryPrefQuery: ", dietaryPrefQuery)
// 		_, err = tx.Exec(dietaryPrefQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting dietary preference: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	// Insert food allergies
// 	for _, allergyID := range foodPreference.FoodAllergies {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, allergyID)
// 	}
// 	if len(placeholders) > 0 {
// 		foodAllergyQuery := fmt.Sprintf(query.InsertUserFoodAllergy, strings.Join(placeholders, ","))
// 		fmt.Println("foodAllergyQuery: ", foodAllergyQuery)
// 		_, err = tx.Exec(foodAllergyQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting food allergy: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	// Insert favorite dishes
// 	for _, dishID := range foodPreference.FavoriteDishes {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, dishID)
// 	}
// 	if len(placeholders) > 0 {
// 		favDishQuery := fmt.Sprintf(query.InsertUserFavoriteDish, strings.Join(placeholders, ","))
// 		fmt.Println("favDishQuery: ", favDishQuery)
// 		_, err = tx.Exec(favDishQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting favorite dish: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	// Insert cooking styles

// 	placeholders = nil
// 	values = nil
// 	// Insert cooking styles
// 	for _, styleID := range foodPreference.CookingStyles {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, styleID)
// 	}
// 	if len(placeholders) > 0 {
// 		cookingStyleQuery := fmt.Sprintf(query.InsertUserCookingStyle, strings.Join(placeholders, ","))
// 		fmt.Println("cookingStyleQuery: ", cookingStyleQuery)
// 		_, err = tx.Exec(cookingStyleQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting cooking style: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	// Insert spice levels
// 	for _, spiceLevelID := range foodPreference.SpiceLevels {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, spiceLevelID)
// 	}
// 	if len(placeholders) > 0 {
// 		spiceToleranceQuery := fmt.Sprintf(query.InsertUserSpiceTolerance, strings.Join(placeholders, ","))
// 		fmt.Println("spiceToleranceQuery: ", spiceToleranceQuery)
// 		_, err = tx.Exec(spiceToleranceQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting spice level: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	// Insert cooking levels
// 	for _, cookingLevelID := range foodPreference.CookingLevels {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, cookingLevelID)
// 	}
// 	if len(placeholders) > 0 {
// 		cookingLevelQuery := fmt.Sprintf(query.InsertUserCookingLevel, strings.Join(placeholders, ","))
// 		fmt.Println("cookingLevelQuery: ", cookingLevelQuery)
// 		_, err = tx.Exec(cookingLevelQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting cooking level: %v", err)
// 		}
// 	}


// 	placeholders = nil
// 	values = nil
// 	// Insert fav drinks
// 	for _, favDrinkID := range foodPreference.FavDrinks {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, favDrinkID)
// 	}
// 	if len(placeholders) > 0 {
// 		favDrinkQuery := fmt.Sprintf(query.InsertUserFavDrink, strings.Join(placeholders, ","))
// 		fmt.Println("favDrinkQuery: ", favDrinkQuery)
// 		_, err = tx.Exec(favDrinkQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting fav drink: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	for _, diningStyleID := range foodPreference.DiningStyles {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, diningStyleID)
// 	}
// 	if len(placeholders) > 0 {
// 		diningStyleQuery := fmt.Sprintf(query.InsertUserDiningStyle, strings.Join(placeholders, ","))
// 		fmt.Println("diningStyleQuery: ", diningStyleQuery)
// 		_, err = tx.Exec(diningStyleQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting dining style: %v", err)
// 		}
// 	}

// 	placeholders = nil
// 	values = nil
// 	for _, foodHabitID := range foodPreference.FoodHabits {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, foodHabitID)
// 	}
// 	if len(placeholders) > 0 {
// 		foodHabitQuery := fmt.Sprintf(query.InsertUserFoodHabit, strings.Join(placeholders, ","))
// 		fmt.Println("foodHabitQuery: ", foodHabitQuery)
// 		_, err = tx.Exec(foodHabitQuery, values...)
// 		if err != nil {
// 			return fmt.Errorf("error inserting food habit: %v", err)
// 		}
// 	}

// 	// Commit the transaction
// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}

// 	return nil
// }

// func mapsFilterMap(strs []string) []int {
// 	ints := make([]int, 0, len(strs))
// 	for _, s := range strs {
// 		if i, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
// 			ints = append(ints, i)
// 		}
// 	}
// 	return ints
// }

// // func GetUserFoodPreference(userID string) (*models.FoodPrefrenceResponseDTO, error) {
// // 	var foodPreference models.UserFoodPreference
// // 	var foodPreferenceResponse models.FoodPrefrenceResponseDTO
// // 	var favoriteCuisines, dietaryPreferences, foodAllergies, favoriteDishes, cookingStyles, spiceLevels, favDrinks, diningStyles, foodHabits, cookingLevels sql.NullString

// // 	err := DB.QueryRow(query.GetUserFoodPreference, userID).Scan(
// // 		&foodPreference.UserID,
// // 		&foodPreference.FavoriteCuisineOther,
// // 		&foodPreference.DietaryPreferenceOther,
// // 		&foodPreference.FoodAllergyOther,
// // 		&foodPreference.FavoriteDishesOther,
// // 		&foodPreference.CookingStyleOther,
// // 		&foodPreference.SpiceToleranceOther,
// // 		&foodPreference.CookingLevelOther,
// // 		&foodPreference.FavDrinkOther,
// // 		&foodPreference.DiningStyleOther,
// // 		&foodPreference.FoodHabitOther,
// // 		&favoriteCuisines,
// // 		&dietaryPreferences,
// // 		&foodAllergies,
// // 		&favoriteDishes,
// // 		&cookingStyles,
// // 		&spiceLevels,
// // 		&cookingLevels,
// // 		&favDrinks,
// // 		&diningStyles,
// // 		&foodHabits,
// // 	)
// // 	if err == sql.ErrNoRows {
// // 		fmt.Println("No food preferences found for user")
// // 		return nil, fmt.Errorf("not found")
// // 	}

// // 	if err != nil {
// // 		fmt.Println("Error getting Food Preference: ", err)
// // 		return nil, err
// // 	}
// // 	foodPreferenceResponse.UserID = foodPreference.UserID

// // 	if foodPreference.FavoriteCuisineOther != "" {
// // 		foodPreferenceResponse.FavoriteCuisineOther = strings.Split(foodPreference.FavoriteCuisineOther, ",")
// // 	}
// // 	if foodPreference.DietaryPreferenceOther != "" {
// // 		foodPreferenceResponse.DietaryPreferenceOther = strings.Split(foodPreference.DietaryPreferenceOther, ",")
// // 	}
// // 	if foodPreference.FoodAllergyOther != "" {
// // 		foodPreferenceResponse.FoodAllergyOther = strings.Split(foodPreference.FoodAllergyOther, ",")
// // 	}
// // 	if foodPreference.FavoriteDishesOther != "" {
// // 		foodPreferenceResponse.FavoriteDishesOther = strings.Split(foodPreference.FavoriteDishesOther, ",")
// // 	}
// // 	if foodPreference.CookingStyleOther != "" {
// // 		foodPreferenceResponse.CookingStyleOther = strings.Split(foodPreference.CookingStyleOther, ",")
// // 	}
// // 	if foodPreference.SpiceToleranceOther != "" {
// // 		foodPreferenceResponse.SpiceToleranceOther = strings.Split(foodPreference.SpiceToleranceOther, ",")
// // 	}
// // 	if foodPreference.CookingLevelOther != "" {
// // 		foodPreferenceResponse.CookingLevelOther = strings.Split(foodPreference.CookingLevelOther, ",")
// // 	}
// // 	if foodPreference.FavDrinkOther != "" {
// // 		foodPreferenceResponse.FavDrinkOther = strings.Split(foodPreference.FavDrinkOther, ",")
// // 	}
// // 	if foodPreference.DiningStyleOther != "" {
// // 		foodPreferenceResponse.DiningStyleOther = strings.Split(foodPreference.DiningStyleOther, ",")
// // 	}
// // 	if foodPreference.FoodHabitOther != "" {
// // 		foodPreferenceResponse.FoodHabitOther = strings.Split(foodPreference.FoodHabitOther, ",")
// // 	}
// // 	if favoriteCuisines.Valid {
// // 		strIDs := strings.Split(favoriteCuisines.String, ",")
// // 		foodPreferenceResponse.FavoriteCuisines = mapsFilterMap(strIDs)
// // 	}
// // 	if dietaryPreferences.Valid {
// // 		strIDs := strings.Split(dietaryPreferences.String, ",")
// // 		foodPreferenceResponse.DietaryPreferences = mapsFilterMap(strIDs)
// // 	}
// // 	if foodAllergies.Valid {
// // 		strIDs := strings.Split(foodAllergies.String, ",")
// // 		foodPreferenceResponse.FoodAllergies = mapsFilterMap(strIDs)
// // 	}
// // 	if favoriteDishes.Valid {
// // 		strIDs := strings.Split(favoriteDishes.String, ",")
// // 		foodPreferenceResponse.FavoriteDishes = mapsFilterMap(strIDs)
// // 	}
// // 	if cookingStyles.Valid {
// // 		strIDs := strings.Split(cookingStyles.String, ",")
// // 		foodPreferenceResponse.CookingStyles = mapsFilterMap(strIDs)
// // 	}
// // 	if spiceLevels.Valid {
// // 		strIDs := strings.Split(spiceLevels.String, ",")
// // 		foodPreferenceResponse.SpiceLevels = mapsFilterMap(strIDs)
// // 	}
// // 	if cookingLevels.Valid {
// // 		strIDs := strings.Split(cookingLevels.String, ",")
// // 		foodPreferenceResponse.CookingLevels = mapsFilterMap(strIDs)
// // 	}
// // 	if favDrinks.Valid {
// // 		strIDs := strings.Split(favDrinks.String, ",")
// // 		foodPreferenceResponse.FavDrinks = mapsFilterMap(strIDs)
// // 	}
// // 	if diningStyles.Valid {
// // 		strIDs := strings.Split(diningStyles.String, ",")
// // 		foodPreferenceResponse.DiningStyles = mapsFilterMap(strIDs)
// // 	}
// // 	if foodHabits.Valid {
// // 		strIDs := strings.Split(foodHabits.String, ",")
// // 		foodPreferenceResponse.FoodHabits = mapsFilterMap(strIDs)
// // 	}
// // 	return &foodPreferenceResponse, nil
// // }

// func UpdateUserFoodPreference(userID string, foodPreference *models.FoodPrefrenceUpdateRequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
	
// 	_, err = tx.Exec(query.UpdateUserFoodPreference,
// 		strings.Join(foodPreference.FavoriteCuisine.Other, ","),
// 		strings.Join(foodPreference.DietaryPreference.Other, ","),
// 		strings.Join(foodPreference.FoodAllergy.Other, ","),
// 		strings.Join(foodPreference.FavoriteDishes.Other, ","),
// 		strings.Join(foodPreference.CookingStyle.Other, ","),
// 		strings.Join(foodPreference.SpiceTolerance.Other, ","),
// 		strings.Join(foodPreference.CookingLevel.Other, ","),
// 		strings.Join(foodPreference.FavDrink.Other, ","),
// 		strings.Join(foodPreference.FoodDiningStyle.Other, ","),
// 		strings.Join(foodPreference.FoodHabit.Other, ","),
// 		userID)

// 	if err != nil {
// 		return fmt.Errorf("error updating food preference: %v", err)
// 	}


	

// 	err = addUserFavouriteCuisine(tx, userID, &foodPreference.FavoriteCuisine)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite cuisine: %v", err)
// 	}
// 	err = removeUserFavouriteCuisine(tx, userID, &foodPreference.FavoriteCuisine)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite cuisine: %v", err)
// 	}

// 	err = addUserDietaryPreference(tx, userID, &foodPreference.DietaryPreference)
// 	if err != nil {
// 		return fmt.Errorf("error updating dietary preference: %v", err)
// 	}
// 	err = removeUserDietaryPreference(tx, userID, &foodPreference.DietaryPreference)
// 	if err != nil {
// 		return fmt.Errorf("error updating dietary preference: %v", err)
// 	}

// 	err = addUserFoodAllergy(tx, userID, &foodPreference.FoodAllergy)
// 	if err != nil {
// 		return fmt.Errorf("error updating food allergy: %v", err)
// 	}

// 	err = removeUserFoodAllergy(tx, userID, &foodPreference.FoodAllergy)
// 	if err != nil {
// 		return fmt.Errorf("error updating food allergy: %v", err)
// 	}

// 	err = addUserFavoriteDish(tx, userID, &foodPreference.FavoriteDishes)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite dish: %v", err)
// 	}
// 	err = removeUserFavoriteDish(tx, userID, &foodPreference.FavoriteDishes)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite dish: %v", err)
// 	}

// 	err = addUserCookingStyle(tx, userID, &foodPreference.CookingStyle)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking style: %v", err)
// 	}

// 	err = removeUserCookingStyle(tx, userID, &foodPreference.CookingStyle)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking style: %v", err)
// 	}

// 	err = addUserSpiceTolerance(tx, userID, &foodPreference.SpiceTolerance)
// 	if err != nil {
// 		return fmt.Errorf("error updating spice tolerance: %v", err)
// 	}
// 	err = removeUserSpiceTolerance(tx, userID, &foodPreference.SpiceTolerance)
// 	if err != nil {
// 		return fmt.Errorf("error updating spice tolerance: %v", err)
// 	}

// 	err = addUserCookingLevel(tx, userID, &foodPreference.CookingLevel)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking level: %v", err)
// 	}

// 	err = removeUserCookingLevel(tx, userID, &foodPreference.CookingLevel)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking level: %v", err)
// 	}

// 	err = addUserFavDrinks(tx, userID, &foodPreference.FavDrink)
// 	if err != nil {
// 		return fmt.Errorf("error updating fav drink: %v", err)
// 	}

// 	err = removeUserFavDrinks(tx, userID, &foodPreference.FavDrink)
// 	if err != nil {
// 		return fmt.Errorf("error updating fav drink: %v", err)
// 	}

// 	err = addUserDiningStyles(tx, userID, &foodPreference.FoodDiningStyle)
// 	if err != nil {
// 		return fmt.Errorf("error updating food dining style: %v", err)
// 	}

// 	err = removeUserDiningStyles(tx, userID, &foodPreference.FoodDiningStyle)
// 	if err != nil {
// 		return fmt.Errorf("error updating food dining style: %v", err)
// 	}

// 	err = addUserFoodHabits(tx, userID, &foodPreference.FoodHabit)
// 	if err != nil {
// 		return fmt.Errorf("error updating food habit: %v", err)
// 	}

// 	err = removeUserFoodHabits(tx, userID, &foodPreference.FoodHabit)
// 	if err != nil {
// 		return fmt.Errorf("error updating food habit: %v", err)
// 	}
	
// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserCookingLevel(tx *sql.Tx, userID string, updateCookingLevel *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, cookingLevelID := range updateCookingLevel.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, cookingLevelID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserCookingLevel, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting cooking level: %v", err)
// 	}
// 	return nil
// }

// func removeUserCookingLevel(tx *sql.Tx, userID string, updateCookingLevel *models.RequestDTO) error {
// 	for _, cookingLevelID := range updateCookingLevel.Remove {
// 		_, err := tx.Exec(query.DeleteUserCookingLevel, userID, cookingLevelID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting cooking level: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserCookingLevel(userID string, updateCookingLevel *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserCookingLevel(tx, userID, updateCookingLevel)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking level: %v", err)
// 	}
// 	err = removeUserCookingLevel(tx, userID, updateCookingLevel)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking level: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserFavouriteCuisine(tx *sql.Tx, userID string, updateCuisine *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, cuisineID := range updateCuisine.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, cuisineID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserFavoriteCuisine, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting favorite cuisine: %v", err)
// 	}
// 	return nil
// }

// func removeUserFavouriteCuisine(tx *sql.Tx, userID string, updateCuisine *models.RequestDTO) error {
// 	for _, cuisineID := range updateCuisine.Remove {
// 		_, err := tx.Exec(query.DeleteUserFavoriteCuisine, userID, cuisineID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting favorite cuisine: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserFavoriteCuisine(userID string, updateCuisine *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserFavouriteCuisine(tx, userID, updateCuisine)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite cuisine: %v", err)
// 	}
// 	err = removeUserFavouriteCuisine(tx, userID, updateCuisine)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite cuisine: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserDietaryPreference(tx *sql.Tx, userID string, updateDietaryPreference *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, prefID := range updateDietaryPreference.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, prefID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserDietaryPreference, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting dietary preference: %v", err)
// 	}
// 	return nil
// }

// func removeUserDietaryPreference(tx *sql.Tx, userID string, updateDietaryPreference *models.RequestDTO) error {
// 	for _, prefID := range updateDietaryPreference.Remove {
// 		_, err := tx.Exec(query.DeleteUserDietaryPreference, userID, prefID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting dietary preference: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserDietaryPreference(userID string, updateDietaryPreference *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserDietaryPreference(tx, userID, updateDietaryPreference)
// 	if err != nil {
// 		return fmt.Errorf("error updating dietary preference: %v", err)
// 	}
// 	err = removeUserDietaryPreference(tx, userID, updateDietaryPreference)
// 	if err != nil {
// 		return fmt.Errorf("error updating dietary preference: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserFoodAllergy(tx *sql.Tx, userID string, updateFoodAllergy *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, allergyID := range updateFoodAllergy.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, allergyID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserFoodAllergy, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting food allergy: %v", err)
// 	}
// 	return nil
// }

// func removeUserFoodAllergy(tx *sql.Tx, userID string, updateFoodAllergy *models.RequestDTO) error {
// 	for _, allergyID := range updateFoodAllergy.Remove {
// 		_, err := tx.Exec(query.DeleteUserFoodAllergy, userID, allergyID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting food allergy: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserFoodAllergy(userID string, updateFoodAllergy *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserFoodAllergy(tx, userID, updateFoodAllergy)
// 	if err != nil {
// 		return fmt.Errorf("error updating food allergy: %v", err)
// 	}
// 	err = removeUserFoodAllergy(tx, userID, updateFoodAllergy)
// 	if err != nil {
// 		return fmt.Errorf("error updating food allergy: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserFavoriteDish(tx *sql.Tx, userID string, updateFavoriteDish *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, dishID := range updateFavoriteDish.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, dishID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserFavoriteDish, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting favorite dish: %v", err)
// 	}
// 	return nil
// }

// func removeUserFavoriteDish(tx *sql.Tx, userID string, updateFavoriteDish *models.RequestDTO) error {
// 	for _, dishID := range updateFavoriteDish.Remove {
// 		_, err := tx.Exec(query.DeleteUserFavoriteDish, userID, dishID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting favorite dish: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserFavoriteDish(userID string, updateFavoriteDish *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserFavoriteDish(tx, userID, updateFavoriteDish)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite dish: %v", err)
// 	}
// 	err = removeUserFavoriteDish(tx, userID, updateFavoriteDish)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite dish: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserCookingStyle(tx *sql.Tx, userID string, updateCookingStyle *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, styleID := range updateCookingStyle.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, styleID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserCookingStyle, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting cooking style: %v", err)
// 	}
// 	return nil
// }

// func removeUserCookingStyle(tx *sql.Tx, userID string, updateCookingStyle *models.RequestDTO) error {
// 	for _, styleID := range updateCookingStyle.Remove {
// 		_, err := tx.Exec(query.DeleteUserCookingStyle, userID, styleID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting cooking style: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserCookingStyle(userID string, updateCookingStyle *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserCookingStyle(tx, userID, updateCookingStyle)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking style: %v", err)
// 	}
// 	err = removeUserCookingStyle(tx, userID, updateCookingStyle)
// 	if err != nil {
// 		return fmt.Errorf("error updating cooking style: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserSpiceTolerance(tx *sql.Tx, userID string, updateSpiceTolerance *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, spiceLevelID := range updateSpiceTolerance.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, spiceLevelID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserSpiceTolerance, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting spice tolerance: %v", err)
// 	}
// 	return nil
// }

// func removeUserSpiceTolerance(tx *sql.Tx, userID string, updateSpiceTolerance *models.RequestDTO) error {
// 	for _, spiceLevelID := range updateSpiceTolerance.Remove {
// 		_, err := tx.Exec(query.DeleteUserSpiceTolerance, userID, spiceLevelID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting spice tolerance: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserSpiceTolerance(userID string, updateSpiceTolerance *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserSpiceTolerance(tx, userID, updateSpiceTolerance)
// 	if err != nil {
// 		return fmt.Errorf("error updating spice tolerance: %v", err)
// 	}
// 	err = removeUserSpiceTolerance(tx, userID, updateSpiceTolerance)
// 	if err != nil {
// 		return fmt.Errorf("error updating spice tolerance: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserFavDrinks(tx *sql.Tx, userID string, updateFavoriteDrinks *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, drinkID := range updateFavoriteDrinks.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, drinkID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserFavDrink, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting favorite drink: %v", err)
// 	}
// 	return nil
// }

// func removeUserFavDrinks(tx *sql.Tx, userID string, updateFavoriteDrinks *models.RequestDTO) error {
// 	for _, drinkID := range updateFavoriteDrinks.Remove {
// 		_, err := tx.Exec(query.DeleteUserFavDrink, userID, drinkID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting favorite drink: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserFavDrinks(userID string, updateFavoriteDrinks *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserFavDrinks(tx, userID, updateFavoriteDrinks)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite drink: %v", err)
// 	}
// 	err = removeUserFavDrinks(tx, userID, updateFavoriteDrinks)
// 	if err != nil {
// 		return fmt.Errorf("error updating favorite drink: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserDiningStyles(tx *sql.Tx, userID string, updateDiningStyles *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, diningStyleID := range updateDiningStyles.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, diningStyleID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserDiningStyle, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting dining style: %v", err)
// 	}
// 	return nil
// }

// func removeUserDiningStyles(tx *sql.Tx, userID string, updateDiningStyles *models.RequestDTO) error {
// 	for _, diningStyleID := range updateDiningStyles.Remove {
// 		_, err := tx.Exec(query.DeleteUserDiningStyle, userID, diningStyleID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting dining style: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserDiningStyles(userID string, updateDiningStyles *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserDiningStyles(tx, userID, updateDiningStyles)
// 	if err != nil {
// 		return fmt.Errorf("error updating dining style: %v", err)
// 	}
// 	err = removeUserDiningStyles(tx, userID, updateDiningStyles)
// 	if err != nil {
// 		return fmt.Errorf("error updating dining style: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }

// func addUserFoodHabits(tx *sql.Tx, userID string, updateFoodHabits *models.RequestDTO) error {
// 	var placeholders []string
// 	var values []interface{}
// 	for _, foodHabitID := range updateFoodHabits.Add {
// 		placeholders = append(placeholders, "(?, ?)")
// 		values = append(values, userID, foodHabitID)
// 	}
// 	if len(values) == 0 {
// 		return nil
// 	}
// 	query := fmt.Sprintf(query.InsertUserFoodHabit, strings.Join(placeholders, ","))
// 	_, err := tx.Exec(query, values...)
// 	if err != nil {
// 		return fmt.Errorf("error inserting food habit: %v", err)
// 	}
// 	return nil
// }

// func removeUserFoodHabits(tx *sql.Tx, userID string, updateFoodHabits *models.RequestDTO) error {
// 	for _, foodHabitID := range updateFoodHabits.Remove {
// 		_, err := tx.Exec(query.DeleteUserFoodHabit, userID, foodHabitID)
// 		if err != nil {
// 			return fmt.Errorf("error deleting food habit: %v", err)
// 		}
// 	}
// 	return nil
// }

// func UpdateUserFoodHabits(userID string, updateFoodHabits *models.RequestDTO) error {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("failed to begin transaction: %v", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	err = addUserFoodHabits(tx, userID, updateFoodHabits)
// 	if err != nil {
// 		return fmt.Errorf("error updating food habit: %v", err)
// 	}
// 	err = removeUserFoodHabits(tx, userID, updateFoodHabits)
// 	if err != nil {
// 		return fmt.Errorf("error updating food habit: %v", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return nil
// }
