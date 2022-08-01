package main

import (
	"net/http"
	"io/ioutil"
	"github.com/yanzay/tbot"
)

func sendAPI(m *tbot.Message) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://api.weatherapi.com/v1/current.json?key=899e2510ac1948588ec165012223107&q=London&aqi=no ", nil)
	CheckError(err)

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	m.Replyf("Weather in, %s, is...", body[:])
}