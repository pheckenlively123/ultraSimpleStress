package stresstests

// Stresser ...This is the common interface implemented by the stressor objects.  (Right now,
// this is not buying me anything, so consider removing it.)
type Stresser interface {
	RunTest()
}

// StressComm ...This is what the RunTest method returns via the channel in it's struct.
type StressComm struct {
	Status bool
}
