package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

// Link представляет запись в таблице links
type Link struct {
	ID           int
	OriginalLink string
	ShortLink    string
}

// Repository содержит подключение к БД
type Repository struct {
	store *sql.DB
}

// NewRepository создает новый экземпляр репозитория
func NewRepository(instance *sql.DB) *Repository {
	return &Repository{store: instance}
}

// Create добавляет новую запись о ссылке (Create)
func (r *Repository) Create(originalLink, shortLink string) (int, error) {
	var id int
	// Используем QueryRow для возврата id вставленной записи
	err := r.store.QueryRow(
		"INSERT INTO links (original_link, short_link) VALUES ($1, $2) RETURNING id",
		originalLink, shortLink,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating link: %v", err)
	}

	return id, nil
}

// GetOriginalLink находит оригинальную ссылку по короткой (Read)
func (r *Repository) GetOriginalLink(shortLink string) (string, error) {
	var originalLink string
	err := r.store.QueryRow(
		"SELECT original_link FROM links WHERE short_link = $1",
		shortLink,
	).Scan(&originalLink)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("link not found")
		}
		return "", fmt.Errorf("error reading link: %v", err)
	}

	return originalLink, nil
}

// GetAllLinks возвращает все ссылки (для демонстрации)
func (r *Repository) GetAllLinks() ([]Link, error) {
	rows, err := r.store.Query("SELECT id, original_link, short_link FROM links")
	if err != nil {
		return nil, fmt.Errorf("error querying links: %v", err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.ID, &link.OriginalLink, &link.ShortLink); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return links, nil
}

// Update обновляет короткую ссылку
func (r *Repository) Update(id int, newShortLink string) error {
	result, err := r.store.Exec(
		"UPDATE links SET short_link = $1 WHERE id = $2",
		newShortLink, id,
	)
	if err != nil {
		return fmt.Errorf("error updating link: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no link found with id %d", id)
	}

	return nil
}

// Delete удаляет запись по ID (Delete)
func (r *Repository) Delete(id int) error {
	result, err := r.store.Exec("DELETE FROM links WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting link: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no link found with id %d", id)
	}

	return nil
}

// Close закрывает подключение к БД
func (r *Repository) Close() error {
	return r.store.Close()
}
