package one

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_CanAddToMake2020(t *testing.T) {
	result := CanAddToMake2020([]int{1721,
		979,
		366,
		299,
		675,
		1456})
	assert.Equal(t, result == 514579, true)

}
