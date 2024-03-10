package store

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ElnurKoke/web-api.git/internal/model"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (a *Store) CreateUser(user model.User) error {
	query := `INSERT INTO user(email, username, password) VALUES ($1, $2, $3);`
	_, err := a.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (a *Store) GetPasswordByUsername(username string) (string, error) {
	query := `SELECT password FROM user WHERE username = $1;`
	row := a.db.QueryRow(query, username)
	var password string
	if err := row.Scan(&password); err != nil {
		return password, err
	}
	return password, nil
}

func (a *Store) SaveToken(token string, expired time.Time, username string) error {
	query := `UPDATE user SET session_token = $1, expiresAt = $2 WHERE username = $3;`
	if _, err := a.db.Exec(query, token, expired, username); err != nil {
		return err
	}
	return nil
}

func (a *Store) DeleteToken(token string) error {
	query := `UPDATE user SET session_token = NULL, expiresAt = NULL WHERE session_token = $1`
	if _, err := a.db.Exec(query, token); err != nil {
		return err
	}
	return nil
}

func (a *Store) CheckUserByNameEmail(email, username string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE email = ? OR username = ?) AS UE_exists;"
	row := a.db.QueryRow(query, email, username)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (a *Store) CheckUser(user model.User) (string, time.Time, error) {
	password, err := a.GetPasswordByUsername(user.Username)
	if err != nil {
		return "", time.Time{}, errors.New(" There is no user with that name <" + user.Username + "> ")
	}
	if err := compareHashAndPassword(password, user.Password); err != nil {
		return "", time.Time{}, err
	}

	token := uuid.NewGen()
	d, err := token.NewV4()
	if err != nil {
		return "", time.Time{}, err
	}
	expired := time.Now().Add(time.Hour * 12)
	if err := a.SaveToken(d.String(), expired, user.Username); err != nil {
		return "", time.Time{}, err
	}
	return d.String(), expired, nil
}

func GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func compareHashAndPassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("password does not match")
	}
	return nil
}

func (a *Store) GetUserByToken(token string) (model.User, error) {
	query := `SELECT 
			id, 
			email, 
			username, 
			role,
			expiresAt
		FROM user 
		WHERE session_token = $1;`
	row := a.db.QueryRow(query, token)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Role, &user.ExpiresAt); err != nil {
		return model.User{}, err
	}
	return user, nil
}
