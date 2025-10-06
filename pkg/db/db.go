// логика подключения к БД
package db

import (
	"database/sql"
	"fmt"

	"github.com/LushnikovSR/url_shortener/internal/config"
)

// название функции Connect заменено на MustInitDB, которе указывает на вызов panic(),
// нет смысла разворачивать дальше сервис если нет соединения с базой данных
func MustInitDB(config *config.Config) *sql.DB {
	//connStr := "user=your_username dbname=your_database password=your_password host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", config.DBHost)
	if err != nil {
		panic(err)
	}
	//defer db.Close() // Закрытие соединения при завершении функции

	// Проверка подключения
	// add retry 3 times with sleep 100 ms
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Успешное подключение к базе данных!")

	return db
}
