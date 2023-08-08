package hackassembler

import (
	"hackassembler/test"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssemble(t *testing.T) {
	type args struct {
		contents io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "should return a buffer of binary instructions",
			args: args{
				contents: test.NewTestReader([]byte("// Comment\n@2\nD=A")),
			},
			want: []byte("0000000000000010\n1110110000010000\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Assemble(tt.args.contents)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
