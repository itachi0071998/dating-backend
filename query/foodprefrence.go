package query


const (
	UpsertUserFoodPreference = `
	INSERT INTO user_food_preferences (
		user_id,
		favorite_cuisine_other,
		dietary_preference_other,
		food_allergy_other,
		favorite_dishes_other,
		cooking_style_other,
		spice_tolerance_other,
		cooking_level_other,
		fav_drink_other,
		food_habit_other
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		favorite_cuisine_other = VALUES(favorite_cuisine_other),
		dietary_preference_other = VALUES(dietary_preference_other),
		food_allergy_other = VALUES(food_allergy_other),
		favorite_dishes_other = VALUES(favorite_dishes_other),
		cooking_style_other = VALUES(cooking_style_other),
		spice_tolerance_other = VALUES(spice_tolerance_other),
		cooking_level_other = VALUES(cooking_level_other),
		fav_drink_other = VALUES(fav_drink_other),
		food_habit_other = VALUES(food_habit_other)
	`

	GetUserFoodPreference = `SELECT 
    ufp.user_id,
    ufp.favorite_cuisine_other,
    ufp.dietary_preference_other,
    ufp.food_allergy_other,
    ufp.favorite_dishes_other,
    ufp.cooking_style_other,
    ufp.spice_tolerance_other,
    ufp.cooking_level_other,
    ufp.fav_drink_other,
    ufp.food_habit_other,

    GROUP_CONCAT(DISTINCT ufc.cuisine_id) AS favorite_cuisine_ids,
    GROUP_CONCAT(DISTINCT udp.dietary_preference_id) AS dietary_preference_ids,
    GROUP_CONCAT(DISTINCT ufa.allergy_id) AS food_allergy_ids,
    GROUP_CONCAT(DISTINCT ufd.dish_id) AS favorite_dish_ids,
    GROUP_CONCAT(DISTINCT ucs.cooking_style_id) AS cooking_style_ids,
    GROUP_CONCAT(DISTINCT usl.spice_level_id) AS spice_level_ids,
    GROUP_CONCAT(DISTINCT ucl.cooking_level_id) AS cooking_level_ids,
    GROUP_CONCAT(DISTINCT ufd2.fav_drink_id) AS fav_drink_ids,
    GROUP_CONCAT(DISTINCT uds.dining_style_id) AS dining_style_ids,
    GROUP_CONCAT(DISTINCT ufhb.food_habit_id) AS food_habit_ids

FROM user_food_preferences ufp

LEFT JOIN user_favorite_cuisines ufc ON ufp.user_id = ufc.user_id
LEFT JOIN user_dietary_preferences udp ON ufp.user_id = udp.user_id
LEFT JOIN user_food_allergies ufa ON ufp.user_id = ufa.user_id
LEFT JOIN user_favorite_dishes ufd ON ufp.user_id = ufd.user_id
LEFT JOIN user_cooking_styles ucs ON ufp.user_id = ucs.user_id
LEFT JOIN user_spice_levels usl ON ufp.user_id = usl.user_id
LEFT JOIN user_cooking_levels ucl ON ufp.user_id = ucl.user_id
LEFT JOIN user_fav_drinks ufd2 ON ufp.user_id = ufd2.user_id
LEFT JOIN user_dining_styles uds ON ufp.user_id = uds.user_id
LEFT JOIN user_food_habits ufhb ON ufp.user_id = ufhb.user_id

WHERE ufp.user_id = ?
GROUP BY ufp.user_id

`

	InsertUserFavoriteCuisine = `
	INSERT INTO user_favorite_cuisines (user_id, cuisine_id)
	VALUES %s
	`

	SelectUserFavoriteCuisine = `
	SELECT 
		ufc.cuisine_id
	FROM user_favorite_cuisines ufc
	WHERE ufc.user_id = ?
	`
	InsertUserDietaryPreference = `
	INSERT INTO user_dietary_preferences (user_id, dietary_preference_id)
	VALUES %s
	`
	SelectUserDietaryPreference = `
	SELECT 
		udp.dietary_preference_id
	FROM user_dietary_preferences udp
	WHERE udp.user_id = ?
	`

	InsertUserFoodAllergy = `
	INSERT INTO user_food_allergies (user_id, allergy_id)
	VALUES %s
	`

	SelectUserFoodAllergy = `
	SELECT 
		ufa.allergy_id
	FROM user_food_allergies ufa
	WHERE ufa.user_id = ?
	`

	InsertUserFavoriteDish = `
	INSERT INTO user_favorite_dishes (user_id, dish_id)
	VALUES %s
	`
	SelectUserFavoriteDish = `
	SELECT 
		ufd.dish_id
	FROM user_favorite_dishes ufd
	WHERE ufd.user_id = ?
	`
	InsertUserCookingStyle = `
	INSERT INTO user_cooking_styles (user_id, cooking_style_id)
	VALUES %s
	`
	SelectUserCookingStyle = `
	SELECT 
		ucs.cooking_style_id
	FROM user_cooking_styles ucs
	WHERE ucs.user_id = ?
	`

	InsertUserSpiceTolerance = `
	INSERT INTO user_spice_levels (user_id, spice_level_id)
	VALUES %s
	`
	SelectUserSpiceTolerance = `
	SELECT 
		usl.spice_level_id
	FROM user_spice_levels usl
	WHERE usl.user_id = ?
	`

	InsertUserCookingLevel = `
	INSERT INTO user_cooking_levels (user_id, cooking_level_id)
	VALUES %s
	`
	SelectUserCookingLevel = `
	SELECT 
		ucl.cooking_level_id
	FROM user_cooking_levels ucl
	WHERE ucl.user_id = ?
	`

	InsertUserFavDrink = `
	INSERT INTO user_fav_drinks (user_id, fav_drink_id)
	VALUES %s
	`
	SelectUserFavDrink = `
	SELECT 
		ufd2.fav_drink_id
	FROM user_fav_drinks ufd2
	WHERE ufd2.user_id = ?
	`

	InsertUserDiningStyle = `
	INSERT INTO user_dining_styles (user_id, dining_style_id)
	VALUES %s
	`
	SelectUserDiningStyle = `
	SELECT 
		uds.dining_style_id
	FROM user_dining_styles uds
	WHERE uds.user_id = ?
	`

	InsertUserFoodHabit = `
	INSERT INTO user_food_habits (user_id, food_habit_id)
	VALUES %s
	`
	SelectUserFoodHabit = `
	SELECT 
		ufhb.food_habit_id
	FROM user_food_habits ufhb
	WHERE ufhb.user_id = ?
	`

	DeleteUserFoodPreference = `
	DELETE FROM user_food_preferences WHERE user_id = ?
	`
	DeleteUserFavoriteCuisine = `
	DELETE FROM user_favorite_cuisines WHERE user_id = ? AND cuisine_id = ?
	`
	DeleteUserDietaryPreference = `
	DELETE FROM user_dietary_preferences WHERE user_id = ? AND dietary_preference_id = ?
	`		
	DeleteUserFoodAllergy = `
	DELETE FROM user_food_allergies WHERE user_id = ? AND allergy_id = ?
	`
	DeleteUserFavoriteDish = `
	DELETE FROM user_favorite_dishes WHERE user_id = ? AND dish_id = ?
	`	
	DeleteUserCookingStyle = `
	DELETE FROM user_cooking_styles WHERE user_id = ? AND cooking_style_id = ?
	`
	DeleteUserSpiceTolerance = `
	DELETE FROM user_spice_levels WHERE user_id = ? AND spice_level_id = ?
	`

	DeleteUserCookingLevel = `
	DELETE FROM user_cooking_levels WHERE user_id = ? AND cooking_level_id = ?
	`

	DeleteUserFavDrink = `
	DELETE FROM user_fav_drinks WHERE user_id = ? AND fav_drink_id = ?
	`

	DeleteUserDiningStyle = `
	DELETE FROM user_dining_styles WHERE user_id = ? AND dining_style_id = ?
	`

	DeleteUserFoodHabit = `
	DELETE FROM user_food_habits WHERE user_id = ? AND food_habit_id = ?
	`

	UpdateUserFoodPreference = `
	UPDATE user_food_preferences SET favorite_cuisine_other = ?, dietary_preference_other = ?, food_allergy_other = ?, favorite_dishes_other = ?, cooking_style_other = ?, spice_tolerance_other = ?, cooking_level_other = ?, fav_drink_other = ?, food_dining_style_other = ?, food_habit_other = ?
	WHERE user_id = ?
	`
)

