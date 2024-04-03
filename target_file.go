package kuu

import (
	"io"
	"os"

	"github.com/google/renameio/v2"
	cp "github.com/otiai10/copy"
	"github.com/winebarrel/kuu/internal/ioutil"
)

var (
	stdin  = os.Stdin
	stdout = os.Stdout
)

type TargetFile struct {
	Fname   string
	Inplace bool
	Backup  string
}

func NewTargetFile(fname string, inplace bool, backupSuffix string) *TargetFile {
	inplace = inplace && fname != ""

	tf := &TargetFile{
		Fname:   fname,
		Inplace: inplace,
	}

	if inplace && backupSuffix != "" {
		tf.Backup = fname + backupSuffix
	}

	return tf
}

func (tf *TargetFile) Filter(proc func(*ioutil.Scanner, io.Writer) error) error {
	var r io.Reader = stdin

	if tf.Fname != "" {
		f, err := os.Open(tf.Fname)

		if err != nil {
			return err
		}

		defer f.Close()
		r = f
	}

	scanner := ioutil.NewScanner(r)

	if !tf.Inplace {
		return proc(scanner, stdout)

	}

	t, err := renameio.NewPendingFile(tf.Fname, renameio.WithExistingPermissions())

	if err != nil {
		return err
	}

	defer t.Cleanup() //nolint:errcheck
	err = proc(scanner, t)

	if err != nil {
		return err
	}

	if tf.Backup != "" {
		err := cp.Copy(tf.Fname, tf.Backup)

		if err != nil {
			return err
		}
	}

	return t.CloseAtomicallyReplace()
}
