package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/glair-ai/glair-vision-go/client"
	"github.com/glair-ai/glair-vision-go/config"
	"github.com/joho/godotenv"
)

func main() {
	// test code here
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load credentials from env file: %s", err)
	}

	config := config.New(os.Getenv("API_USERNAME"), os.Getenv("API_PASSWORD"), os.Getenv("API_KEY"))
	client := client.New(config)

	file, _ := os.Open("fixtures/sample/ktp.jpeg")

	result, err := client.Ocr.KtpWithQuality(file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
