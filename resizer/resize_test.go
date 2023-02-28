package resizer

import (
	"reflect"
	"testing"
)

func Test_resizeImage(t *testing.T) {
	type args struct {
		inputPath  string
		outputPath string
		maxWidth   uint
		maxHeight  uint
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 int
		want2 int
	}{
		{
			name: "ThumbnailImage",
			args: args{
				inputPath:  "../testdata/resizer/resizer.jpeg",
				outputPath: "../testdata/resizer/resizer_output.jpeg",
				maxWidth:   500,
				maxHeight:  500,
			},
			want:  nil,
			want1: 375,
			want2: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ThumbnailImage(tt.args.inputPath, tt.args.outputPath, tt.args.maxWidth, tt.args.maxHeight)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThumbnailImage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThumbnailImage() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ThumbnailImage() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
