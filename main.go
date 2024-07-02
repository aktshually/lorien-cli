package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("./lorien.config.json"); os.IsNotExist(err) {
		os.Create("./lorien.config.json")
	}

	dataInBytes, _ := os.ReadFile("~/lorien.config.json")
	configFileData := ConfigFile{}
	json.Unmarshal(dataInBytes, &configFileData)

	if configFileData.Dir == "" {
		appDirectory := ""
		fmt.Println("What's Lorien's file directory? ")
		fmt.Scanln(&appDirectory)

		fmt.Println(appDirectory)

		_, err := os.Stat(appDirectory)
		fmt.Println(os.IsNotExist(err))
	}
}
