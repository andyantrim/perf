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
