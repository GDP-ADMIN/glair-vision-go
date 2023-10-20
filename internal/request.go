package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/glair-ai/glair-vision-go"
)

// MakeRequest creates and sends HTTP request to a specified
// GLAIR Vision service endpoint
func MakeRequest[T any](
	config *glair.Config,
	url *url.URL,
	file *os.File,
) (T, error) {
	var obj T

	header, body, err := createRequestPayload(file)
	if err != nil {
		return obj, err
	}

	req, _ := http.NewRequest("POST", url.String(), body)
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("x-api-key", config.ApiKey)
	for key, value := range header {
		req.Header.Set(key, value)
	}

	res, err := config.Client.Do(req)
	if err != nil {
		return obj, glair.ErrBadClient
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return obj, glair.ErrInvalidResponseBody
		}

		return obj, glair.RequestError{
			StatusCode:   res.StatusCode,
			ResponseBody: bytes,
		}
	}

	err = json.NewDecoder(res.Body).Decode(&obj)
	if err != nil {
		return obj, glair.ErrInvalidResponseBody
	}

	return obj, nil
}

func createRequestPayload(file *os.File) (map[string]string, *bytes.Buffer, error) {
	header := map[string]string{}
	body := &bytes.Buffer{}

	var bytes []byte
	file.Read(bytes)

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	if err != nil {
		return header, nil, glair.ErrInvalidFormData
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return header, nil, glair.ErrInvalidFile
	}

	writer.Close()

	header["Content-Type"] = writer.FormDataContentType()

	return header, body, nil
}
