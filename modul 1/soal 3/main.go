package main

import "fmt"

func swap(name1, name2 *string){
	*name1, *name2 = *name2, *name1
}

func updateSlice(slice *[]string, newValue string) {
	*slice = append(*slice, newValue)
}

func main() {
	name1 := "yaemiko"
	name2 := "jinhsi"
	fmt.Println("sebelum ditukar:", name1, name2)
	swap(&name1, &name2)
	fmt.Println("setelah ditukar:", name1, name2)

	characters := []string{"yaemiko", "jinhsi"}
	fmt.Println("sebelum update slice:", characters)
	updateSlice(&characters, "cantarella")
	fmt.Println("setelah update slice:", characters)
}