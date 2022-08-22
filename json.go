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

	res_str := make([]string, 2)
	res_str[0] = w.City
	res_str[1] = w.Country

	if w.City == "" {
		return "Введите правильное название города"
	}
	new_str := sendApiToTranslate(res_str[0]+" "+res_str[1], "ru")
	res_str_new := getWords(new_str)
	w.City = res_str_new[0]
	w.Country = res_str_new[1]

	result := "Погода в " + w.City + ", " + w.Country + " на " + w.Time + ":\n\nТемпература: " + w.Tempreture_c + "℃\nУсловия: " + w.Conditions + "\nСкорость ветра: " + w.Wind + " км/ч\nДавление: " + w.Pressure + " дюйм р.ст.\nВлажность: " + w.Humidity + "%\nТемпература ощущается как: " + w.Feels + "℃\nИндекс ультрафиолета: " + w.UV + "\n"

	sendRequestToDB(me, result)
	return result
}
