package hard

//type strNode126 struct {
//    str         string
//    pre         []string
//    next        []*strNode126
//    level       int
//}
//
//func findLadders(beginWord string, endWord string, wordList []string) [][]string {
//    wordn := len(beginWord)
//    dict := map[string]*strNode126{}
//    ans := [][]string{}
//    beginExist, endExist := false, false
//    for _, word := range wordList {
//        if word == beginWord {
//            beginExist = true
//        }
//        if word == endWord {
//            endExist = true
//        }
//    }
//    if endExist == false {
//        return ans
//    }
//    if !beginExist {
//        wordList = append(wordList, beginWord)
//    }
//    for _, word := range wordList {
//        wordByte := []byte(word)
//        neighSNs := []*strNode126{}
//        for i:=0; i<wordn; i++ {
//            tmp := wordByte[i]
//            wordByte[i] = '*'
//            neighWord := string(wordByte)
//            if _, exist := dict[neighWord]; !exist {
//                dict[neighWord] = &strNode126{str:neighWord}
//            }
//            neighSNs = append(neighSNs, dict[neighWord])
//            wordByte[i] = tmp
//        }
//        dict[word] = &strNode126{str:word, next:neighSNs}
//        for _, neighSN := range neighSNs {
//            neighSN.next = append(neighSN.next, dict[word])
//        }
//    }
//    // fmt.Println(*dict[beginWord])
//    // fmt.Println(*dict[endWord])
//    currAns := []string{}
//    existDict := map[string]bool{}
//    var dfs func(dfsStr string, isString bool)
//    dfs = func(dfsStr string, isString bool) {
//        // fmt.Println(dfsStr)
//        // fmt.Println(dict[dfsStr].pre)
//        if dfsStr == beginWord {
//            currAnsn := len(currAns)
//            tmpAns := make([]string, currAnsn+1)
//            for i:=currAnsn-1; i>=0; i-- {
//                tmpAns[currAnsn-i] = currAns[i]
//            }
//            tmpAns[0] = beginWord
//            ans = append(ans, tmpAns)
//            return
//        }
//        if isString == true {
//            existDictKey := ""
//            for i:=0; i<len(currAns); i++ {
//                existDictKey += currAns[i]
//            }
//            existDictKey += dfsStr
//            if existDict[existDictKey] {
//                return
//            }
//            existDict[existDictKey] = true
//            currAns = append(currAns, dfsStr)
//        }
//        for _, preStr := range dict[dfsStr].pre {
//            dfs(preStr, !isString)
//        }
//        if isString == true {
//            currAns = currAns[:len(currAns)-1]
//        }
//    }
//    queue := []*strNode126{dict[beginWord]}
//    level := 1
//    for len(queue) != 0 {
//        // fmt.Println(queue)
//        curr := len(queue)
//        found := false
//        for i:=0; i<curr; i++ {
//            currStr := queue[i].str
//            for _, next := range queue[i].next {
//                if next.level > 0 && next.level != level {
//                    continue
//                } else {
//                    if next.str == endWord {
//                        found = true
//                    }
//                    next.level = level
//                    next.pre = append(next.pre, currStr)
//                    queue = append(queue, next)
//                }
//            }
//        }
//        queue = queue[curr:]
//        if found == true {
//            dfs(endWord, true)
//            break
//        }
//        level += 1
//    }
//    return ans
//}

//你的解法思路是通过构建单词间的邻接关系（利用通配符中间态），再通过 BFS 找最短路径，最后用 DFS 收集所有路径，方向是对的。
//但代码存在一些可以优化的点，主要集中在数据结构设计、BFS 效率、去重逻辑三个方面，优化后能显著提升性能。
//
//核心问题分析
//
//数据结构冗余：定义 strNode126 包含 pre、next、level 等字段，且用 dict 存储所有节点（包括原始单词和通配符中间态），
//导致内存占用大，且操作复杂（需要区分原始单词和通配符）。实际上，通配符只是构建邻接关系的工具，无需长期存储节点。

