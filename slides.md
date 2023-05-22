---
title: PPROF
author: Andy Blair
styles:
  table:
      column_spacing: 3
extensions:
  - terminal
  - file_loader
---
# The PPROF is in the pudding
_Performance testing of go applications_


# What we'll cover 
_if there's time_

### Basic performance testing of a function
### How to micro optimise on a function by function basis
### How to monitor performance, and debug production performance

 
# What this talk isn't
 _I am not an expert_

 This is a **Beginners** guide, it's not going to cover everything

 
# Lets get started

 To start off we're going to compare the performance of two implementations of fibinanci

## Recursive
## Iterative

 And see which is better!

# What we're testing
Incase you don't know how to implement fibinanci from memory
## Recursive
```file
path: fib/runner.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 2
    end: 7
```

## Iterative
```file
path: fib/runner.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 9
    end: 18
```

# Okay, what now? 
As good devs, we always test our stuff! Right?

<!-- stop -->

So here's what the test looks like

```file
path: fib/runner_test.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 7
    end: 27
```

# And it runs? 

```terminal-ex
command: bash
rows: 30
init_text: cd fib/; go test -v ./runner_test.go 
init_wait: '$ '
init_codeblock_lang: bash
```

# The good stuff
Now to actually benchmark this!

We'll simply copy and update the test, to match the naming convention `BenchmarkXXX` and use testing.B instead of testing.T

<!-- stop -->

Tada!
```file
path: fib/runner_test.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 45
    end: 67
```

# Now lets run it!
```terminal-ex
command: bash
rows: 30
init_text: cd fib/; go test -v ./runner_test.go -bench=.
init_wait: '$ '
init_codeblock_lang: bash
```

# Using PPROF to optimize functions
Benchmarking is a great way of analyising code execution

This is commonly refered to as micro-optimisation: taking one part of your code and optimising it as much as possible

This can be a good way to spot common pitfalls, and useful for high throughput/volume code paths

# A very unrepresentitive view

```file
path: load/runner.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 12
    end: 18
```

```file
path: load/runner.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 37
    end: 54
```

# Some added feature of pprof
PPROF doesn't just run the benchmark tests, it can also

### Generate CPU and memory profiles of the execution
### Configure how long to run the benchmark for
### go test can run multiple times, and give you an average run information

# Lets go!
Using these three points, we can run a 5 passes of a benchmark test, each test lasting 5 seconds

We'll save the memory and CPU profiles

```terminal-ex
command: bash
rows: 30
init_text: cd load/; go test ./... -bench=BenchmarkV1 -cpuprofile cpu.pprof -memprofile=mem.pprof -benchtime=5s -count=5 > v1.out
init_wait: '$ '
init_codeblock_lang: bash
```

# Lets take a look! 
Go has ways to let you run certain tools, we can run `go tool pprof` to access pprof, giving it mem.pprof or cpu.pprof


```terminal-ex
command: bash
rows: 30
init_text: cd load/
init_wait: '$ '
init_codeblock_lang: bash
```

# We will rebuild
Lets make this better

```file
path: load/runner.go
lang: go
transform: sed 's/\t/    /g'
lines:
    start: 55
    end: null
```

# Re-run

```terminal-ex
command: bash
rows: 30
init_text: cd load/; go test ./... -bench=BenchmarkV2 -cpuprofile cpu2.pprof -memprofile=mem2.pprof -benchtime=5s -count=5 > v2.out
init_wait: '$ '
init_codeblock_lang: bash
```

# Compare and contrast
Go also provide a tool at `golang.org/x/perf/cmd/benchstat` 
We can use that to compare the output of the two benchmarks
