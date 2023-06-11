package model

import (
	"github.com/google/uuid"
)

type RegUnion struct {
	RegExp
	Left  RegExp
	Right RegExp
}

func (ru RegUnion) ToStates(startId string) ([]State, string, error) {
	if ru.Left == nil || ru.Right == nil {
		panic("invalid regex.")
	}
	// Left, Right毎にoutputのidが指定されるので、イプシロン動作でいい感じに辻褄合わせる
	states := []State{}
	leftStart, _ := uuid.NewRandom()
	rightStart, _ := uuid.NewRandom()
	start := State{
		Id:    startId,
		IsEnd: false,
		Moves: []Move{
			{
				IsEpsilon: true,
				MoveTo:    leftStart.String(),
				Input:     "",
			},
			{
				IsEpsilon: true,
				MoveTo:    rightStart.String(),
				Input:     "",
			},
		},
	}
	states = append(states, start)

	// TODO: エラー処理ちゃんとする
	leftStates, leftGoal, _ := ru.Left.ToStates(leftStart.String())
	states = append(states, leftStates...)
	rightStates, rightGoal, _ := ru.Right.ToStates(rightStart.String())
	states = append(states, rightStates...)

	goalUUID, _ := uuid.NewRandom()

	leftGoalState := State{
		Id:    leftGoal,
		IsEnd: false,
		Moves: []Move{
			{
				IsEpsilon: true,
				Input:     "",
				MoveTo:    goalUUID.String(),
			},
		},
	}
	rightGoalState := State{
		Id: rightGoal,
		IsEnd: false,
		Moves: []Move{
			{
				IsEpsilon: true,
				Input: "",
				MoveTo: leftGoal,
			},
		},
	}
	states = append(states, leftGoalState, rightGoalState)

	return states, goalUUID.String(), nil
}
