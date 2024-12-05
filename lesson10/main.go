package main

import (
	"encoding/json"
	"fmt"
)

type Resp struct {
	Posts              []Post
	Total, Skip, Limit int
}

type Post struct {
	Id        int
	Title     string
	Body      string
	Tags      []string
	Reactions map[string]int
	Views     int
	UserId    int
}

var data = `{
	"posts": [
	  {
		"id": 1,
		"title": "His mother had always taught him",
		"body": "His mother had always taught him not to ever think of himself as better than others. He'd tried to live by this motto. He never looked down on those who were less fortunate or who had less money than him. But the stupidity of the group of people he was talking to made him change his mind.",
		"tags": [
		  "history",
		  "american",
		  "crime"
		],
		"reactions": {
		  "likes": 192,
		  "dislikes": 25
		},
		"views": 305,
		"userId": 121
	  },
	  {
		"id": 2,
		"title": "He was an expert but not in a discipline",
		"body": "He was an expert but not in a discipline that anyone could fully appreciate. He knew how to hold the cone just right so that the soft server ice-cream fell into it at the precise angle to form a perfect cone each and every time. It had taken years to perfect and he could now do it without even putting any thought behind it.",
		"tags": [
		  "french",
		  "fiction",
		  "english"
		],
		"reactions": {
		  "likes": 859,
		  "dislikes": 32
		},
		"views": 4884,
		"userId": 91
	  }
	],
	"total": 251,
	"skip": 0,
	"limit": 30
  }`

type Computer struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	Year        int    `json:"year"`
	HasDocument bool   `json:"has_document"`
}

func main() {
	// c := Computer{"Macbook", "Apple", 2024}
	var r Resp

	err := json.Unmarshal([]byte(data), &r)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
