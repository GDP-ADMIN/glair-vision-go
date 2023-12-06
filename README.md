<p align="center">
  <a href="https://docs.glair.ai/vision" target="_blank">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://glair-chart.s3.ap-southeast-1.amazonaws.com/images/glair-horizontal-logo-blue.png">
      <source media="(prefers-color-scheme: light)" srcset="https://glair-chart.s3.ap-southeast-1.amazonaws.com/images/glair-horizontal-logo-color.png">
      <img alt="GLAIR" src="https://glair-chart.s3.ap-southeast-1.amazonaws.com/images/glair-horizontal-logo-color.png" width="180" height="60" style="max-width: 100%;">
    </picture>
  </a>
</p>

<p align="center">
  GLAIR Vision Go SDK
<p>

<p align="center">
    <a href="https://github.com/glair-ai/glair-vision-go/blob/main/LICENSE"><img src="https://img.shields.io/npm/l/@glair/vision" alt="License"></a>
</p>

## Requirement

- Go 1.18 or later with Go modules.

## Installation

You can import the SDK in your Go files with `import`:

```go
import (
    "github.com/glair-ai/glair-vision-go"
    "github.com/glair-ai/glair-vision-go/client"
)
```

After that, you can run `go` commands and let the Go toolchain resolve and fetch the SDK module automatically.

Alternatively, you can also run `go get` to explicitly resolve and fetch the SDK module:

```bash
go get -u github.com/glair-ai/glair-vision-go
```

## Usage

The package needs to be configured with your credentials, see [here](https://docs.glair.ai/vision/authentication) for more detailed instructions

```go
package main

import (
    "github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
    config := glair.NewConfig("<username>", "<password>", "<api_key>")

    client := client.New(config)
}
```

The configuration object will be initialized with the following values:

| Option       | Default                          | Description                                                                 |
| ------------ | -------------------------------- | --------------------------------------------------------------------------- |
| `BaseUrl`    | `https://api.vision.glair.ai`    | Base URL for GLAIR Vision API                                               |
| `ApiVersion` | `v1`                             | GLAIR Vision API version to be used                                         |
| `Client`     | Default Go HTTP client           | HTTP Client to be used when sending request to GLAIR Vision API             |
| `Logger`     | `LeveledLogger` with `LevelNone` | Logger instace to be used to log errors, information, or debugging messages |

You can change the above values using the provided `With<Option>` method of the configuration object, for example:

```go
package main

import (
    "github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
    config := glair.NewConfig("<username>", "<password>", "<api_key>")
    // set the base url to `http://localhost:3000` 
    config = config.WithBaseURL("http://localhost:3000")

    client := client.New(config)
}
```

Afterwards, you can use the provided functions to access GLAIR Vision API.

## Documentation

For comprehensive list of available API provided by GLAIR Vision Go SDK, check out the [API Documentation](https://docs.glair.ai/vision). You can also see the runnable examples in the [examples folder](./examples). For details on all the functionality in this library, see the Go documentation.

Below are a few simple usage examples:

### Perform OCR on KTP

```go
package main

