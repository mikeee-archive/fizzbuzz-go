package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mikeee/fizzbuzz-go"
)

func main() {
	var input string
	fmt.Println(`Fizz🥤, Buzz⚡, FizzBuzz🥤⚡
	1) Number range of your choice
	2) Specify a number
	3) Random number chance!`)

	fmt.Print("Make a selection (enter the number then hit enter): ")
	fmt.Scanln(&input)
	choiceint, _ := strconv.Atoi(input)
	switch choiceint {
	case 1:

		fmt.Print("What number do you want me to start from? ")
		fmt.Scanln(&input)
		m, _ := strconv.Atoi(input)
		fmt.Print("What number do you want me to end on? ")
		fmt.Scanln(&input)
		n, _ := strconv.Atoi(input)
		type result struct {
			value  int
			result string
		}
		var rArray []result
		for i := m; i < n; i++ {
			rArray = append(rArray, result{value: i, result: fizzbuzz.Check(i)})
		}
		fmt.Println(rArray)
	case 2:
		fmt.Print("What's the lucky number you'd like me to check? ")
		fmt.Scanln(&input)
		x, _ := strconv.Atoi(input)
		fmt.Println(fizzbuzz.Check(x))
	case 3:
		seed := time.Now().Unix()
		random := rand.NewSource(seed)
		getRandom := rand.New(random)
		number := getRandom.Intn(9999)
		fmt.Printf("[%d] %s \n", number, fizzbuzz.Check(number))
	default:
		fmt.Println("Invalid option selected... exiting application")
	}
	fmt.Println("Au Revoir... until next time 👋")
}
