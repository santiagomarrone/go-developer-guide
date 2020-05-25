package main

import "fmt"

type deck []string 

func newDeck() deck {

	d := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {

			d = append(d, value+" of "+suit)
		}
	}

	return d
}

func (d deck) print() {

	for _, card := range d {
		
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}