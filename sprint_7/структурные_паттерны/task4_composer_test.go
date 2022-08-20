//go:build ignore
// +build ignore

package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	root := &Oper{
		Type: Div,
		Left: &Oper{
			Type: Mul,
			Left: &Oper{
				Type:  Add,
				Left:  &Number{Value: 2},
				Right: &Number{Value: 3},
			},
			Right: &Oper{
				Type:  Sub,
				Left:  &Number{Value: 77},
				Right: &Number{Value: 55},
			},
		},
		Right: &Number{Value: 2},
	}
	if root.Calculate() != 55 {
		t.Errorf(`get %d want %d`, root.Calculate(), 77)
	}
}
