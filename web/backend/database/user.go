package database

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	DB *sql.DB
}

func (d *UserDB) AddUser(username, password string) error {
	createTime := time.Now().Format("2006-01-02 15:04:05")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO employees(username, password, create_time) VALUES ($1, $2, $3)`
	_, err = d.DB.Exec(stmt, username, hashedPassword, createTime)
	if err != nil {
		return err
	}

	return nil
}

func (d *UserDB) Authenticate(username, password string) (int, error) {
	var id int
	var dbPassword []byte
	stmt := `SELECT id, password FROM employess WHERE username=$1`
	row := d.DB.QueryRow(stmt, username)
	if err := row.Scan(&id, &dbPassword); err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword(dbPassword, []byte(password)); err != nil {
		return 0, err
	}

	return id, nil
}
