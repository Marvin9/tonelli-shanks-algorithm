package algo_test

import (
	"testing"

	"github.com/Marvin9/tonelli-shanks-algorithm/algo"
)

type input struct {
	n          int64
	p          int64
	expectedOp []int64
}

func contains(slc []int64, foo interface{}) bool {
	for _, val := range slc {
		if val == foo {
			return true
		}
	}
	return false
}

func expectedErr(expected []int64, got interface{}, t *testing.T) {
	if !contains(expected, got) {
		t.Errorf("Expected one of %v, got %v\n", expected, got)
	}
}

func TestTonelliShanks(t *testing.T) {
	tests := []input{
		input{
			n:          5,
			p:          41,
			expectedOp: []int64{28},
		},
		input{
			n:          2,
			p:          113,
			expectedOp: []int64{62},
		},
		input{
			n:          2,
			p:          7,
			expectedOp: []int64{3, 4},
		},
		input{
			n:          2,
			p:          5,
			expectedOp: []int64{-1},
		},
		input{
			n:          29,
			p:          53,
			expectedOp: []int64{33, 20},
		},
		input{
			n:          37,
			p:          137,
			expectedOp: []int64{41, 96},
		},
		input{
			n:          11,
			p:          257,
			expectedOp: []int64{36, 221},
		},
		input{
			n:          92,
			p:          101,
			expectedOp: []int64{71, 30},
		},
		input{
			n:          92,
			p:          193,
			expectedOp: []int64{51, 142},
		},
		input{
			n:          6,
			p:          37,
			expectedOp: []int64{-1},
		},
		input{
			n:          21,
			p:          37,
			expectedOp: []int64{13, 24},
		},
		input{
			n:          25,
			p:          127,
			expectedOp: []int64{5, 122},
		},
		input{
			n:          35,
			p:          127,
			expectedOp: []int64{17, 110},
		},
		input{
			n:          6,
			p:          37,
			expectedOp: []int64{-1},
		},
		input{
			n:          21,
			p:          37,
			expectedOp: []int64{13, 24},
		},
	}

	for _, test := range tests {
		op := algo.TonelliShanks(test.n, test.p)
		expectedErr(test.expectedOp, op, t)
	}
}
