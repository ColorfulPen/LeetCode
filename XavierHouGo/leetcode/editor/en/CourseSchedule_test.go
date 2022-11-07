package cn

import (
	"testing"
)

func TestCourseSchedule(t *testing.T) {
	t.Log(canFinish(3, [][]int{{1, 0}, {1, 2}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		indeg = make([]int, numCourses)

		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return len(result) == numCourses
}

//leetcode submit region end(Prohibit modification and deletion)
