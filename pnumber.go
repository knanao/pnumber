package main

import (
	"context"
	"math/rand"
	"runtime"
	"sync"
)

func Perm(n int, m int64) []int64 {
	ctx := context.Background()
	defer ctx.Done()

	doFunc := func() interface{} { return rand.Int63n(m) }
	randCh := toInt64(ctx, generateFn(ctx, doFunc))
	fs := make([]<-chan interface{}, runtime.NumCPU())
	for i := 0; i < len(fs); i++ {
		fs[i] = primeFinder(ctx, randCh)
	}

	retCh := do(ctx, fanIn(ctx, fs...), n)
	ret := make([]int64, 0, n)
	for i := range retCh {
		ret = append(ret, i.(int64))
	}
	return ret
}

func toInt64(ctx context.Context, valCh <-chan interface{}) <-chan int64 {
	ret := make(chan int64)
	go func() {
		defer close(ret)
		for v := range valCh {
			select {
			case <-ctx.Done():
				return
			case ret <- v.(int64):
			}
		}
	}()
	return ret
}

func generateFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ret := make(chan interface{})
	go func() {
		defer close(ret)
		for {
			select {
			case <-ctx.Done():
				return
			case ret <- fn():
			}
		}
	}()
	return ret
}

func primeFinder(ctx context.Context, valCh <-chan int64) <-chan interface{} {
	ret := make(chan interface{})
	go func() {
		defer close(ret)
		for i := range valCh {
			select {
			case <-ctx.Done():
				return
			default:
			}
			if isPrime(i) {
				ret <- i
			}
		}
	}()
	return ret
}

func isPrime(n int64) bool {
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func do(ctx context.Context, valCh <-chan interface{}, n int) <-chan interface{} {
	ret := make(chan interface{})
	go func() {
		defer close(ret)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case ret <- <-valCh:
			}
		}
	}()
	return ret
}

func fanIn(ctx context.Context, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	ret := make(chan interface{})
	mux := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case ret <- i:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go mux(c)
	}
	go func() {
		wg.Wait()
		close(ret)
	}()
	return ret
}
