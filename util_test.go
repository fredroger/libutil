package libutil_test

import (
	"testing"

	"github.com/fredroger/libutil"
	"github.com/stretchr/testify/assert"
)

func TestMyFirstFunction(t *testing.T) {
	err := libutil.MyFirstFunction("patate")

	assert.Equal(t, nil, err)
}

func TestSecondFunction(t *testing.T) {
	value := libutil.MySecondFunction(0)
	assert.Equal(t, 100, value)

	value = libutil.MySecondFunction(-1)
	assert.Equal(t, -100, value)

	value = libutil.MySecondFunction(1)
	assert.Equal(t, -100, value)

}
