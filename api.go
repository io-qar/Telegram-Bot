package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func sendAPI(m *tbot.Message, loc string) {
	clnt := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://api.weatherapi.com/v1/current.json?key=899e2510ac1948588ec165012223107&q="+loc+"&aqi=no&lang=ru", nil)
	CheckError(err)

	resp, err := clnt.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	resp.Body.Close()

	client.SendMessage(m.Chat.ID, encode(string(body[:]), m))
}
func sendApiToTranslate(str string, lg string) string {
	//body_test := `[{"detectedLanguage":{"language":"ru","score":1.0},"translations":[{"text":"How to translate you into a string","to":"en"}]}]`

	url := "https://microsoft-translator-text.p.rapidapi.com/translate?to%5B0%5D=" + lg + "&api-version=3.0&profanityAction=NoAction&textType=plain"

	payload := strings.NewReader("[\r{\r	\"Text\": \"" + str + "\"\r}\r]")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "0458650784msh4824a26ccdfb8f3p177b8ejsn5d76519c4c1e")
	req.Header.Add("X-RapidAPI-Host", "microsoft-translator-text.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	new_body := string(body)
	new_body = strings.ReplaceAll(new_body, "[", "")
	new_body = strings.ReplaceAll(new_body, "]", "")

	result := gjson.Get(new_body, "translations.text")
	fmt.Println(result)
	return result.String()
}
