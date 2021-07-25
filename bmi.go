package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator!")
	fmt.Println("--------------------")

	fmt.Println("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Println("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')

	fmt.Print(weightInput)
	fmt.Print(heightInput)
}