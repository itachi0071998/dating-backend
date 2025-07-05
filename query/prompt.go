package query

const(
	InsertUserPrompt = `
	INSERT INTO user_prompts (user_id, prompt_id, prompt)
	VALUES %s
	`

	DeleteUserPrompt = `
	DELETE FROM user_prompts WHERE user_id = ? AND prompt_id = ?
	`

	UpdateUserPrompt = `
	UPDATE user_prompts SET prompt = ? WHERE user_id = ? AND prompt_id = ?
	`

	GetUserPrompt = `
	SELECT 
		up.prompt_id,
		up.prompt
	FROM user_prompts up
	WHERE up.user_id = ?`

	SelectUserPrompt = `
	SELECT 
		up.prompt_id
	FROM user_prompts up
	WHERE up.user_id = ?`
)
