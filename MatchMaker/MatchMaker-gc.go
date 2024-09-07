package main

import (
	"fmt"
	"strconv"
)

var desiredAnswers = []int{1, 5, 5, 5, 1}

// Function to validate user input
func checkResponse(response string) int {
	num, err := strconv.Atoi(response)
	for err != nil || num < 1 || num > 5 {
		fmt.Print("*!ERROR!* Input must be between 1 through 5: ")
		fmt.Scanln(&response)
		num, err = strconv.Atoi(response)
	}
	return num
}

// Function that asks the user the question and collects the response
func askQuestion(prompt string, index int) int {
	fmt.Print(prompt)
	var response string
	fmt.Scanln(&response)
	responseInt := checkResponse(response)
	score := 5 - abs(responseInt-desiredAnswers[index])
	return score
}

// function to get the absolute value of a number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(`
		~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		*                   *Instructions*                     *
		*  For each question, enter a number between 1 and 5.  *
		*                                                      *
		~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`)

	scores := []int{
		askQuestion("\n-Question 1- \nTrying new things is scary, doing the same comfortable things is best: ", 0),
		askQuestion("\n-Question 2- \nSpending time in nature is essential: ", 1),
		askQuestion("\n-Question 3- \nVolleyball is the best sport: ", 2),
		askQuestion("\n-Question 4- \nWinter is the best season: ", 3),
		askQuestion("\n-Question 5- \nLearning about history is boring: ", 4),
	}

	// Scoring summary and final message
	fmt.Println("\nSCORING SUMMARY:")
	for i := 0; i < len(scores); i++{
		fmt.Print("\nQuestion ", i + 1, " Score: ", scores[i] * 4)
	}

	totalScore := (scores[0] + scores[1] + scores[2] + scores[3] + scores[4]) * 4
	fmt.Printf("\n<YOU SCORED A %d OUT OF 100>\n", totalScore)

	if totalScore == 100 {
		fmt.Println("\n*!PERFECT SCORE. YOU'RE THE ONE!*")
	} else if totalScore >= 75 {
		fmt.Println("\nPretty close to perfect!")
	} else if totalScore >= 55 {
		fmt.Println("\nWe could make this work.")
	} else if totalScore >= 35 {
		fmt.Println("\nFriends?")
	} else {
		fmt.Println("\nNot meant to be :(")
	}
	fmt.Println()
}