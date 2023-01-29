package modules_separator

import (
	"github.com/spf13/afero"
	"log"
	"testing"
)

func TestModulesSeparator_ProceedInputFile(t *testing.T) {
	type fields struct {
		FS        afero.Fs
		Time      TimeInterface
		outWriter afero.File
		errWriter afero.File
		errLogger *log.Logger
	}
	type args struct {
		filename   string
		targetPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "success",
			fields: fields{FS: DefaultFS, Time: DefaultTime, outWriter: DefaultOutWriter, errWriter: DefaultErrWriter},
			args:   args{filename: "test_data.cs", targetPath: DefaultTargetPath}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			separator := NewModulesSeparator(
				tt.fields.FS, tt.fields.Time, tt.fields.outWriter, tt.fields.errWriter, tt.fields.errLogger)
			separator.ProceedInputFile(tt.args.filename, tt.args.targetPath)
		})
	}
}
