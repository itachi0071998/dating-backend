package query

const (
	UpsertUser = `
	INSERT INTO users (
		id,
		mobile,
		email,
		dob,
		first_name,
		last_name,
		gender,
		url
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		mobile = VALUES(mobile),
		email = VALUES(email),
		dob = VALUES(dob),
		first_name = VALUES(first_name),
		last_name = VALUES(last_name),
		gender = VALUES(gender),
		url = VALUES(url)
	`


	UserExists = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`

	UpdateUser = `
		UPDATE users SET id=?, mobile=?, email=?, dob=?, first_name=?, last_name=?, gender=?, url=?
		WHERE id = ?`

	GetUserById = `
		SELECT id, mobile, email, dob, first_name, last_name, gender, url
		FROM users WHERE id = ?`
)
