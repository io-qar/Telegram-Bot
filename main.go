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
		m.Replyf("", m.From.UserName, v.idUser, v.request)
	}
}

func startHandler(m *tbot.Message) {
	m.Reply("Hello!\nHere are commands:\n/start\n/weather City_Name\n/results\n results returns you the last 20 requests")
	sendUserInfoToBD(m)
}

func weatherHandler(m *tbot.Message) {
	sendAPI(m, m.Vars["city"])
	sendRequestToDB(m)
}

func unmatchedHandler(m *tbot.Message) {
	m.Reply("Извините, вы ввели недопустимую команду.\nПожалуйста, используйте клавиатуру бота.")
}
