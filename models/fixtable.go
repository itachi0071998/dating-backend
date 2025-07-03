package models

type Cuisine struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DietaryPreference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FoodAllergy struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Dish struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CookingStyle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SpiceLevel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GalleryTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DiningStyle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FoodHabit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FavDrink struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CookingLevel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Prompt struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
