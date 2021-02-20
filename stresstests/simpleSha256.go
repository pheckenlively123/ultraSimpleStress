package stresstests

import (
	"crypto/sha256"
	"fmt"
)

type ShaTest struct {
	Done       chan StressComm
	Iterations int
}

func doSha256(inString string) (shaStr string) {
	h := sha256.New()
	h.Write([]byte(inString))
	rv := fmt.Sprintf("%x\n", h.Sum(nil))

	return rv
}

func GetShaTest(iterations int, comm chan StressComm) ShaTest {

	rv := new(ShaTest)
	rv.Done = comm
	rv.Iterations = iterations

	return *rv
}

// RunTest ...This runs the test for this stressor.
func (s ShaTest) RunTest() {

	foo := "something wonderful"

	for i := 0; i < s.Iterations; i++ {
		sum := doSha256(foo)
		sum += ""
		foo += fmt.Sprint(":%d", i)
	}

	rv := new(StressComm)
	rv.Status = true

	s.Done <- *rv
}
