package kuu_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/kuu"
)

func TestFilterFile(t *testing.T) {
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

	tf := kuu.NewTargetFile(fin.Name(), true, "")
	err := kuu.FilterFile(tf, true)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out))
}

func TestFilterFile_WithoutHeadEmpty(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString(`foo
bar

zoo


baz


`)
	fin.Sync()

	tf := kuu.NewTargetFile(fin.Name(), true, "")
	err := kuu.FilterFile(tf, true)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out))
}

func TestFilterFile_WithoutTailEmpty(t *testing.T) {
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

	tf := kuu.NewTargetFile(fin.Name(), true, "")
	err := kuu.FilterFile(tf, true)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out))
}

func TestFilterFile_WithoutTailNewLine(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fin, _ := os.CreateTemp("", "")
	defer os.Remove(fin.Name())
	fin.WriteString(`

foo
bar

zoo


baz`)
	fin.Sync()

	tf := kuu.NewTargetFile(fin.Name(), true, "")
	err := kuu.FilterFile(tf, true)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar

zoo

baz
`, string(out))
}

func TestFilterFile_WithoutContentNewLines(t *testing.T) {
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

	tf := kuu.NewTargetFile(fin.Name(), true, "")
	err := kuu.FilterFile(tf, true)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar
zoo
baz
`, string(out))
}

func TestFilterFile_NotUnique(t *testing.T) {
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

	tf := kuu.NewTargetFile(fin.Name(), true, "")
	err := kuu.FilterFile(tf, false)
	require.NoError(err)

	out, _ := os.ReadFile(fin.Name())
	assert.Equal(`foo
bar

zoo


baz
`, string(out))
}
