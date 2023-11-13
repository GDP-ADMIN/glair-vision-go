package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
	ctx := context.Background()

	config := glair.NewConfig("", "", "")
	client := client.New(config)

	file, _ := os.Open("../images/receipt.jpg")

	result, err := client.Ocr.Receipt(ctx, glair.OCRInput{
		File: file,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
