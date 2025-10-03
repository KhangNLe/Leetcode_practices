package solution

type Data struct {
	val int
	row int
	col int
}

type MinHeap struct {
	data    []Data
	visited map[[2]int]bool
}

func (m *MinHeap) Insert(val Data) {
	if ok, _ := m.visited[[2]int{val.row, val.col}]; ok {
		return
	} else {
		m.visited[[2]int{val.row, val.col}] = true
	}

	m.data = append(m.data, val)
	m.heapifyUp(len(m.data) - 1)
}

func (m *MinHeap) isEmpty() bool {
	return len(m.data) == 0
}

func (m *MinHeap) getMin() Data {
	smallest := m.data[0]
	last := len(m.data) - 1
	m.data[0] = m.data[last]
	m.data = m.data[:last]

	m.heapifyDown(0)
	return smallest
}

func (m *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentNode := (index - 1) / 2
		if m.data[index].val < m.data[parentNode].val {
			m.data[index], m.data[parentNode] =
				m.data[parentNode], m.data[index]

			index = parentNode
		} else {
			break
		}
	}
}

func (m *MinHeap) heapifyDown(index int) {
	lastNode := len(m.data) - 1
	for {
		leftNode := 2*index + 1
		rightNode := 2*index + 2
		smallest := index

		if leftNode <= lastNode &&
			m.data[leftNode].val < m.data[smallest].val {
			smallest = leftNode
		}
		if rightNode <= lastNode &&
			m.data[rightNode].val < m.data[smallest].val {
			smallest = rightNode
		}

		if smallest != index {
			m.data[index], m.data[smallest] =
				m.data[smallest], m.data[index]
			index = smallest
		} else {
			break
		}
	}
}

func TrapRainWater(heightMap [][]int) int {
	rows := len(heightMap)
	cols := len(heightMap[0])

	if rows <= 2 || cols <= 2 {
		return 0
	}

	minHeap := MinHeap{
		data:    make([]Data, 0),
		visited: make(map[[2]int]bool),
	}

	for i := 1; i < rows-1; i++ {
		minHeap.Insert(Data{heightMap[i][0], i, 0})
		minHeap.Insert(Data{heightMap[i][cols-1], i, cols - 1})
	}

	for j := range cols {
		minHeap.Insert(Data{heightMap[0][j], 0, j})
		minHeap.Insert(Data{heightMap[rows-1][j], rows - 1, j})
	}

	highestColumn := 0
	totalWater := 0
	var current Data
	for !minHeap.isEmpty() {
		current = minHeap.getMin()
		newDirections := getNeightbors(current, rows, cols)
		highestColumn = max(highestColumn, current.val)
		totalWater += highestColumn - current.val

		for _, dir := range newDirections {
			if minHeap.visited[[2]int{dir[0], dir[1]}] {
				continue
			}

			val := heightMap[dir[0]][dir[1]]
			minHeap.Insert(Data{val, dir[0], dir[1]})
		}
	}

	return totalWater
}

func getNeightbors(data Data, maxRow int, maxCol int) [][2]int {
	ans := make([][2]int, 0)

	direction := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for _, dir := range direction {
		newRow := data.row + dir[0]
		newCol := data.col + dir[1]

		if 0 <= newRow && newRow < maxRow && 0 <= newCol && newCol < maxCol {
			ans = append(ans, [2]int{newRow, newCol})
		}
	}

	return ans
}
