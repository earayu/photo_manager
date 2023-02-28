package cutter

import (
	"reflect"
	"testing"
)

func Test_cutImage(t *testing.T) {
	type args struct {
		inputPath    string
		outputPath   string
		targetWidth  int
		targetHeight int
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 int
		want2 int
	}{
		{
			name: "cutImage",
			args: args{
				inputPath:    "../testdata/resizer/resizer.jpeg",
				outputPath:   "../testdata/resizer/resizer_output.jpeg",
				targetWidth:  500,
				targetHeight: 500,
			},
			want:  nil,
			want1: 500,
			want2: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := cutImage(tt.args.inputPath, tt.args.outputPath, tt.args.targetWidth, tt.args.targetHeight)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cutImage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("cutImage() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("cutImage() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_cutImageByRatio(t *testing.T) {
	type args struct {
		inputPath    string
		outputPath   string
		widthWeight  int
		heightWeight int
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 int
		want2 int
	}{
		{
			name: "CutImageByRatio",
			args: args{
				inputPath:    "../testdata/resizer/resizer.jpeg",
				outputPath:   "../testdata/resizer/resizer_output.jpeg",
				widthWeight:  1,
				heightWeight: 1,
			},
			want:  nil,
			want1: 3024,
			want2: 3024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := CutImageByRatio(tt.args.inputPath, tt.args.outputPath, tt.args.widthWeight, tt.args.heightWeight)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CutImageByRatio() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CutImageByRatio() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CutImageByRatio() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
