package repository

import (
	"database/sql"
	"fmt"
	"foodieMatch/db"
	"foodieMatch/models"
	"foodieMatch/query"
	"log"
	"strconv"
	"strings"
)

func InsertUserFavoriteCuisine(userID string, favoriteCuisines []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserFavoriteCuisine started")
	for _, cuisineID := range favoriteCuisines {
		values = append(values, userID, cuisineID)
		placeholders = append(placeholders, "(?, ?)")
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserFavoriteCuisine, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting favorite cuisine: %v", err)
	}
	log.Println("InsertUserFavoriteCuisine success")
	return nil
}

func ExcludeUserFavoriteCuisine(userID string, favoriteCuisines []int) error {
	for _, cuisineID := range favoriteCuisines {
		_, err := db.DB.Exec(query.DeleteUserFavoriteCuisine, userID, cuisineID)
		if err != nil {
			return fmt.Errorf("error deleting favorite cuisine: %v", err)
		}
	}
	return nil
}

func GetUserFavoriteCuisine(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserFavoriteCuisine, userID)
	if err != nil {
		return nil, fmt.Errorf("error selecting favorite cuisine: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning favorite cuisine: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertDietaryPreference(userID string, dietaryPreferences []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertDietaryPreference started")
	for _, dietaryPreferenceID := range dietaryPreferences {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, dietaryPreferenceID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserDietaryPreference, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting dietary preference: %v", err)
	}
	log.Println("InsertDietaryPreference success")
	return nil
}

func ExcludeUserDietaryPreference(userID string, dietaryPreferences []int) error {
	for _, dietaryPreferenceID := range dietaryPreferences {
		_, err := db.DB.Exec(query.DeleteUserDietaryPreference, userID, dietaryPreferenceID)
		if err != nil {
			return fmt.Errorf("error deleting dietary preference: %v", err)
		}
	}
	return nil
}

func GetUserDietaryPreference(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserDietaryPreference, userID)
	log.Println("GetUserDietaryPreference started")
	if err != nil {
		return nil, fmt.Errorf("error selecting dietary preference: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning dietary preference: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserFoodAllergy(userID string, foodAllergies []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserFoodAllergy started")
	for _, allergyID := range foodAllergies {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, allergyID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserFoodAllergy, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting food allergy: %v", err)
	}
	log.Println("InsertUserFoodAllergy success")
	return nil
}

func ExcludeUserFoodAllergy(userID string, foodAllergies []int) error {
	for _, allergyID := range foodAllergies {
		_, err := db.DB.Exec(query.DeleteUserFoodAllergy, userID, allergyID)
		if err != nil {
			return fmt.Errorf("error deleting food allergy: %v", err)
		}
	}
	return nil

}

func GetUserFoodAllergy(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserFoodAllergy, userID)
	log.Println("GetUserFoodAllergy started")
	if err != nil {
		return nil, fmt.Errorf("error selecting food allergy: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning food allergy: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserFavoriteDish(userID string, favoriteDishes []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserFavoriteDish started")
	for _, favoriteDishID := range favoriteDishes {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, favoriteDishID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserFavoriteDish, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting favorite dish: %v", err)
	}
	log.Println("InsertUserFavoriteDish success")
	return nil
}

func ExcludeUserFavoriteDish(userID string, favoriteDishes []int) error {
	for _, favoriteDishID := range favoriteDishes {
		_, err := db.DB.Exec(query.DeleteUserFavoriteDish, userID, favoriteDishID)
		if err != nil {
			return fmt.Errorf("error deleting favorite dish: %v", err)
		}
	}
	return nil
}

