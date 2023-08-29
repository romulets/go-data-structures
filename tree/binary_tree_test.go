package tree_test

import (
	"github.com/romulets/go-data-structures/tree"
	"testing"
)

func TestVal(t *testing.T) {
	testCases := []struct {
		desc  string
		input interface{}
	}{
		{
			desc:  "Integer value",
			input: 1,
		},
		{
			desc:  "String value",
			input: "abc",
		},
		{
			desc: "Struct value",
			input: struct {
				a string
				b int
			}{
				a: "a",
				b: 1,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			n := tree.NewBinaryNode(c.input)
			if n.Value() != c.input {
				t.Errorf("Expected %s but got %s", c.input, n.Value())
			}

			if n.Left() != nil {
				t.Error("Expected nil")
			}

			if n.Right() != nil {
				t.Error("Expected nil")
			}
		})
	}
}

func TestPutLeft(t *testing.T) {
	n := tree.NewBinaryNode(2)
	n.PutLeft(1)

	if n.Left() == nil {
		t.Error("Expected 1 but got nil")
	}

	if n.Right() != nil {
		t.Error("Expected nil")
	}

	if n.Left().Value() != 1 {
		t.Errorf("Expected 1 but got %d", n.Left().Value())
	}

	if n.Left().Left() != nil {
		t.Error("Expected nil")
	}

	if n.Left().Right() != nil {
		t.Error("Expected nil")
	}
}

func TestPutRight(t *testing.T) {
	n := tree.NewBinaryNode(2)
	n.PutRight(1)

	if n.Right() == nil {
		t.Error("Expected 1 but got nil")
	}

	if n.Left() != nil {
		t.Error("Expected nil")
	}

	if n.Right().Value() != 1 {
		t.Errorf("Expected 1 but got %d", n.Left().Value())
	}

	if n.Right().Left() != nil {
		t.Error("Expected nil")
	}

	if n.Right().Right() != nil {
		t.Error("Expected nil")
	}
}
