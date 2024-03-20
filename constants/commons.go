package constants

type CMDType string

const (
	MenuList string = "Menu:\n1. Is Circular Palindrome\n2. Calculate Diagonal Sum\n3. Decode Encoded Message\n4. Add emoji to dictionary\n5. Encode Message\n"
)

// List of Commands
const (
	CMDCircularPalindrome        CMDType = "1"
	CMDCalculateDiagonalSum      CMDType = "2"
	CMDDecodeEmojiMessage        CMDType = "3"
	CMDSpecialDecodeEmojiMessage CMDType = "140.85"
	CMDAddEmojiDictionary        CMDType = "4"
	CMDSpecialAddEmojiDictionary CMDType = "140.96"
	CMDEncodeMessage             CMDType = "5"
	CMDSpecialEncodeMessage      CMDType = "148.41"
)

// Dictionary JSON Filenames
const (
	EmojiDictionaryFN         string = "./emoji_dictionary.json"
	ReversedEmojiDictionaryFN string = "./reversed_emoji_dictionary.json"
)
