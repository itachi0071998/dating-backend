package db

import (
	"database/sql"
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

	result, err := tx.Exec(query.InsertUserPost, userID, userPostRequestDTO.Caption, userPostRequestDTO.Location, userPostRequestDTO.IsPublic, time.Now())
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

	placeholders = nil	
	values = nil

	for _, tag := range userPostRequestDTO.Tags {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, userPostID, tag)
	}
	tagQuery := fmt.Sprintf(query.InsertPostTag, strings.Join(placeholders, ","))
	_, err = tx.Exec(tagQuery, values...)
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
		var tags sql.NullString

		err := rows.Scan(&userPost.PostID, &userPost.UserID, &userPost.Caption, &userPost.Location, &userPost.IsPublic, &userPost.CreatedAt, &userPost.UpdatedAt, &userPost.MediaURL, &userPost.FileType, &tags)
		if err != nil {
			return nil, err
		}
		if tags.Valid {
			userPost.Tags = strings.Split(tags.String, ",")
		}
		userPosts = append(userPosts, userPost)
	}

	return userPosts, nil
}

func GetUserPostById(postID string) (models.UserPostResponseDTO, error) {
	var userPost models.UserPostResponseDTO
	var tags sql.NullString
	err := DB.QueryRow(query.GetUserPostById, postID).Scan(&userPost.PostID, &userPost.UserID, &userPost.Caption, &userPost.Location, &userPost.IsPublic, &userPost.CreatedAt, &userPost.UpdatedAt, &userPost.MediaURL, &userPost.FileType, &tags)
	if err==sql.ErrNoRows {
		return userPost, fmt.Errorf("post not found")
	}
	if err != nil {
		return userPost, err
	}
	if tags.Valid {
		userPost.Tags = strings.Split(tags.String, ",")
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
	
func UpdateUserPostTag(updatePostTag *models.UpdatePostTagRequestDTO) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}	
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var placeholders []string		
	var values []interface{}

	for _, tagID := range updatePostTag.AddTags {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, updatePostTag.PostID, tagID)
	}
	if len(placeholders) > 0 {
		tagQuery := fmt.Sprintf(query.InsertPostTag, strings.Join(placeholders, ","))
		fmt.Println("tagQuery: ", tagQuery)
		_, err = tx.Exec(tagQuery, values...)
		if err != nil {
			return fmt.Errorf("error inserting tag: %v", err)
		}
	}	

	for _, tagID := range updatePostTag.RemoveTags {
		_, err = tx.Exec(query.DeletePostTag, updatePostTag.PostID, tagID)
		if err != nil {
			return fmt.Errorf("error deleting tag: %v", err)
		}
	}
	
	_, err = tx.Exec(query.UpdateUserPostUpdatedTime, time.Now(), updatePostTag.PostID)
	if err != nil {
		return fmt.Errorf("error updating user post updated time: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}

func UpdateUserPost(userPostRequestDTO models.UpdateUserPostRequestDTO) error {
	_, err := DB.Exec(query.UpdateUserPost, userPostRequestDTO.Caption, userPostRequestDTO.Location, userPostRequestDTO.IsPublic, time.Now(), userPostRequestDTO.PostID)	
	if err != nil {
		return err
	}
	return nil
}
	
