package main

import "fmt"

func main() {
	arrA := []string{"king", " jin", "akuma", "sergei"}
	arrB := []string{"akuma", "eddie", "steve", "king"}

	merged := ArrayMerge(arrA, arrB)
	fmt.Println(merged)

	//Soal No. 2 menghitung string yang sama

	arraySlice := []string{"asd", "qwe", "asd", "adi", "qwe", "qwe", "awa"}
	counts := make(map[string]int)

	for _, arraySlice := range arraySlice {
		counts[arraySlice]++
	}

	for arraySlice, count := range counts {
		fmt.Printf("%s: %d\n", arraySlice, count)
	}

}
func ArrayMerge(arrA, arrB []string) []string {
	// membuat map kosong untuk menyimpan elemen
	mergedMap := make(map[string]bool)

	// menyimpan elemen dari array pertama ke dalam map
	for _, elem := range arrA {
		mergedMap[elem] = true
	}

	// menyimpan elemen dari array kedua ke dalam map, jika belum ada di map
	for _, elem := range arrB {
		if _, exists := mergedMap[elem]; !exists {
			mergedMap[elem] = true
		}
	}

	// mengembalikan elemen-elemen yang tersisa di map sebagai slice
	var merged []string
	for elem := range mergedMap {
		merged = append(merged, elem)
	}

	return merged
}
