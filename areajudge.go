package wordproperty

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Area struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	PID       uint   `json:"pid"`
	ShortName string `json:"short_name"`
	Level     int    `json:"level"`
	CityCode  string `json:"city_code"`
	YZCode    string `json:"yz_code"`
	MergeName string `json:"merge_name"`
	Lng       string `json:"lng"`
	Lat       string `json:"lat"`
	PinYin    string `json:"pin_yin"`
}

func IsChinaArea(key string) (bool, float64, []Area) {
	confidence := 1.0
	isArea := false
	var jAreas []Area
	for _, area := range areas {
		if strings.Contains(key, area.ShortName) {
			jAreas = append(jAreas, area)
			isArea = true
		}
	}
	return isArea, confidence, jAreas
}

var areas []Area

func init() {
	data, _ := ioutil.ReadFile("area.json")
	_ = json.Unmarshal(data, &areas)
}
