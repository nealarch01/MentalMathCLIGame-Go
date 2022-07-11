package main

import (
	"fmt"
	// "log"
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

func queueAdderThread(exprQueue *ExprQueue, level int, active *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// Code block for thread testing
	// lf, lf_err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 644)
	// if lf_err != nil {
	// 	fmt.Println(lf_err.Error())
	// 	*active = false
	// }
	// defer lf.Close()
	// log.SetOutput(lf)

	currentTime := time.Now().Unix()
	addTime := currentTime + 3
	endTime := currentTime + 60

	for currentTime < endTime && *active == true {
		if currentTime >= addTime {
			addTime = currentTime + 3
			addToQueue(exprQueue, level)
			// log.Println("Added expression to queue") // For thread testing
		}
		currentTime = time.Now().Unix()
	}
	// log.Println("Thread operations finished") // Also for thread testing
}

func answerThread(exprQueue *ExprQueue, correct *int, incorrect *int, active *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	currentTime := time.Now().Unix()
	endTime := time.Now().Unix() + 60

	var userInput string = ""

	for currentTime < endTime && *active == true {
		queueCount := exprQueue.Count()
		if queueCount >= 50 || queueCount == 0 {
			*active = false
			return
		}

		topExpr, err := exprQueue.Top()

		if err != nil {
			*active = false
			return
		}

		fmt.Print(topExpr.Display(), " = ")
		fmt.Scanln(&userInput)

		if !isNumber(userInput) {
			fmt.Println("Invalid input. Numbers only")
		}

		i, _ := strconv.Atoi(userInput)

		if i == topExpr.CalcResult() {
			*correct++
			removeFromQueue(exprQueue)
		} else {
			*incorrect++
		}

		currentTime = time.Now().Unix()
	}

}

func playGame(level int) bool {
	var exprQueue ExprQueue
	var randomExpr Expression

	for i := 0; i < 3; i++ {
		randomExpr.Init(level)
		exprQueue.Push(randomExpr)
	}

	var correct int = 0
	var incorrect int = 0

	var wg sync.WaitGroup

	var active bool = true

	wg.Add(2)

	go answerThread(&exprQueue, &correct, &incorrect, &active, &wg)
	go queueAdderThread(&exprQueue, level, &active, &wg)

	wg.Wait()

	fmt.Printf("Correct answers: %d", correct)
	fmt.Print("\n")
	fmt.Printf("Incorrect answers: %d", incorrect)
	fmt.Printf("\n")

	if exprQueue.Count() >= 12 {
		fmt.Println("Game over :(")
		return false
	}

	fmt.Println("You won!")
	return true
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
