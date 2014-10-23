package things

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func MakeItem(wrappedItem map[string]interface{}) (thing Itemlike) {
	var itemType string
	var data map[string]interface{}

	if _, ok := wrappedItem["type"]; ok {
		itemType, _ = wrappedItem["type"].(string)
	} else {
		return nil
	}

	// TODO: ability to add a "file" field to load json from there
	// Important: ensure no circular loading is allowed, so track it
	// Maybe keep an array of files loaded from but instantiate it
	// somehow

	data, _ = wrappedItem["data"].(map[string]interface{})

	var preThing Populatable
	switch itemType {
	// TODO: add player here
	case "area":
		preThing = new(Area)
	case "item":
		preThing = new(Item)
	default:
		preThing = new(Item)
	}

	preThing.Populate(data)

	thing, _ = preThing.(Itemlike)
	return
}

func strEndsWith(str string, end string) bool {
	return strings.LastIndex(str, end) == len(str)-len(end)
}

func ParseJson(filename string) (map[string]interface{}, error) {
	if !strEndsWith(filename, ".json") {
		return nil, errorString{"Filename should be a .json"}
	}

	if bytes, err := ioutil.ReadFile(filename); err == nil {
		rawMap := map[string]interface{}{}

		err := json.Unmarshal(bytes, &rawMap)

		if err != nil {
			return nil, errorString{err.Error()}
		}

		return rawMap, nil

	} else {
		return nil, errorString{err.Error()}
	}
}

func MakeFromFile(filename string) (Itemlike, error) {
	var rawThing map[string]interface{}

	rawThing, err := ParseJson(filename)

	if err != nil {
		return nil, err
	}

	return MakeItem(rawThing), nil
}
