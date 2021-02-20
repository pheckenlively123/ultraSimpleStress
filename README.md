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
