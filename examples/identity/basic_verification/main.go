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

	result, err := client.Identity.BasicVerification(ctx, glair.BasicVerificationInput{
		Nik:    "",
		Name:   glair.String(""),
		Gender: glair.String(""),
	})

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
