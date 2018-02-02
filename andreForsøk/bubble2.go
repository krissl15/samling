package andreForsøk

func BubbleSort(arr[] int)[]int {
	for i:=1; i< len(arr); i++ {
		for j:=0; j < len(arr)-i; j++ {
			if (arr[j] > arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
