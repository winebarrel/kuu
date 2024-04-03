package ioutil

import (
	"bufio"
	"io"
)

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
