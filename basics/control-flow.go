package main

import "fmt"

func ifLoop() {
	var x int = 5
	if x > 5 {
		fmt.Println(x)
	}
}

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

//no break in the end, it'll break automatically
func switchFlow(x int) {
	switch x {
	case 1:
		fmt.Printf("case1")
	case -1:
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
	}
}

func multiSwitch(score int) {
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
}

//tagless switch
func taglessSwitch(x int) {
	switch {
	case x > 1:
		fmt.Printf("case1")
	case x < -1:
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
	}
}
