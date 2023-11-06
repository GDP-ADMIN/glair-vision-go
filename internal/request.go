package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"

	"github.com/glair-ai/glair-vision-go"
)

// MakeRequest creates and sends HTTP request to a specified
// GLAIR Vision service endpoint.
//
// This function is not meant to be used outside GLAIR Vision SDK
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
			Message: "Invalid base URL is provided in configuration.",
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
			Message: "Bad HTTP client is provided in configuration.",
			Err:     err,
		}
	}
	defer res.Body.Close()

	str, _ := httputil.DumpResponse(res, true)
	fmt.Println(string(str))

	if res.StatusCode != http.StatusOK {
		var resBody map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&resBody)

		if err != nil {
			// TODO: telemetry?
			return response, &glair.Error{
				Code:    glair.ErrorCodeInvalidResponse,
				Message: "Failed to parse API response. Please contact us about this error.",
				Err:     err,
				Response: glair.Response{
					Code: res.StatusCode,
				},
			}
		}

		glairErr := &glair.Error{
			Code:    glair.ErrorCodeAPIError,
			Message: "GLAIR API returned non-OK response. Please check the Response property for more detailed explanation.",
			Response: glair.Response{
				Code: res.StatusCode,
			},
		}

		reason, ok := resBody["reason"].(string)
		if ok {
			glairErr.Response.Reason = reason
		}

		message, ok := resBody["message"].(string)
		if glairErr.Response.Reason == "" && ok {
			glairErr.Response.Reason = message
		}

		status, ok := resBody["status"].(string)
		if ok {
			glairErr.Response.Status = status
		}

		return response, glairErr
	}

	// we don't need to check the error here
	json.NewDecoder(res.Body).Decode(&response)

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
			Message: "Failed to append file into request body.",
			Err:     err,
		}
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return header, nil, &glair.Error{
			Code:    glair.ErrorCodeFileCorrupted,
			Message: "Failed to parse image data.",
			Err:     err,
		}
	}

	writer.Close()

	header["Content-Type"] = writer.FormDataContentType()

	return header, body, nil
}
