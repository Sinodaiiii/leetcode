package hot100

import "sort"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	idleList := make([][2]int, n)
	taskList := make([][2]int, 0)
	dict := map[int]int{}
	for i := 0; i < len(tasks); i++ {
		dict[int(tasks[i]-'A'+1)] += 1
	}
	for k, v := range dict {
		taskList = append(taskList, [2]int{k, v})
	}
	ans := 0
	remain := len(taskList)
	for remain > 0 {
		sort.Slice(taskList, func(i, j int) bool { return taskList[i][1] > taskList[j][1] })
		ans += 1
		tmp := idleList[0]
		if len(taskList) > 0 {
			curr := taskList[0]
			taskList = taskList[1:]
			if curr[1] != 1 {
				idleList = append(idleList[1:], [2]int{curr[0], curr[1] - 1})
			} else {
				remain -= 1
				idleList = append(idleList[1:], [2]int{})
			}
		} else {
			idleList = append(idleList[1:], [2]int{})
		}
		if tmp[1] != 0 {
			taskList = append(taskList, tmp)
		}
	}
	return ans
}
