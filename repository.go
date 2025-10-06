package db

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	store *sql.DB
}

func NewRepository(instance *sql.DB) *Repository {
	return &Repository{store: instance}
}

// added returning error so we have conflict with contracts ##
func (r *Repository) Set(key, value string) error {
	_, err := r.store.Exec("SELECT set($1, $2)", key, value)
	if err != nil {
		fmt.Printf("Error setting key %s: %v", key, err)
		return err
	}
	return nil
}

func (r *Repository) Get(key string) (string, bool) {
	var value string

	// Call the get_key_value stored function and scan the result.
	err := r.store.QueryRow("SELECT value_text FROM get($1)", key).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			// Key was not found in the database.
			return "", false
		}
		// An actual error occurred during query execution.
		fmt.Printf("Error getting key %s: %v", key, err)
		return "", false
	}

	return value, true
}