func GetUserFavoriteDish(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserFavoriteDish, userID)
	log.Println("GetUserFavoriteDish started")
	if err != nil {
		return nil, fmt.Errorf("error selecting favorite dish: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning favorite dish: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}
func InsertUserCookingStyle(userID string, cookingStyles []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserCookingStyle started")
	for _, cookingStyleID := range cookingStyles {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, cookingStyleID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserCookingStyle, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting cooking style: %v", err)
	}
	log.Println("InsertUserCookingStyle success")
	return nil
}

func ExcludeUserCookingStyle(userID string, cookingStyles []int) error {
	for _, cookingStyleID := range cookingStyles {
		_, err := db.DB.Exec(query.DeleteUserCookingStyle, userID, cookingStyleID)
		if err != nil {
			return fmt.Errorf("error deleting cooking style: %v", err)
		}
	}
	return nil
}

func GetUserCookingStyle(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserCookingStyle, userID)
	log.Println("GetUserCookingStyle started")
	if err != nil {
		return nil, fmt.Errorf("error selecting cooking style: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning cooking style: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserSpiceTolerance(userID string, spiceTolerances []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserSpiceTolerance started")
	for _, spiceToleranceID := range spiceTolerances {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, spiceToleranceID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserSpiceTolerance, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting spice tolerance: %v", err)
	}
	log.Println("InsertUserSpiceTolerance success")
	return nil
}

func ExcludeUserSpiceTolerance(	userID string, spiceTolerances []int) error {
	for _, spiceToleranceID := range spiceTolerances {
		_, err := db.DB.Exec(query.DeleteUserSpiceTolerance, userID, spiceToleranceID)
		if err != nil {
			return fmt.Errorf("error deleting spice tolerance: %v", err)
		}
	}
	return nil
}

func GetUserSpiceTolerance(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserSpiceTolerance, userID)
	log.Println("GetUserSpiceTolerance started")
	if err != nil {
		return nil, fmt.Errorf("error selecting spice tolerance: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning spice tolerance: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserCookingLevel(userID string, cookingLevels []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserCookingLevel started")
	for _, cookingLevelID := range cookingLevels {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, cookingLevelID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserCookingLevel, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting cooking level: %v", err)
	}
	log.Println("InsertUserCookingLevel success")
	return nil
}

func ExcludeUserCookingLevel(userID string, cookingLevels []int) error {
	for _, cookingLevelID := range cookingLevels {
		_, err := db.DB.Exec(query.DeleteUserCookingLevel, userID, cookingLevelID)
		if err != nil {
			return fmt.Errorf("error deleting cooking level: %v", err)
		}
	}
	return nil
}

func GetUserCookingLevel(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserCookingLevel, userID)
	log.Println("GetUserCookingLevel started")
	if err != nil {
		return nil, fmt.Errorf("error selecting cooking level: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning cooking level: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserFavDrink(userID string, favDrinks []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserFavDrink started")
	for _, favDrinkID := range favDrinks {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, favDrinkID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserFavDrink, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting fav drink: %v", err)
	}
	log.Println("InsertUserFavDrink success")
	return nil
}

func ExcludeUserFavDrink(	userID string, favDrinks []int) error {
	for _, favDrinkID := range favDrinks {
		_, err := db.DB.Exec(query.DeleteUserFavDrink, userID, favDrinkID)
		if err != nil {
			return fmt.Errorf("error deleting fav drink: %v", err)
		}
	}
	return nil
}

func GetUserFavDrink(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserFavDrink, userID)
	log.Println("GetUserFavDrink started")
	if err != nil {
		return nil, fmt.Errorf("error selecting fav drink: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning fav drink: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserDiningStyle(userID string, diningStyles []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserDiningStyle started")
	for _, diningStyleID := range diningStyles {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, diningStyleID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserDiningStyle, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting dining style: %v", err)
	}
	log.Println("InsertUserDiningStyle success")
	return nil
}

func ExcludeUserDiningStyle(userID string, diningStyles []int) error {
	for _, diningStyleID := range diningStyles {
		_, err := db.DB.Exec(query.DeleteUserDiningStyle, userID, diningStyleID)
		if err != nil {
			return fmt.Errorf("error deleting dining style: %v", err)
		}
	}
	return nil
}

func GetUserDiningStyle(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserDiningStyle, userID)
	log.Println("GetUserDiningStyle started")
	if err != nil {
		return nil, fmt.Errorf("error selecting dining style: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning dining style: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func InsertUserFoodHabit(userID string, foodHabits []int) error {
	var placeholders []string
	var values []interface{}
	log.Println("InsertUserFoodHabit started")
	for _, foodHabitID := range foodHabits {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, foodHabitID)
	}
	if len(values) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserFoodHabit, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting food habit: %v", err)
	}
	log.Println("InsertUserFoodHabit success")
	return nil
}

func ExcludeUserFoodHabit(userID string, foodHabits []int) error {
	for _, foodHabitID := range foodHabits {
		_, err := db.DB.Exec(query.DeleteUserFoodHabit, userID, foodHabitID)
		if err != nil {
			return fmt.Errorf("error deleting food habit: %v", err)
		}
	}
	return nil
}

func GetUserFoodHabit(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserFoodHabit, userID)
	log.Println("GetUserFoodHabit started")
	if err != nil {
		return nil, fmt.Errorf("error selecting food habit: %v", err)
	}
	defer rows.Close()

	var prefrenceTableResponse []int
	for rows.Next() {
		var prefrenceID int
		if err := rows.Scan(&prefrenceID); err != nil {
			return nil, fmt.Errorf("error scanning food habit: %v", err)
		}
		prefrenceTableResponse = append(prefrenceTableResponse, prefrenceID)
	}
	return prefrenceTableResponse, nil
}

