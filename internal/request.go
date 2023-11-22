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
	"time"

	"github.com/glair-ai/glair-vision-go"
)

type RequestParameters struct {
	Url       string
	RequestID string
	Payload   map[string]*os.File
}

// MakeRequest creates and sends HTTP request to a specified
// GLAIR Vision service endpoint.
//
// This function is not meant to be used outside GLAIR Vision SDK
func MakeRequest[T any](
	ctx context.Context,
	payload RequestParameters,
	config *glair.Config,
) (T, error) {
	var response T

	res, status, err := sendRequest(ctx, payload, config)
	if err != nil {
		return response, err
	}

	config.Logger.Debugf("API Response: %s", string(res))

	if status != http.StatusOK {
		var resBody map[string]interface{}
		err = json.Unmarshal(res, &resBody)

		if err != nil {
			config.Logger.Errorf("Failed to parse API response due to %v", err)
			return response, &glair.Error{
				Code:    glair.ErrorCodeInvalidResponse,
				Message: "Failed to parse API response. Please contact us about this error.",
				Err:     err,
				Response: glair.Response{
					Code: status,
				},
			}
		}

		glairErr := &glair.Error{
			Code:    glair.ErrorCodeAPIError,
			Message: "GLAIR API returned non-OK response. Please check the Response property for more detailed explanation.",
			Response: glair.Response{
				Code: status,
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
	json.Unmarshal(res, &response)

	return response, nil
}

func sendRequest(
	ctx context.Context,
	payload RequestParameters,
	config *glair.Config,
) ([]byte, int, error) {
	header, body, err := createRequestPayload(payload.Payload, config.Logger)
	if err != nil {
		return []byte{}, 0, err
	}

	req, err := http.NewRequest("POST", payload.Url, body)
	if err != nil {
		return []byte{}, 0, &glair.Error{
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

	if payload.RequestID != "" {
		req.Header.Set("x-request-id", payload.RequestID)
	}
	req.Header.Set("User-Agent", "go/GLAIR-Vision-SDK")

	config.Logger.Infof("Calling GLAIR Vision API at %s", payload.Url)

	start := time.Now()

	res, err := config.Client.Do(req)
	if err != nil {
		return []byte{}, 0, &glair.Error{
			Code:    glair.ErrorCodeBadClient,
			Message: "Bad HTTP client is provided in configuration.",
			Err:     err,
		}
	}
	defer res.Body.Close()

	elapsed := time.Since(start)

	config.Logger.Infof("Request handled in %.2f second(s)", elapsed.Seconds())

	str, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, 0, &glair.Error{
			Code:    glair.ErrorCodeInvalidResponse,
			Message: "Failed to parse API response. Please contact us about this error.",
			Err:     err,
			Response: glair.Response{
				Code: res.StatusCode,
			},
		}
	}

	return str, res.StatusCode, nil
}

func createRequestPayload(
	payload map[string]*os.File,
	logger glair.Logger,
) (map[string]string, *bytes.Buffer, error) {
	header := map[string]string{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for field, file := range payload {
		var bytes []byte
		file.Read(bytes)

		part, err := writer.CreateFormFile(field, filepath.Base(file.Name()))
		if err != nil {
			logger.Errorf("Failed when appending file to request body: %v", err.Error())
			return header, nil, &glair.Error{
				Code:    glair.ErrorCodeFileCorrupted,
				Message: "Failed to append file into request body.",
				Err:     err,
			}
		}

		_, err = io.Copy(part, file)
		if err != nil {
			logger.Errorf("Failed to copy image data to request body: %v", err.Error())
			return header, nil, &glair.Error{
				Code:    glair.ErrorCodeFileCorrupted,
				Message: "Failed to parse image data.",
				Err:     err,
			}
		}
	}

	writer.Close()

	header["Content-Type"] = writer.FormDataContentType()

	return header, body, nil
}
