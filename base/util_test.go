package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomStr(t *testing.T) {
	str := RandomStr()
	assert.Len(t, str, 10)
}
