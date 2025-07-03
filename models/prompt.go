package models

type PromptDTO struct {
	PromptID int `json:"prompt_id"`
	Answer string `json:"answer"`
}


type PromptUpdateRequestDTO struct {
	Add []PromptDTO `json:"add"`
	Remove []int `json:"remove"`
	Update []PromptDTO `json:"update"`
}

type Prompts struct {
	Prompts []PromptDTO `json:"prompts"`
}

type PromptResponseDTO struct {
	UserId string `json:"user_id"`
	Prompts []PromptDTO `json:"prompts"`
}