package hard

func stringIndices260528(wordsContainer []string, wordsQuery []string) []int {
	type node struct {
		minIndex int
		next     []*node
	}
	head := &node{0, make([]*node, 26)}
	n := len(wordsContainer)
	for index, word := range wordsContainer {
		m := len(word)
		if m < len(wordsContainer[head.minIndex]) {
			head.minIndex = index
		}
		curr := head
		for i := m - 1; i >= 0; i-- {
			if curr.next[word[i]-'a'] == nil {
				curr.next[word[i]-'a'] = &node{index, make([]*node, 26)}
			}
			curr = curr.next[word[i]-'a']
			if m < len(wordsContainer[curr.minIndex]) {
				curr.minIndex = index
			}
		}
	}
	n = len(wordsQuery)
	ans := make([]int, n)
	for i, word := range wordsQuery {
		m := len(word)
		curr := head
		for j := m - 1; j >= 0 && curr.next[word[j]-'a'] != nil; j-- {
			curr = curr.next[word[j]-'a']
		}
		ans[i] = curr.minIndex
	}
	return ans
}
