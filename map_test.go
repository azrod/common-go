package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {

	slice := []string{"a", "b", "c", "d"}
	val := "c"
	index, found := Find(slice, val)
	assert.Equal(t, 3, index)
	assert.True(t, found)

	sliceNotFound := []string{"a", "b", "c", "d"}
	valNotFound := "e"
	indexNotFound, notFound := Find(sliceNotFound, valNotFound)
	assert.Equal(t, -1, indexNotFound)
	assert.False(t, notFound)

}
