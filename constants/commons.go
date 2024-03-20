package constants

type CMDType string

const (
	MenuList string = "Menu:\n1. Is Circular Palindrome\n2. Calculate Diagonal Sum\n140.85 Decode Encoded Message\n140.96 Add emoji to dictionary\n"
)

// List of Commands
const (
	CMDCircularPalindrome   CMDType = "1"
	CMDCalculateDiagonalSum CMDType = "2"
	CMDDecodeEmojiMessage   CMDType = "140.85"
	CMDAddEmojiDictionary   CMDType = "140.96"
)
