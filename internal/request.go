package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/glair-ai/glair-vision-go/config"
)

// MakeRequest creates and sends HTTP request to a specified
// GLAIR Vision service endpoint
//
// Not intended to be used outside the SDK
func MakeRequest[T any](
	config config.Config,
	endpoint string,
	file *os.File,
) (T, error) {
	var responseObj T

	url := config.GetEndpointURL("ocr", endpoint)
	fmt.Println(url.String())

	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	defer writer.Close()

	var fileContent []byte
	file.Read(fileContent)

	fileWriter, _ := writer.CreateFormFile("image", file.Name())
	fileWriter.Write(fileContent)

	req, err := http.NewRequest("POST", url.String(), buffer)
	if err != nil {
		return responseObj, err
	}

	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("x-api-key", config.ApiKey)

	res, err := config.Client.Do(req)
	if err != nil {
		return responseObj, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errorBody responseBody
		_ = json.NewDecoder(req.Body).Decode(&errorBody)

		return responseObj, RequestError{
			StatusCode: res.StatusCode,
			Body:       errorBody,
		}
	}

	err = json.NewDecoder(req.Body).Decode(&responseObj)
	if err != nil {
		return responseObj, err
	}

	return responseObj, nil
}
