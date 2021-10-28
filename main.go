package main

import (
	"DragonRealm/gojaygo"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DisplayIntro() {
	fmt.Println("You are in a land full of dragons. In front of you, you see two caves. \n" +
		"In one cave, the dragon is friendly and will share his treasure with you. \n" +
		"The other dragon is greedy and hungry, and will eat you on sight.")
}

func ChooseCave() int {
	cave := ""

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Which cave would you like to enter? (1 or 2)")
		caveEntry, _ := reader.ReadString('\n')
		caveEntry = strings.TrimRight(caveEntry, "\r\n")

		if caveEntry == "1" || caveEntry == "2" {
			cave = caveEntry
			break
		} else {
			fmt.Println("Incorrect Entry. Try Again.")
		}
	}

	caveNum, err := strconv.Atoi(cave)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Something went wrong, let's just go with cave #1")
		caveNum = 1
	}

	return caveNum
}

func CheckCave(cave int) {
	random := gojaygo.NewRandomGenerator()

	fmt.Printf("You approach cave #%v...\n", cave)
	time.Sleep(2 * time.Second)

	fmt.Println("It is dark and spooky")
	time.Sleep(2 * time.Second)

	fmt.Println("A large dragon jumps out at you and...")
	time.Sleep(2 * time.Second)

	friendlyCave := random.NextInt(1, 2)

	if friendlyCave == cave {
		fmt.Println("Gives you his treasure!")
	} else {
		fmt.Println("Gobbles you down in one bite!")
	}
}

func main() {
	playAgain := "Y"

	reader := bufio.NewReader(os.Stdin)

	for ; playAgain == "Y"; {
		DisplayIntro()
		cave := ChooseCave()
		CheckCave(cave)

		fmt.Println("\nDo you want to play again? (Y or N)")
		playAgain, _ = reader.ReadString('\n')
		playAgain = strings.TrimRight(playAgain,"\r\n")
	}
}