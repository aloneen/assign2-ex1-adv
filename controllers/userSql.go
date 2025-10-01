package controllers

import (
	"fmt"
	"log"

	"github.com/aloneen/assign2-ex1-adv/initializers"
	"github.com/aloneen/assign2-ex1-adv/models"
)

func CreateUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		age INT NOT NULL
	);`
	_, err := initializers.DB.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	fmt.Println("Users table created")
}

func InsertUsers(users []models.User) {
	tx, err := initializers.DB.Begin()
	if err != nil {
		log.Fatal("Cannot begin tx:", err)
	}

	stmt, err := tx.Prepare("INSERT INTO users(name, age) VALUES($1, $2)")
	if err != nil {
		tx.Rollback()
		log.Fatal("Prepare failed:", err)
	}
	defer stmt.Close()

	for _, u := range users {
		_, err = stmt.Exec(u.Name, u.Age)
		if err != nil {
			tx.Rollback()
			log.Fatal("Insert failed:", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("Commit failed:", err)
	}
	fmt.Println("Users inserted in transaction")
}

func QueryUsers(minAge, page, pageSize int) {
	offset := (page - 1) * pageSize

	rows, err := initializers.DB.Query(
		"SELECT id, name, age FROM users WHERE age >= $1 ORDER BY id LIMIT $2 OFFSET $3",
		minAge, pageSize, offset,
	)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

	fmt.Printf("\nUsers (page %d):\n", page)
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatal("Scan failed:", err)
		}
		fmt.Printf("ID=%d, Name=%s, Age=%d\n", u.ID, u.Name, u.Age)
	}
}

func UpdateUser(id int, name string, age int) {
	_, err := initializers.DB.Exec(
		"UPDATE users SET name=$1, age=$2 WHERE id=$3",
		name, age, id,
	)
	if err != nil {
		log.Fatal("Update failed:", err)
	}
	fmt.Printf("User %d updated\n", id)
}

func DeleteUser(id int) {
	_, err := initializers.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Fatal("Delete failed:", err)
	}
	fmt.Printf("User %d deleted\n", id)
}
