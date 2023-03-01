package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync/atomic"
	"testing"
)

func TestFileSystemSource_Next(t *testing.T) {
	type fields struct {
		inputDir  string
		outputDir string
		files     []os.DirEntry
		index     atomic.Int32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  bool
	}{
		{
			name: "test1",
			fields: fields{
				inputDir:  "../testdata/resizer",
				outputDir: "../testdata/resizer",
			},
			want:  "resizer.jpeg",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileSystemSource{
				InputDir:  tt.fields.inputDir,
				OutputDir: tt.fields.outputDir,
			}
			f.Open()
			{
				got, got1 := f.Next()
				assert.Equal(t, "resizer.jpeg", got)
				assert.Equal(t, true, got1)
			}
			{
				got, got1 := f.Next()
				assert.Equal(t, "resizer_output.jpeg", got)
				assert.Equal(t, false, got1)
			}
		})
	}
}
