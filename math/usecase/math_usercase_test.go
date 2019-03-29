package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMathUseCase(t *testing.T) {
	mathUseCase:=NewSimpleMathUseCase()
	assert.True(t,mathUseCase.IsMax(10))
	assert.True(t,mathUseCase.IsMax(20))
	assert.False(t,mathUseCase.IsMax(20))
	assert.False(t,mathUseCase.IsMax(-100))
	mathUseCase.Reset()
	assert.True(t,mathUseCase.IsMax(-100))
}
