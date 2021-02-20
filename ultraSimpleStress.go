package main

import (
	"fmt"
	"ultraSimpleStress/commandline"
	"ultraSimpleStress/stresstests"
)

func main() {

	opts := commandline.GetOpts()

	comm := make(chan stresstests.StressComm)

	for i := 0; i < opts.Threads; i++ {

		if opts.DoShaTest {
			myStress := stresstests.GetShaTest(opts.Iterations, comm)
			go myStress.RunTest()
		}

		if opts.DoFloatTest {
			myStress := stresstests.GetFloatTest(opts.Iterations, comm)
			go myStress.RunTest()
		}
	}

	for i := 0; i < opts.Threads; i++ {
		done := <-comm

		fmt.Print(done)
	}

	fmt.Print("\n")
}
