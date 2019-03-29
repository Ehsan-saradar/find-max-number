package benchmarks

import (
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
	"testing"
)
const (
	message = "Test message"
)
func BenchmarkGenerateEd25519Key(b *testing.B) {
	var ed25519 usecase.Ed25519
	ed25519.GenerateKey()
}

func BenchmarkSignEd25519(b *testing.B) {
	var ed25519 usecase.Ed25519
	ed25519.GenerateKey()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ed25519.Sign(message)
	}
}

func BenchmarkVerifyEd25519(b *testing.B) {
	var ed25519 usecase.Ed25519
	ed25519.GenerateKey()
	signature,_:=ed25519.Sign(message)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ed25519.Verify(message,signature)
	}
}