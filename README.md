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
    <a href="https://github.com/glair-ai/glair-vision-java/blob/main/LICENSE"><img src="https://img.shields.io/npm/l/@glair/vision" alt="License"></a>
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

The configuration object will be initialized with the following options:

| Option       | Default                       | Description                         |
| ------------ | ----------------------------- | ----------------------------------- |
| `BaseUrl`    | `https://api.vision.glair.ai` | Base URL for GLAIR Vision API                |
| `apiVersion` | `v1`                          | GLAIR Vision API version to be used |

Afterwards, you can use the provided functions to access GLAIR Vision API:

1. [OCR](#ocr)

## API

### OCR



## Examples

You can see the examples in the [examples folder](./examples) 