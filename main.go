package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	product := "Облачное хранилище"
	price := 300

	// название продукта и цена передаются через параметры
	res, err := db.Exec("INSERT INTO products (product, price) VALUES (:product, :price)",
		sql.Named("product", product),
		sql.Named("price", price))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Вызов функции LastInsertId() у объекта типа sql.Result получаем последний добавленный идентификатор.
	fmt.Println(res.LastInsertId())
	// Функция RowsAffected() говорит, сколько строк было затронуто.
	// В данном случае добавилась всего одна строка, значит функция вернёт 1.
	fmt.Println(res.RowsAffected())

}
