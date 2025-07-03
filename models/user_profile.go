package models

type UserProfileRequest struct {
	User User `json:"user"`
	FoodPreference FoodPrefrenceRequestDTO `json:"food_preference"`
	Prompts []PromptDTO `json:"prompts"`
}

type UserProfileResponse struct {
	User User `json:"user"`
	FoodPreference FoodPrefrenceRequestDTO `json:"food_preference"`
	Prompts []PromptDTO `json:"prompts"`
}


func (u *UserProfileRequest) Validate() error {
	if u.User.Validate() != nil {
		return u.User.Validate()
	}
	return nil
}

