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

	regex := regexp.MustCompile(`^[1|2|3]$`)

	if regex.MatchString(userInput) {
		i, _ := strconv.Atoi(userInput)
		return i
	}
	return -1
}

func addToQueue(exprQueue *ExprQueue) {
	var expr Expression
	expr.Init()
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

func playGame() bool {
	var exprQueue ExprQueue
	var randomExpr Expression

	for i := 0; i < 3; i++ {
		randomExpr.Init()
		exprQueue.Push(randomExpr)
	}

	currentTime := time.Now().Unix()
	addTime := currentTime + 5
	endTime := time.Now().Unix() + 30

	var count int = 0
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
				addToQueue(&exprQueue)
			}
		}()

		go func() {
			defer wg.Done()

			topExpr, err := exprQueue.Top()
			if err != nil {
				active = false
				return
			}

			if exprQueue.Count() >= 12 {
				active = false
				return
			}

			fmt.Print(topExpr.Display(), " = ")
			fmt.Scanln(&userInput)

			if !isNumber(userInput) {
				fmt.Println("Incorrect")
				return
			}

			i, _ := strconv.Atoi(userInput)

			if i == topExpr.CalcResult() {
				count++
				removeFromQueue(&exprQueue)
			}
		}()

		wg.Wait()
		currentTime = time.Now().Unix()
	}

	// for currentTime < endTime && active == true {
	// go func() {
	// 	currentTime = time.Now().Unix()
	// 	if currentTime >= addTime {
	// 		addTime = currentTime + 5
	// 		fmt.Println("Added to queue")
	// 		addToQueue(&exprQueue)
	// 	}
	// }()

	// 	if exprQueue.Count() >= 10 {
	// 		won = false
	// 		active = false
	// 		return
	// 		// break
	// 	}

	// 	topExpr, err := exprQueue.Top()
	// 	if err != nil {
	// 		won = true
	// 		active = false
	// 		return
	// 		// break
	// 	}

	// 	fmt.Print(topExpr.Display(), " = ")
	// 	fmt.Scanln(&userInput)

	// 	if !isNumber(userInput) {
	// 		fmt.Println("Incorrect")
	// 		// return
	// 	}

	// 	i, _ := strconv.Atoi(userInput)

	// 	if i == topExpr.CalcResult() {
	// 		count++
	// 		removeFromQueue(&exprQueue)
	// 	}
	// }()
	// }

	if won == true {
		fmt.Println("You've won!")
		fmt.Printf("Score: %d", count)
		return true
	}

	fmt.Println("Game over :(, practice makes perfect!")
	return false
}

func main() {
	fmt.Println("Welcome to the terminal game!")
	fmt.Println("[1.] Start Game")
	fmt.Println("[2.] How To Play")
	fmt.Println("[3.] Exit")
	selected := askUserInput()
	for selected == -1 {
		fmt.Println("Invalid input. Enter a number between 1 and 3")
		selected = askUserInput()
	}

	switch selected {
	case 1:
		playGame()
		break

	case 2:
		fmt.Println("1. You will be shown a math expression")
		fmt.Println("2. Every few seconds, a new expression will be added into a queue")
		fmt.Println("3. If the queue reaches its max size (you did not answer fast enough, the game will end)")
		fmt.Println("4. To win, you must empty the queue completely or prevent the queue from filling for 60 seconds")
		break

	default:
		os.Exit(0)
	}
}
