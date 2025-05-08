package models

import "time"

type UserPost struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Caption   string    `json:"caption"`
	Location  string    `json:"location"`
	IsPublic  bool      `json:"is_public" default:"true"`
	IsActive  bool      `json:"is_active" default:"false"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostMedia struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	MediaURL  string    `json:"media_url"`
	FileType  string    `json:"file_type"`
}

type PostTag struct {
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}

