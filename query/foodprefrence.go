package query


const (
	InsertUserFoodPreference = `
	INSERT INTO user_food_preferences (user_id, favorite_cuisine_other, dietary_preference_other, food_allergy_other, favorite_dishes_other, cooking_style_other, spice_tolerance_other)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	GetUserFoodPreference = `
	SELECT ufp.user_id,
    ufp.favorite_cuisine_other,
    ufp.dietary_preference_other,
    ufp.food_allergy_other,
    ufp.favorite_dishes_other,
    ufp.cooking_style_other,
    ufp.spice_tolerance_other,
    GROUP_CONCAT(DISTINCT c.name) AS favorite_cuisines,
    GROUP_CONCAT(DISTINCT dp.name) AS dietary_preferences,
    GROUP_CONCAT(DISTINCT fa.name) AS food_allergies,
    GROUP_CONCAT(DISTINCT d.name) AS favorite_dishes,
    GROUP_CONCAT(DISTINCT cs.name) AS cooking_styles,
    GROUP_CONCAT(DISTINCT sl.name) AS spice_levels
	FROM user_food_preferences ufp
	LEFT JOIN user_favorite_cuisines ufc ON ufp.user_id = ufc.user_id
	LEFT JOIN cuisines c ON ufc.cuisine_id = c.id
	LEFT JOIN user_dietary_preferences udp ON ufp.user_id = udp.user_id
	LEFT JOIN dietary_preferences dp ON udp.dietary_preference_id = dp.id
	LEFT JOIN user_food_allergies ufa ON ufp.user_id = ufa.user_id
	LEFT JOIN food_allergies fa ON ufa.allergy_id = fa.id
	LEFT JOIN user_favorite_dishes ufd ON ufp.user_id = ufd.user_id
	LEFT JOIN dishes d ON ufd.dish_id = d.id
	LEFT JOIN user_cooking_styles ucs ON ufp.user_id = ucs.user_id
	LEFT JOIN cooking_styles cs ON ucs.cooking_style_id = cs.id
	LEFT JOIN user_spice_levels usl ON ufp.user_id = usl.user_id
	LEFT JOIN spice_levels sl ON usl.spice_level_id = sl.id
	WHERE ufp.user_id = ?
	GROUP BY ufp.user_id;`

	InsertUserFavoriteCuisine = `
	INSERT INTO user_favorite_cuisines (user_id, cuisine_id)
	VALUES %s
	`
	InsertUserDietaryPreference = `
	INSERT INTO user_dietary_preferences (user_id, dietary_preference_id)
	VALUES %s
	`
	InsertUserFoodAllergy = `
	INSERT INTO user_food_allergies (user_id, allergy_id)
	VALUES %s
	`

	InsertUserFavoriteDish = `
	INSERT INTO user_favorite_dishes (user_id, dish_id)
	VALUES %s
	`
	InsertUserCookingStyle = `
	INSERT INTO user_cooking_styles (user_id, cooking_style_id)
	VALUES %s
	`

	InsertUserSpiceTolerance = `
	INSERT INTO user_spice_levels (user_id, spice_level_id)
	VALUES %s
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
	DeleteUserSpiceLevel = `
	DELETE FROM user_spice_levels WHERE user_id = ? AND spice_level_id = ?
	`
	UpdateUserFoodPreference = `
	UPDATE user_food_preferences SET favorite_cuisine_other = ?, dietary_preference_other = ?, food_allergy_other = ?, favorite_dishes_other = ?, cooking_style_other = ?, spice_tolerance_other = ? 
	WHERE user_id = ?
	`
)

