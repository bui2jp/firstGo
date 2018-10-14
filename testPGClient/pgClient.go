package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func selectUser() {
	fmt.Println("start selectUser")

	type User struct {
		Id   int
		Name string
	}

	db, err := sql.Open("postgres",
		"user=takuya password=pw dbname=takuya sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(`SELECT * FROM "users"`)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	users := []User{}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)

}
