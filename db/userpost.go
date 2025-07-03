package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"foodieMatch/models"
	"foodieMatch/query"
	"strings"
	"time"
)

func CreateUserPost(userID string, userPostRequestDTO models.UserPostRequestDTO) (int, error) {
	tx, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var userPostID int

	result, err := tx.Exec(query.InsertUserPost, userID, userPostRequestDTO.Caption, userPostRequestDTO.IsPublic, time.Now(), userPostRequestDTO.RestaurantName, userPostRequestDTO.RestaurantLocation)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	userPostID = int(id)

	var placeholders []string
	var values []interface{}

	for _, media := range userPostRequestDTO.Media {
		placeholders = append(placeholders, "(?, ?, ?)")
		values = append(values, userPostID, media.URL, media.Type)
	}	
	mediaQuery := fmt.Sprintf(query.InsertPostMedia, strings.Join(placeholders, ","))
	_, err = tx.Exec(mediaQuery, values...)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return userPostID, nil
}

func GetAllUserPost(userID string) ([]models.UserPostResponseDTO, error) {
	rows, err := DB.Query(query.GetUserPosts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userPosts []models.UserPostResponseDTO

	for rows.Next() {
		var userPost models.UserPostResponseDTO
		var mediaJSON sql.NullString

		err := rows.Scan(&userPost.PostID, &userPost.UserID, 
			&userPost.Caption, &userPost.IsPublic, 
			&userPost.CreatedAt, &userPost.UpdatedAt, 
			&userPost.RestaurantName, &userPost.RestaurantLocation, &mediaJSON)
		if err != nil {
			return nil, err
		}
		if mediaJSON.Valid {
			err := json.Unmarshal([]byte(mediaJSON.String), &userPost.Media)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal media JSON: %w", err)
			}
		}

		userPosts = append(userPosts, userPost)
	}

	return userPosts, nil
}

func GetUserPostById(postID string) (models.UserPostResponseDTO, error) {
	var userPost models.UserPostResponseDTO
	var mediaJSON sql.NullString
	err := DB.QueryRow(query.GetUserPostById, postID).Scan(&userPost.PostID, &userPost.UserID, &userPost.Caption, &userPost.IsPublic, &userPost.CreatedAt, &userPost.UpdatedAt, &userPost.RestaurantName, &userPost.RestaurantLocation, &mediaJSON)
	if err==sql.ErrNoRows {
		return userPost, fmt.Errorf("post not found")
	}
	if err != nil {
		return userPost, err
	}
	if mediaJSON.Valid {
		err := json.Unmarshal([]byte(mediaJSON.String), &userPost.Media)
		if err != nil {
			return userPost, fmt.Errorf("failed to unmarshal media JSON: %w", err)
		}
	}

	return userPost, nil
}

func DeleteUserPost(postID string) error {
	_, err := DB.Exec(query.DeleteUserPost, postID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPost(userPostRequestDTO models.UserPostRequestDTO, postID string) error {
	_, err := DB.Exec(query.UpdateUserPost, 
		userPostRequestDTO.Caption, 
		userPostRequestDTO.IsPublic, 
		time.Now(), userPostRequestDTO.RestaurantName, 
		userPostRequestDTO.RestaurantLocation, postID)	
	if err != nil {
		return err
	}
	return nil
}
	
