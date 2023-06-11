package model

import (
	"github.com/google/uuid"
)

type RegString struct {
	RegExp
	Content string
}

func (rs RegString) ToStates(startId string) ([]State, string, error) {
	states := []State{}

	var currentId, nextId string
	currentId = startId
	for _, c := range rs.Content {
		rawId, err := uuid.NewRandom()
		if err != nil {
			return nil, "", err
		}
		nextId = rawId.String()
		moves := []Move{
			{
				IsEpsilon: false,
				Input:     string(c),
				MoveTo:    nextId,
			},
		}
		states = append(states, State{
			Id:    currentId,
			Moves: moves,
			IsEnd: false,
		})
		currentId = nextId
	}

	return states, currentId, nil
}
