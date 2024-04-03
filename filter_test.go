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

func TestFilter_MultiFiles(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin1, _ := os.CreateTemp("", "")
	defer os.Remove(fin1.Name())
	fin1.WriteString(`

foo
bar

zoo


baz


`)
	fin1.Sync()

	fin2, _ := os.CreateTemp("", "")
	defer os.Remove(fin2.Name())
	fin2.WriteString(`

 xfoo
 xbar

 xzoo


 xbaz


`)
	fin2.Sync()

	options := &kuu.Options{
		Inplace: true,
	}

	err := kuu.Filter([]string{fin1.Name(), fin2.Name()}, options)
	require.NoError(err)

	out1, _ := os.ReadFile(fin1.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out1))

	out2, _ := os.ReadFile(fin2.Name())
	assert.Equal(` xfoo
 xbar

 xzoo

 xbaz
`, string(out2))
}
