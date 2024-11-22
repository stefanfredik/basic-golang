package main

func main() {
	println("Hello Word")

	var arr = []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	var sorted = sortArray(arr)

	print(sorted)
}

// buatkan sebuah fungsi sederhana untuk mengurutkan aarray
func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// buatkan sebuah variabel array
