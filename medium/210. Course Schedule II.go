package medium

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 1 {
		return []int{0}
	}
	ans := []int{}
	prev := make([]int, numCourses)
	post_course := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		prev[prerequisite[0]] += 1
		post_course[prerequisite[1]] = append(post_course[prerequisite[1]], prerequisite[0])
	}
	stack := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if prev[i] == 0 {
			stack = append(stack, i)
		}
	}
	for len(stack) != 0 {
		course := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, course)
		postl := len(post_course[course])
		for i := 0; i < postl; i++ {
			prev[post_course[course][i]] -= 1
			if prev[post_course[course][i]] == 0 {
				stack = append(stack, post_course[course][i])
			}
		}
	}
	if len(ans) != numCourses {
		return nil
	}
	return ans
}
