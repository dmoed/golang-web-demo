package model

import "database/sql"

//User model
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Enabled      bool   `json:"enabled"`
	PasswordHash string
	CreatedAt    string `json:"created_at"`
}

func FindUserByID(db *sql.DB, ID int) (*User, error) {

	result := db.QueryRow("select id, username, email, enabled, password, created_at from user where id = ? ", ID)

	//todo switch error result

	user := &User{}

	err := result.Scan(&user.ID, &user.Username, &user.Email, &user.Enabled, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByEmail(db *sql.DB, email string) (*User, error) {

	result := db.QueryRow("select id, username, email, enabled, password, created_at from user where email = ? ", email)

	//todo switch error result

	user := &User{}

	err := result.Scan(&user.ID, &user.Username, &user.Email, &user.Enabled, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
