package utils

const (
	// USER
	INSERT_USER          = "INSERT INTO users (id, name, email, password, is_admin) VALUES (:id, :name, :email, :password, :is_admin)"
	SELECT_USER_ID       = "SELECT * FROM users WHERE id = $1"
	SELECT_USER_EMAIL    = "SELECT * FROM users WHERE email = $1"
	UPDATE_USER          = "UPDATE users SET name = :name, email = :email WHERE id = :id"
	UPDATE_USER_PASSWORD = "UPDATE users SET password = :password WHERE id = :id"
	CHECK_ADMIN          = "SELECT COUNT(*) FROM users WHERE is_admin = true"
	IS_ADMIN             = "SELECT is_admin FROM users WHERE id = $1"
	CHECK_EMAIL          = "SELECT COUNT(*) FROM users WHERE email = $1"
	// END USER

	// PIGGY BANK
	INSERT_PIGGY_BANK         = "INSERT INTO piggy_bank (id, user_id, piggy_bank_name, type) VALUES (:id, :user_id, :piggy_bank_name, :type)"
	SELECT_PIGGY_BANK         = "SELECT * FROM piggy_bank WHERE user_id = $1"
	SELECT_PIGGY_BANK_ID      = "SELECT * FROM piggy_bank WHERE id = $1"
	UPDATE_PIGGY_BANK         = "UPDATE piggy_bank SET piggy_bank_name = :piggy_bank_name WHERE id = :id"
	CHECK_MAIN_PIGGY_BANK     = "SELECT COUNT(*) FROM piggy_bank WHERE user_id = $1"
	CHECK_PIGGY_BANK_NAME     = "SELECT COUNT(*) FROM piggy_bank WHERE piggy_bank_name = $1 AND user_id = $2"
	SELECT_PIGGY_BANK_USER_ID = "SELECT user_id FROM piggy_bank WHERE id = $1"
	INSERT_PIGGY_BANK_TRANSACTION = "INSERT INTO piggy_bank_transaction (id, piggy_bank_id, transaction_name, amount, status, date) VALUES (:id, :piggy_bank_id, :transaction_name, :amount, :status, :date)"
	SELECT_PIGGY_BANK_TRANSACTION = "SELECT * FROM piggy_bank_transaction WHERE piggy_bank_id = $1 ORDER BY date DESC LIMIT $2 OFFSET $3"
	SELECT_PIGGY_BANK_AMOUNT      = "SELECT amount FROM piggy_bank_transaction WHERE piggy_bank_id = $1"

	// END PIGGY BANK

	// WHISLIST
	INSERT_WISHLIST     = "INSERT INTO wishlist (id, user_id, wishlist_name, wishlist_target, progress) VALUES (:id, :user_id, :wishlist_name, :wishlist_target, :progress)"
	SELECT_WISHLIST     = "SELECT * FROM wishlist WHERE user_id = $1"
	SELECT_WISHLIST_ID  = "SELECT * FROM wishlist WHERE id = $1"
	CHECK_WISHLIST_NAME = "SELECT COUNT(*) FROM wishlist WHERE wishlist_name = $1 and user_id = $2"
	UPDATE_WISHLIST     = "UPDATE wishlist SET (wishlist_name) VALUES (:wishlist_name) WHERE id = :id"
	// END WHISLIST

)
