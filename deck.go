package main

import "fmt"

// create a new type of `deck`
// which is a slice of strings
type deck []string

func newDeck() deck {
  cards := deck{}

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace", "Two", "Three", "Four"}

  // for variables created but not used, use '_'
  // so that there won't be any compilation warnings
  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, value+" of "+suit)
    }
  }

  return cards
}

/*
  Receiver function
  The variable 'd' is the same as
  having the variable 'cards', but 
  by convention it is used to first
  one or two letters of the variable
  type; in this case, 'd' for type 'deck'
*/
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}
