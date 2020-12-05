package main

import "fmt"

var (
	height, width, opc int
	exit               bool
)

func main() {
	for exit == false {
		fmt.Println("Welcome to system")
		fmt.Println("1. Print square")
		fmt.Println("2. Print triangle")
		fmt.Println("3. Print trapezoid")
		fmt.Println("4. Exit")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			square()
		case 2:
			triangle()
		case 3:
			trapezoid()
		case 4:
			fmt.Println("E X I T . . .")
			exit = true
		default:
			fmt.Println("Option not valid")
		}
	}

}

func square() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}

func triangle() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}

func trapezoid() {
	defer square()
	triangle()
}
