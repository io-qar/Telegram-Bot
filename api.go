package main

import (
	"io/ioutil"
	"net/http"

	"github.com/yanzay/tbot"
)

func sendAPI(m *tbot.Message, city string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://api.weatherapi.com/v1/current.json?key=899e2510ac1948588ec165012223107&q="+city+"&aqi=no ", nil)
	CheckError(err)

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	resp.Body.Close()

	m.Reply(encode(string(body[:]), m))
}
