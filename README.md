# Cards Game in Golang

The code is for the Go-Cards project, first project from the Udemy course [Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)

## Output

The application returns a randomized list of strings corresponding to cards (limited to the values of `Ace, one, two, three`)

### Example

```bash
0 Three of Diamonds
1 Ace of Diamonds
2 Three of Clubs
3 Four of Hearts
4 Ace of Spades
5 Four of Clubs
6 Four of Spades
7 Two of Spades
8 Two of Diamonds
9 Ace of Clubs
10 Two of Hearts
11 Three of Spades
12 Two of Clubs
13 Four of Diamonds
14 Three of Hearts
15 Ace of Hearts

## How-To

### Run the Application

```bash
go build main.go deck.go
./main
```

### Test

```bash
go test
```

## Notes

- The fact that the cards are a string concatenation of their suit and value makes it harder to assertain their value or suit separately without more convoluted string manipulation. The use of ´structs` would be a good idea here.
