package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	// x := 10

	// changeValue(&x)
	// fmt.Println(x)

	// print format
	// printByTypes()

	// arrays
	// var EvenNum[5] int
	EvenNum := [5]int{10, 20, 30, 40, 50}
	fmt.Println(EvenNum)

	for _, value := range EvenNum {
		fmt.Println(value)
	}

	// imprime o index (i)
	for i, value := range EvenNum {
		fmt.Println(value, i)
	}

	// corta um intervalo
	numSlice := []int{1, 2, 3, 4, 5, 6}
	sliced := numSlice[3:6]
	fmt.Println(sliced)

	slice2 := make([]int, 5, 10)
	fmt.Println(slice2) // pq não imprime os valores?

	copy(slice2, numSlice)
	fmt.Println(slice2)

	slice3 := append(numSlice, 3, 0, 15)
	fmt.Println(slice3)

	// map
	CatsAge := make(map[string]int)

	CatsAge["Auri"] = 10
	CatsAge["Quitu"] = 6
	CatsAge["Nina"] = 2

	fmt.Println(CatsAge)
	fmt.Println(CatsAge["Quitu"])

	// mapSuperhero()

	fmt.Println(add(1, 2))

	// defer, recover and panic
	defer FirstRun()
	SecondRun()

	fmt.Println((div(3, 0)))
	fmt.Println((div(15, 3)))
	demPanic()

	// structures and interfaces
	rect1 := Rectangle{10, 5}
	circ := Circle{7}
	fmt.Println("Height", rect1.height)
	fmt.Println("Width", rect1.width)

	fmt.Println("Area of rectangle is: ", rect1.area())
	fmt.Println("Area of circle is: ", circ.area())

	// IO
	writeFile()
}

func changeValue(x *int) {
	*x = 7
}

func printByTypes() {
	var name string = "Auri Costa Baltezan"
	fmt.Println(name + " is a little cow ;)")
	fmt.Println(len(name))    //imprime o tamanho de uma string ?
	fmt.Printf("%T \n", name) //imprime o type

	const pi float64 = 3.142435
	const number int = 1234
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%e \n", pi) // imprime notação científica

	win := true
	fmt.Printf("%t \n", win)
	fmt.Println(win)

	x := 5
	fmt.Printf("%d \n", x)

	fmt.Printf("%b \n", 2)  // imprime o binário
	fmt.Printf("%c \n", 2)  // imprime o código ASCI
	fmt.Printf("%x \n", 15) // imprime o hexadecimal
}

func mapSuperhero() {
	// ????????????
	type hero map[string]string

	superhero := map[string]hero{
		"Wonder Woman": {
			"RealName": "Diana",
			"City":     "",
		},
		"Batman": {
			"RealName": "Bruce Wayne",
			"City":     "Gotham",
		},
	}

	if temp, hero := superhero["Wonder Woman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}

func add(num1, num2 int) int {
	return num1 + num2
}

func FirstRun() {
	fmt.Println("Executed firstRun")
}

func SecondRun() {
	fmt.Println("Executed SecondRun")
}

func div(num1 int, num2 int) int {
	defer func() {
		fmt.Println("div recover", recover())
	}()
	result := num1 / num2
	return result
}

func demPanic() {
	defer func() {
		fmt.Println("demPanic recover", recover())
	}()
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 *Rectangle) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func writeFile() {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal()
	}

	file.WriteString("Hello")
	file.Close()

	stream, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal()
	}

	s1 := string(stream)

	println(s1)
}
