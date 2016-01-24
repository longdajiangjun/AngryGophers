package engine

var SpeedGround1 []int = []int{10, 0, 10, 30, 3}
var Mapa1 [][]int = [][]int{
	[]int{2, 2, 2, 0, 0, 2, 0, 1, 2, 0, 0, 2, 0, 2, 0, 2},
	[]int{0, 2, 0, 0, 2, 0, 2, 1, 2, 2, 0, 2, 0, 2, 2, 0},
	[]int{0, 2, 0, 0, 2, 2, 2, 1, 2, 0, 2, 2, 0, 2, 0, 2},
	[]int{0, 2, 0, 0, 2, 0, 2, 1, 2, 0, 0, 2, 0, 2, 0, 2},
	[]int{3, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 0, 0, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 0, 0, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 3},
	[]int{3, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 3},
	[]int{0, 2, 0, 0, 0, 2, 0, 1, 2, 0, 0, 2, 0, 2, 2, 2},
	[]int{2, 0, 0, 0, 2, 0, 2, 1, 2, 2, 2, 2, 0, 2, 0, 0},
	[]int{2, 2, 2, 0, 2, 2, 2, 1, 2, 0, 0, 2, 0, 2, 2, 2},
	[]int{0, 2, 2, 0, 2, 0, 2, 1, 2, 0, 0, 2, 0, 2, 2, 2},
}

const FramePerSec int64 = 30
const TimePerFrame int64 = 1000000 / FramePerSec