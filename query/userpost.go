package query


const (
	InsertUserPost = `
	INSERT INTO user_posts (user_id, caption, location, is_public, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`

	InsertPostMedia = `
	INSERT INTO post_media (post_id, media_url, file_type)
	VALUES %s
	`

	InsertPostTag = `
	INSERT INTO post_tags (post_id, tag_id)
	VALUES %s
	`

	UpdateUserPost = `
	UPDATE user_posts SET caption = ?, location = ?, is_public = ?, updated_at = ?
	WHERE id = ?
	`

	UpdateUserPostUpdatedTime = `
	UPDATE user_posts SET updated_at = ?
	WHERE id = ?
	`
	
	DeletePostTag = `
	DELETE FROM post_tags WHERE post_id = ? AND tag_id = ?
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
    up.location,
    up.is_public,
    up.created_at,
    up.updated_at,
    ANY_VALUE(pm.media_url) AS media_url,
    ANY_VALUE(pm.file_type) AS file_type,
    GROUP_CONCAT(gt.name) AS tags
	FROM user_posts up
	LEFT JOIN post_media pm ON up.id = pm.post_id
	LEFT JOIN post_tags pt ON up.id = pt.post_id
	LEFT JOIN gallery_tags gt ON pt.tag_id = gt.id
	WHERE up.user_id = ?
	AND up.is_active = TRUE
	GROUP BY up.id
	ORDER BY up.created_at DESC;
	`

	GetUserPostById = `
	SELECT up.id AS post_id,
    up.user_id,
    up.caption,
    up.location,
    up.is_public,
    up.created_at,
    up.updated_at,
    ANY_VALUE(pm.media_url) AS media_url,
    ANY_VALUE(pm.file_type) AS file_type,
    GROUP_CONCAT(gt.name) AS tags
	FROM user_posts up
	LEFT JOIN post_media pm ON up.id = pm.post_id
	LEFT JOIN post_tags pt ON up.id = pt.post_id
	LEFT JOIN gallery_tags gt ON pt.tag_id = gt.id
	WHERE up.id = ?	
	AND up.is_active = TRUE
	GROUP BY up.id
	ORDER BY up.created_at DESC;
	`
)