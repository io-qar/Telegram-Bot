package main

import "github.com/tidwall/gjson"

func encode(s string) string {
	m := gjson.GetMany(s, "location.name", "location.country", "current.temp_c", "current.condition.text", "current.wind_kph", "current.pressure_mb", "current.humidity", "current.feelslike_c", "current.uv")

	var w = Weather {
		City: m[0].String(),
		Country: m[1].String(),
		Tempreture_c: m[2].String(),
		Conditions: m[3].String(),
		Wind: m[4].String(),
		Pressure: m[5].String(),
		Humidity: m[6].String(),
		Feels: m[7].String(),
		UV: m[8].String(),
	}

	return "Погода в " + w.City + ", " + w.Country + ":\n" + "Температура: " + w.Tempreture_c + "\nУсловия: " + w.Conditions + "\nСкорость ветра: " + w.Wind + "\nДавление: " + w.Pressure + "\nВлажность: " + w.Humidity + "\nТемпература ощущается как: " + w.Feels + "\nУровень ультрафиолета: " + w.UV + "\n"
}