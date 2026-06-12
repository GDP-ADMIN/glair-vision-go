package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/examples/config"
	"github.com/glair-ai/glair-vision-go/ocr"
)

func main() {
	ctx := context.Background()

	client := config.NewClient()

	file, _ := os.Open("../images/skpr.jpg")

	var result ocr.SKPR

	err := client.Ocr.SKPR(
		ctx,
		glair.OCRInput{
			Image: file,
		},
		&result,
	)

	if err != nil {
		if glairErr, ok := err.(*glair.Error); ok {
			switch glairErr.Code {
			case glair.ErrorCodeAPIError:
				log.Printf("API Error: %v\n", glairErr.Response)
			default:
				log.Printf("Error: %v\n", glairErr.Code)
			}
		} else {
			log.Printf("Unexpected Error: %v\n", err)
		}

		os.Exit(1)
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
