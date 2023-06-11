package model

import "github.com/google/uuid"

type RegStar struct {
	RegExp
	Content RegExp
}

func (rs RegStar) ToStates(startId string) ([]State, string, error) {
	nextId, err := uuid.NewRandom()
	if err != nil {
		panic("uuid generation error.")
	}
	contentStart, err := uuid.NewRandom()
	if err != nil {
		panic("uuid generation error.")
	}
	cs, cout, err := rs.Content.ToStates(contentStart.String())
	if err != nil {
		return nil, "", err
	}
	entry := State{
		Id: startId,
		IsEnd: false,
		Moves: []Move{
			{
				IsEpsilon: true,
				Input: "",
				MoveTo: contentStart.String(),
			},
			{
				IsEpsilon: true,
				Input: "",
				MoveTo: cout,
			},
		},
	}
	root := State{
		Id:    cout,
		IsEnd: false,
		Moves: []Move{
			{
				IsEpsilon: true,
				Input: "",
				MoveTo: contentStart.String(),
			},
			{
				IsEpsilon: true,
				Input: "",
				MoveTo: nextId.String(),
			},
		},
	}
	states := []State{}
	states = append(states, cs...)
	states = append(states, entry, root)
	return states, nextId.String(), nil
}
