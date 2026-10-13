package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entry struct {
	Timestamp time.Time
	Title string
	Content string
}

func saveDiary(data Entry) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("マーシャルに失敗しました", err)
		return
	}
	fmt.Println(string(jsonData))
}