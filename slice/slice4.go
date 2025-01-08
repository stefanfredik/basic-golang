// slice refferncy ke array
package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	slice[0] = 99                // Mengubah elemen pertama slice
	fmt.Println("Array:", array) // Output: [1 99 3 4 5]
	fmt.Println("Slice:", slice) // Output: [99 3 4]
}
