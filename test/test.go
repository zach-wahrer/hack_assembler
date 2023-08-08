package test

import "io"

type (
	TestReader struct {
		contents []byte
	}
)

func NewTestReader(contents []byte) io.Reader {
	return &TestReader{contents: contents}
}

func (r *TestReader) Read(p []byte) (n int, err error) {
	toCopyLen := len(p)
	if toCopyLen > len(r.contents) {
		toCopyLen = len(r.contents)
	}
	copiedLen := copy(p, r.contents[:toCopyLen])

	r.contents = r.contents[toCopyLen:]

	return copiedLen, nil
}
