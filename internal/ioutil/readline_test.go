package ioutil_test

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/kuu/internal/ioutil"
)

func TestReadLine(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	f, _ := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	f.WriteString("foo\n")  //nolint:errcheck
	f.WriteString("bar\n")  //nolint:errcheck
	f.WriteString("zoo\n")  //nolint:errcheck
	f.Sync()                //nolint:errcheck
	f.Seek(0, io.SeekStart) //nolint:errcheck

	buf := bufio.NewReader(f)

	for _, expected := range []string{"foo", "bar", "zoo"} {
		line, err := ioutil.ReadLine(buf)
		require.NoError(err)
		assert.Equal(expected, line)
	}
}

func TestReadLine_WithoutTailNewLine(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	f, _ := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	f.WriteString("foo\n")  //nolint:errcheck
	f.WriteString("bar\n")  //nolint:errcheck
	f.WriteString("zoo")    //nolint:errcheck
	f.Sync()                //nolint:errcheck
	f.Seek(0, io.SeekStart) //nolint:errcheck

	buf := bufio.NewReader(f)

	for _, expected := range []string{"foo", "bar", "zoo"} {
		line, err := ioutil.ReadLine(buf)
		require.NoError(err)
		assert.Equal(expected, line)
	}
}

func TestReadLine_EmptyFile(t *testing.T) {
	assert := assert.New(t)

	f, _ := os.CreateTemp("", "")
	buf := bufio.NewReader(f)

	_, err := ioutil.ReadLine(buf)
	assert.Equal(io.EOF, err)
}

func TestScanner(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	f, _ := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	f.WriteString("foo\n")  //nolint:errcheck
	f.WriteString("bar\n")  //nolint:errcheck
	f.WriteString("zoo\n")  //nolint:errcheck
	f.Sync()                //nolint:errcheck
	f.Seek(0, io.SeekStart) //nolint:errcheck

	scanner := ioutil.NewScanner(f)

	for _, expected := range []string{"foo", "bar", "zoo"} {
		line, err := scanner.Scan()
		require.NoError(err)
		assert.Equal(expected, line)
	}
}

func TestScanner_WithoutTailNewLine(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	f, _ := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	f.WriteString("foo\n")  //nolint:errcheck
	f.WriteString("bar\n")  //nolint:errcheck
	f.WriteString("zoo")    //nolint:errcheck
	f.Sync()                //nolint:errcheck
	f.Seek(0, io.SeekStart) //nolint:errcheck

	scanner := ioutil.NewScanner(f)

	for _, expected := range []string{"foo", "bar", "zoo"} {
		line, err := scanner.Scan()
		require.NoError(err)
		assert.Equal(expected, line)
	}
}

func TestScanner_EmptyFile(t *testing.T) {
	assert := assert.New(t)

	f, _ := os.CreateTemp("", "")
	scanner := ioutil.NewScanner(f)

	_, err := scanner.Scan()
	assert.Equal(io.EOF, err)
}
