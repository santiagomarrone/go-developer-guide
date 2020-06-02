package main

func main() {

	cards := newDeck()
	cards.saveToFile("cardsFile")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}