func UpsertUserFoodPreference(userID string, foodPreference *models.FoodPrefrenceRequestDTO) error {
	_, err := db.DB.Exec(query.UpsertUserFoodPreference, userID, strings.Join(foodPreference.FavoriteCuisineOther, ","), strings.Join(foodPreference.DietaryPreferenceOther, ","),
		strings.Join(foodPreference.FoodAllergyOther, ","), strings.Join(foodPreference.FavoriteDishesOther, ","), strings.Join(foodPreference.CookingStyleOther, ","),
		strings.Join(foodPreference.SpiceToleranceOther, ","), strings.Join(foodPreference.CookingLevelOther, ","), strings.Join(foodPreference.FavDrinkOther, ","), strings.Join(foodPreference.FoodHabitOther, ","))

	if err != nil {
		return fmt.Errorf("error updating user food preference: %v", err)
	}
	return nil
}

func GetUserFoodPreference(userID string) (*models.FoodPrefrenceRequestDTO, error) {
	var foodPreference models.UserFoodPreference
	var foodPreferenceResponse models.FoodPrefrenceRequestDTO
	var favoriteCuisines, dietaryPreferences, foodAllergies, favoriteDishes, cookingStyles, spiceLevels, favDrinks, diningStyles, foodHabits, cookingLevels sql.NullString

	err := db.DB.QueryRow(query.GetUserFoodPreference, userID).Scan(
		&foodPreference.FavoriteCuisineOther,
    &foodPreference.DietaryPreferenceOther, 
    &foodPreference.FoodAllergyOther,       
    &foodPreference.FavoriteDishesOther,    
    &foodPreference.CookingStyleOther,      
    &foodPreference.SpiceToleranceOther,    
    &foodPreference.CookingLevelOther,      
    &foodPreference.FavDrinkOther,          
    &foodPreference.FoodHabitOther,         
    &favoriteCuisines,                      
    &dietaryPreferences,                    
    &foodAllergies,                         
    &favoriteDishes,                        
    &cookingStyles,                         
    &spiceLevels,                           
    &cookingLevels,                         
    &favDrinks,                             
    &diningStyles,                          
    &foodHabits,              
	)
	if err == sql.ErrNoRows {
		fmt.Println("No food preferences found for user")
		return nil, fmt.Errorf("not found")
	}

	if err != nil {
		fmt.Println("Error getting Food Preference: ", err)
		return nil, err
	}
	if foodPreference.FavoriteCuisineOther != "" {
		foodPreferenceResponse.FavoriteCuisineOther = strings.Split(foodPreference.FavoriteCuisineOther, ",")
	}
	if foodPreference.DietaryPreferenceOther != "" {
		foodPreferenceResponse.DietaryPreferenceOther = strings.Split(foodPreference.DietaryPreferenceOther, ",")
	}
	if foodPreference.FoodAllergyOther != "" {
		foodPreferenceResponse.FoodAllergyOther = strings.Split(foodPreference.FoodAllergyOther, ",")
	}
	if foodPreference.FavoriteDishesOther != "" {
		foodPreferenceResponse.FavoriteDishesOther = strings.Split(foodPreference.FavoriteDishesOther, ",")
	}
	if foodPreference.CookingStyleOther != "" {
		foodPreferenceResponse.CookingStyleOther = strings.Split(foodPreference.CookingStyleOther, ",")
	}
	if foodPreference.SpiceToleranceOther != "" {
		foodPreferenceResponse.SpiceToleranceOther = strings.Split(foodPreference.SpiceToleranceOther, ",")
	}
	if foodPreference.CookingLevelOther != "" {
		foodPreferenceResponse.CookingLevelOther = strings.Split(foodPreference.CookingLevelOther, ",")
	}
	if foodPreference.FavDrinkOther != "" {
		foodPreferenceResponse.FavDrinkOther = strings.Split(foodPreference.FavDrinkOther, ",")
	}
	if foodPreference.FoodHabitOther != "" {
		foodPreferenceResponse.FoodHabitOther = strings.Split(foodPreference.FoodHabitOther, ",")
	}
	if favoriteCuisines.Valid {
		strIDs := strings.Split(favoriteCuisines.String, ",")
		foodPreferenceResponse.FavoriteCuisines = mapsFilterMap(strIDs)
	}
	if dietaryPreferences.Valid {
		strIDs := strings.Split(dietaryPreferences.String, ",")
		foodPreferenceResponse.DietaryPreferences = mapsFilterMap(strIDs)
	}
	if foodAllergies.Valid {
		strIDs := strings.Split(foodAllergies.String, ",")
		foodPreferenceResponse.FoodAllergies = mapsFilterMap(strIDs)
	}
	if favoriteDishes.Valid {
		strIDs := strings.Split(favoriteDishes.String, ",")
		foodPreferenceResponse.FavoriteDishes = mapsFilterMap(strIDs)
	}
	if cookingStyles.Valid {
		strIDs := strings.Split(cookingStyles.String, ",")
		foodPreferenceResponse.CookingStyles = mapsFilterMap(strIDs)
	}
	if spiceLevels.Valid {
		strIDs := strings.Split(spiceLevels.String, ",")
		foodPreferenceResponse.SpiceLevels = mapsFilterMap(strIDs)
	}
	if cookingLevels.Valid {
		strIDs := strings.Split(cookingLevels.String, ",")
		foodPreferenceResponse.CookingLevels = mapsFilterMap(strIDs)
	}
	if favDrinks.Valid {
		strIDs := strings.Split(favDrinks.String, ",")
		foodPreferenceResponse.FavDrinks = mapsFilterMap(strIDs)
	}
	if diningStyles.Valid {
		strIDs := strings.Split(diningStyles.String, ",")
		foodPreferenceResponse.DiningStyles = mapsFilterMap(strIDs)
	}
	if foodHabits.Valid {
		strIDs := strings.Split(foodHabits.String, ",")
		foodPreferenceResponse.FoodHabits = mapsFilterMap(strIDs)
	}
	return &foodPreferenceResponse, nil
}

func mapsFilterMap(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		if i, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}
