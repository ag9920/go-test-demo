package testifyexp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSomethingRequire(t *testing.T) {
	require := require.New(t)
	expectResult := int(5)
	result := dummyFn()
	require.Equal(expectResult, result)
}

func TestSomethingV2(t *testing.T) {
	assert := assert.New(t)
	expectResult := int(5)
	result := dummyFn()
	assert.Equal(expectResult, result)
}

func TestSomething(t *testing.T) {
	expectResult := int(5)
	result := dummyFn()
	assert.Equal(t, expectResult, result)
}

func dummyFn() int {
	return 7
}
