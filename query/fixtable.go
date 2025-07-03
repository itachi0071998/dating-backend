package query


const (
	GetCusines = `SELECT id, name FROM cuisines`

	GetDietyPrefrence = `SELECT id, name FROM dietary_preferences`

	GetFoodAllergy = `SELECT id, name FROM food_allergies`

	GetDishes = `SELECT id, name FROM dishes`

	GetCookingStyle = `SELECT id, name FROM cooking_styles`

	GetSpiceLevel = `SELECT id, name FROM spice_levels`

	GetCookingLevel = `SELECT id, name FROM cooking_levels`

	GetFavDrink = `SELECT id, name FROM fav_drinks`

	GetDiningStyle = `SELECT id, name FROM dining_styles`

	GetFoodHabit = `SELECT id, name FROM food_habits`

	GetPrompt = `SELECT id, name FROM prompts`

	GetGalleryTag = `SELECT id, name FROM gallery_tags`
)