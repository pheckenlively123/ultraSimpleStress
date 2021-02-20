package commandline

import "flag"

// CmdOpts - All of the options provided from the command line.
type CmdOpts struct {
	Iterations  int
	Threads     int
	DoShaTest   bool
	DoFloatTest bool
}

// GetOpts - Return the command line arguments in a CmdOpts struct
func GetOpts() CmdOpts {

	rv := new(CmdOpts)

	iterations := flag.Int("iterations", 100000, "Number of iterations of the stress test to run.")
	threads := flag.Int("threads", 8, "The number of simultaneous stress tests to run.")
	doShaTest := flag.Bool("doShaTest", false, "Run the sha256 stress test.")
	doFloatTest := flag.Bool("doFloatTest", false, "Run the floating point stress test.")

	flag.Parse()

	if !*doShaTest && !*doFloatTest {
		flag.PrintDefaults()
		panic("-doShaTest and -doFloatTest must not both be false.")
	}

	if *doShaTest && *doFloatTest {
		flag.PrintDefaults()
		panic("-doShaTest and -doFloatTest are mutually exclusive.")
	}

	if *iterations <= 0 {
		flag.PrintDefaults()
		panic("-iterations must be a positive integer.")
	}

	if *threads <= 0 {
		flag.PrintDefaults()
		panic("-threads must be a positive integer")
	}

	rv.Iterations = *iterations
	rv.Threads = *threads
	rv.DoShaTest = *doShaTest
	rv.DoFloatTest = *doFloatTest

	return *rv
}
