package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var Db *sql.DB

type configuration struct {
	Driver string
	User   string
	Name   string
	Passwd string
}

type Todo struct {
	ID           int    `json:"id"`
	Description  string `json:"description"`
	Created_at   string `json:"created_at"`
	Completed_at string `json:"completed_at"`
}

var ErrNotF error = errors.New("not found")

// InitDB initialises db
func InitDB() error {
	// Reading configurations for db
	var err error
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	config := configuration{}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	//Setting up database
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s  sslmode=disable", config.User, config.Passwd, config.Name)
	Db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		return err
	}

	if err = Db.Ping(); err != nil {
		return err
	}

	// Creating tables in database
	file2, err := os.Open("migrations.sql")
	if err != nil {
		return err
	}

	defer file2.Close()

	data, err := ioutil.ReadAll(file2)
	if err != nil {
		return err
	}

	if _, err = Db.Exec(string(data)); err != nil {
		return err
	}

	return nil
}

//List ...
func List() (*[]Todo, error) {
	rows, err := Db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}

	list := []Todo{}
	for rows.Next() {
		item := Todo{}
		var complete interface{}
		err := rows.Scan(&item.ID, &item.Description, &item.Created_at, &complete)
		if err != nil {
			return nil, err
		}

		ok := false
		item.Completed_at, ok = complete.(string)
		if !ok {
			item.Completed_at = ""
		}

		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &list, nil
}

//PostItem ...
func (td *Todo) PostItem() error {
	row := Db.QueryRow(`
	INSERT INTO todo (description, created_at) 
	VALUES ($1, $2) RETURNING id`, td.Description, td.Created_at)
	err := row.Scan(&td.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateItem ...
func (td *Todo) UpdateItem(id *int) error {
	t := time.Now().Format("2006-02-01 15:04:05")
	row := Db.QueryRow(`
	UPDATE todo
	SET completed_at = $1
	WHERE id = $2
	RETURNING description, created_at`, t, *id)
	err := row.Scan(&td.Description, &td.Created_at)
	if err != nil {
		return fmt.Errorf("UpdateItem: %w", ErrNotF)
	}

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

//DeleteItem ...
func DeleteItem(id *int) error {
	_, err := Db.Exec(`DELETE FROM todo WHERE id = $1`, *id)
	if err != nil {
		return err
	}

	return nil
}
