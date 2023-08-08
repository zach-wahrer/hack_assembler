package hackassembler

import (
	"bytes"
	"io"
)

func Assemble(contents io.Reader) ([]byte, error) {
	parser := NewParser(contents)

	buf := bytes.NewBuffer([]byte{})
	for ins := range parser.Parse() {
		trans, err := Translate(ins)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write([]byte(trans + "\n")); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