import (
	"context"
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

	file, _ := os.Open("path/to/image.jpg")

	result, err := client.Ocr.KTP(ctx, glair.OCRInput{
		Image: file,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

  	fmt.Println(result.Read.Nama)
}
```

### Perform OCR on Receipt by providing path to the image file

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
	ctx := context.Background()

	config := glair.NewConfig("", "", "")
	client := client.New(config)

	result, err := client.Ocr.Receipt(ctx, glair.OCRInput{
		Image: "path/to/image.jpg",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

  	fmt.Println(result.Read.Nama)
}
```

### Perform OCR on KTP with timeout

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

func main() {
	baseContext := context.Background()
	contextWithTimeout, cancel := context.WithTimeout(baseContext, 100*time.Millisecond)
	defer cancel()

	config := glair.NewConfig("", "", "")
	client := client.New(config)

	file, _ := os.Open("../images/ktp.jpeg")

	result, err := client.Ocr.Ktp(contextWithTimeout, glair.OCRInput{
		Image: file,
	})

	if err != nil {
		if glairErr, ok := err.(*glair.Error); ok {
			switch glairErr.Code {
			case glair.ErrorCodeTimeout:
				log.Printf("Request timed out")
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

```

### Using custom HTTP client to intercept HTTP requests

```go
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/glair-ai/glair-vision-go"
)

// MyClient is a HTTP client that adds `x-powered-by`
// header to a normal HTTP request.
//
// It wraps the default HTTP client
type MyClient struct {
	client *http.Client
}

func (c *MyClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Powered-By", "GLAIR")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	ctx := context.Background()

	config := glair.NewConfig("", "", "").WithClient(&MyClient{client: http.DefaultClient})
	client := New(config)

	file, _ := os.Open("../images/ktp.jpeg")

	result, err := client.Ocr.KTP(ctx, glair.OCRInput{
		Image: file,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}

```

### Perform face verification using GLAIR Vision Face Verification API

```go
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

	image, _ := os.Open("path/to/image.jpg")

	result, err := client.FaceBio.FaceMatching(ctx, glair.FaceMatchingInput{
		StoredImage:   image,
		CapturedImage: image,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}
```

### Perform KTP data verification using GLAIR Identity Verification API

```go
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
		DateOfBirth: ""
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	beautified, _ := json.MarshalIndent(result, "", "  ")

	fmt.Println(string(beautified))
}

```

## Error Handling

Whenever an error occurs, GLAIR Vision Go SDK will wrap the error into a `glair.Error` object that contains the following properties

| Property   | Type        | Description                                                                                                   |
| ---------- | ----------- | ------------------------------------------------------------------------------------------------------------- |
| `Code`     | `ErrorCode` | Unique identifier that distinguish errors                                                                     |
| `Message`  | `string`    | Human-readable error message. Contains basic information of error cause                                       |
| `Err`      | `error`     | The original error object returned by the SDK                                                                 |
| `Response` | `Response`  | GLAIR Vision API response body. Only available if the request has been successfully sent the GLAIR Vision API |

It's recommended to assert the error to `glair.Error` whenever an error is returned, for example:

```go
package main

import (
	"context"
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

	file, _ := os.Open("path/to/image.jpg")

	result, err := client.Ocr.KTP(ctx, glair.OCRInput{
		Image: file,
	})

	if err != nil {
    		// is a glair.Error, assert the error code
	  	if glairErr, ok := err.(*glair.Error); ok {
      		switch glairErr.Code {
        		case glair.ErrorCodeFileError:
          			fmt.Println("Cannot read input file correctly")
        		case glair.ErrorCodeNetworkError:
          			fmt.Println("There are problems while connecting to GLAIR Vision API")
        		default:
          			fmt.Printf("GLAIR SDK returns error code: %d", glairErr.Code)
      			}
    		} else {
      			fmt.Printf("Unexpected error occured: %w", err)
    		}
	}
}
```

### Error Code

To make debugging errors easier, GLAIR Vision Go SDK provides error code to all `glair.Error` objects. Below are the list of error codes that are returned by GLAIR Vision Go SDK

| Error Code                 | Reason                                                                                                                                                                                                                                                          |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ErrorCodeFileError`       | The SDK encounters an error when processing the image file. It's possible that the file doesn't exist, the SDK cannot access the file from the given path, or the file is corrupted. This code also returned when incorrect representation of file is provided. |
| `ErrorCodeNetworkError`    | The SDK fails to complete the HTTP request with the given HTTP client                                                                                                                                                                                           |
| `ErrorCodeTimeout`         | The network request sent by the SDK has timed out. This error is fixable by increasing the request timeout duration or by completely removing the timeout from the context                                                                                      |
| `ErrorCodeForbidden`       | The SDK attempts to access an API endpoint with insufficient credentials. Please contact us if you think that this is a mistake                                                                                                                                 |
| `ErrorCodeAPIError`        | GLAIR Vision API returns a non-OK response. Please inspect the `Response` object for more detailed explanation if this code is returned                                                                                                                         |
| `ErrorCodeInvalidResponse` | GLAIR Vision API returns an unexpected response. Please contact us if you receive this error code                                                                                                                                                               |

### `Response`

When error with code `ErrorCodeAPIError` is returned, GLAIR Vision SDK with return additional context of the failure encapsulated in the `Response` object. The `Response` object has the following properties.

| Property | Description                                    |
| -------- | ---------------------------------------------- |
| `Code`   | HTTP Status code returned by GLAIR Vision API  |
| `Body`   | Raw response body returned by GLAIR Vision API |

## Logging

By default, GLAIR Vision Go SDK does not log anything regardless of severity. However, you can enable logging implementing the `Logger` interface from the main package and add it to the configuration object with `WithLogger` method.

```go
package main

import (
	"fmt"

	"github.com/glair-ai/glair-vision-go"
	"github.com/glair-ai/glair-vision-go/client"
)

type MyLogger struct {}

func (l MyLogger) Debugf(format string, val ...interface{}) {
	// do not log debug messages
}

func (l MyLogger) Infof(format string, val ...interface{}) {
	fmt.Printf("[GLAIR - Information] " + format, val)
}

func (l MyLogger) Warnf(format string, val ...interface{}) {
	fmt.Printf("[GLAIR - Warning] " + format, val)
}

func (l MyLogger) Errorf(format string, val ...interface{}) {
	fmt.Printf("[GLAIR - Debug] " + format, val)
}

func main() {
	config := glair.NewConfig("<username>", "<password>", "<api_key>").
		WithLogger(MyLogger{})
}
```

The `Logger` interface has the following signature

```go
type Logger interface {
	Debugf(format string, val ...interface{})
	Infof(format string, val ...interface{})
	Warnf(format string, val ...interface{})
	Errorf(format string, val ...interface{})
}
```

Alternatively, the SDK provides convenient `LeveledLogger` struct that implements the
`Logger` interface.

```go
type LeveledLogger struct {
	Level LogLevel
}
```

`LeveledLogger` accepts `Level` property that determines what messages should be logged. Below is the list of available `Level` for `LeveledLogger`

| Level        | Value | Description                                                |
| ------------ | ----- | ---------------------------------------------------------- |
| `LevelNone`  | `0`   | Do not log anything                                        |
| `LevelError` | `1`   | Log all error messages and output them to `stderr`         |
| `LevelWarn`  | `2`   | Log all warning messages and output them to `stdout`       |
| `LevelInfo`  | `3`   | Log all informational messages and output them to `stdout` |
| `LevelDebug` | `4`   | Log all debugging messages and output them to `stdout`     |

All `Level` property also logs any messages below their `Level`. For example, `LeveledLogger` with `LevelInfo` will log informational, warnings, and error messages.

## License

This project is licensed under [the MIT License](./LICENSE)
