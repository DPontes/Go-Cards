package main

import "fmt"

// create a new type of `deck`
// which is a slice of strings
type deck []string

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
