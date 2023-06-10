package model

import (
	"testing"
)

func TestRegStringToStates(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			"aaaから正しく[]Stateを生成可能",
			"aaa",
		},
		{
			"hogehogeから正しく[]Stateを生成可能",
			"hogehoge",
		},
		{
			"空文字から正しく[]Stateを生成可能",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := RegString{Content: tt.input}
			states, output, err := rs.ToStates("hoge")
			if err != nil {
				t.Error("error found after ToStates\n")
			}
			if output == "" {
				t.Errorf("output was empty\n")
			}
			if len(states) != len(tt.input) {
				t.Errorf("states length is expected to be%v, but got %v\n", len(states), len(tt.input))
			}

			// idの整合性確認
			headedTo := "hoge"
			for _, state := range states {
				if headedTo != state.Id {
					t.Errorf("destination id should be %v, but state's id was %v\n", headedTo, state.Id)
				}
				if len(state.Moves) != 1 {
					t.Errorf("state.Moves length should be 1, but got %v\n", len(state.Moves))
				}
				headedTo = state.Moves[0].MoveTo
			}
			if headedTo != output {
				t.Errorf("wrong destination given as return value of ToStates. %v != %v", output, headedTo)
			}

			// Movesのinput確認
			for i, c := range tt.input {
				if states[i].Moves[0].Input != string(c) {
					t.Errorf("at %v state, input should be %v, but got %v\n", i, c, states[i].Moves[0].Input)
				}
			}
		})
	}
}
