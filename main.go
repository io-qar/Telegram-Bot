package main

import (
	"fmt"
	"github.com/yanzay/tbot"
	_ "github.com/yanzay/tbot/v2"
	"log"
)

const token = "5442667303:AAGZej_QAla_ii5f8X66-hoCC3weuNBYOog"

var bot *tbot.Server

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// type application struct {
// 	client *tbot.Client
// }

// var (
// 	app application
// 	bot *tbot.Server
// 	token string
// )

func main() {
	bot, err := tbot.NewServer(token)
	CheckError(err)

	bot.HandleFunc("/start", startHandler)
	bot.HandleFunc("/weather {city}", weatherHandler)
	bot.HandleFunc("/results", ResultHandler)
	//unmatched input
	bot.HandleDefault(unmatchedHandler)

	err = bot.ListenAndServe()
	log.Fatal(err)
}

func ResultHandler(m *tbot.Message) {
	res, err := getResultsFromDB(m)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range res {
		m.Replyf("", v)
	}

}

func startHandler(m *tbot.Message) {
	m.Reply("Hello!")
	sendUserInfoToBD(m)
	// buttons := [][]string{
	// 	{"Show weather in London", "Test", "Buttons"},
	// 	{"Another", "Row"},
	// }
	// m.ReplyKeyboard("Choose funcs below", buttons)
}

func weatherHandler(m *tbot.Message) {
	sendAPI(m, m.Vars["city"])
	sendRequestToDB(m)

}

// func KeyboardHandler(m *tbot.Message) {
// 	buttons := [][]string{
// 		{"Show weather in London", "Test", "Buttons"},
// 		{"Another", "Row"},
// 	}
// 	m.ReplyKeyboard("", buttons)
// 	m.Reply("Sending API...")
// 	// println(sendAPI())
// }

func unmatchedHandler(m *tbot.Message) {
	m.Reply(`Sorry, you've just entered incorrect command.
Please, use buttons below.`)
}
