package main

import "fmt"

func main() {
	var name string = "Akhil"
	lastname := "Singh"
	fmt.Print(name)
	fmt.Print(lastname + "\n")

	age := 25

	fmt.Printf("Hi there, my name is %v %v and I am %0.0f years old.", name, lastname, float64(age))

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	names := []string{"akhil", "rina", "ansh"}

	for index, value := range names {
		fmt.Printf("Hi there, this name '%v' is at index '%v'.\n", value, index)
	}

	fmt.Println(add(2, 5))
	fmt.Println(sub(2, 5))
	fmt.Println(abs(2, 5))

	for i := 0; i < len(name); i++ {
		print(name[i])
	}
}
