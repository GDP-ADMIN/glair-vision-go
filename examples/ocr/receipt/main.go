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
		if glairErr, ok := err.(*glair.Error); ok {
			switch glairErr.Code {
			case glair.ErrorCodeAPIError:
				log.Printf("API Error: %v\n", glairErr.Response)
			default:
				log.Printf("Error: %v\n", glairErr.Code)
			}
		}

		log.Printf("Unexpected Error: %v\n", err)
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
