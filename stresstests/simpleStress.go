package stresstests

// FloatTest ...Float test struct
type FloatTest struct {
	Done       chan StressComm
	Iterations int
}

// GetFloatTest ...Get a FloatTest object.
func GetFloatTest(iterations int, comm chan StressComm) FloatTest {

	rv := new(FloatTest)
	rv.Done = comm
	rv.Iterations = iterations

	return *rv
}

// RunTest ...Run the stress test.
func (f FloatTest) RunTest() {

	for i := 0; i < f.Iterations; i++ {
		foo := float64(i) / float64((i + 1))
		foo += 0
	}

	rv := new(StressComm)
	rv.Status = true

	f.Done <- *rv
}
