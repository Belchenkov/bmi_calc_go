package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator!"
const separator = "--------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightPrompt = "Please enter your height (m): "

func main() {
	fmt.Println(mainTitle)
	fmt.Println(separator)

	fmt.Println(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Println(heightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi)
}
