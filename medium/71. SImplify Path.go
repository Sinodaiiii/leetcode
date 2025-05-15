package medium

func simplifyPath(path string) string {
	stack := make([]string, 0)
	//stack = append(stack, "/")
	n := len(path)
	for i := 0; i < n; i++ {
		if path[i] == '/' {
			for i+1 < n && path[i+1] == '/' {
				i++
			}
			continue
		} else if path[i] == '.' {
			if i+1 < n {
				if path[i+1] == '.' {
					if i+2 >= n || path[i+2] == '/' {
						if len(stack) >= 1 {
							stack = stack[:len(stack)-1]
						}
						continue
					}
				} else if path[i+1] == '/' {
					continue
				}
			} else {
				break
			}
		}
		j := i
		for j < n && path[j] != '/' {
			j++
		}
		stack = append(stack, path[i:j])
		i = j - 1
	}
	ans := ""
	for _, filename := range stack {
		ans += "/" + filename
	}
	if len(ans) == 0 {
		ans = "/"
	}
	return ans
}
