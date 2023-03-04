package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you are opens")
	case 2:
		fmt.Println("You can move 2 places")
	case 3:
		fmt.Println("You can move 3 places")
		fallthrough
	case 4:
		fmt.Println("you can move 4 places")
		fallthrough
	case 5:
		fmt.Println("you can move 5 places")
	case 6:
		fmt.Println("You can move 6 places and re-roll the dice")
	default:
		fmt.Println("What was that!!")
	}
}
