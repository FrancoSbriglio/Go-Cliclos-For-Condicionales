package main

import "fmt"

func main() {
	numbers := []int{0, 5, 9, 10}

	fruits := []string{"manzana", "Banano", "Sandia", "Pera"}

	numbers = append(numbers, 8, 3)
	fmt.Println(numbers)

	//slices
	numbers = numbers[1:5]
	//numbers = numbers[:5]
	//numbers = numbers[1:]
	fmt.Println(numbers)

	//formas de realizar for

	/*for  {
		fmt.Println("El ciclo no va a terminar nunca")
	}
	*/
	/*for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}*/

	for i := range numbers {
		fmt.Println(numbers[i])
	}

	// si no necesitamos el indice podemos usar _
	for i, number := range numbers {

		//condicionales
		if 1 == 1 {

		}

		fmt.Println("Index", i, "Value", number)

	}

	for _, fruit := range fruits {
		if fruit != "Banano" {
			continue
		}
		fmt.Println("Fruit", fruit)

		index := len(fruit) - 1
		letter := fruit[index:]

		if letter == "a" {
			continue
		}
		fmt.Println("Fruit", fruit)

	}

}
