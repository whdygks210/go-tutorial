package the_bomberman_game

import (
	"fmt"
	"testing"
)

type testCase struct {
	n      int32
	grid   []string
	result []string
}

func Test_bomberMan(t *testing.T) {
	testCaseArr := []testCase{
		{
			n: 3,
			grid: []string{
				".......",
				"...O...",
				"....O..",
				".......",
				"OO.....",
				"OO.....",
			},
			result: []string{
				"OOO.OOO",
				"OO...OO",
				"OOO...O",
				"..OO.OO",
				"...OOOO",
				"...OOOO",
			},
		},
		{
			n: 5,
			grid: []string{
				".......",
				"...O.O.",
				"....O..",
				"..O....",
				"OO...OO",
				"OO.O...",
			},
			result: []string{
				".......",
				"...O.O.",
				"...OO..",
				"..OOOO.",
				"OOOOOOO",
				"OOOOOOO",
			},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", bomberMan(tc.n, tc.grid))

		if res == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}