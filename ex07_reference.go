package main

import "fmt"

func main() {
	name := "Sebastian"
	course := "Mastering Go"

	fmt.Println("\nHi", name, "you're currently watching", course)

	fmt.Println("address", &course)
	changeCourse(&course)

	fmt.Println("value of course after passing to function", course)
}

// tak podajemy wskaźnik jako typ
func changeCourse(courseVal *string) {

	// sama zmienna courseVal zawiera adres, musimy użyć dereferencji
	fmt.Println("address in changeCourse function", courseVal)

	// mamy teraz dostęp do zmiennej w danej komórce pamięci
	*courseVal = "Mastering Docker"
	fmt.Println("changed value in function", *courseVal)

	// zwracanie nie jest do niczego potrzebne, ponieważ operujemy bezpośrednio wartością zapisaną w konkretnej komórce pamięci
	//return *courseVal
}
