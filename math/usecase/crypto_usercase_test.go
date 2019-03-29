package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRSA(t *testing.T) {
	var rsa Rsa
	key,err:=rsa.GenerateKey()
	assert.NoError(t,err)
	assert.NotEmpty(t,key.PublicKey)
	assert.NotEmpty(t,key.PrivateKey)
	sign,err:=rsa.Sign("test message")
	assert.NoError(t,err)
	assert.Equal(t,true,rsa.Verify("test message",sign))
}
func TestED25519(t *testing.T) {
	var ed25519 Ed25519
	key,err:=ed25519.GenerateKey()
	assert.NoError(t,err)
	assert.NotEmpty(t,key.PublicKey)
	assert.NotEmpty(t,key.PrivateKey)
	sign,err:=ed25519.Sign("test message")
	assert.NoError(t,err)
	assert.Equal(t,true,ed25519.Verify("test message",sign))
}