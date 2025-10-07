// логика подключения к БД
package db

import (
	"database/sql"
	"fmt"

	"github.com/LushnikovSR/url_shortener/internal/config"

	"github.com/lib/pq"
)

// название функции Connect заменено на MustInitDB, которе указывает на вызов panic(),
// нет смысла разворачивать дальше сервис если нет соединения с базой данных
func MustInitDB(config *config.Config) *sql.DB {

	connector, err := pq.NewConnector(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	))
	if err != nil {
		panic(err)
	}

	// Создание DB с помощью коннектора
	db := sql.OpenDB(connector)

	// Настройка пула подключений
	db.SetMaxOpenConns(config.MaxConns)
	db.SetMaxIdleConns(config.IdleConns)

	//defer db.Close() // Закрытие соединения при завершении функции в main.go

	// Проверка подключения
	// add retry 3 times with sleep 100 ms
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Успешное подключение к базе данных!")

	return db
}
