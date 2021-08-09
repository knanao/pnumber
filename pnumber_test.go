package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerm(t *testing.T) {
	const max = 50000000
	v := Perm(10, max)
	require.Equal(t, 10, len(v))

	p := v[0]
	require.True(t, isPrime(p))

	v = Perm(100, max)
	require.Equal(t, 100, len(v))
}

func BenchmarkPerm10(b *testing.B) {
	const max = 50000000
	for i := 0; i < b.N; i++ {
		_ = Perm(10, max)
	}
}

func BenchmarkPerm100x(b *testing.B) {
	const max = 50000000
	for i := 0; i < b.N; i++ {
		_ = Perm(100, max)
	}
}

func BenchmarkPerm1000(b *testing.B) {
	const max = 50000000
	for i := 0; i < b.N; i++ {
		_ = Perm(1000, max)
	}
}
