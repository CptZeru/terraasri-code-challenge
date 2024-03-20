package challenges

import (
	"CptZeru/terraasri-code-challenge/constants"
	"CptZeru/terraasri-code-challenge/repository"
	"fmt"
	"regexp"
	"strings"
)

// DecodeMessage decode string contained emoji using emojiDecoder.
func DecodeMessage(message string) string {
	regex := regexp.MustCompile(string(constants.EmojiV140Std))
	return regex.ReplaceAllStringFunc(message, func(match string) string {
		return basicEmojiDecoder(match)
	})
}

// EncodeMessage encode plain string to emoji string.
func EncodeMessage(message string) string {
	words := strings.Fields(message)
	var encodedWords []string
	var skipCurrentLoop bool
	reversedEmojiDictionary := repository.GetReversedEmojiDictionary()
	for index, word := range words {
		// skip loop if word used previously for combinedWords
		if skipCurrentLoop {
			skipCurrentLoop = false
			continue
		}
		// checks for one word key
		val, ok := reversedEmojiDictionary[word]
		if ok {
			word = val
		} else {
			// checks for combined words key
			combinedWords := fmt.Sprintf("%v %v", word, words[index+1])
			val, ok = reversedEmojiDictionary[combinedWords]
			if ok {
				word = val
				skipCurrentLoop = true
			}
		}
		encodedWords = append(encodedWords, word)
	}
	return strings.Join(encodedWords[:], " ")
}

// UpdateEmojiDictionary update emoji dictionary.
func UpdateEmojiDictionary(emoji string, meaning string) string {
	err := repository.UpdateEmojiDictionaries(emoji, meaning)
	if err != nil {
		fmt.Println(err.Error())
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
