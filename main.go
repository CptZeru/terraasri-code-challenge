package main

import (
	"CptZeru/terraasri-code-challenge/challenges"
	"CptZeru/terraasri-code-challenge/constants"
	"CptZeru/terraasri-code-challenge/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input string
	for input != "exit" {
		switch input {
		case string(constants.CMDDecodeEmojiMessage):
			input = utils.ReadStringInput(reader, "Enter Encoded Message: ")
			res := challenges.DecodeMessage(input)
			fmt.Println(res)
		case string(constants.CMDEncodeMessage):
			input = utils.ReadStringInput(reader, "Enter Message to Encode: ")
			res := challenges.EncodeMessage(input)
			fmt.Println(res)
		case string(constants.CMDAddEmojiDictionary):
			input = utils.ReadStringInput(reader, "Enter Emoji to dictionary: ")
			emoji := input
			input = utils.ReadStringInput(reader, "Enter Emoji meaning: ")
			meaning := input
			res := challenges.UpdateEmojiDictionary(emoji, meaning)
			fmt.Println(res)
		case string(constants.CMDCircularPalindrome):
			input = utils.ReadStringInput(reader, "Enter string to check if it's circular palindrome: ")
			res := challenges.IsCircularPalindrome(input)
			fmt.Printf("is %v a circular palindrome ? %v\n", input, res)
		case string(constants.CMDCalculateDiagonalSum):
			input = utils.ReadStringInput(reader, "Enter odd number for to sum diagonal spiral matrix: ")
			intInput, err := strconv.Atoi(input)
			if err != nil {
				fmt.Print(err)
			}
			res := challenges.CalculateSpiralDiagonalSum(intInput)
			fmt.Printf("diagonal sum of %vx%v spiral matrix is %v\n", input, input, res)
		default:
			fmt.Print(constants.MenuList)
		}
		input = utils.ReadStringInput(reader, "")
	}
	fmt.Println("Program exited.")
}
