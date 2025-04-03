package top100

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make(map[int]int)
	adjList := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
		adjList[i] = make([]int, 0)
	}

	for _, p := range prerequisites {
		inDegree[p[0]]++
		adjList[p[1]] = append(adjList[p[1]], p[0])
	}

	queue := make([]int, 0)
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	cnt := 0
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		cnt++
		for _, v := range adjList[c] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return cnt == numCourses
}
