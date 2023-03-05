# Simpan Uang

Simpan Uang helps you achieve, allocate, and set targets for your savings in a piggy bank. this project uses the gin-gonic framework and some other golang libraries.

## Database

This project uses PostgreSQL as the database

```sql
CREATE TABLE users (
	id VARCHAR(50) NOT NULL,
	name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password CHAR(76) NOT NULL,
	is_admin BOOL NOT NULL,
	avatar VARCHAR(50) NOT NULL,

	PRIMARY KEY(id)
);

CREATE TABLE piggy_bank (
	id VARCHAR(50) NOT NULL,
	user_id VARCHAR(50) NOT NULL,
	piggy_bank_name VARCHAR(15) NOT NULL,
	type BOOL NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE piggy_bank_transaction (
	id VARCHAR(50) NOT NULL,
	piggy_bank_id VARCHAR(50) NOT NULL,
	transaction_name VARCHAR(15) NOT NULL,
	amount NUMERIC(15, 2),
	status BOOL NOT NULL,
	date INT NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(piggy_bank_id) REFERENCES piggy_bank(id) ON DELETE CASCADE
);

CREATE TABLE wishlist (
	id VARCHAR(50) NOT NULL,
	user_id VARCHAR(50) NOT NULL,
	wishlist_name VARCHAR(15) NOT NULL,
	wishlist_target NUMERIC(15,2) NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE wishlist_transaction (
	id VARCHAR(50) NOT NULL,
	wishlist_id VARCHAR(50) NOT NULL,
	transaction_name VARCHAR(15) NOT NULL,
	amount NUMERIC(15, 2),
	status BOOL NOT NULL,
	date INT NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(wishlist_id) REFERENCES wishlist(id) ON DELETE CASCADE
);
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file, or simply change env.example to .env

`DB_HOST`

`DB_PORT`

`DB_USER`

`DB_PASSWORD`

`DB_DRIVER`

`DB_NAME`

`API_HOST`

`API_PORT`

`MAIL_HOST`

`MAIL_PORT`

`MAIL_SENDER`

`MAIL_USERNAME`

`MAIL_PASSWORD`

## Run Locally

Clone the project

```bash
  git clone https://github.com/St0rage/Simpan-Uang
```

Go to the project directory

```bash
  cd simpan-uang
```

Start the server

```bash
  go run main.go
```
