package inspect

import "fmt"

func InspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d %p\n", name, len(slice), cap(slice), &slice)
	fmt.Println(slice)
	fmt.Println()
}
