package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yanzay/tbot"
)

type requestStruct struct {
	idUser  int
	request string
}

func sendUserInfoToBD(m *tbot.Message) {
	name := m.From.UserName
	id := m.From.ID
	connStr := "user=postgres dbname=tg_bot password=1111 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	insert := fmt.Sprintf("INSERT INTO users (id,username) SELECT %d, '%s' WHERE NOT EXISTS (SELECT id FROM users WHERE id = %d)", id, name, id)
	fmt.Println(insert)
	_, err = db.Exec(insert)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}

func sendRequestToDB(m *tbot.Message, req string) {

	id := m.From.ID
	connStr := "user=postgres dbname=tg_bot password=1111 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	insert := fmt.Sprintf("INSERT INTO requests (id_user,request) VALUES (%d,'%s')", id, req)
	fmt.Println(insert)

	_, err = db.Exec(insert)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}

func getResultsFromDB(m *tbot.Message) ([]requestStruct, error) {
	id := m.From.ID
	connStr := "user=postgres dbname=tg_bot password=1111 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully connected to database!\n")

	rows, err := db.Query(fmt.Sprintf("SELECT id_user, request FROM requests WHERE id_user = %d ORDER BY id_request DESC LIMIT 20", id))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var resSlize []requestStruct
	for rows.Next() {
		r := requestStruct{}
		err := rows.Scan(&r.idUser, &r.request)
		CheckError(err)
		resSlize = append(resSlize, r)
	}

	return resSlize, nil
}

func getWords(str string) []string {
	result := []string{}
	word := ""
	for _, v := range str {
		if v != ' ' {
			word += string(v)
		} else {
			if len(word) != 0 {
				result = append(result, word)
				word = ""
			}
		}
	}
	if len(word) != 0 {
		result = append(result, word)
	}
	return result
}
