package service

import (
	"fmt"
	"foodieMatch/db"
	"foodieMatch/models"
	"foodieMatch/repository"
	"log"
	"sync"
)

func UserExists(userID string) (bool, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return false, fmt.Errorf("Error beginning transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	return repository.CheckUserExists(tx, userID)
}

func InsertUserProfile(userID string, userProfile *models.UserProfileRequest) error {
	log.Printf("Request started for user: %v", userID)
	userProfile.User.ID = userID

	if err := repository.InsertUser(&userProfile.User); err != nil {
		return err
	}
	log.Printf("InsertUser success for user ID: %v", userProfile.User.ID)

	if err := insertUserFoodPreference(userID, &userProfile.FoodPreference); err != nil {
		return err
	}
	if err := repository.InsertUserPrompt(userID, &userProfile.Prompts); err != nil {
		return err
	}
	return nil
}

func UpdateUserProfile(userID string, userProfile *models.UserProfileRequest) error {
	if err := repository.UpsertUser(&userProfile.User); err != nil {
		return err
	}
	if err := updateUserFoodPreference(userID, &userProfile.FoodPreference); err != nil {
		return err
	}
	if err := updateUserPrompts(userID, &userProfile.Prompts); err != nil {
		return err
	}
	return nil
}

func GetUserProfile(userID string) (*models.UserProfileResponse, error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	foodPreference, err := repository.GetUserFoodPreference(userID)
	if err != nil {
		return nil, err
	}
	prompts, err := repository.GetUserPrompts(userID)
	if err != nil {
		return nil, err
	}
	userProfile := &models.UserProfileResponse{
		User:           *user,
		FoodPreference: *foodPreference,
		Prompts:        prompts,
	}
	return userProfile, nil
}

func updateUserFoodPreference(userID string, userFoodPreference *models.FoodPrefrenceRequestDTO) error {

	wg := sync.WaitGroup{}
	errCh := make(chan error, 11)

	wg.Add(11)

	go func() {
		defer wg.Done()
		if err := repository.UpsertUserFoodPreference(userID, userFoodPreference); err != nil {
			errCh <- fmt.Errorf("UpsertUserFoodPreference: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserFavoriteCuisine(userID, userFoodPreference.FavoriteCuisines); err != nil {
			errCh <- fmt.Errorf("UpdateUserFavoriteCuisine: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserDietaryPreference(userID, userFoodPreference.DietaryPreferences); err != nil {
			errCh <- fmt.Errorf("UpdateUserDietaryPreference: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserFoodAllergy(userID, userFoodPreference.FoodAllergies); err != nil {
			errCh <- fmt.Errorf("UpdateUserFoodAllergy: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserFavoriteDish(userID, userFoodPreference.FavoriteDishes); err != nil {
			errCh <- fmt.Errorf("UpdateUserFavoriteDish: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserCookingStyle(userID, userFoodPreference.CookingStyles); err != nil {
			errCh <- fmt.Errorf("UpdateUserCookingStyle: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserSpiceTolerance(userID, userFoodPreference.SpiceLevels); err != nil {
			errCh <- fmt.Errorf("UpdateUserSpiceTolerance: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserCookingLevel(userID, userFoodPreference.CookingLevels); err != nil {
			errCh <- fmt.Errorf("UpdateUserCookingLevel: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserFavDrink(userID, userFoodPreference.FavDrinks); err != nil {
			errCh <- fmt.Errorf("UpdateUserFavDrink: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserDiningStyle(userID, userFoodPreference.DiningStyles); err != nil {
			errCh <- fmt.Errorf("UpdateUserDiningStyle: %w", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserFoodHabit(userID, userFoodPreference.FoodHabits); err != nil {
			errCh <- fmt.Errorf("UpdateUserFoodHabit: %w", err)
		}
	}()

	wg.Wait()
	close(errCh)
	var _err error
	for err := range errCh {
		_err = fmt.Errorf("%w\nupdateUserFoodPreference: %w", _err, err)
		log.Println("Error in updateUserFoodPreference: ", err)
	}
	return _err
}

func insertUserFoodPreference(userID string, userFoodPreference *models.FoodPrefrenceRequestDTO) error {
	wg := sync.WaitGroup{}
	errCh := make(chan error, 11) // 10 associative inserts

	wg.Add(11)

	go func() {
		log.Println("UpsertUserFoodPreference")
		defer wg.Done()
		if err := repository.UpsertUserFoodPreference(userID, userFoodPreference); err != nil {
			errCh <- fmt.Errorf("UpsertUserFoodPreference: %w", err)
		}
		log.Println("UpsertUserFoodPreference success")
	}()

	go func(){
		defer wg.Done()
		if err := repository.InsertUserFavoriteCuisine(userID, userFoodPreference.FavoriteCuisines); err != nil {
			errCh <- fmt.Errorf("InsertUserFavoriteCuisine: %w", err)
		}
	}()

	go func() {
		log.Println("InsertDietaryPreference")
		defer wg.Done()
		if err := repository.InsertDietaryPreference(userID, userFoodPreference.DietaryPreferences); err != nil {
			errCh <- fmt.Errorf("InsertDietaryPreference: %w", err)
		}
		log.Println("InsertDietaryPreference success")
	}()

	go func() {
		log.Println("InsertUserFoodAllergy")
		defer wg.Done()
		if err := repository.InsertUserFoodAllergy(userID, userFoodPreference.FoodAllergies); err != nil {
			errCh <- fmt.Errorf("InsertUserFoodAllergy: %w", err)
		}
		log.Println("InsertUserFoodAllergy success")
	}()

	go func() {
		log.Println("InsertUserFavoriteDish")
		defer wg.Done()
		if err := repository.InsertUserFavoriteDish(userID, userFoodPreference.FavoriteDishes); err != nil {
			errCh <- fmt.Errorf("InsertUserFavoriteDish: %w", err)
		}
		log.Println("InsertUserFavoriteDish success")
	}()

	go func() {
		log.Println("InsertUserCookingStyle")
		defer wg.Done()
		if err := repository.InsertUserCookingStyle(userID, userFoodPreference.CookingStyles); err != nil {
			errCh <- fmt.Errorf("InsertUserCookingStyle: %w", err)
		}
		log.Println("InsertUserCookingStyle success")
	}()

	go func() {
		log.Println("InsertUserSpiceTolerance")
		defer wg.Done()
		if err := repository.InsertUserSpiceTolerance(userID, userFoodPreference.SpiceLevels); err != nil {
			errCh <- fmt.Errorf("InsertUserSpiceTolerance: %w", err)
		}
		log.Println("InsertUserSpiceTolerance success")
	}()

	go func() {
		log.Println("InsertUserCookingLevel")
		defer wg.Done()
		if err := repository.InsertUserCookingLevel(userID, userFoodPreference.CookingLevels); err != nil {
			errCh <- fmt.Errorf("InsertUserCookingLevel: %w", err)
		}
		log.Println("InsertUserCookingLevel success")
	}()

	go func() {
		log.Println("InsertUserFavDrink")
		defer wg.Done()
		if err := repository.InsertUserFavDrink(userID, userFoodPreference.FavDrinks); err != nil {
			errCh <- fmt.Errorf("InsertUserFavDrink: %w", err)
		}
		log.Println("InsertUserFavDrink success")
	}()

	go func() {
		log.Println("InsertUserDiningStyle")
		defer wg.Done()
		if err := repository.InsertUserDiningStyle(userID, userFoodPreference.DiningStyles); err != nil {
			errCh <- fmt.Errorf("InsertUserDiningStyle: %w", err)
		}
		log.Println("InsertUserDiningStyle success")
	}()

	go func() {
		log.Println("InsertUserFoodHabit")
		defer wg.Done()
		if err := repository.InsertUserFoodHabit(userID, userFoodPreference.FoodHabits); err != nil {
			errCh <- fmt.Errorf("InsertUserFoodHabit: %w", err)
		}
		log.Println("InsertUserFoodHabit success")
	}()
	wg.Wait()
	close(errCh)
	for err := range errCh {
		log.Println("Error in insertUserFoodPreference: ", err)
		return err
	}

	return nil
}

func updateUserDietaryPreference(userID string, prefrences []int) error {
	currentDietaryPreference, err := repository.GetUserDietaryPreference(userID)
	if err != nil {
		return err
	}
	currentDietaryPreferenceIds, err := getFoodPrefrenceIds(currentDietaryPreference, prefrences)
	if err != nil {
		return err
	}
	if len(currentDietaryPreferenceIds.Add) > 0 {
		err = repository.InsertDietaryPreference(userID, currentDietaryPreferenceIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentDietaryPreferenceIds.Remove) > 0 {
		err = repository.ExcludeUserDietaryPreference(userID, currentDietaryPreferenceIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserFoodAllergy(userID string, prefrences []int) error {
	currentFoodAllergy, err := repository.GetUserFoodAllergy(userID)
	if err != nil {
		return err
	}
	currentFoodAllergyIds, err := getFoodPrefrenceIds(currentFoodAllergy, prefrences)
	if err != nil {
		return err
	}
	if len(currentFoodAllergyIds.Add) > 0 {
		err = repository.InsertUserFoodAllergy(userID, currentFoodAllergyIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentFoodAllergyIds.Remove) > 0 {
		err = repository.ExcludeUserFoodAllergy(userID, currentFoodAllergyIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserFavoriteDish(userID string, prefrences []int) error {
	currentFavoriteDish, err := repository.GetUserFavoriteDish(userID)
	if err != nil {
		return err
	}
	currentFavoriteDishIds, err := getFoodPrefrenceIds(currentFavoriteDish, prefrences)
	if err != nil {
		return err
	}
	if len(currentFavoriteDishIds.Add) > 0 {
		err = repository.InsertUserFavoriteDish(userID, currentFavoriteDishIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentFavoriteDishIds.Remove) > 0 {
		err = repository.ExcludeUserFavoriteDish(userID, currentFavoriteDishIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserCookingStyle(userID string, prefrences []int) error {
	currentCookingStyle, err := repository.GetUserCookingStyle(userID)
	if err != nil {
		return err
	}
	currentCookingStyleIds, err := getFoodPrefrenceIds(currentCookingStyle, prefrences)
	if err != nil {
		return err
	}
	if len(currentCookingStyleIds.Add) > 0 {
		err = repository.InsertUserCookingStyle(userID, currentCookingStyleIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentCookingStyleIds.Remove) > 0 {
		err = repository.ExcludeUserCookingStyle(userID, currentCookingStyleIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserSpiceTolerance(userID string, prefrences []int) error {
	currentSpiceTolerance, err := repository.GetUserSpiceTolerance(userID)
	if err != nil {
		return err
	}
	currentSpiceToleranceIds, err := getFoodPrefrenceIds(currentSpiceTolerance, prefrences)
	if err != nil {
		return err
	}
	if len(currentSpiceToleranceIds.Add) > 0 {
		err = repository.InsertUserSpiceTolerance(userID, currentSpiceToleranceIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentSpiceToleranceIds.Remove) > 0 {
		err = repository.ExcludeUserSpiceTolerance(userID, currentSpiceToleranceIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserCookingLevel(userID string, prefrences []int) error {
	currentCookingLevel, err := repository.GetUserCookingLevel(userID)
	if err != nil {
		return err
	}
	currentCookingLevelIds, err := getFoodPrefrenceIds(currentCookingLevel, prefrences)
	if err != nil {
		return err
	}
	if len(currentCookingLevelIds.Add) > 0 {
		err = repository.InsertUserCookingLevel(userID, currentCookingLevelIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentCookingLevelIds.Remove) > 0 {
		err = repository.ExcludeUserCookingLevel(userID, currentCookingLevelIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserFavoriteCuisine(userID string, prefrences []int) error {
	currentFavoriteCuisine, err := repository.GetUserFavoriteCuisine(userID)
	if err != nil {
		return err
	}
	currentFavoriteCuisineIds, err := getFoodPrefrenceIds(currentFavoriteCuisine, prefrences)
	if err != nil {
		return err
	}
	if len(currentFavoriteCuisineIds.Add) > 0 {
		err = repository.InsertUserFavoriteCuisine(userID, currentFavoriteCuisineIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentFavoriteCuisineIds.Remove) > 0 {
		err = repository.ExcludeUserFavoriteCuisine(userID, currentFavoriteCuisineIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserFavDrink(userID string, prefrences []int) error {
	currentFavDrink, err := repository.GetUserFavDrink(userID)
	if err != nil {
		return err
	}
	currentFavDrinkIds, err := getFoodPrefrenceIds(currentFavDrink, prefrences)
	if err != nil {
		return err
	}
	if len(currentFavDrinkIds.Add) > 0 {
		err = repository.InsertUserFavDrink(userID, currentFavDrinkIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentFavDrinkIds.Remove) > 0 {
		err = repository.ExcludeUserFavDrink(userID, currentFavDrinkIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserDiningStyle(userID string, prefrences []int) error {
	currentDiningStyle, err := repository.GetUserDiningStyle(userID)
	if err != nil {
		return err
	}
	currentDiningStyleIds, err := getFoodPrefrenceIds(currentDiningStyle, prefrences)
	if err != nil {
		return err
	}
	if len(currentDiningStyleIds.Add) > 0 {
		err = repository.InsertUserDiningStyle(userID, currentDiningStyleIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentDiningStyleIds.Remove) > 0 {
		err = repository.ExcludeUserDiningStyle(userID, currentDiningStyleIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateUserFoodHabit(userID string, prefrences []int) error {
	currentFoodHabit, err := repository.GetUserFoodHabit(userID)
	if err != nil {
		return err
	}
	currentFoodHabitIds, err := getFoodPrefrenceIds(currentFoodHabit, prefrences)
	if err != nil {
		return err
	}
	if len(currentFoodHabitIds.Add) > 0 {
		err = repository.InsertUserFoodHabit(userID, currentFoodHabitIds.Add)
		if err != nil {
			return err
		}
	}
	if len(currentFoodHabitIds.Remove) > 0 {
		err = repository.ExcludeUserFoodHabit(userID, currentFoodHabitIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}
func getFoodPrefrenceIds(current []int, to_update []int) (*models.PrefrenceTableIds, error) {
	var add []int
	var remove []int
	for _, id := range to_update {
		if !contains(current, id) {
			add = append(add, id)
		}
	}
	for _, id := range current {
		if !contains(to_update, id) {
			remove = append(remove, id)
		}
	}
	return &models.PrefrenceTableIds{Add: add, Remove: remove}, nil
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}


func updateUserPrompts(userID string, prompts *[]models.PromptDTO) error {
	currentPrompts, err := repository.GetUserPrompt(userID)
	toUpdatePromptIds := make([]int, len(*prompts))
	for i, prompt := range *prompts {
		toUpdatePromptIds[i] = prompt.PromptID
	}
	if err != nil {
		return err
	}
	currentPromptIds, err := getFoodPrefrenceIds(currentPrompts, toUpdatePromptIds)
	if err != nil {
		return err
	}
	err = repository.UpdateUserPrompt(userID, prompts)
	if err != nil {
		return err
	}
	if len(currentPromptIds.Remove) > 0 {
		err = repository.ExcludeUserPrompt(userID, currentPromptIds.Remove)
		if err != nil {
			return err
		}
	}
	return nil
}

