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

	file, _ := os.Open("../images/bpkb.pdf")

	result, err := client.Ocr.BPKB(ctx, glair.OCRInput{
		File: file,
	})

	if err != nil {
		log.Fatalln(err.(*glair.Error).Response)
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
