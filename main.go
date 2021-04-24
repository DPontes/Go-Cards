package main

import "fmt"

func main() {
  /*
    The following two lines are equivalent,
    but the first one will throw a "sugestion"
    in modern IDEs as it considers redundant to
    assign `string` to the variable as it's
    inferred by the RHL
  */
  // var card string= newCard()
  card := newCard()   // := operator can only be used during new variable assignment

  fmt.Println(card)
}

func newCard() string {
  return "Five of Diamonds"
}
