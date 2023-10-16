package internal

import (
	"bytes"
	"encoding/json"
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

	err = json.NewDecoder(req.Body).Decode(&responseObj)
	if err != nil {
		return responseObj, err
	}

	return responseObj, nil
}
