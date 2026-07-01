package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/glair-ai/glair-vision-go"
)

const defaultMaxResponseBytes = 50 << 20 // 50 MB

type RequestParameters struct {
	Url       string
	RequestID string
	Body      map[string]any
}

type RequestPayload struct {
	RequestParameters
	Header map[string]string
	Body   io.Reader
}

// MakeMultipartRequest creates and sends multipart/formdata request and
// unmarshals the response into the provided type.
func MakeMultipartRequest[T any](
	ctx context.Context,
	params RequestParameters,
	config *glair.Config,
) (T, error) {
	var result T

	header, body, err := createMultipartPayload(params.Body, config.Logger)
	if err != nil {
		return result, err
	}

	res, status, err := sendRequest(ctx, RequestPayload{
		RequestParameters: params,
		Header:            header,
		Body:              body,
	}, config)
	if err != nil {
		return result, err
	}

	if err := responseError(res, status, config); err != nil {
		return result, err
	}

	if len(res) == 0 {
		return result, nil
	}

	if err := json.Unmarshal(res, &result); err != nil {
		return result, &glair.Error{
			Code:    glair.ErrorCodeInvalidResponse,
			Message: "Failed to parse API response.",
			Err:     err,
		}
	}

	return result, nil
}

func MakeMultipartRequestRaw(
	ctx context.Context,
	params RequestParameters,
	config *glair.Config,
) ([]byte, error) {
	header, body, err := createMultipartPayload(params.Body, config.Logger)
	if err != nil {
		return nil, err
	}

	res, status, err := sendRequest(ctx, RequestPayload{
		RequestParameters: params,
		Header:            header,
		Body:              body,
	}, config)
	if err != nil {
		return nil, err
	}

	if err := responseError(res, status, config); err != nil {
		return nil, err
	}

	return res, nil
}

// MakeJSONRequest creates and sends application/json request to a specified
// GLAIR Vision service endpoint.
func MakeJSONRequest[T any](
	ctx context.Context,
	params RequestParameters,
	config *glair.Config,
) (T, error) {
	var result T

	body, err := json.Marshal(params.Body)
	if err != nil {
		return result, &glair.Error{
			Code:    glair.ErrorCodeInvalidRequest,
			Message: "Failed to serialize request body.",
			Err:     err,
		}
	}

	reader := bytes.NewReader(body)

	res, status, err := sendRequest(ctx, RequestPayload{
		RequestParameters: params,
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: reader,
	}, config)
	if err != nil {
		return result, err
	}

	if err := responseError(res, status, config); err != nil {
		return result, err
	}

	if len(res) == 0 {
		return result, nil
	}

	if err := json.Unmarshal(res, &result); err != nil {
		return result, &glair.Error{
			Code:    glair.ErrorCodeInvalidResponse,
			Message: "Failed to parse API response.",
			Err:     err,
		}
	}

	return result, nil
}

