package operator

import (
	"testing"
)

func TestDefaultOperator_Open(t *testing.T) {
	type args struct {
		inputPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Open",
			args: args{
				inputPath: "../testdata/resizer/resizer.jpeg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DefaultOperator{}
			got, err := d.Open(tt.args.inputPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Open() got nil, want not nil")
			}
		})
	}
}
