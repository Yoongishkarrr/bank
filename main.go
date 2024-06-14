package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear") 
	cmd.Stdout = os.Stdout	
	cmd.Run()
}

func animateText(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func main() {
	animateText("Welcome to Yoongishkarrr's bank")

	fmt.Println("Enter the amount on the card in sum: ")
	var amt int
	fmt.Scan(&amt)

	if amt > 100000000 {
		fmt.Println("The minimum amount should not exceed 100000000")
	} else {
		fmt.Println("Enter the currency (Dollar, Euro, Yuan): ")
		var cur string
		fmt.Scan(&cur)

		fmt.Println("Enter the sum of conversion: ")
		var num float32
		fmt.Scan(&num)

		var exchangeRate float32

		if num > float32(amt) {
			fmt.Println("Currency value cannot be greater than the amount.")
		} else {
			if cur == "Dollar" {
				exchangeRate = 12250.0
				fmt.Printf("%.2f Dollar\n", num/exchangeRate)
			} else if cur == "Euro" {
				exchangeRate = 13333.12
				fmt.Printf("%.2f Euro\n", num/exchangeRate)
			} else if cur == "Yuan" {
				exchangeRate = 1718.29
				fmt.Printf("%.2f Yuan\n", num/exchangeRate)
			} else {
				fmt.Println("Invalid Currency!")
			}
		}
		var remainingAmount = float32(amt) - num
		fmt.Printf("\nRemaining Amount is %f \n", remainingAmount)
	}
	fmt.Println("Thank you  for using our service!")
	time.Sleep(4 * time.Second)
	clearScreen()
	main()
}
