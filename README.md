# ultraSimpleStress

Now that go 1.16 supports M1 natively, I thought it might be interesting to write a dirt simple stress tester to see how performance on M1 compares to various x86_64 based machines running Linux.

The benchmark uses the flags package, so it understands -h.  The usage declaration is provided below.

```
Patricks-MacBook-Air:ultraSimpleStress patrickheckenlively$ ./ultraSimpleStress -h
Usage of ./ultraSimpleStress:
  -doFloatTest
        Run the floating point stress test.
  -doShaTest
        Run the sha256 stress test.
  -iterations int
        Number of iterations of the stress test to run. (default 100000)
  -threads int
        The number of simultaneous stress tests to run. (default 8)
```

These options are described in more detail below.

## -doFloatTest

This option causes the benchmark to run a bunch of floating point division operations.

## -doShaTest

This option causes the benchmark to run a bunch of sha256 checksum operations.

## -iterations

This is how many iterations of the selected benchmark to run in each thread.

## -threads

This is the number of go routines to spawn while running the benchmark.

# Real World Usage

I recommend using this small benchmark with the UNIX time command.
The example below shows how this is useful.

```
Patricks-MacBook-Air:ultraSimpleStress patrickheckenlively$ time ./ultraSimpleStress -doShaTest -iterations 50000 -threads 8
{true}{true}{true}{true}{true}{true}{true}{true}

real    0m11.766s
user    0m34.714s
sys     0m3.931s
```

# Sample Real World Results

The table below provides some real world results from some of the systems in my house.

| Processor | OS | Arguments | Real | User | Sys |
|:----------|:---|:----------|:-----|:-----|:----|
| M1 | MacOS | -doShaTest -iterations 50000 -threads 8 | 0m11.766s | 0m34.714s | 0m3.931s |
| AMD Ryzen 7 3700X 8-Core Processor | Linux | -doShaTest -iterations 50000 -threads 8 | 0m42.945s | 3m58.577s | 0m11.129s |
| Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz | Linux | -doShaTest -iterations 50000 -threads 8 | 0m49.934s | 3m50.659s | 0m6.199s |
| M1 | MacOS | -doFloatTest -iterations 10000000000 -threads 8 | 0m4.491s | 0m33.660s | 0m0.093s |
| AMD Ryzen 7 3700X 8-Core Processor | Linux | -doFloatTest -iterations 10000000000 -threads 8 | 0m2.390s | 0m19.022s | 0m0.000s |
| Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz | Linux | -doFloatTest -iterations 10000000000 -threads 8 | 0m5.289s | 0m41.094s | 0m0.048s |

The M1 processor's sha256 performance, is IMO pretty surprising.  The
x86_64 processor family has a native sha256 instruction.  Also, I have
read that the x86_64 implementation of the Go crypto library contains
lots of assembly optimizations.  Based on the numbers above, I
speculate that perhaps the M1 implementation (of at least sha256) uses
the GPU built into the M1 processor.