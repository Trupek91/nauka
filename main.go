package main // paczka komend golang

import "fmt" //komenda do importowania paczki

var trupka = "Trupek"

// fmt moze byc zastapione yy (zapomnialem w jakiej sytuacji)
// Test jest zmienna ktora pomaga okreslic cos
const Test = "Pogromca, Yakuzy"
const jedynka = 6

// const nie pozwala nadpisac zmiennej
func main() { // glowna funkcja ktora zawsze jest wyswietlana i wykonywana
	var niraziel string // var okresla zmienna a string definiuje zmienna jako litery
	niraziel = "niraziel"
	fmt.Println(niraziel)
	fmt.Println(trupka, Test)
	fmt.Println(dodawanie(3, 1))
	fmt.Println(xx())

}

// dowanie(a,b)  a ib to arugmenty i musza byc zawsze podane
// dodawanie()
func dodawanie(a int, b int) (int, int) { // int okresla zmienna jako liczbe
	a = a * 2
	b = b * 2
	return a, b
}

func dodowanieBezWracania(a int, b int) int {
	fmt.Println(a + b) // funkcja do wyswietlania
	fmt.Println("Test pod koniec funkcji", Test)
	return a + b // zadanie co ma funkcja zwrocic po wykonaniu czynnosci
}
