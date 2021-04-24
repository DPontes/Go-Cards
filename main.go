package main

import "fmt"

func main() {
  /*
    The following two lines do the same,
    but the first one will throw a "sugestion"
    in modern IDEs as it considers redundant to
    assign `string` to the variable as it's
    inferred by the RHL
  */
  // var card string= "Ace of Spades"
  card := "Ace of Spades"
  fmt.Println(card)
}
