package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/glair-ai/glair-vision-go"
)

// MakeRequest creates and sends HTTP request to a specified
// GLAIR Vision service endpoint
func MakeRequest[T any](
	ctx context.Context,
	url string,
	config *glair.Config,
	file *os.File,
) (T, error) {
	var response T

	header, body, err := createRequestPayload(file)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return response, &glair.Error{
			Code:    glair.ErrorCodeInvalidURL,
			Message: "invalid base url provided in configuration",
			Err:     err,
		}
	}

	req = req.WithContext(ctx)

	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("x-api-key", config.ApiKey)
	for key, value := range header {
		req.Header.Set(key, value)
	}

	res, err := config.Client.Do(req)
	if err != nil {
		return response, &glair.Error{
			Code:    glair.ErrorCodeBadClient,
			Message: "bad http client provided in configuration",
			Err:     err,
		}
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		var body map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&body)

		glairErr := &glair.Error{
			Code:    glair.ErrorCodeInvalidResponse,
			Message: "failed to parse http response",
			Err:     err,
		}

		if err != nil {
			return response, glairErr
		}

		glairErr.Code = body["status"].(glair.ErrorCode)
		glairErr.Message = body["reason"].(string)
		glairErr.Body = glair.ResponseBody{
			Status: body["status"].(string),
			Reason: body["reason"].(string),
		}

		return response, glairErr
	}

	return response, nil
}

func createRequestPayload(file *os.File) (map[string]string, *bytes.Buffer, error) {
	header := map[string]string{}
	body := &bytes.Buffer{}

	var bytes []byte
	file.Read(bytes)

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	if err != nil {
		return header, nil, &glair.Error{
			Code:    glair.ErrorCodeFileCorrupted,
			Message: "failed to append image to request",
			Err:     err,
		}
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return header, nil, &glair.Error{
			Code:    glair.ErrorCodeFileCorrupted,
			Message: "failed to parse image data",
			Err:     err,
		}
	}

	writer.Close()

	header["Content-Type"] = writer.FormDataContentType()

	return header, body, nil
}