func sendRequest(
	ctx context.Context,
	payload RequestPayload,
	config *glair.Config,
) ([]byte, int, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", payload.Url, payload.Body)

	if err != nil {
		return []byte{}, 0, &glair.Error{
			Code:    glair.ErrorCodeNetworkError,
			Message: fmt.Sprintf("Invalid base URL %s is provided in configuration.", config.BaseUrl),
			Err:     err,
		}
	}

	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("x-api-key", config.ApiKey)
	for key, value := range payload.Header {
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
		if strings.HasSuffix(err.Error(), "context deadline exceeded") {
			return []byte{}, 0, &glair.Error{
				Code:    glair.ErrorCodeTimeout,
				Message: fmt.Sprintf("Request to %s timed out", payload.Url),
			}
		}

		return []byte{}, 0, &glair.Error{
			Code:    glair.ErrorCodeNetworkError,
			Message: "Failed to send HTTP request due to network error.",
			Err:     err,
		}
	}
	defer res.Body.Close()

	elapsed := time.Since(start)

	config.Logger.Infof("Request handled in %.2f second(s)", elapsed.Seconds())

	maxResponseBytes := config.MaxResponseBytes
	if maxResponseBytes <= 0 {
		// No limit if 0 or negative
		str, err := io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, 0, &glair.Error{
				Code:    glair.ErrorCodeInvalidResponse,
				Message: "Failed to parse API response. Please contact us about this error.",
				Err:     err,
				Response: glair.Response{
					Status: res.StatusCode,
				},
			}
		}
		return str, res.StatusCode, nil
	}

	maxResponseBytesInt := int(maxResponseBytes)

	str, err := io.ReadAll(io.LimitReader(res.Body, int64(maxResponseBytesInt)+1))
	if err != nil {
		return []byte{}, 0, &glair.Error{
			Code:    glair.ErrorCodeInvalidResponse,
			Message: "Failed to parse API response. Please contact us about this error.",
			Err:     err,
			Response: glair.Response{
				Status: res.StatusCode,
			},
		}
	}

	if len(str) > maxResponseBytesInt {
		return []byte{}, 0, &glair.Error{
			Code:    glair.ErrorCodeInvalidResponse,
			Message: "API response exceeded maximum allowed size.",
			Response: glair.Response{
				Status: res.StatusCode,
			},
		}
	}

	config.Logger.Debugf("API Response: %d bytes (status: %d)", len(str), res.StatusCode)

	return str, res.StatusCode, nil
}

func createMultipartPayload(
	payload map[string]any,
	logger glair.Logger,
) (map[string]string, io.Reader, error) {
	header := map[string]string{}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var writtenBytes int64 = 0

	for field, value := range payload {
		switch val := value.(type) {
		case string:
			writer, _ := writer.CreateFormField(field)
			reader := strings.NewReader(val)

			bytes, _ := io.Copy(writer, reader)

			writtenBytes += bytes
		case *string:
			if val != nil {
				writer, _ := writer.CreateFormField(field)
				reader := strings.NewReader(*val)

				bytes, _ := io.Copy(writer, reader)

				writtenBytes += bytes
			}
		case *int:
			if val != nil {
				writer, _ := writer.CreateFormField(field)
				numStr := strconv.Itoa(*val)
				reader := strings.NewReader(numStr)

				bytes, _ := io.Copy(writer, reader)

				writtenBytes += bytes
			}
		case *os.File:
			writer, err := writer.CreateFormFile(field, filepath.Base(val.Name()))
			if err != nil {
				logger.Errorf("Failed to append file %s to request body: %v", field, err.Error())
				return header, nil, &glair.Error{
					Code:    glair.ErrorCodeFileError,
					Message: fmt.Sprintf("Failed to append file %s into request body.", field),
					Err:     err,
				}
			}

			bytes, err := io.Copy(writer, val)
			if err != nil {
				logger.Errorf("Failed to write field %s to request body: %v", field, err.Error())
				return header, nil, &glair.Error{
					Code:    glair.ErrorCodeFileError,
					Message: fmt.Sprintf("Failed to write field %s into request body.", field),
					Err:     err,
				}
			}

			writtenBytes += bytes
		}
	}

	writer.Close()

	header["Content-Type"] = writer.FormDataContentType()
	header["Content-Length"] = fmt.Sprintf("%d", writtenBytes)

	return header, body, nil
}

func responseError(res []byte, status int, config *glair.Config) error {
	if status == http.StatusForbidden {
		return &glair.Error{
			Code:    glair.ErrorCodeForbidden,
			Message: "Insufficient access to the API endpoint",
		}
	}

	if status >= 400 {
		var resBody map[string]any

		if err := json.Unmarshal(res, &resBody); err != nil {
			config.Logger.Errorf("Failed to parse API response due to %v", err)

			return &glair.Error{
				Code:    glair.ErrorCodeInvalidResponse,
				Message: "Failed to parse API response. Please contact us about this error.",
				Err:     err,
				Response: glair.Response{
					Status: status,
				},
			}
		}

		return &glair.Error{
			Code:    glair.ErrorCodeAPIError,
			Message: "GLAIR API returned non-OK response. Please check the Response property for more detailed explanation.",
			Response: glair.Response{
				Status: status,
				Body:   resBody,
			},
		}
	}

	return nil
}
