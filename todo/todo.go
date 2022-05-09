package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Item struct {
	Text string
}

func ReadItems(filename string) ([]Item, error) {
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		return []Item{}, nil
	}

	var items []Item

	if err := json.Unmarshal(input, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}

func SaveItems(filename string, items []Item) error {
	output, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, output, 0644)

	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
