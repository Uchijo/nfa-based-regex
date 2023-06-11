package main

import (
	"fmt"

	"github.com/uchijo/nfa-based-regex/model"
)

func main() {
	hoge := model.RegApp{
		Contents: []model.RegExp{
			model.RegString{Content: "a"},
			model.RegString{Content: "b"},
			model.RegString{Content: "c"},
		},
	}

	states, out, err := hoge.ToStates("hoge")
	if err != nil {
		panic("something went wrong")
	}

	endState := model.State{
		Id: out,
		Moves: []model.Move{},
		IsEnd: true,
	}
	states = append(states, endState)
	for i, v := range states {
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Println("")

	input := "abc"
	result := search(input, "hoge", true, states)
	fmt.Printf("result: %v\n", result)

	// input := "aaa"
	// result := search(input, "1", true, states)
	// fmt.Printf("result: %v\n", result)
}

type log struct {
	id     string
	buffer string
	moveTo string
}

func (l log) alike(log log) bool {
	return l.buffer == log.buffer && l.id == log.id && l.moveTo == log.moveTo
}

// 既に訪れた状態を記憶して無限ループを防ぐ
var tried = []log{}

func search(input, currentStateId string, first bool, states []model.State) bool {
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
	// 入力なしでゴール状態じゃなかったらイプシロンづたいにゴールにたどり着けるか確認
	if input == "" {
		return searchEpsilonGoal(currentStateId, states, false)
	}

	// 行ける可能性があるところをリストアップ
	for _, move := range currentState.Moves {
		if move.IsEpsilon || move.Input == input[:1] {
			// if true {
			// 	fmt.Printf("trying move: %v\n", move.MoveTo)
			// 	fmt.Printf("is epsilon : %v\n", move.IsEpsilon)
			// 	fmt.Println(" === ")
			// }

			result := false
			if move.IsEpsilon {
				nextLog := log{currentStateId, input, move.MoveTo}
				for _, l := range tried {
					if l.alike(nextLog) {
						return false
					}
				}
				tried = append(tried, log{currentStateId, input, move.MoveTo})
				result = search(input, move.MoveTo, false, states)
			} else {
				nextLog := log{currentStateId, input, move.MoveTo}
				for _, l := range tried {
					if l.alike(nextLog) {
						return false
					}
				}
				tried = append(tried, log{currentStateId, input, move.MoveTo})
				result = search(input[1:], move.MoveTo, false, states)
			}
			if result {
				return true
			}
			continue
		}
	}
	return false
}

// 既に訪れた状態を記憶して無限ループを防ぐ
var triedEpsilon = []log{}

func searchEpsilonGoal(currentId string, states []model.State, recursiveCall bool) bool {
	// ログ初期化
	if !recursiveCall {
		triedEpsilon = []log{}
	}

	// 現在の状態がゴールならtrueで終わり
	var currentState model.State
	for _, v := range states {
		if v.Id == currentId {
			currentState = v
			break
		}
	}
	if currentState.IsEnd {
		return true
	}

	// ゴールじゃないならIsEpsilonなmoveを探索
	for _, v := range currentState.Moves {
		if !v.IsEpsilon {
			continue
		}

		currentLog := log{id: currentId, buffer: "", moveTo: v.MoveTo}
		for _, vv := range triedEpsilon {
			if vv.alike(currentLog) {
				return false
			}
		}
		triedEpsilon = append(triedEpsilon, currentLog)
		hasGoalAhead := searchEpsilonGoal(v.MoveTo, states, true)
		if hasGoalAhead {
			return true
		} else {
			// 次のループへ
		}
	}

	// 上に該当しないならこの状態からはゴールに辿り着けない
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
