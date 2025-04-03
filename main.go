package main

import "fmt"

func main(){
	var ages = [5]int{17,16,20,40,50}
	fmt.Println(ages)
	// slice
	var nomes =[]string {"Leticia","Yasmim","Amanda", "Ana","Fabiano" }
	fmt.Println(nomes)
	fmt.Println(nomes, len(nomes), cap(nomes))
	rangeOne := nomes[:2]
	fmt.Println(rangeOne)	
	rangeTwo := nomes[3:]
	fmt.Println(rangeTwo)
	rangeThree := (nomes[2])
	fmt.Println(rangeThree)

}

