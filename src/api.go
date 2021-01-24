package main

import (
	"db"
	"fmt"
)

// UserInfo type
type UserInfo struct {
	App    string `json:"app"`
	Device string `json:"device"`
	Bot    bool   `json:"bot"`
}

func isBot(element db.DbElement) bool {
	_, found := element["bot"]
	return found
}

func main() {
	var dataAccessor db.DataAccessor

	err := dataAccessor.LoadJSON("user-agents.json")
	if err != nil {
		fmt.Println(err)
	}
	result := dataAccessor.FindAll(isBot)
	
	for _, e := range result {
		fmt.Println(e)
	}
}