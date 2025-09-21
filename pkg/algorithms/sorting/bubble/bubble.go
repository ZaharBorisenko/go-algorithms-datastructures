package bubble

func Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		isSwapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
	return arr
}
