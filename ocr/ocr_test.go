package ocr

import (
	"testing"

	"github.com/glair-ai/glair-vision-go"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	config := glair.NewConfig("", "", "")
	ocr := New(config)

	assert.Equal(t, config, ocr.config)
}
