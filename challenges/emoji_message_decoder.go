package challenges

import (
	"CptZeru/terraasri-code-challenge/constants"
	"CptZeru/terraasri-code-challenge/repository"
	"regexp"
)

// DecodeMessage decode emoji contained string using emojiDecoder.
func DecodeMessage(message string) string {
	regex := regexp.MustCompile(string(constants.EmojiV140Std))
	return regex.ReplaceAllStringFunc(message, func(match string) string {
		return basicEmojiDecoder(match)
	})
}

// UpdateEmojiDictionary update emoji dictionary.
func UpdateEmojiDictionary(emoji string, meaning string) string {
	err := repository.UpdateEmojiDictionary(emoji, meaning)
	if err != nil {
		return "Unable to update dictionary"
	}
	return "Dictionary Updated"
}

// basicEmojiDecoder basic emoji decoder replace emoji with its true meaning.
func basicEmojiDecoder(emoji string) string {
	emojiDictionary := repository.GetEmojiDictionary()
	val, ok := emojiDictionary[emoji]
	if ok {
		return val
	}
	return emoji
}
