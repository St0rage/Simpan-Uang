package utils

const (
	// USER
	INSERT_USER          = "INSERT INTO users (id, name, email, password) VALUES (:id, :name, :email, :password)"
	SELECT_USER_ID       = "SELECT * FROM users WHERE id = $1"
	SELECT_USER_EMAIL    = "SELECT * FROM users WHERE email = $1"
	UPDATE_USER          = "UPDATE users SET (name, email) VALUES (:name, :email) WHERE id = :id"
	UPDATE_USER_PASSWORD = "UPDATE users SET (password) VALUES (:password) WHERE id = :id"
	// END USER

	// PIGGY BANK
	// edit here
	// END PIGGY BANK

	// WHISLIST
	// edit here
	// END WHISLIST

)
