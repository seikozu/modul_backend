package main

import "fmt"

type student struct {
	name string
	id int 
	grade  float64
	isActive bool
}

func (s student) printInfo() {
	fmt.Printf("name: %s\nid: %d\ngrade: %.2f\nisActive: %t\n", s.name, s.id, s.grade, s.isActive)
}

func (s *student) updateGrade(newGrade float64) {
	s.grade = newGrade
}

func (s *student) deactivate() {
	s.isActive = false
}

func main() {
	var yaeMiko = student{
		name: "Yae Miko",
		id: 1,
		grade: 85.5,
		isActive: true,
	}

	jinhsi := student{
		name: "Jinhsi",
		id: 2,
		grade: 90.0,
		isActive: true,
	}
	cantarella := student{
		name: "Cantarella",
		id: 3,
		grade: 88.0,
		isActive: true,
	}
	yaeMiko.printInfo()

	jinhsi.updateGrade(90.2)
	fmt.Println("setelah update grade:")
	jinhsi.printInfo()
	
	cantarella.deactivate()
	fmt.Println("setelah deactivate:")
	cantarella.printInfo()
}
