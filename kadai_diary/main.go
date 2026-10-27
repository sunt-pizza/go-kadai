package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const JsonFilePath = "./diary.json"

type Entry struct {
	Title string
	Content string
	Timestamp time.Time
}

func saveDiary(data []Entry) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return  err
	}
	return os.WriteFile(JsonFilePath, jsonData, 0644)
}

func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile(JsonFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, err	
	}
	
	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err	
	}
	return entries, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("コマンドを指定してください: add / list")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
			if len(os.Args) < 4 {
				fmt.Println("使い方: add タイトル 本文")
				return
			}
			entries, err := loadDiary()
			if err != nil {
				fmt.Println("読み込みに失敗しました", err)
				return
			}
			entries = append(entries, Entry{
				Title: os.Args[2],
				Content: os.Args[3],
				Timestamp: time.Now(),
			})
			if err := saveDiary(entries); err != nil {
				fmt.Println("保存に失敗しました:", err)
			}
	case "list":
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("読み込みに失敗しました", err)
			return
		}
		for _, e := range entries {
			fmt.Printf("--- 日記データ ---\n日時:%s\nタイトル:%s\n本文:\n%s",e.Timestamp.Format("2006-01-02 15:04:05"), e.Title, e.Content)
		}
	default:
			fmt.Println("不明なコマンドです:", os.Args[1])
	}
}