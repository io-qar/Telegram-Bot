package main

import (
	"log"
	// "os"

	// "github.com/joho/godotenv"
	// "github.com/yanzay/tbot/v2"
	"github.com/yanzay/tbot"
)

const token = "5471768780:AAEAbreeE6DDECknHmMrlD2Mfvedb5GIQ-w"
var bot *tbot.Server

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bot, err := tbot.NewServer(token)
	CheckError(err)

	bot.HandleFunc("/start", startHandler)
	bot.HandleFunc("/weather {city}", weatherHandler)
	//unmatched input
	bot.HandleDefault(unmatchedHandler)

	err = bot.ListenAndServe()
	log.Fatal(err)
}

func startHandler(m *tbot.Message) {
	m.Reply("Hello!")
	// buttons := [][]string{
	// 	{"Show weather in London", "Test", "Buttons"},
	// 	{"Another", "Row"},
	// }
	// m.ReplyKeyboard("Choose funcs below", buttons)
}

func weatherHandler(m *tbot.Message) {
	sendAPI(m, m.Vars["city"])
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