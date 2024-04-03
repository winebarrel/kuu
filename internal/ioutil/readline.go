package ioutil

import (
	"bufio"
	"bytes"
	"io"
)

func ReadLine(r *bufio.Reader) (string, error) {
	var buf bytes.Buffer

	for {
		line, isPrefix, err := r.ReadLine()
		n := len(line)

		if n > 0 {
			buf.Write(line)
		}

		if !isPrefix || err != nil {
			return buf.String(), err
		}
	}
}

type Scanner struct {
	buf bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	scanner := &Scanner{
		buf: *bufio.NewReader(r),
	}

	return scanner
}

func (scanner *Scanner) Scan() (string, error) {
	return ReadLine(&scanner.buf)
}
