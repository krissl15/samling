package OBLIGGGG

import "fmt"

// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	var sorted bool = false
	var lastUnsorted int = len(list)-1
	for !sorted{
		sorted = true
		for i := 0; i<lastUnsorted; i++{
			if list[1]>list[i+1]{
				i, i+1 = swap(i+1,i)
				sorted = false
			}
		}
	}
}


func sap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func bubbleSort(arrayzor []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor) - 1; i++ {
			if arrayzor[i + 1] < arrayzor[i] {
				sap(arrayzor, i, i + 1)
				swapped = true
			}
		}
	}
}

func main() {

	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arrayzor)
	bubbleSort(arrayzor)
	fmt.Println("Sorted array: ", arrayzor)
}