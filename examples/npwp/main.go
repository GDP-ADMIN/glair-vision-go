package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
	ctx := context.Background()

	config := glair.NewConfig("", "", "")
	client := client.New(config)

	result, err := client.Ocr.NPWP(ctx, "../images/npwp.jpg")

	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
