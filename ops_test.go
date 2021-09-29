package truthy_test

import (
	"testing"

	"github.com/carlmjohnson/truthy"
)

func TestOps(t *testing.T) {
	for name, tc := range map[string]struct {
		F   func(int, int) bool
		ops [4][3]int
	}{
		"and": {truthy.And[int, int], [4][3]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 1},
		}},
		"or": {truthy.Or[int, int], [4][3]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}},
		"xor": {truthy.Xor[int, int], [4][3]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 0},
		}},
		"nor": {truthy.Nor[int, int], [4][3]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 0},
		}},
		"nand": {truthy.Nand[int, int], [4][3]int{
			{0, 0, 1},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 0},
		}},
		"xnor": {truthy.Xnor[int, int], [4][3]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 1},
		}},
	} {
		t.Run(name, func(t *testing.T) {
			for _, op := range tc.ops {
				a := op[0]
				b := op[1]
				r := op[2] == 1
				if tc.F(a, b) != r {
					t.Fatalf("%d %d", a, b)
				}
			}
		})
	}
}
