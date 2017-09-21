package main

func distributeCandies(candies []int) int {
	length := len(candies)
	kinds := 0
	candiesMap := make(map[int]int)

	// 遍历所有的candies
	for _, v := range candies {
		_, ok := candiesMap[v]
		// 将新发现的candy种类加入Map中
		// candy种类加1
		if !ok {
			candiesMap[v] = 1
			kinds++
		}

		// 判断当前计算出的candy种类数据是否已经大于所有candy数量的一半
		if kinds > length/2 {
			return length / 2
		}
	}

	return kinds
}
