package repository

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestNewMathFileRepository(t *testing.T) {
	var b bytes.Buffer
	assert.Equal(t,nil,generateRandom(&b))
	t.Run("Test newMathFileRepository", func(t *testing.T) {
		var mathRepo MathRepository
		var err error
		var cnt int
		mathRepo,err = ReadFileRepository(&b)
		assert.NoError(t,err)
		for cnt=0;err==nil;cnt++{
			_,err=mathRepo.GetNextNumber()
		}
		assert.Equal(t,io.EOF,err)
		assert.Equal(t,RandomNumbers+1,cnt)
	})
}

func TestGenerateRandom(t *testing.T) {
	var b bytes.Buffer
	assert.NoError(t,generateRandom(&b))
	assert.Equal(t,RandomNumbers, len(strings.Split(b.String(), "\r\n")))
}
