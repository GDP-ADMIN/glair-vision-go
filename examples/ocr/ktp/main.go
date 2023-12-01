package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
	baseContext := context.Background()
	contextWithTimeout, cancel := context.WithTimeout(baseContext, 100*time.Millisecond)
	defer cancel()

	config := glair.NewConfig("", "", "")
	client := client.New(config)

	file, _ := os.Open("../images/ktp.jpeg")

	result, err := client.Ocr.Ktp(contextWithTimeout, glair.OCRInput{
		Image: file,
	})

	if err != nil {
		if glairErr, ok := err.(*glair.Error); ok {
			switch glairErr.Code {
			case glair.ErrorCodeTimeout:
				log.Printf("Request timed out")
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
