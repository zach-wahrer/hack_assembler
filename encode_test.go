package hackassembler_test

import (
	"hackassembler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	type args struct {
		ins hackassembler.Instruction
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "should return binary of A instruction",
			args: args{
				ins: hackassembler.Instruction{
					Type:   hackassembler.AInstruction,
					Symbol: "2",
				},
			},
			want: "0000000000000010",
		},
		{
			name: "should return binary of larger A instruction",
			args: args{
				ins: hackassembler.Instruction{
					Type:   hackassembler.AInstruction,
					Symbol: "20",
				},
			},
			want: "0000000000010100",
		},
		{
			name: "should return binary of C instruction",
			args: args{
				ins: hackassembler.Instruction{
					Type: hackassembler.CInstruction,
					Dest: "M",
					Comp: "M+1",
					Jump: "JMP",
				},
			},
			want: "1111110111001111",
		},
		{
			name: "should return binary of complex C instruction",
			args: args{
				ins: hackassembler.Instruction{
					Type: hackassembler.CInstruction,
					Dest: "MD",
					Comp: "M+1",
				},
			},
			want: "1111110111011000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := hackassembler.Translate(tt.args.ins)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestTranslateDest(t *testing.T) {
	type args struct {
		dest string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should translate empty dest",
			args: args{
				dest: "",
			},
			want: "000",
		},
		{
			name: "should translate full dest",
			args: args{
				dest: "ADM",
			},
			want: "111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hackassembler.TranslateDest(tt.args.dest))
		})
	}
}

func TestTranslateComp(t *testing.T) {
	type args struct {
		comp string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should translate 0 comp",
			args: args{
				comp: "0",
			},
			want: "0101010",
		},
		{
			name: "should translate D|M comp",
			args: args{
				comp: "D|M",
			},
			want: "1010101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hackassembler.TranslateComp(tt.args.comp))
		})
	}
}

func TestTranslateJump(t *testing.T) {
	type args struct {
		jump string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should translate empty jump",
			args: args{
				jump: "",
			},
			want: "000",
		},
		{
			name: "should translate unconditional jump",
			args: args{
				jump: "JMP",
			},
			want: "111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hackassembler.TranslateJump(tt.args.jump))
		})
	}
}
