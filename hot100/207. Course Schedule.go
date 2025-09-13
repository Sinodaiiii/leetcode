package hot100

func canFinish(numCourses int, prerequisites [][]int) bool {
	postCourses := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		postCourses[i] = make([]int, 0)
	}
	preNum := make([]int, numCourses)
	n := len(prerequisites)
	for i := 0; i < n; i++ {
		postCourses[prerequisites[i][1]] = append(postCourses[prerequisites[i][1]], prerequisites[i][0])
		preNum[prerequisites[i][0]] += 1
	}
	queue := []int{}
	total := numCourses
	for i := 0; i < total; i++ {
		if preNum[i] == 0 {
			numCourses -= 1 // ##
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for i := 0; i < len(postCourses[curr]); i++ {
			preNum[postCourses[curr][i]] -= 1
			if preNum[postCourses[curr][i]] == 0 {
				numCourses -= 1 // ##
				queue = append(queue, postCourses[curr][i])
			}
		}
	}
	return numCourses == 0
}
