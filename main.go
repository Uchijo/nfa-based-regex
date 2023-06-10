package main

import (
	"fmt"

	"github.com/uchijo/nfa-based-regex/model"
)

func main() {
	input := "ab"
	result := search(input, "1", true)
	fmt.Printf("result: %v\n", result)
}

type log struct {
	id     string
	buffer string
	moveTo string
}

func (l log) alike(log log) bool {
	return l.buffer == log.buffer && l.id == log.id && l.moveTo == log.moveTo
}

var tried = []log{}

func search(input, currentStateId string, first bool) bool {
	// fmt.Printf("current buffer: %v\n", input)
	// fmt.Printf("current stateId: %v\n", currentStateId)

	var currentState model.State
	for _, s := range states {
		if s.Id == currentStateId {
			currentState = s
			break
		}
	}
	if currentState.Id == "" {
		panic("no such state")
	}

	// 入力なしでゴール状態なら終わり
	// 入力なしでゴール状態じゃなかったら終わり
	if input == "" {
		return currentState.IsEnd
	}

	// 行ける可能性があるところをリストアップ
	for _, move := range currentState.Moves {
		if move.IsEpsilon || move.Input == input[:1] {
			if first {
				fmt.Println(" === ")
				fmt.Printf("trying move: %v\n", move.MoveTo)
				fmt.Printf("is epsilon : %v\n", move.IsEpsilon)
			}

			result := false
			if move.IsEpsilon {
				nextLog := log{currentStateId, input, move.MoveTo}
				for _, l := range tried {
					if l.alike(nextLog) {
						return false
					}
				}
				tried = append(tried, log{currentStateId, input, move.MoveTo})
				result = search(input, move.MoveTo, false)
			} else {
				nextLog := log{currentStateId, input, move.MoveTo}
				for _, l := range tried {
					if l.alike(nextLog) {
						return false
					}
				}
				tried = append(tried, log{currentStateId, input, move.MoveTo})
				result = search(input[1:], move.MoveTo, false)
			}
			if result {
				return true
			}
			continue
		}
	}
	return false
}

var states = []model.State{
	{
		Id: "1",
		Moves: []model.Move{
			{
				IsEpsilon: true,
				MoveTo:    "2",
			},
			{
				IsEpsilon: true,
				MoveTo:    "3",
			},
			{
				IsEpsilon: false,
				MoveTo:    "6",
				Input:     "a",
			},
		},
		IsEnd: false,
	},
	{
		Id: "2",
		Moves: []model.Move{
			{
				IsEpsilon: false,
				MoveTo:    "2",
				Input:     "a",
			},
			{
				IsEpsilon: true,
				MoveTo:    "5",
			},
		},
		IsEnd: false,
	},
	{
		Id: "3",
		Moves: []model.Move{
			{
				IsEpsilon: true,
				MoveTo:    "5",
			},
			{
				IsEpsilon: false,
				MoveTo:    "4",
				Input:     "a",
			},
		},
		IsEnd: false,
	},
	{
		Id: "4",
		Moves: []model.Move{
			{
				IsEpsilon: false,
				MoveTo:    "3",
				Input:     "b",
			},
		},
		IsEnd: false,
	},
	{
		Id: "5",
		Moves: []model.Move{
			{
				IsEpsilon: true,
				MoveTo:    "1",
			},
		},
		IsEnd: false,
	},
	{
		Id: "6",
		Moves: []model.Move{
			{
				IsEpsilon: false,
				MoveTo:    "7",
				Input:     "a",
			},
		},
		IsEnd: false,
	},
	{
		Id: "7",
		Moves: []model.Move{
			{
				IsEpsilon: false,
				MoveTo:    "8",
				Input:     "a",
			},
		},
		IsEnd: false,
	},
	{
		Id:    "8",
		Moves: []model.Move{},
		IsEnd: true,
	},
}
