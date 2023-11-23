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

	image, _ := os.Open("../images/face.jpeg")

	result, err := client.FaceBio.ActiveLiveness(ctx, glair.ActiveLivenessInput{
		Image:       image,
		GestureCode: "HAND_00000",
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
