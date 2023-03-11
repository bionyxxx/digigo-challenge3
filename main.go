package main

import "fmt"

func main() {
	//DUTTA FACHREZY

	//variabel yang berisi kalimat apa saja
	sentence := "digitalent kominfo"

	//loop untuk menampilkan setiap karakter yang ada di variable sentence
	for _, char := range sentence {
		fmt.Printf("%c\n", char)
	}

	//map untuk menghitung kemunculan setiap karakter
	charCount := make(map[rune]int)

	//looping untuk menghitung kemunculan setiap karakter
	for _, char := range sentence {
		charCount[char]++
	}

	//menampilkan hasil perhitungan dengan format map[karakter:jumlah]
	fmt.Print("map[ ")
	for char, count := range charCount {
		fmt.Printf("%c:%d ", char, count)
	}
	fmt.Print("]\n")
}
