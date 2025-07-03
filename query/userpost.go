package query


const (
	InsertUserPost = `
	INSERT INTO user_posts (user_id, caption, is_public, updated_at, restaurant_name, restaurant_location)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	InsertPostMedia = `
	INSERT INTO post_media (post_id, media_url, file_type)
	VALUES %s
	`

	UpdateUserPost = `
	UPDATE user_posts SET caption = ?, is_public = ?, updated_at = ?, restaurant_name = ?, restaurant_location = ?
	WHERE id = ?
	`

	UpdateUserPostUpdatedTime = `
	UPDATE user_posts SET updated_at = ?
	WHERE id = ?
	`

	DeleteUserPost = `
	UPDATE user_posts SET is_active = FALSE
	WHERE id = ?
	`

	GetUserPosts = `
	SELECT 
    up.id AS post_id,
	up.user_id,
    up.caption,
    up.is_public,
    up.created_at,
    up.updated_at,
	up.restaurant_name,
	up.restaurant_location,
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'url', pm.media_url,
            'type', pm.file_type
        )
    ) AS media
	FROM user_posts up
	LEFT JOIN post_media pm ON up.id = pm.post_id
	WHERE up.user_id = ?
	AND up.is_active = TRUE
	GROUP BY up.id
	ORDER BY up.created_at DESC;
	`

	GetUserPostById = `
	SELECT up.id AS post_id,
    up.user_id,
    up.caption,
    up.is_public,
    up.created_at,
    up.updated_at,
	up.restaurant_name,
	up.restaurant_location,
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'url', pm.media_url,
            'type', pm.file_type
        )
    ) AS media
	FROM user_posts up
	LEFT JOIN post_media pm ON up.id = pm.post_id
	WHERE up.id = ?	
	AND up.is_active = TRUE
	GROUP BY up.id
	ORDER BY up.created_at DESC;
	`
)