package testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	//rst := square(9)
	//if rst != 81 {
	//	t.Errorf("square(9) = %d, want 81", rst)
	//}
	assert1 := assert.New(t)
	assert1.Equal(81, square(9), "\"square(9) = %d, want 81\"")
}

func TestSquare2(t *testing.T) {
	assert2 := assert.New(t)
	assert2.Equal(9, square(3), "\"square(3) = %d, want 3\"")
}
