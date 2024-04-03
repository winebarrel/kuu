package kuu

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/winebarrel/kuu/internal/ioutil"
)

type Options struct {
	Inplace       bool   `kong:"short='i',xor='inplace,inplace-backup',help='Edit files in place.'"`
	InplaceBackup string `kong:"short='b',xor='inplace,inplace-backup',help='Edit files in place (makes backup with specified suffix).'"`
	Unique        bool   `kong:"short='u',xor='inplace,inplace-backup',help='Make duplicate blank lines unique.'"`
}

func Filter(files []string, options *Options) error {
	if len(files) == 0 {
		files = []string{""}
	}

	inplace := options.Inplace || options.InplaceBackup != ""

	for _, fname := range files {
		tf := NewTargetFile(fname, inplace, options.InplaceBackup)
		err := FilterFile(tf, options.Unique)

		if err != nil {
			return err
		}
	}

	return nil
}

func FilterFile(tf *TargetFile, unique bool) error {
	return tf.Filter(func(scanner *ioutil.Scanner, w io.Writer) error {
		var buf strings.Builder
		out := bufio.NewWriter(w)
		defer out.Flush()
		hasAny := false
		prevEmpty := false

		for {
			line, err := scanner.Scan()

			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}

			empty := line == ""
			hasAny = hasAny || !empty

			// Skip head empty lines
			if !hasAny {
				continue
			}

			// Skip duplicated empty lines
			if unique && empty && prevEmpty {
				continue
			}

			fmt.Fprintln(&buf, line)

			if !empty {
				out.WriteString(buf.String()) //nolint:errcheck
				buf.Reset()
			}

			prevEmpty = empty
		}

		return nil
	})
}
