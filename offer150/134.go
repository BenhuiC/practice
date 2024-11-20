package offer150

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	left, right := 0, 1%n
	curGas := gas[0] - cost[0]
	for left != right {
		if curGas < 0 {
			left = (left - 1 + n) % n
			curGas += gas[left] - cost[left]
		} else {
			curGas += gas[right] - cost[right]
			right = (right + 1) % n
		}
	}
	if curGas < 0 {
		return -1
	}
	return left
}
