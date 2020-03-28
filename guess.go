package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getNumberToGuess() int {
	rand.Seed(time.Now().Unix())

	return rand.Intn(10)
}

func getGuessedNumber() int {
	reader := bufio.NewReader(os.Stdin)
	readNumber, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	guessedNumber, err := strconv.Atoi(strings.TrimSpace(readNumber))

	if err != nil {
		log.Fatal(err)
	}

	return guessedNumber
}

func compare(numberToGuess, guessedNumber int) bool {
	result := false

	if guessedNumber > numberToGuess {
		fmt.Println("That's too much!")
	} else if guessedNumber < numberToGuess {
		fmt.Println("WROOOONG! Too low.")
	} else {
		fmt.Println("You did it!")
		result = true
	}

	return result
}

func main() {
	numberToGuess := getNumberToGuess()

	for true {
		guessedNumber := getGuessedNumber()

		if compare(numberToGuess, guessedNumber) == true {
			break
		}

	}
}
