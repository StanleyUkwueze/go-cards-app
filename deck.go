package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck{
	cards:= deck{}

	cardsSuits := []string{"Spades","Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two","Three","Four"}


	for _, suit :=range cardsSuits{
		for _, value:= range cardValues {

			cards =append(cards, value + " of "+ suit)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i,card)
	}
}

func deal(d deck, handsize int)(deck,deck){
     
   return d[:handsize],d[handsize:]
}

func (d deck) toString() string{
 return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck{
	bs, err := ioutil.ReadFile(fileName)

	if err != nil{
		fmt.Println("Error:", err)

		os.Exit(1)
	}

	s:= strings.Split(string(bs), ",")
     
	return deck(s)
}

func (d deck) shuffle(){
	for i := range d{
		newPosition := rand.Intn(len(d)-1)

		d[i],d[newPosition] = d[newPosition], d[i]
	}
}