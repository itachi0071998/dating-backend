package query


const (
	InsertUserFoodPreference = `
	INSERT INTO user_food_preferences (user_id, favorite_cuisine_other, dietary_preference_other, food_allergy_other, favorite_dishes_other, cooking_style_other, spice_tolerance_other, cooking_level_other, fav_drink_other, food_dining_style_other, food_habit_other)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
	ufp.food_dining_style_other,
	ufp.food_habit_other,

    GROUP_CONCAT(DISTINCT c.name) AS favorite_cuisines,
    GROUP_CONCAT(DISTINCT dp.name) AS dietary_preferences,
    GROUP_CONCAT(DISTINCT fa.name) AS food_allergies,
    GROUP_CONCAT(DISTINCT d.name) AS favorite_dishes,
    GROUP_CONCAT(DISTINCT cs.name) AS cooking_styles,
    GROUP_CONCAT(DISTINCT sl.name) AS spice_levels,
    GROUP_CONCAT(DISTINCT cl.name) AS cooking_levels,
    GROUP_CONCAT(DISTINCT fd.name) AS fav_drinks,
    GROUP_CONCAT(DISTINCT ds.name) AS dining_styles,
    GROUP_CONCAT(DISTINCT fhb.name) AS food_habits,

    (
        SELECT JSON_ARRAYAGG(obj) 
        FROM (
            SELECT DISTINCT JSON_OBJECT('prompt', p.name, 'answer', up.prompt) AS obj
            FROM user_prompts up
            JOIN prompts p ON up.prompt_id = p.id
            WHERE up.user_id = ufp.user_id
        ) AS uniq_prompts
    ) AS prompts

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

LEFT JOIN user_cooking_levels ucl ON ufp.user_id = ucl.user_id
LEFT JOIN cooking_levels cl ON ucl.cooking_level_id = cl.id

LEFT JOIN user_fav_drinks ufd2 ON ufp.user_id = ufd2.user_id
LEFT JOIN fav_drinks fd ON ufd2.fav_drink_id = fd.id

LEFT JOIN user_dining_styles uds ON ufp.user_id = uds.user_id
LEFT JOIN dining_styles ds ON uds.dining_style_id = ds.id

LEFT JOIN user_food_habits ufhb ON ufp.user_id = ufhb.user_id
LEFT JOIN food_habits fhb ON ufhb.food_habit_id = fhb.id

WHERE ufp.user_id = ?
GROUP BY ufp.user_id

`

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

	InsertUserCookingLevel = `
	INSERT INTO user_cooking_levels (user_id, cooking_level_id)
	VALUES %s
	`

	InsertUserFavDrink = `
	INSERT INTO user_fav_drinks (user_id, fav_drink_id)
	VALUES %s
	`

	InsertUserPrompt = `
	INSERT INTO user_prompts (user_id, prompt_id, prompt)
	VALUES %s
	`

	InsertUserDiningStyle = `
	INSERT INTO user_dining_styles (user_id, dining_style_id)
	VALUES %s
	`

	InsertUserFoodHabit = `
	INSERT INTO user_food_habits (user_id, food_habit_id)
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

	DeleteUserCookingLevel = `
	DELETE FROM user_cooking_levels WHERE user_id = ? AND cooking_level_id = ?
	`

	DeleteUserFavDrink = `
	DELETE FROM user_fav_drinks WHERE user_id = ? AND fav_drink_id = ?
	`
	DeleteUserPrompt = `
	DELETE FROM user_prompts WHERE user_id = ? AND prompt_id = ?
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

