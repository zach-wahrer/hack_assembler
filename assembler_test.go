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
		{
			name: "should return a buffer of binary instructions for a full program",
			args: args{
				contents: test.NewTestReader(fullRectProgramWithoutSymbolicRefs()),
			},
			want: []byte("0000000000000000\n1111110000010000\n0000000000010111\n1110001100000110\n0000000000010000\n1110001100001000\n0100000000000000\n1110110000010000\n0000000000010001\n1110001100001000\n0000000000010001\n1111110000100000\n1110111010001000\n0000000000010001\n1111110000010000\n0000000000100000\n1110000010010000\n0000000000010001\n1110001100001000\n0000000000010000\n1111110010011000\n0000000000001010\n1110001100000001\n0000000000010111\n1110101010000111\n"),
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

func fullRectProgramWithoutSymbolicRefs() []byte {
	return []byte("@0\nD=M\n@23\nD;JLE\n@16\nM=D\n@16384\nD=A\n@17\nM=D\n@17\nA=M\nM=-1\n@17\nD=M\n@32\nD=D+A\n@17\nM=D\n@16\nMD=M-1\n@10\nD;JGT\n@23\n0;JMP")

}