//BFS 逻辑低效：
//未严格区分 “原始单词” 和 “通配符中间态”，导致 BFS 队列中混入大量通配符节点，增加遍历成本。
//level 标记逻辑不清晰，可能重复处理节点。
//去重逻辑复杂：DFS 中用 existDictKey（拼接路径字符串）来去重，效率极低（字符串拼接耗时，且哈希表键值过长）。
//优化方案
//
//1. 简化数据结构，分离 “单词邻接表” 和 “路径追踪”
//
//用 map[string][]string 直接存储原始单词的邻接表（跳过通配符节点，仅在构建时使用通配符）。
//用 map[string][]string 存储 BFS 中的 pre 前驱关系（替代 strNode126.pre），用 map[string]int 存储 level（替代节点字段）。
//2. 优化 BFS 流程
//
//仅将原始单词加入队列，通配符仅用于构建邻接表，不参与 BFS。
//严格按层级遍历，避免重复访问同一单词（通过 level 记录是否已访问）。
//3. 简化 DFS 去重逻辑
//
//由于 BFS 已保证路径是最短的，且 pre 中只记录同一层级的前驱，DFS 时无需额外去重，直接回溯即可。

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 1. 构建单词集合，快速判断是否存在
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] { // 终点不在列表中，直接返回
		return nil
	}
	// 确保起点在集合中（方便统一处理）
	wordSet[beginWord] = true

	// 2. 构建邻接表：key是单词，value是所有可能的邻居（相差一个字符）
	adj := make(map[string][]string)
	wordLen := len(beginWord)
	// 用通配符构建邻接关系（如 "hot" → "*ot", "h*t", "ho*"）
	patternMap := make(map[string][]string)
	for w := range wordSet {
		for i := 0; i < wordLen; i++ {
			// 生成通配符模式（如 "hot" 在i=0时为 "*ot"）
			pattern := w[:i] + "*" + w[i+1:]
			patternMap[pattern] = append(patternMap[pattern], w)
		}
	}
	// 为每个单词收集邻居（所有共享同一通配符模式的单词）
	for w := range wordSet {
		neighbors := make([]string, 0)
		for i := 0; i < wordLen; i++ {
			pattern := w[:i] + "*" + w[i+1:]
			for _, neighbor := range patternMap[pattern] {
				if neighbor != w { // 排除自身
					neighbors = append(neighbors, neighbor)
				}
			}
		}
		adj[w] = neighbors
	}

	// 3. BFS 寻找最短路径，记录每个单词的前驱节点（pre）和层级（level）
	pre := make(map[string][]string) // 前驱节点列表
	level := make(map[string]int)    // 记录每个单词的层级（用于控制BFS顺序）
	queue := []string{beginWord}
	level[beginWord] = 0
	found := false

	for len(queue) > 0 && !found {
		size := len(queue)
		// 同一层级的节点一起处理
		for i := 0; i < size; i++ {
			curr := queue[i]
			for _, neighbor := range adj[curr] {
				if _, ok := level[neighbor]; !ok { // 未访问过
					level[neighbor] = level[curr] + 1
					pre[neighbor] = []string{curr}
					queue = append(queue, neighbor)
					if neighbor == endWord {
						found = true
					}
				} else if level[neighbor] == level[curr]+1 { // 已访问过但同层级，添加前驱
					pre[neighbor] = append(pre[neighbor], curr)
				}
			}
		}
		queue = queue[size:] // 移除当前层级的节点
	}

	// 4. DFS 回溯收集所有最短路径
	var result [][]string
	if !found {
		return result
	}
	// 从终点回溯到起点
	var dfs func(word string, path []string)
	dfs = func(word string, path []string) {
		if word == beginWord {
			// 反转路径（因为是从终点回溯的）
			reversed := make([]string, len(path))
			for i := 0; i < len(path); i++ {
				reversed[i] = path[len(path)-1-i]
			}
			result = append(result, reversed)
			return
		}
		// 遍历所有前驱，继续回溯
		for _, p := range pre[word] {
			dfs(p, append(path, p))
		}
	}
	// 初始路径为 [endWord]，从终点开始回溯
	dfs(endWord, []string{endWord})

	return result
}
