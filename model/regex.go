package model

type RegExp interface {
	// ToStates creates state list from RegExp and returns it.
	//
	// startId specifies entry node's Id.
	// 2nd return value is state id which the sequence of states' output is headed to.
	ToStates(startId string) ([]State, string, error)
}
