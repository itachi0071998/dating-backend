package repository

import (
	"foodieMatch/db"
	"foodieMatch/models"
	"fmt"
	"strings"
	"foodieMatch/query"
)

func InsertUserPrompt(userID string, prompts *[]models.PromptDTO) error {
	placeholders := make([]string, 0)
	values := make([]interface{}, 0)
	for _, prompt := range *prompts {
		placeholders = append(placeholders, "(?, ?, ?)")
		values = append(values, userID, prompt.PromptID, prompt.Answer)
	}
	if len(placeholders) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserPrompt, strings.Join(placeholders, ","))
	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting prompts: %v", err)
	}
	return nil
}

func ExcludeUserPrompt(userID string, prompts []int) error {
	for _, promptID := range prompts {
		_, err := db.DB.Exec(query.DeleteUserPrompt, userID, promptID)
		if err != nil {
			return fmt.Errorf("error deleting prompt: %v", err)
		}
	}
	return nil
}

func GetUserPrompt(userID string) ([]int, error) {
	rows, err := db.DB.Query(query.SelectUserPrompt, userID)
	if err != nil {
		return nil, fmt.Errorf("error selecting prompts: %v", err)
	}
	defer rows.Close()

	var prompts []int
	for rows.Next() {
		var promptID int
		if err := rows.Scan(&promptID); err != nil {
			return nil, fmt.Errorf("error scanning prompts: %v", err)
		}
		prompts = append(prompts, promptID)
	}
	return prompts, nil
}

func UpdateUserPrompt(userID string, prompts *[]models.PromptDTO) error {
	for _, prompt := range *prompts {
		_, err := db.DB.Exec(query.UpdateUserPrompt, userID, prompt.PromptID, prompt.Answer)
		if err != nil {
			return fmt.Errorf("error updating prompt: %v", err)
		}
	}
	return nil
}

func GetUserPrompts(userID string) ([]models.PromptDTO, error) {
	rows, err := db.DB.Query(query.GetUserPrompt, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting prompts: %v", err)
	}
	defer rows.Close()

	prompts := make([]models.PromptDTO, 0)
	for rows.Next() {
		var prompt models.PromptDTO
		if err := rows.Scan(&prompt.PromptID, &prompt.Answer); err != nil {
			return nil, fmt.Errorf("error scanning prompts: %v", err)
		}
		prompts = append(prompts, prompt)
	}
	return prompts, nil
}
