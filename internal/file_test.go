package internal

import (
	"os"
	"testing"

	"github.com/glair-ai/glair-vision-go"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	file, _ := os.Open("../examples/ocr/images/ktp.jpeg")

	tests := []struct {
		name    string
		file    interface{}
		want    bool
		wantErr glair.ErrorCode
	}{
		{
			name:    "should do nothing on real file",
			file:    file,
			want:    true,
			wantErr: "",
		},
		{
			name:    "should read file from path",
			file:    "../examples/ocr/images/ktp.jpeg",
			want:    true,
			wantErr: "",
		},
		{
			name:    "should return invalid file error when file is not expected type",
			file:    12321321,
			want:    false,
			wantErr: glair.ErrorCodeFileError,
		},
		{
			name:    "should return error if filepath is invalid",
			file:    "not_exist.jpg",
			want:    false,
			wantErr: glair.ErrorCodeFileError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			file, err := ReadFile(tc.file)

			assert.Equal(t, tc.want, file != nil)
			if glairErr, ok := err.(*glair.Error); ok {
				assert.Equal(t, tc.wantErr, glairErr.Code)
			}
		})
	}
}
