package main

import "fmt"

func main () {

	// x := 10

	// changeValue(&x)
	// fmt.Println(x)
	

	// printByTypes()

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
	numSlice := []int{1,2,3,4,5,6}
	sliced := numSlice[3:6]
	fmt.Println(sliced)

	slice2 := make([]int,5,10)
	fmt.Println(slice2) // pq não imprime os valores?

	copy(slice2,numSlice)
	fmt.Println(slice2)

	slice3 := append(numSlice, 3, 0,15)
	fmt.Println(slice3)

	CatsAge := make(map[string] int)

	CatsAge["Auri"] = 10
	CatsAge["Quitu"] = 6
	CatsAge["Nina"] = 2

	fmt.Println(CatsAge)
	fmt.Println(CatsAge["Quitu"])

	mapSuperhero()

	fmt.Println(add(1,2))

}

func changeValue(x *int) {
	*x = 7
}

func printByTypes() {
	var name string = "Auri Costa Baltezan"
	fmt.Println(name + " is a little cow ;)")
	fmt.Println(len(name)); //imprime o tamanho de uma string ?
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

	fmt.Printf("%b \n", 2) // imprime o binário
	fmt.Printf("%c \n", 2) // imprime o código ASCI
	fmt.Printf("%x \n", 15) // imprime o hexadecimal
}


func mapSuperhero() {
	// ????????????
	superhero := map[string]map[string]string {
		"Wonder Woman": map[string]string {
			"RealName": "Diana",
			"City": "",
		},
		"Batman": map[string] string{
			"RealName": "Bruce Wayne",
			"City": "Gotham",
		},
	}

	if temp, hero := superhero["Wonder Woman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}

func add (num1, num2 int) int {
	return num1 + num2
}