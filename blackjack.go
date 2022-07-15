package main

import(
	"fmt"
	"math/rand"
)

type Card struct {
	value string
	suit string
}

var deck[] Card

func shuffleDeck() {
	for i := len(deck) - 1; i > 0: i-- {
		index := rand.Intn(i + 1)
		temp := deck[i]
		deck[i] = deck[index]
		deck[index] = temp
	}
}

func generateDeck() {
	deck = append(deck, Card{"Ace", "Spades"})
	deck = append(deck, Card{"One", "Spades"})
	deck = append(deck, Card{"Two", "Spades"})
	deck = append(deck, Card{"Three", "Spades"})
	deck = append(deck, Card{"Four", "Spades"})
	deck = append(deck, Card{"Five", "Spades"})
	deck = append(deck, Card{"Six", "Spades"})
	deck = append(deck, Card{"Seven", "Spades"})
	deck = append(deck, Card{"Eight", "Spades"})
	deck = append(deck, Card{"Nine", "Spades"})
	deck = append(deck, Card{"Ten", "Spades"})
	deck = append(deck, Card{"Jack", "Spades"})
	deck = append(deck, Card{"Queen", "Spades"})
	deck = append(deck, Card{"King", "Spades"})
	deck = append(deck, Card{"Ace", "Hearts"})
	deck = append(deck, Card{"One", "Hearts"})
	deck = append(deck, Card{"Two", "Hearts"})
	deck = append(deck, Card{"Three", "Hearts"})
	deck = append(deck, Card{"Four", "Hearts"})
	deck = append(deck, Card{"Five", "Hearts"})
	deck = append(deck, Card{"Six", "Hearts"})
	deck = append(deck, Card{"Seven", "Hearts"})
	deck = append(deck, Card{"Eight", "Hearts"})
	deck = append(deck, Card{"Nine", "Hearts"})
	deck = append(deck, Card{"Ten", "Hearts"})
	deck = append(deck, Card{"Jack", "Hearts"})
	deck = append(deck, Card{"Queen", "Hearts"})
	deck = append(deck, Card{"King", "Hearts"})
	deck = append(deck, Card{"Ace", "Clubs"})
	deck = append(deck, Card{"One", "Clubs"})
	deck = append(deck, Card{"Two", "Clubs"})
	deck = append(deck, Card{"Three", "Clubs"})
	deck = append(deck, Card{"Four", "Clubs"})
	deck = append(deck, Card{"Five", "Clubs"})
	deck = append(deck, Card{"Six", "Clubs"})
	deck = append(deck, Card{"Seven", "Clubs"})
	deck = append(deck, Card{"Eight", "Clubs"})
	deck = append(deck, Card{"Nine", "Clubs"})
	deck = append(deck, Card{"Ten", "Clubs"})
	deck = append(deck, Card{"Jack", "Clubs"})
	deck = append(deck, Card{"Queen", "Clubs"})
	deck = append(deck, Card{"King", "Clubs"})
	deck = append(deck, Card{"Ace", "Diamonds"})
	deck = append(deck, Card{"One", "Diamonds"})
	deck = append(deck, Card{"Two", "Diamonds"})
	deck = append(deck, Card{"Three", "Diamonds"})
	deck = append(deck, Card{"Four", "Diamonds"})
	deck = append(deck, Card{"Five", "Diamonds"})
	deck = append(deck, Card{"Six", "Diamonds"})
	deck = append(deck, Card{"Seven", "Diamonds"})
	deck = append(deck, Card{"Eight", "Diamonds"})
	deck = append(deck, Card{"Nine", "Diamonds"})
	deck = append(deck, Card{"Ten", "Diamonds"})
	deck = append(deck, Card{"Jack", "Diamonds"})
	deck = append(deck, Card{"Queen", "Diamonds"})
	deck = append(deck, Card{"King", "Diamonds"})
}

func main() {

}