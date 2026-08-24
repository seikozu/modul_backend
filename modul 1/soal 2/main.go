package main

import "fmt"

func main() {
	var (
		name     = "Yae Miko"
		age      = 25
		ipk      = 3.85
		isActive = true
		nilai    = []int{85, 90, 78}
	)
	fmt.Printf("name: %s\n age: %d\n ipk: %.2f\n isActive: %t\n nilai: %v\n", name, age, ipk, isActive, nilai)

	mahasiswa := map[string]int{
		"Yae Miko": 85,
		"Daji":     95,
	}
	mahasiswa["Jinhsi"] = 90
	fmt.Println(mahasiswa)

	if n, ada := mahasiswa["Changli"]; ada {
		fmt.Println("Changli:", n)
	} else {
		fmt.Println("Changli belum punya nilai")
	}
	delete(mahasiswa, "Jinhsi")
	fmt.Println("daftar mahasiswa:", mahasiswa)
	fmt.Println("Jumlah mahasiswa:", len(mahasiswa))
}
