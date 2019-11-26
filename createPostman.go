package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
)

var col Collection

func generatePostman(name string, items *[]Item) {
	info := createInfo(name)
	varHost := createVariable("1100e68c-e506-41f3-998b-d5c188e80c19", "host", "localhost:8080")
	varID := createVariable("1100e68c-e506-41f3-998b-d5c188e80c20", "id", "1")
	vars := []Variable{*varHost, *varID}
	jsonCollection, _ := json.Marshal(createCollection(info, items, &vars))
	// fmt.Println(string(jsonCollection))

	filename := folder + "/" + name + ".postman_collection.json"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	jsonCollection, _ = prettyprint(jsonCollection)
	f.Write(jsonCollection)
}
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func createCollection(info *Info, items *[]Item, variable *[]Variable) *Collection {
	return &Collection{
		Info:     *info,
		Item:     *items,
		Variable: *variable,
	}
}

func createInfo(name string) *Info {
	return &Info{
		PostID: "1998c3bf-f7ac-4893-892d-2d7a2928dd38",
		Name:   name,
		Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
	}
}
func createVariable(id, key, value string) *Variable {
	return &Variable{
		ID:    id,
		Key:   key,
		Value: value,
	}
}

func createItems(name string, path string, queries *[]Query) (items *[]Item) {
	itemReadAll := Item{
		Name: "read_all_" + name,
		Request: Request{
			Method: "GET",
			Header: []string{},
			Body:   Body{Mode: "raw", Raw: ""},
			URL: URL{
				Raw:   "",
				Host:  []string{"{{host}}"},
				Path:  strings.Split(path, "/"),
				Query: []Query{},
			},
		},
		Response: []string{},
	}

	itemReadOne := Item{
		Name: "read_one_" + name,
		Request: Request{
			Method: "GET",
			Header: []string{},
			Body:   Body{Mode: "raw", Raw: ""},
			URL: URL{
				Raw:   "{{host}}/" + path + "/{{id}}",
				Host:  []string{"{{host}}"},
				Path:  append(strings.Split(path, "/"), "{{id}}"),
				Query: []Query{},
			},
		},
		Response: []string{},
	}

	itemCreateOne := Item{
		Name: "create_one_" + name,
		Request: Request{
			Method: "POST",
			Header: []string{},
			Body:   Body{Mode: "raw", Raw: ""},
			URL: URL{
				Raw:   "{{host}}/" + path,
				Host:  []string{"{{host}}"},
				Path:  strings.Split(path, "/"),
				Query: *queries,
			},
		},
		Response: []string{},
	}

	itemUpdateOne := Item{
		Name: "update_one_" + name,
		Request: Request{
			Method: "PUT",
			Header: []string{},
			Body:   Body{Mode: "raw", Raw: ""},
			URL: URL{
				Raw:   "{{host}}/" + path + "/{{id}}",
				Host:  []string{"{{host}}"},
				Path:  append(strings.Split(path, "/"), "{{id}}"),
				Query: *queries,
			},
		},
		Response: []string{},
	}
	itemDeleteOne := Item{
		Name: "delete_one_" + name,
		Request: Request{
			Method: "DELETE",
			Header: []string{},
			Body:   Body{Mode: "raw", Raw: ""},
			URL: URL{
				Raw:   "{{host}}/" + path + "/{{id}}",
				Host:  []string{"{{host}}"},
				Path:  append(strings.Split(path, "/"), "{{id}}"),
				Query: []Query{},
			},
		},
		Response: []string{},
	}
	return &[]Item{itemReadAll, itemReadOne, itemCreateOne, itemUpdateOne, itemDeleteOne}
}
