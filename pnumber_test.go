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
}

func BenchmarkPerm(b *testing.B) {
	const max = 50000000
	for i := 0; i < b.N; i++ {
		_ = Perm(10, max)
	}
}
