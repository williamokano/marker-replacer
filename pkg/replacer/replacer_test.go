package replacer

import (
	"testing"

	"github.com/spf13/afero"
)

func TestFileReplacer_Replace(t *testing.T) {
	type fields struct {
		Filename string
		fs       afero.Fs
	}
	type args struct {
		marker     string
		newContent string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FileReplacer{
				Filename: tt.fields.Filename,
				fs:       tt.fields.fs,
			}
			got, err := r.Replace(tt.args.marker, tt.args.newContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Replace() got = %v, want %v", got, tt.want)
			}
		})
	}
}
