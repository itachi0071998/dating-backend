package query

const (

	InsertUser = `
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
	`
	UpsertUser = `
	REPLACE INTO users (
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
	`


	UserExists = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`

	UpdateUser = `
		UPDATE users SET mobile=?, email=?, dob=?, first_name=?, last_name=?, gender=?, url=?
		WHERE id = ?`

	GetUserById = `
		SELECT id, mobile, email, dob, first_name, last_name, gender, url
		FROM users WHERE id = ?`
)
