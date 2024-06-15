package replacer

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/williamokano/marker-replacer/internal/helpers"
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
		name         string
		fields       fields
		args         args
		fileContents string
		want         string
		wantErr      bool
	}{
		{
			name: "should replace if marker found",
			args: args{
				marker:     "command",
				newContent: "new value 1\nnew value 2\nnew value 3",
			},
			fields: fields{
				Filename: "whatever",
				fs:       afero.NewMemMapFs(),
			},
			fileContents: originalInput,
			want:         expectedInput,
			wantErr:      false,
		},
		{
			name: "should replace if inline marker found",
			args: args{
				marker:     "marker",
				newContent: "Our cool written",
			},
			fields: fields{
				Filename: "whatever",
				fs:       afero.NewMemMapFs(),
			},
			fileContents: inlineOriginalInput,
			want:         inlineExpectedInput,
			wantErr:      false,
		},
		{
			name: "should replace if mixed marker found",
			args: args{
				marker:     "marker",
				newContent: "another message",
			},
			fields: fields{
				Filename: "whatever",
				fs:       afero.NewMemMapFs(),
			},
			fileContents: mixedOriginalInput,
			want:         mixedExpectedInput,
			wantErr:      false,
		},
		{
			name: "should keep original text if marker not found",
			args: args{
				marker:     "something",
				newContent: "new value 1\nnew value 2\nnew value 3",
			},
			fields: fields{
				Filename: "whatever",
				fs:       afero.NewMemMapFs(),
			},
			fileContents: originalInput,
			want:         originalInput,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := helpers.MockFilePath(tt.fields.fs, tt.fields.Filename, tt.fileContents)
			if err != nil {
				t.Errorf("MockFilePath() error = %v", err)
			}
			r := NewFileReplacer(tt.fields.fs, tt.fields.Filename)
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

var originalInput = `Regular text
should ignore this part
<!--command-->
old value 1
old value 2
old value 3
<!--/command-->

this should still be here`

var expectedInput = `Regular text
should ignore this part
<!--command-->
new value 1
new value 2
new value 3
<!--/command-->

this should still be here`

var inlineOriginalInput = `Hello World! <!--marker-->My awesome text<!--/marker-->`
var inlineExpectedInput = `Hello World! <!--marker-->Our cool written<!--/marker-->`

var mixedOriginalInput = `Regular text<!--marker-->
some text<!--/marker-->`
var mixedExpectedInput = `Regular text<!--marker-->
another message<!--/marker-->`
