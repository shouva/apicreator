package main

// Collection :
type Collection struct {
	Info     Info       `json:"info"`
	Item     []Item     `json:"item"`
	Variable []Variable `json:"variable"`
}

// Info :
type Info struct {
	PostID string `json:"_postman_id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

// Item :
type Item struct {
	Name     string   `json:"name"`
	Request  Request  `json:"request"`
	Response []string `json:"response"`
}

// {
// 	"id": "1100e68c-e506-41f3-998b-d5c188e80c19",
// 	"key": "andon19_token",
// 	"value": "",
// 	"type": "string"
// },

// Variable :
type Variable struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// Request :
type Request struct {
	Method string   `json:"method"`
	Header []string `json:"header"`
	Body   Body     `json:"body"`
	URL    URL      `json:"url"`
}

// Body :
type Body struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw"`
}

// URL :
type URL struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Port     string   `json:"port"`
	Path     []string `json:"path"`
	Query    []Query  `json:"query"`
}

// Query :
type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Response :
type Response struct {
	Name string `json:"name"`
}
