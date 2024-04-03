package kuu_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/kuu"
)

func TestFilter(t *testing.T) {
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

	options := &kuu.Options{
		Inplace: true,
	}

	err := kuu.Filter([]string{fin.Name()}, options)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out))
}
