package one

func CanAddToMake2020(nums []int) int {
	mapper := make(map[int]bool)
	for _, currentNum := range nums {
		tt := 2020
		mapNum := tt - currentNum
		if mapper[mapNum] {
			return mapNum * currentNum
		}
		mapper[currentNum] = true

	}
	return 0
}
