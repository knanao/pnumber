# pnumber
generate random prime numbers

# Benchmark
## n=10
```
BenchmarkPerm-12    	       3	 393436424 ns/op	   14477 B/op	     252 allocs/op
PASS
ok  	github.com/knanao/pnumber	1.866s
```

## n=100
```
BenchmarkPerm100-12    	       1	2838983383 ns/op	   36536 B/op	    2567 allocs/op
PASS
ok  	github.com/knanao/pnumber	5.974s
```

## n=1000
```
BenchmarkPerm1000-12    	       1	23896238763 ns/op	  175504 B/op	   18747 allocs/op
ok  	github.com/knanao/pnumber	28.916s
```
