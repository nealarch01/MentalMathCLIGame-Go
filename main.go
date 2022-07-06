package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var mutex sync.Mutex

func askUserInput() int {
	var userInput string
	fmt.Print("Enter choice []: ")
	fmt.Scanln(&userInput)

	regex := regexp.MustCompile(`^[1-9]$`)

	if regex.MatchString(userInput) {
		i, _ := strconv.Atoi(userInput)
		return i
	}
	return -1
}

func addToQueue(exprQueue *ExprQueue, level int) {
	var expr Expression
	expr.Init(level)
	exprQueue.Push(expr)
}

func removeFromQueue(exprQueue *ExprQueue) {
	mutex.Lock()
	err := exprQueue.Pop()
	mutex.Unlock()

	if err != nil {
		fmt.Println("Queue Overflow")
		os.Exit(1)
	}
}

func isNumber(str string) bool {
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(str)
}

func playGame(level int) bool {
	var exprQueue ExprQueue
	var randomExpr Expression

	for i := 0; i < 3; i++ {
		randomExpr.Init(level)
		exprQueue.Push(randomExpr)
	}

	currentTime := time.Now().Unix()
	addTime := currentTime + 2
	endTime := time.Now().Unix() + 20

	var correct int = 0
	var incorrect int = 0
	var userInput string = ""
	var active bool = true
	var won bool = false

	var wg sync.WaitGroup

	for currentTime < endTime && active == true {
		wg.Add(2)

		go func() {
			defer wg.Done()

			if currentTime >= addTime {
				addTime = currentTime + 5
				addToQueue(&exprQueue, level)
			}
		}()

		go func() {
			defer wg.Done()

			if exprQueue.Count() >= 12 {
				active = false
				won = false
				return
			}

			topExpr, err := exprQueue.Top()
			if err != nil { // Queue underflow
				active = false
				won = true
				return
			}

			fmt.Print(topExpr.Display(), " = ")
			fmt.Scanln(&userInput)

			if !isNumber(userInput) {
				fmt.Println("Invalid input. Numbers only")
				return
			}

			i, _ := strconv.Atoi(userInput)

			if i == topExpr.CalcResult() {
				correct++
				removeFromQueue(&exprQueue)
			} else {
				incorrect++
			}
		}()

		wg.Wait()
		currentTime = time.Now().Unix()
	}

	fmt.Printf("Correct answers: %d", correct)
	fmt.Print("\n")
	fmt.Printf("Incorrect answers: %d", incorrect)
	fmt.Printf("\n")

	if won == true {
		fmt.Println("You won!")
		return true
	}

	fmt.Println("Game over :(")
	return false
}

func main() {
	fmt.Println("Welcome to the terminal game!")
	fmt.Println("[1.] Start Game")
	fmt.Println("[2.] How To Play")
	fmt.Println("[3.] Exit")
	selected := askUserInput()
	for selected > 3 || selected < 1 {
		fmt.Println("Invalid input. Enter a number between 1 and 3")
		selected = askUserInput()
	}

	var level int = 0

	if selected == 1 {
		fmt.Println("Difficulty (1 - easy, 2 - medium, 3 - hard)")
		level = askUserInput()
		for level > 3 || level < 1 {
			fmt.Println("Invalid input. Enter a number between 1 and 3")
			level = askUserInput()
		}
		playGame(level)
	} else if selected == 2 {
		fmt.Println("1. You will be shown a math expression")
		fmt.Println("2. Every few seconds, a new expression will be added into a queue")
		fmt.Println("3. If the queue reaches its max size (you did not answer fast enough, the game will end)")
		fmt.Println("4. To win, you must empty the queue completely or prevent the queue from filling for 60 seconds")
	} else {
		os.Exit(0)
	}

}
