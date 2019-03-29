package usecase

import (
	"github.com/Ehsan-saradar/find-max-number/math/repository"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestSimpleInputUseCase(t *testing.T) {
	var err error
	numbers:=[]int{10,20,15,30,25}
	mathRepo,err := repository.ReadFileRepository(strings.NewReader(strings.Trim(strings.Join(strings.Split(fmt.Sprint(numbers), " "), "\r\n"), "[]")))
	assert.NoError(t,err)
	inputUseCase:=NewSimpleInputUseCase(mathRepo)
	for i:=0;i<len(numbers);i++{
		number,err:=inputUseCase.GetNextNumber()
		assert.Equal(t,numbers[i],number)
		assert.NoError(t,err)
	}
	_,err=inputUseCase.GetNextNumber()
	assert.Equal(t,io.EOF,err)
}
