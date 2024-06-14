package strings

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/williamokano/marker-replacer/internal/helpers"
)

func TestReadAllLines(t *testing.T) {
	type args struct {
		fs       afero.Fs
		filename string
		expected string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should read all lines",
			args: args{
				fs:       afero.NewMemMapFs(),
				filename: "testdata/readAllLines.txt",
				expected: input,
			},
			want:    input,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := helpers.MockFilePath(tt.args.fs, tt.args.filename, tt.args.expected)
			if err != nil {
				t.Errorf("MockFilePath() error = %v", err)
			}

			got, err := ReadAllLines(tt.args.fs, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAllLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ReadAllLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var input = `hello my
awesome and beautiful
world that does exist`
