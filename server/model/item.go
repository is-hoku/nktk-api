package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/is-hoku/nktk-api/server/util"
)

type Item struct {
	Id            int    `json:"id"`
	Text          string `json:"text"`
	Tag           string `json:"tag"`
	Understanding bool   `json:"understanding"`
}

type ItemModel struct{}

func (r ItemModel) GetRandom() (Item, error) {
	jsonFile, err := ioutil.ReadFile("../items/index.json")
	if err != nil {
		return Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(jsonFile, &items); err != nil {
		return Item{}, err
	}
	index := util.GenarateRandomIntForJSON(len(items))
	item := items[index]
	return item, nil
}
