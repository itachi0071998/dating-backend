package db

import (
	"database/sql"
	"fmt"
	"foodieMatch/models"
	"foodieMatch/query"
	"strings"
)

func InsertUserPrompt(userID string, prompts []models.PromptDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err := insertUserPrompt(tx, userID, prompts); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func insertUserPrompt(tx *sql.Tx, userID string, prompts []models.PromptDTO) error {
	placeholders := make([]string, 0)
	values := make([]interface{}, 0)
	for _, prompt := range prompts {
		placeholders = append(placeholders, "(?, ?, ?)")
		values = append(values, userID, prompt.PromptID, prompt.Answer)
	}
	if len(placeholders) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.InsertUserPrompt, strings.Join(placeholders, ","))
	_, err := tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting prompts: %v", err)
	}
	return nil
}

func deleteUserPrompt(tx *sql.Tx, userID string, prompts []int) error {
	placeholders := make([]string, 0)
	values := make([]interface{}, 0)
	for _, prompt := range prompts {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userID, prompt)
	}
	if len(placeholders) == 0 {
		return nil
	}
	query := fmt.Sprintf(query.DeleteUserPrompt, strings.Join(placeholders, ","))
	_, err := tx.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error deleting prompts: %v", err)
	}
	return nil
}

func updateUserPrompts(tx *sql.Tx, userID string, prompts []models.PromptDTO) error {
	for _, prompt := range prompts {
		_, err := tx.Exec(query.UpdateUserPrompt, prompt.Answer, userID, prompt.PromptID)
		if err != nil {
			return fmt.Errorf("error updating prompts: %v", err)
		}
	}
	return nil
}

func UpdateUserPrompt(userID string, prompts models.PromptUpdateRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err := updateUserPrompts(tx, userID, prompts.Update); err != nil {
		return err
	}
	if err := deleteUserPrompt(tx, userID, prompts.Remove); err != nil {
		return err
	}
	if err := insertUserPrompt(tx, userID, prompts.Add); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func getUserPrompt(tx *sql.Tx, userID string)	(models.PromptResponseDTO, error) {
	rows, err := tx.Query(query.GetUserPrompt, userID)
	if err != nil {
		return models.PromptResponseDTO{}, fmt.Errorf("error getting prompts: %v", err)
	}
	defer rows.Close()

	prompts := make([]models.PromptDTO, 0)
	for rows.Next() {
		var prompt models.PromptDTO
		if err := rows.Scan(&prompt.PromptID, &prompt.Answer); err != nil {
			return models.PromptResponseDTO{}, fmt.Errorf("error scanning prompts: %v", err)
		}
		prompts = append(prompts, prompt)
	}
	return models.PromptResponseDTO{UserId: userID, Prompts: prompts}, nil
}

func GetUserPrompt(userID string) (models.PromptResponseDTO, error) {
	tx, err := DB.Begin()
	if err != nil {
		return models.PromptResponseDTO{}, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	return getUserPrompt(tx, userID)
}


