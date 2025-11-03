package graph

import "testing"

func Test_entryPoint(t *testing.T) {
	testMap := map[int][]int{
		0: {1, 3},
		1: {2},
		2: {3},
		3: {4},
		4: {1, 2},
	}

	entryPoint(testMap)
}
