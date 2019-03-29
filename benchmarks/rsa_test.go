package benchmarks

import (
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
	"testing"
)

func BenchmarkGenerateRsaKey(b *testing.B) {
	var rsa usecase.Rsa
	rsa.GenerateKey()
}

func BenchmarkSignRsa(b *testing.B) {
	var rsa usecase.Rsa
	rsa.GenerateKey()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rsa.Sign(message)
	}
}
func BenchmarkVerifyRsa(b *testing.B) {
	var rsa usecase.Rsa
	rsa.GenerateKey()
	signature,_:=rsa.Sign(message)
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rsa.Verify(message,signature)
	}
}