package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/glair-ai/glair-vision-go/config"
)

// MakeRequest creates and sends HTTP request to a specified
// GLAIR Vision service endpoint
func MakeRequest[T any](
	config *config.Config,
	endpoint string,
	file *os.File,
) (T, error) {
	var responseObj T

	url := config.GetEndpointURL("ocr", endpoint)
	header, body := createRequestPayload(file)

	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return responseObj, err
	}

	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("x-api-key", config.ApiKey)

	for key, value := range header {
		req.Header.Set(key, value)
	}

	res, err := config.Client.Do(req)
	if err != nil {
		return responseObj, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errorBody responseBody

		_ = json.NewDecoder(res.Body).Decode(&errorBody)

		return responseObj, RequestError{
			StatusCode: res.StatusCode,
			Body:       errorBody,
		}
	}

	err = json.NewDecoder(res.Body).Decode(&responseObj)
	if err != nil {
		return responseObj, err
	}

	return responseObj, nil
}

func createRequestPayload(file *os.File) (map[string]string, *bytes.Buffer) {
	header := map[string]string{}
	body := &bytes.Buffer{}

	var bytes []byte
	file.Read(bytes)

	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.Close()

	header["Content-Type"] = writer.FormDataContentType()

	return header, body
}
