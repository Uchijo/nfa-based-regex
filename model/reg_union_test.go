package model

import "testing"

func TestRegUnion(t *testing.T) {
	tests := []struct {
		name  string
		input RegUnion
	}{
		{
			name: "a, bのunionを正しく生成可能",
			input: RegUnion{
				Left:  RegString{Content: "a"},
				Right: RegString{Content: "b"},
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			states, out, err := v.input.ToStates("hoge")

			if out == "" {
				t.Errorf("output id was empty")
			}

			if err != nil {
				t.Errorf("something went wrong; error reported.\n")
			}

			l, _, _ := v.input.Left.ToStates("hoge")
			r, _, _ := v.input.Left.ToStates("hoge")
			expectedStLen := len(l) + len(r) + 3
			if expectedStLen != len(states) {
				t.Errorf("expected to get %v states, but got %v\n", expectedStLen, len(states))
			}
		})
	}
}
