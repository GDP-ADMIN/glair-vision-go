package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
	config := glair.NewConfig("USERNAME", "PASSWORD", "API_KEY")
	client := client.New(config)

	file, _ := os.Open("./images/ktp.jpeg")

	result, err := client.Ocr.Ktp(file)

	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
