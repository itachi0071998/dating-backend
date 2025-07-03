package db

import (
	"foodieMatch/models"
	"foodieMatch/query"
)



func GetCusines() ([]models.Cuisine, error) {
	rows, err := DB.Query(query.GetCusines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cuisines []models.Cuisine
	for rows.Next() {
		var cuisine models.Cuisine
		if err := rows.Scan(&cuisine.ID, &cuisine.Name); err != nil {
			return nil, err
		}
		cuisines = append(cuisines, cuisine)
	}
	return cuisines, nil
}	

func GetDietyPrefrence() ([]models.DietaryPreference, error) {
	rows, err := DB.Query(query.GetDietyPrefrence)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dietaryPreferences []models.DietaryPreference
	for rows.Next() {
		var dietaryPreference models.DietaryPreference
		if err := rows.Scan(&dietaryPreference.ID, &dietaryPreference.Name); err != nil {
			return nil, err
		}
		dietaryPreferences = append(dietaryPreferences, dietaryPreference)
	}
	return dietaryPreferences, nil
}

func GetFoodAllergy() ([]models.FoodAllergy, error) {
	rows, err := DB.Query(query.GetFoodAllergy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foodAllergies []models.FoodAllergy
	for rows.Next() {
		var foodAllergy models.FoodAllergy
		if err := rows.Scan(&foodAllergy.ID, &foodAllergy.Name); err != nil {
			return nil, err
		}
		foodAllergies = append(foodAllergies, foodAllergy)
	}
	return foodAllergies, nil
}

func GetDishes() ([]models.Dish, error) {
	rows, err := DB.Query(query.GetDishes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []models.Dish
	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.ID, &dish.Name); err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func GetCookingStyle() ([]models.CookingStyle, error) {
	rows, err := DB.Query(query.GetCookingStyle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cookingStyles []models.CookingStyle
	for rows.Next() {
		var cookingStyle models.CookingStyle
		if err := rows.Scan(&cookingStyle.ID, &cookingStyle.Name); err != nil {
			return nil, err
		}
		cookingStyles = append(cookingStyles, cookingStyle)
	}
	return cookingStyles, nil
}

func GetSpiceLevel() ([]models.SpiceLevel, error) {
	rows, err := DB.Query(query.GetSpiceLevel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spiceLevels []models.SpiceLevel
	for rows.Next() {
		var spiceLevel models.SpiceLevel
		if err := rows.Scan(&spiceLevel.ID, &spiceLevel.Name); err != nil {
			return nil, err
		}
		spiceLevels = append(spiceLevels, spiceLevel)
	}
	return spiceLevels, nil
}

func GetGalleryTag() ([]models.GalleryTag, error) {
	rows, err := DB.Query(query.GetGalleryTag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var galleryTags []models.GalleryTag
	for rows.Next() {
		var galleryTag models.GalleryTag
		if err := rows.Scan(&galleryTag.ID, &galleryTag.Name); err != nil {
			return nil, err
		}
		galleryTags = append(galleryTags, galleryTag)
	}
	return galleryTags, nil
}		

func GetCookingLevel() ([]models.CookingLevel, error) {
	rows, err := DB.Query(query.GetCookingLevel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cookingLevels []models.CookingLevel
	for rows.Next() {
		var cookingLevel models.CookingLevel
		if err := rows.Scan(&cookingLevel.ID, &cookingLevel.Name); err != nil {
			return nil, err
		}
		cookingLevels = append(cookingLevels, cookingLevel)
	}
	return cookingLevels, nil
}

func GetPrompt() ([]models.Prompt, error) {
	rows, err := DB.Query(query.GetPrompt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prompts []models.Prompt
	for rows.Next() {
		var prompt models.Prompt
		if err := rows.Scan(&prompt.ID, &prompt.Name); err != nil {
			return nil, err
		}
		prompts = append(prompts, prompt)
	}
	return prompts, nil
}

func GetDiningStyle() ([]models.DiningStyle, error) {
	rows, err := DB.Query(query.GetDiningStyle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var diningStyles []models.DiningStyle
	for rows.Next() {
		var diningStyle models.DiningStyle
		if err := rows.Scan(&diningStyle.ID, &diningStyle.Name); err != nil {
			return nil, err
		}
		diningStyles = append(diningStyles, diningStyle)
	}
	return diningStyles, nil
}

func GetFoodHabit() ([]models.FoodHabit, error) {
	rows, err := DB.Query(query.GetFoodHabit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foodHabits []models.FoodHabit
	for rows.Next() {
		var foodHabit models.FoodHabit
		if err := rows.Scan(&foodHabit.ID, &foodHabit.Name); err != nil {
			return nil, err
		}
		foodHabits = append(foodHabits, foodHabit)
	}
	return foodHabits, nil
}

func GetFavDrink() ([]models.FavDrink, error) {
	rows, err := DB.Query(query.GetFavDrink)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favDrinks []models.FavDrink
	for rows.Next() {
		var favDrink models.FavDrink
		if err := rows.Scan(&favDrink.ID, &favDrink.Name); err != nil {
			return nil, err
		}
		favDrinks = append(favDrinks, favDrink)
	}
	return favDrinks, nil
}
