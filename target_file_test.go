package kuu

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/kuu/internal/ioutil"
)

func TestNewTargetFile(t *testing.T) {
	assert := assert.New(t)

	tf := NewTargetFile("foo.txt", true, ".bak")
	assert.Equal("foo.txt", tf.Fname)
	assert.True(tf.Inplace)
	assert.Equal("foo.txt.bak", tf.Backup)
}

func TestNewTargetFile_Stdin(t *testing.T) {
	assert := assert.New(t)

	{
		tf1 := NewTargetFile("", true, ".bak")
		assert.Equal("", tf1.Fname)
		assert.False(tf1.Inplace)
		assert.Empty(tf1.Backup)
	}

	{
		tf2 := NewTargetFile("", false, "")
		assert.Equal("", tf2.Fname)
		assert.False(tf2.Inplace)
		assert.Empty(tf2.Backup)
	}
}

func TestNewTargetFile_NotInplace(t *testing.T) {
	assert := assert.New(t)

	{
		tf1 := NewTargetFile("foo.txt", false, ".bak")
		assert.Equal("foo.txt", tf1.Fname)
		assert.False(tf1.Inplace)
		assert.Empty(tf1.Backup)
	}

	{
		tf2 := NewTargetFile("foo.txt", false, "")
		assert.Equal("foo.txt", tf2.Fname)
		assert.False(tf2.Inplace)
		assert.Empty(tf2.Backup)
	}
}

func TestTargetFileFilter_Stdout(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString("foo\n")
	fin.WriteString("bar\n")
	fin.WriteString("zoo\n")
	fin.Sync()

	fout, _ := os.CreateTemp("", "")
	defer os.Remove(fout.Name())
	origStdout := stdout
	stdout = fout
	t.Cleanup(func() { stdout = origStdout })

	tf := NewTargetFile(fin.Name(), false, "")

	tf.Filter(func(scanner *ioutil.Scanner, w io.Writer) error {
		for _, expected := range []string{"foo", "bar", "zoo"} {
			line, err := scanner.Scan()
			require.NoError(err)
			assert.Equal(expected, line)
		}

		_, err := scanner.Scan()
		assert.Equal(io.EOF, err)
		fmt.Fprint(w, "aaa\nbbb\nccc\n")
		return nil
	})

	fout.Seek(0, io.SeekStart)
	out, _ := io.ReadAll(fout)
	assert.Equal("aaa\nbbb\nccc\n", string(out))
}

func TestTargetFileFilter_Inplace(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString("foo\n")
	fin.WriteString("bar\n")
	fin.WriteString("zoo\n")
	fin.Sync()

	tf := NewTargetFile(fin.Name(), true, "")

	tf.Filter(func(scanner *ioutil.Scanner, w io.Writer) error {
		for _, expected := range []string{"foo", "bar", "zoo"} {
			line, err := scanner.Scan()
			require.NoError(err)
			assert.Equal(expected, line)
		}

		_, err := scanner.Scan()
		assert.Equal(io.EOF, err)
		fmt.Fprint(w, "aaa\nbbb\nccc\n")
		return nil
	})

	out, _ := os.ReadFile(fin.Name())
	assert.Equal("aaa\nbbb\nccc\n", string(out))
}

func TestTargetFileFilter_InplaceBackup(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString("foo\n")
	fin.WriteString("bar\n")
	fin.WriteString("zoo\n")
	fin.Sync()

	tf := NewTargetFile(fin.Name(), true, ".bak")

	tf.Filter(func(scanner *ioutil.Scanner, w io.Writer) error {
		for _, expected := range []string{"foo", "bar", "zoo"} {
			line, err := scanner.Scan()
			require.NoError(err)
			assert.Equal(expected, line)
		}

		_, err := scanner.Scan()
		assert.Equal(io.EOF, err)
		fmt.Fprint(w, "aaa\nbbb\nccc\n")
		return nil
	})

	out, _ := os.ReadFile(fin.Name())
	assert.Equal("aaa\nbbb\nccc\n", string(out))

	bak, _ := os.ReadFile(fin.Name() + ".bak")
	assert.Equal("foo\nbar\nzoo\n", string(bak))
	os.Remove(fin.Name() + ".bak")
}

func TestTargetFileFilter_Stdin(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString("foo\n")
	fin.WriteString("bar\n")
	fin.WriteString("zoo\n")
	fin.Sync()
	fin.Seek(0, io.SeekStart)

	origStdin := stdin
	stdin = fin
	t.Cleanup(func() { stdin = origStdin })

	fout, _ := os.CreateTemp("", "")
	defer os.Remove(fout.Name())
	origStdout := stdout
	stdout = fout
	t.Cleanup(func() { stdout = origStdout })

	tf := NewTargetFile("", false, "")

	tf.Filter(func(scanner *ioutil.Scanner, w io.Writer) error {
		for _, expected := range []string{"foo", "bar", "zoo"} {
			line, err := scanner.Scan()
			require.NoError(err)
			assert.Equal(expected, line)
		}

		_, err := scanner.Scan()
		assert.Equal(io.EOF, err)
		fmt.Fprint(w, "aaa\nbbb\nccc\n")
		return nil
	})

	out, _ := os.ReadFile(fout.Name())
	assert.Equal("aaa\nbbb\nccc\n", string(out))
}
