package ioutil_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/kuu/internal/ioutil"
)

func TestScanner(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	f, _ := os.CreateTemp("", "")
	defer os.Remove(f.Name())
	f.WriteString("foo\n")
	f.WriteString("bar\n")
	f.WriteString("zoo\n")
	f.Sync()
	f.Seek(0, io.SeekStart)

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
	f.WriteString("foo\n")
	f.WriteString("bar\n")
	f.WriteString("zoo")
	f.Sync()
	f.Seek(0, io.SeekStart)

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
