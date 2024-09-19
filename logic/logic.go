package logic

// FindPairs takes an array of integers and a target and returns all pairs of indices that sum up to the target
func FindPairs(numbers []int, target int) [][]int {
	seen := make(map[int][]int)
	var solutions [][]int

	for i, num := range numbers {
		complement := target - num
		if indices, found := seen[complement]; found {
			for _, j := range indices {
				solutions = append(solutions, []int{j, i})
			}
		}
		seen[num] = append(seen[num], i)
	}

	if len(solutions) == 0 {
		return [][]int{}
	}

	return solutions
}