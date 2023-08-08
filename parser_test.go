package hackassembler_test

import (
	"hackassembler"
	"hackassembler/test"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	type args struct {
		contents io.Reader
	}
	tests := []struct {
		name string
		args args
		want *hackassembler.Parser
	}{
		{
			name: "should return a parser",
			args: args{
				contents: test.NewTestReader([]byte("test")),
			},
			want: &hackassembler.Parser{
				Contents: test.NewTestReader([]byte("test")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hackassembler.NewParser(tt.args.contents))
		})
	}
}

func TestParser_Parse(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	tests := []struct {
		name             string
		fields           fields
		wantInstructions []hackassembler.Instruction
	}{
		{
			name: "should return A and C instructions",
			fields: fields{
				Contents: test.NewTestReader([]byte("// Comment\n@2\nD=A\n@3\n0;JMP")),
			},
			wantInstructions: []hackassembler.Instruction{
				{
					Type:   hackassembler.AInstruction,
					Symbol: "2",
				},
				{
					Type: hackassembler.CInstruction,
					Dest: "D",
					Comp: "A",
				},
				{
					Type:   hackassembler.AInstruction,
					Symbol: "3",
				},
				{
					Type: hackassembler.CInstruction,
					Comp: "0",
					Jump: "JMP",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{
				Contents: tt.fields.Contents,
			}
			got := p.Parse()
			assert.NotNil(t, got)

			gotIns := []hackassembler.Instruction{}
			for ins := range got {
				gotIns = append(gotIns, ins)
			}

			assert.Equal(t, tt.wantInstructions, gotIns)
		})
	}
}

func TestParser_SkipLine(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should skip commented lines",
			args: args{
				line: "// Skip me",
			},
			want: true,
		},
		{
			name: "should skip empty lines",
			args: args{
				line: "",
			},
			want: true,
		},
		{
			name: "should not skip line",
			args: args{
				line: "D=A",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{}
			assert.Equal(t, tt.want, p.SkipLine(tt.args.line))
		})
	}
}

func TestParser_GetInstructionType(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   hackassembler.InstructionType
	}{
		{
			name: "should return an A instruction type",
			args: args{
				line: "@2",
			},
			want: hackassembler.AInstruction,
		},
		{
			name: "should return a L instruction type",
			args: args{
				line: "(LOOP)",
			},
			want: hackassembler.LInstruction,
		},
		{
			name: "should return a C instruction type",
			args: args{
				line: "D=D+A",
			},
			want: hackassembler.CInstruction,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{}
			got := p.GetInstructionType(tt.args.line)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParser_GetSymbol(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should return decimal symbol",
			args: args{
				line: "@2",
			},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{}
			assert.Equal(t, tt.want, p.GetSymbol(tt.args.line))
		})
	}
}

func TestParser_GetDest(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should return a dest when = is present",
			args: args{
				line: "D=D+A",
			},
			want: "D",
		},
		{
			name: "should return empty when no = is present",
			args: args{
				line: "@2",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{}
			assert.Equal(t, tt.want, p.GetDest(tt.args.line))
		})
	}
}

func TestParser_GetComp(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should return comp when there is a jump instruction",
			args: args{
				line: "0;JMP",
			},
			want: "0",
		},
		{
			name: "should return comp when there is a dest instruction",
			args: args{
				line: "D=D+A",
			},
			want: "D+A",
		},
		{
			name: "should return comp when there is a dest and jump instruction",
			args: args{
				line: "D=D+A;JMP",
			},
			want: "D+A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{}
			assert.Equal(t, tt.want, p.GetComp(tt.args.line))
		})
	}
}

func TestParser_GetJump(t *testing.T) {
	type fields struct {
		Contents io.Reader
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should return jump when ; is present",
			args: args{
				line: "0;JMP",
			},
			want: "JMP",
		},
		{
			name: "should return empty when no ; is present",
			args: args{
				line: "D=D+A",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &hackassembler.Parser{}
			assert.Equal(t, tt.want, p.GetJump(tt.args.line))
		})
	}
}
