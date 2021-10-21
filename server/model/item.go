package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/is-hoku/nktk-api/server/util"
)

type Item struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	Class         string `json:"class"`
	Understanding bool   `json:"understanding"`
}

func (r *Item) GetRandom(class string) error {
	jsonFile, err := ioutil.ReadFile("./items/index.json")
	if err != nil {
		return err
	}
	var items []Item
	if err := json.Unmarshal(jsonFile, &items); err != nil {
		return err
	}

	if class == "" { // 条件なしランダム
		index := util.GenarateRandomIntForJSON(len(items))
		r.ID = items[index].ID
		r.Text = items[index].Text
		r.Class = items[index].Class
		r.Understanding = items[index].Understanding
		return nil
	} else { // class が指定されたランダム
		var searchedItems []Item
		for _, v := range items {
			if v.Class == class {
				searchedItems = append(searchedItems, v)
			}
		}
		index := util.GenarateRandomIntForJSON(len(searchedItems))
		r.ID = searchedItems[index].ID
		r.Text = searchedItems[index].Text
		r.Class = searchedItems[index].Class
		r.Understanding = searchedItems[index].Understanding
		return nil
	}
}

func (r *Item) GetByID(id int) error {
	jsonFile, err := ioutil.ReadFile("./items/index.json")
	if err != nil {
		return err
	}
	var items []Item
	if err := json.Unmarshal(jsonFile, &items); err != nil {
		return err
	}
	index := id - 1
	r.ID = items[index].ID
	r.Text = items[index].Text
	r.Class = items[index].Class
	r.Understanding = items[index].Understanding
	return nil
}
