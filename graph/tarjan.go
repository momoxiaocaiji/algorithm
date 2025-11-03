package graph

/*
寻找有向图中的最大强连通子图
复杂度 O(V+E)
*/

func entryPoint(root map[int][]int) {
	nodeNum := len(root) // 这里必须是图的邻接矩阵, 也就是说如果元素没有出度, 那在map中也得有一个key

	// 关键的数据结构
	dfn := make([]int, nodeNum) // 用于记录图中节点的访问顺序
	low := make([]int, nodeNum) // 用于记录节点能从栈中最早返回的节点 -> 字面上老tm难理解了
	index := 1                  // 用于区别默认值0

	stack := make([]int, 0)          // 用于辅助记录当前的链路
	onStack := make([]bool, nodeNum) // 用于记录当前的元素是否在栈内

	cmmList := make([][]int, 0)

	var tarjan func(int)
	tarjan = func(u int) {
		dfn[u] = index
		low[u] = index
		index++

		stack = append(stack, u)
		onStack[u] = true

		for _, v := range root[u] {
			if dfn[v] == 0 {
				tarjan(v)
				low[u] = min(low[u], low[v])
			} else if onStack[v] {
				low[u] = min(low[u], dfn[v])
			}
		}

		if dfn[u] == low[u] {
			cmm := make([]int, 0)
			for {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[top] = false
				cmm = append(cmm, top)

				if top == u {
					cmmList = append(cmmList, cmm)
					break
				}
			}
		}
	}

	for u := range root {
		if dfn[u] == 0 {
			tarjan(u)
		}
	}
	return
}
