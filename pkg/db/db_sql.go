// логика подключения к БД
package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Conect(connStr string) {
	//connStr := "user=your_username dbname=your_database password=your_password host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Закрытие соединения при завершении функции

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение к базе данных!")
}
