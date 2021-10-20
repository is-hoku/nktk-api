package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/is-hoku/nktk-api/server/util"
)

type Item struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	Tag           string `json:"tag"`
	Understanding bool   `json:"understanding"`
}

func (r *Item) GetRandom() error {
	jsonFile, err := ioutil.ReadFile("./items/index.json")
	if err != nil {
		return err
	}
	var items []Item
	if err := json.Unmarshal(jsonFile, &items); err != nil {
		return err
	}
	index := util.GenarateRandomIntForJSON(len(items))
	r.ID = items[index].ID
	r.Text = items[index].Text
	r.Tag = items[index].Tag
	r.Understanding = items[index].Understanding
	return nil
}
