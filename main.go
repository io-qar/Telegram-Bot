package main

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"strings"
)

const token = ""

var (
	bot    *tbot.Server
	client *tbot.Client
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bot = tbot.New(token)
	client = bot.Client()

	bot.HandleMessage("/start", startHandler)
	bot.HandleMessage("/weather .+", weatherHandler)
	bot.HandleMessage("/results", ResultHandler)
	bot.HandleMessage("", locHandler)

	err := bot.Start()
	log.Fatal(err)
}

func ResultHandler(m *tbot.Message) {
	res, err := getResultsFromDB(m)
	CheckError(err)
	for _, v := range res {
		client.SendMessage(m.Chat.ID, v.request)
	}
}

func startHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Hello!\nHere are commands:\n/start\n/weather City_Name\n/results\n results returns you the last 20 requests")
	sendUserInfoToBD(m)
}

func weatherHandler(m *tbot.Message) {
	city := strings.TrimPrefix(m.Text, "/weather ")
	city = strings.TrimSpace(city)
	city = sendApiToTranslate(city, "en")
	sendAPI(m, city)
}

func locHandler(m *tbot.Message) {
	if m.Location != nil {
		sendAPI(m, fmt.Sprint(m.Location.Latitude)+","+fmt.Sprint(m.Location.Longitude))
	} else {
		client.SendMessage(m.Chat.ID, "Похоже, вы не прикрепили местоположение")
	}
}

func unmatchedHandler(m *tbot.Message) {
	client.SendMessage(m.Chat.ID, "Sorry, you entered an invalid command.\nPlease use the bot keyboard.")
}
