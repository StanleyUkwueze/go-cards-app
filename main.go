package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()

	//cards.Print()

	//check for odd and even numbers
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, ar := range arr {
		if i%2 == 0 {
			fmt.Printf("%v is even \n", ar)
		} else {
			fmt.Printf("%v is odd \n", ar)
		}
	}
}
