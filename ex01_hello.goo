// potrzebujemy zmiennej GOPATH, w której będzie ścieżka do przestrzeni roboczej
// go env - wypisuje zmienne środowiskowe związane z Go
// cd %GOPATH%
// go run program.go - plik wykonywalny będzie w temp

// Go jest case-sensitive
// whitespaces nie mają znaczenia

// dzięki package name będziemy mieli program wykonywalny, a nie bibliotekę

package main

// pakiety z biblioteki standardowej, importowane za pomocą krótkiej nazwy, bez ścieżki

import (
    "fmt"
    "runtime"
)

// w Go funkcje są first-class
// main() jest entry point, biblioteki go nie potrzebują
// zwraca kod 0 po poprawnym wykonaniu (właściwie po zakończeniu całego programu)

func main() {

    // w kompilatorze jest też funkcja println, ale pod kątem środowiska produkcyjnego lepiej używać idiomatycznego fmt.Println

    // funkcje pisane z wielkiej litery będą eksponowane poza pakietem, w którym są zdefiniowane

    // będzie spacja między wypisanymi argumentami

    fmt.Println("Hello from", runtime.GOOS)
}
