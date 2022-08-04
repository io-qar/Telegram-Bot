package main

import "github.com/tidwall/gjson"

func encode(s string) map[string]string {
	m := gjson.GetMany(s, "location.name", "location.country")
	// m, _ := gjson.Get(s, "location.name").Value().(string)

	var weather = map[string]string {
		"Город": m[0].String(),
		"Страна": m[1].String(),
	}
	return weather
}