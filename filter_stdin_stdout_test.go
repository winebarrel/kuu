package kuu

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilter_Stdin(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString(`

foo
bar

zoo


baz


`)
	fin.Sync()
	fin.Seek(0, io.SeekStart)

	origStdin := stdin
	stdin = fin
	t.Cleanup(func() { stdin = origStdin })

	options := &Options{
		Inplace: false,
	}

	fout, _ := os.CreateTemp("", "")
	defer os.Remove(fout.Name())
	origStdout := stdout
	stdout = fout
	t.Cleanup(func() { stdout = origStdout })

	err := Filter([]string{""}, options)
	require.NoError(err)

	out, _ := os.ReadFile(fout.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out))
}
