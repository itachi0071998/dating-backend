package models

import "time"

type UserPost struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Caption   string    `json:"caption"`
	IsPublic  bool      `json:"is_public" default:"true"`
	IsActive  bool      `json:"is_active" default:"false"`
	UpdatedAt time.Time `json:"updated_at"`
	RestaurantName string    `json:"restaurant_name"`
	RestaurantLocation string    `json:"restaurant_location"`
}

type PostMedia struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	MediaURL  string    `json:"media_url"`
	FileType  string    `json:"file_type"`
}

type MediaItem struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}


type UserPostRequestDTO struct {
	Caption   string `json:"caption"`
	IsPublic  bool   `json:"is_public"`
	Media     []MediaItem 	`json:"media"`
	Tags      []int  `json:"tags"`
	RestaurantName string `json:"restaurant_name"`
	RestaurantLocation string `json:"restaurant_location"`
}

type UserPostResponseDTO struct {
	PostID                   int   `json:"post_id"`
	UserID                   string   `json:"user_id"`
	Caption                  string   `json:"caption"`
	IsPublic                 bool     `json:"is_public"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	Media     				[]MediaItem 	  `json:"media"`
	RestaurantName           string   `json:"restaurant_name"`
	RestaurantLocation       string   `json:"restaurant_location"`
}