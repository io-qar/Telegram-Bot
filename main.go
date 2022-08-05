package main

import (
	"github.com/yanzay/tbot"
	"log"
)

const token = "5442667303:AAGZej_QAla_ii5f8X66-hoCC3weuNBYOog"

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
	bot.HandleFunc("/results", ResultHandler)
	//unmatched input
	bot.HandleDefault(unmatchedHandler)

	err = bot.ListenAndServe()
	log.Fatal(err)
}

func ResultHandler(m *tbot.Message) {
	res, err := getResultsFromDB(m)
	CheckError(err)
	for _, v := range res {
		m.Reply(v.request)
	}
}

func startHandler(m *tbot.Message) {
	m.Reply("Hello!")
	sendUserInfoToBD(m)
}

func weatherHandler(m *tbot.Message) {
	sendAPI(m, m.Vars["city"])
	sendRequestToDB(m)
}

func unmatchedHandler(m *tbot.Message) {
	m.Reply("Извините, вы ввели недопустимую команду.\nПожалуйста, используйте клавиатуру бота.")
}