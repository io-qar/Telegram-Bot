package main

import (
	"github.com/tidwall/gjson"
	"github.com/yanzay/tbot/v2"
)

func encode(s string, me *tbot.Message) string {
	m := gjson.GetMany(s, "location.name", "location.country", "location.localtime", "current.temp_c", "current.condition.text", "current.wind_kph", "current.pressure_mb", "current.humidity", "current.feelslike_c", "current.uv")
	var w = Weather{
		City:         m[0].String(),
		Country:      m[1].String(),
		Time:         m[2].String(),
		Tempreture_c: m[3].String(),
		Conditions:   m[4].String(),
		Wind:         m[5].String(),
		Pressure:     m[6].String(),
		Humidity:     m[7].String(),
		Feels:        m[8].String(),
		UV:           m[9].String(),
	}
	req_word := w.City + " " + w.Country + " " + w.Conditions
	req_word = sendApiToTranslate(req_word, "ru")
	ret_words := getWords(req_word)
	w.City = ret_words[0]
	w.Country = ret_words[1]
	w.Conditions = ret_words[2]
	if w.City == "" {

		return "Введите правильное название города"
	}
	sendRequestToDB(me, w.City)
	return "Погода в " + w.City + ", " + w.Country + " на " + w.Time + ":\n\nТемпература: " + w.Tempreture_c + "℃\nУсловия: " + w.Conditions + "\nСкорость ветра: " + w.Wind + " км/ч\nДавление: " + w.Pressure + " дюйм р.ст.\nВлажность: " + w.Humidity + "%\nТемпература ощущается как: " + w.Feels + "℃\nИндекс ультрафиолета: " + w.UV + "\n"

}
