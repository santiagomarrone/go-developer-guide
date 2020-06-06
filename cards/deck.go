package main

import(

	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func newDeckFromFile(file string) deck {

	content, err := ioutil.ReadFile(file)

	if err != nil {

		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	
	s := strings.Split(string(content),",")
	
	return deck(s)
}

func (d deck) print() {

	for _, card := range d {
		
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(file string) error {

	return ioutil.WriteFile(file, []byte(d.toString()), 0666)
}