package model

import "github.com/google/uuid"

type RegExp interface {
	// ToStates creates state list from RegExp and returns it.
	//
	// startId specifies entry node's Id.
	// 2nd return value is state id which the sequence of states' output is headed to.
	ToStates(startId string) ([]State, string, error)
}

// CreateCompleteStates creates goal-included slice of States.
//
// ToStates creates slice of States, but it doesn't include end state.
// This function takes [re] as argument and creates goal-included slice of States with it.
func CreateCompleteStates(re RegExp) ([]State, string, error) {
	startUuid, err := uuid.NewRandom()
	if err != nil {
		panic("uuid generation error")
	}

	states := []State{}
	rawStates, outId, err := re.ToStates(startUuid.String())
	states = append(states, rawStates...)

	endState := State{
		Id: outId,
		IsEnd: true,
		Moves: []Move{},
	}
	states = append(states, endState)

	return states, startUuid.String(), nil
}
